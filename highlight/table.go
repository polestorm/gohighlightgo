package highlight

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// table opisuje język, który mieści się w skanerze tabelowym (ADR 0001, punkt 5; ADR 0002, punkt 6).
type table struct {
	words         map[string]class // słowa kluczowe, typy i nazwy wbudowane; przy foldCase małymi literami
	foldCase      bool             // słowa bez rozróżniania wielkości liter (SQL, Pascal)
	identExtra    string           // dodatkowe bajty identyfikatora poza literami, cyframi i _ (np. $ w JS)
	lineComments  []string         // przedrostki komentarzy do końca wiersza
	blockComments []blockComment
	quotes        []quote
	prefixes      []prefix                    // przedrostki doklejone do napisu: u8"…", r"…", b'…'
	escape        func(src []byte, i int) int // długość sekwencji ucieczki od src[i] == '\\'; co najmniej 1
	numbers       numberRule
	operators     string // znaki, które dostają hl-op
	calls         bool   // identyfikator przed ( dostaje hl-fn
	macros        bool   // identyfikator przed ! (nie !=) dostaje hl-fn (makra w Ruście)
	decorators    bool   // @nazwa dostaje hl-meta (adnotacje i dekoratory)
	templates     bool   // `…${…}…` z wyrażeniami wewnątrz (JS, TS)
	keys          bool   // napis przed dwukropkiem dostaje hl-attr (klucze JSON)

	// special obsługuje konstrukcje spoza tabeli (preprocesor, napisy surowe z ogranicznikiem, czasy życia).
	// Woła się go na początku każdego tokenu; zwraca koniec tego, co zapisał, albo i, gdy nic nie rozpoznał.
	special func(t *table, w *writer, src []byte, i int) int
}

type blockComment struct {
	open, close string
	nested      bool
}

// quote to rodzaj napisu.
type quote struct {
	open, close string
	escapes     bool // \ rozpoczyna sekwencję ucieczki
	doubled     bool // podwojony ogranicznik zamykający to znak ucieczki ('' w SQL i Pascalu)
	multiline   bool // napis może przejść przez koniec wiersza
	ident       bool // to nazwa w cudzysłowie, nie napis: bez klasy ("kolumna" w SQL)
}

// prefix to przedrostek napisu; raw wyłącza sekwencje ucieczki (\ chroni tylko następny znak, jak w Pythonie).
type prefix struct {
	text string
	raw  bool
}

// numberRule opisuje różnice w zapisie liczb między językami.
type numberRule struct {
	noLeadingDot bool // liczba nie może zaczynać się od kropki (Rust, Pascal)
	dotIdent     bool // kropka przed identyfikatorem kończy liczbę: 1.max(2) w Ruście
	quoteSep     bool // apostrof między cyframi jest separatorem (C23, C++14)
}

// wordGroup łączy słowa rozdzielone odstępami z ich klasą; words buduje z grup mapę dla table.words.
type wordGroup struct {
	class class
	words string
}

func words(groups ...wordGroup) map[string]class {
	m := map[string]class{}
	for _, g := range groups {
		for _, w := range strings.Fields(g.words) {
			m[w] = g.class
		}
	}
	return m
}

// foldWords zamienia klucze mapy na małe litery, dla table.foldCase.
func foldWords(m map[string]class) map[string]class {
	out := make(map[string]class, len(m))
	for k, v := range m {
		out[strings.ToLower(k)] = v
	}
	return out
}

// scan koloruje src w jednym przebiegu. Każda gałąź zużywa co najmniej jeden bajt, a każde spojrzenie
// w przód kończy się na bajtach, które ta sama gałąź zużywa albo które są potem zużywane raz (ADR 0002, punkt 7).
func (t *table) scan(w *writer, src []byte) {
	// templates: dla każdego otwartego ${ w szablonie liczba otwartych w nim nawiasów {.
	var templates []int
	for i := 0; i < len(src); {
		b := src[i]
		if isSpace(b) {
			j := i + 1
			for j < len(src) && isSpace(src[j]) {
				j++
			}
			w.emit(none, src[i:j])
			i = j
			continue
		}
		if t.special != nil {
			if j := t.special(t, w, src, i); j > i {
				i = j
				continue
			}
		}
		if j := t.comment(src, i); j > i {
			w.emit(comment, src[i:j])
			i = j
			continue
		}
		if t.templates {
			switch {
			case b == '`':
				var open bool
				i, open = t.scanTemplate(w, src, i, i+1)
				if open {
					templates = append(templates, 0)
				}
				continue
			case b == '{' && len(templates) > 0:
				templates[len(templates)-1]++
			case b == '}' && len(templates) > 0:
				if top := len(templates) - 1; templates[top] > 0 {
					templates[top]--
				} else {
					templates = templates[:top]
					w.emit(escape, src[i:i+1])
					var open bool
					i, open = t.scanTemplate(w, src, i+1, i+1)
					if open {
						templates = append(templates, 0)
					}
					continue
				}
			}
		}
		if q := t.quoteAt(src, i); q != nil {
			c := str
			if t.keys && colonAfter(src, t.scanQuote(nil, src, i, i, q, false, str)) {
				c = attribute
			}
			i = t.scanQuote(w, src, i, i, q, false, c)
			continue
		}
		if isDigit(b) || !t.numbers.noLeadingDot && b == '.' && i+1 < len(src) && isDigit(src[i+1]) {
			j := scanNumber(src, i, t.numbers)
			w.emit(number, src[i:j])
			i = j
			continue
		}
		if t.decorators && b == '@' && i+1 < len(src) {
			if j := t.identEnd(src, i+1); j > i+1 {
				w.emit(meta, src[i:j])
				i = j
				continue
			}
		}
		if j := t.identEnd(src, i); j > i {
			if len(t.prefixes) > 0 {
				if q := t.quoteAt(src, j); q != nil {
					if p := t.prefixOf(src[i:j]); p != nil {
						i = t.scanQuote(w, src, i, j, q, p.raw, str)
						continue
					}
				}
			}
			c := t.wordClass(src[i:j])
			if c == none && t.calls && parenAfter(src, j) {
				c = call
			}
			if c == none && t.macros && j+1 < len(src) && src[j] == '!' && src[j+1] != '=' {
				c = call
			}
			w.emit(c, src[i:j])
			i = j
			continue
		}
		if b < utf8.RuneSelf && strings.IndexByte(t.operators, b) >= 0 {
			w.emit(operator, src[i:i+1])
			i++
			continue
		}
		_, size := utf8.DecodeRune(src[i:])
		w.emit(none, src[i:i+size])
		i += size
	}
}

// identEnd zwraca koniec identyfikatora od src[i] albo i, gdy go tam nie ma.
func (t *table) identEnd(src []byte, i int) int {
	r, size := utf8.DecodeRune(src[i:])
	if !isIdentStart(r) && !t.isExtra(r) {
		return i
	}
	j := i + size
	for j < len(src) {
		r, size := utf8.DecodeRune(src[j:])
		if !isIdentPart(r) && !t.isExtra(r) {
			break
		}
		j += size
	}
	return j
}

func (t *table) isExtra(r rune) bool {
	return r < utf8.RuneSelf && t.identExtra != "" && strings.IndexByte(t.identExtra, byte(r)) >= 0
}

// wordClass zwraca klasę słowa z tabeli. Przy foldCase zamienia litery ASCII na małe w buforze na stosie:
// słowa z tabel są krótsze niż 32 bajty i całe w ASCII.
func (t *table) wordClass(word []byte) class {
	if !t.foldCase {
		return t.words[string(word)]
	}
	var buf [32]byte
	if len(word) > len(buf) {
		return none
	}
	for k, c := range word {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		buf[k] = c
	}
	return t.words[string(buf[:len(word)])]
}

func (t *table) prefixOf(word []byte) *prefix {
	for k := range t.prefixes {
		if string(word) == t.prefixes[k].text {
			return &t.prefixes[k]
		}
	}
	return nil
}

// comment zwraca koniec komentarza zaczynającego się w src[i] albo i, gdy go tam nie ma.
func (t *table) comment(src []byte, i int) int {
	rest := src[i:]
	for _, p := range t.lineComments {
		if hasPrefix(rest, p) {
			return lineEnd(src, i)
		}
	}
	for _, c := range t.blockComments {
		if hasPrefix(rest, c.open) {
			return blockEnd(src, i+len(c.open), c.open, c.close, c.nested)
		}
	}
	return i
}

// blockEnd zwraca koniec bloku od src[j] (za ogranicznikiem otwierającym): za pierwszym close, a przy
// nested za close, które wyrównuje licznik otwarć. Niezamknięty blok sięga końca danych.
func blockEnd(src []byte, j int, open, close string, nested bool) int {
	if !nested {
		if end := bytes.Index(src[j:], []byte(close)); end >= 0 {
			return j + end + len(close)
		}
		return len(src)
	}
	for depth := 1; j < len(src); {
		switch {
		case hasPrefix(src[j:], close):
			j += len(close)
			if depth--; depth == 0 {
				return j
			}
		case hasPrefix(src[j:], open):
			j += len(open)
			depth++
		default:
			j++
		}
	}
	return len(src)
}

// lineEnd zwraca pozycję najbliższego \n od src[i] albo koniec danych.
func lineEnd(src []byte, i int) int {
	if nl := bytes.IndexByte(src[i:], '\n'); nl >= 0 {
		return i + nl
	}
	return len(src)
}

func (t *table) quoteAt(src []byte, i int) *quote {
	for k := range t.quotes {
		if hasPrefix(src[i:], t.quotes[k].open) {
			return &t.quotes[k]
		}
	}
	return nil
}

// scanQuote zapisuje napis, który zaczyna się w src[start] (przedrostek), a ogranicznikiem w src[i], i zwraca
// jego koniec. Niezamknięty napis jednowierszowy kończy się przed \n, wielowierszowy na końcu danych.
// Przy w == nil tylko mierzy. Nazwa w cudzysłowie (quote.ident) zostaje bez klasy.
func (t *table) scanQuote(w *writer, src []byte, start, i int, q *quote, raw bool, c class) int {
	if q.ident {
		c = none
	}
	emit := func(c class, text []byte) {
		if w != nil {
			w.emit(c, text)
		}
	}
	j := i + len(q.open)
	for j < len(src) {
		if hasPrefix(src[j:], q.close) {
			if q.doubled && hasPrefix(src[j+len(q.close):], q.close) {
				emit(c, src[start:j])
				emit(escape, src[j:j+2*len(q.close)])
				j += 2 * len(q.close)
				start = j
				continue
			}
			j += len(q.close)
			emit(c, src[start:j])
			return j
		}
		switch src[j] {
		case '\n':
			if !q.multiline {
				emit(c, src[start:j])
				return j
			}
		case '\\':
			if !q.escapes {
				break
			}
			if raw {
				if j+1 < len(src) && src[j+1] != '\n' {
					j++ // w napisie surowym \ chroni następny bajt, ale nie jest sekwencją ucieczki
				}
				break
			}
			n := t.escape(src, j)
			emit(c, src[start:j])
			emit(escape, src[j:j+n])
			j += n
			start = j
			continue
		}
		j++
	}
	emit(c, src[start:j])
	return j
}

// scanTemplate zapisuje szablon JS od src[start] (od ` albo od miejsca za } zamykającym wyrażenie), z treścią
// od src[j]. Zwraca koniec i open == true, gdy zatrzymał się na ${, czyli dalej jest wyrażenie.
func (t *table) scanTemplate(w *writer, src []byte, start, j int) (end int, open bool) {
	for j < len(src) {
		switch {
		case src[j] == '`':
			w.emit(str, src[start:j+1])
			return j + 1, false
		case src[j] == '$' && j+1 < len(src) && src[j+1] == '{':
			w.emit(str, src[start:j])
			w.emit(escape, src[j:j+2])
			return j + 2, true
		case src[j] == '\\':
			n := t.escape(src, j)
			w.emit(str, src[start:j])
			w.emit(escape, src[j:j+n])
			j += n
			start = j
			continue
		}
		j++
	}
	w.emit(str, src[start:j])
	return j, false
}

// scanNumber zwraca koniec liczby od src[i] (cyfra albo kropka przed cyfrą): przedrostki 0x, 0b, 0o,
// podkreślenia między cyframi, część ułamkowa, wykładnik e albo p (szesnastkowy) i przyklejony przyrostek
// z liter i cyfr (i w Go, u i L w C, f32 w Ruście, n w JS, j w Pythonie).
func scanNumber(src []byte, i int, rule numberRule) int {
	digits := func(j int, ok func(byte) bool) int {
		for j < len(src) {
			switch {
			case ok(src[j]) || src[j] == '_':
			case rule.quoteSep && src[j] == '\'' && j > i && j+1 < len(src) && ok(src[j+1]):
			default:
				return j
			}
			j++
		}
		return j
	}
	j := i
	kind := byte(0)
	if src[j] == '0' && j+1 < len(src) {
		kind = src[j+1] | 0x20
	}
	switch kind {
	case 'x':
		j = digits(j+2, isHex)
		if j < len(src) && src[j] == '.' {
			j = digits(j+1, isHex)
		}
		j = exponent(src, j, 'p')
	case 'b', 'o':
		j = digits(j+2, isDigit)
	default:
		j = digits(j, isDigit)
		if j < len(src) && src[j] == '.' && fractionAfter(src, j, rule) {
			j = digits(j+1, isDigit)
		}
		j = exponent(src, j, 'e')
	}
	for j < len(src) && (isDigit(src[j]) || src[j]|0x20 >= 'a' && src[j]|0x20 <= 'z' || src[j] == '_') {
		j++
	}
	return j
}

// fractionAfter mówi, czy kropka w src[j] należy do liczby: nie, gdy to początek .. (zakres) albo, przy
// dotIdent, gdy po niej stoi identyfikator.
func fractionAfter(src []byte, j int, rule numberRule) bool {
	if j+1 >= len(src) {
		return true
	}
	next := src[j+1]
	if next == '.' {
		return false
	}
	return !rule.dotIdent || !(next == '_' || next|0x20 >= 'a' && next|0x20 <= 'z')
}

// exponent zużywa wykładnik (litera e lub p bez względu na wielkość, znak, cyfry), jeśli stoi w src[j].
func exponent(src []byte, j int, letter byte) int {
	if j >= len(src) || src[j]|0x20 != letter {
		return j
	}
	k := j + 1
	if k < len(src) && (src[k] == '+' || src[k] == '-') {
		k++
	}
	if k >= len(src) || !isDigit(src[k]) {
		return j
	}
	for k < len(src) && (isDigit(src[k]) || src[k] == '_') {
		k++
	}
	return k
}

// parenAfter mówi, czy po src[j:] (po spacjach i tabulatorach) stoi nawias otwierający.
func parenAfter(src []byte, j int) bool {
	for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
		j++
	}
	return j < len(src) && src[j] == '('
}

// colonAfter mówi, czy po src[j:] (po spacjach i tabulatorach) stoi dwukropek.
func colonAfter(src []byte, j int) bool {
	for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
		j++
	}
	return j < len(src) && src[j] == ':'
}

// atLineStart mówi, czy przed src[i] w tym wierszu stoją tylko spacje i tabulatory. Cofa się tylko przez
// odstęp, który właśnie został zużyty jako jeden token, więc łączny koszt jest liniowy.
func atLineStart(src []byte, i int) bool {
	for i > 0 && (src[i-1] == ' ' || src[i-1] == '\t') {
		i--
	}
	return i == 0 || src[i-1] == '\n'
}

// run liczy co najwyżej max bajtów od src[i] spełniających ok.
func run(src []byte, i, max int, ok func(byte) bool) int {
	n := 0
	for n < max && i+n < len(src) && ok(src[i+n]) {
		n++
	}
	return n
}

func hasPrefix(b []byte, s string) bool { return len(b) >= len(s) && string(b[:len(s)]) == s }

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '\f' || b == '\v'
}

func isDigit(b byte) bool { return b >= '0' && b <= '9' }

func isOctal(b byte) bool { return b >= '0' && b <= '7' }

func isBinary(b byte) bool { return b == '0' || b == '1' }

func isHex(b byte) bool { return isDigit(b) || b|0x20 >= 'a' && b|0x20 <= 'f' }

func isIdentStart(r rune) bool {
	if r < utf8.RuneSelf {
		return r == '_' || r|0x20 >= 'a' && r|0x20 <= 'z'
	}
	return unicode.IsLetter(r)
}

func isIdentPart(r rune) bool {
	if r < utf8.RuneSelf {
		return r == '_' || r >= '0' && r <= '9' || r|0x20 >= 'a' && r|0x20 <= 'z'
	}
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
