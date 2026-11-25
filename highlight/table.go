package highlight

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// table opisuje język, który mieści się w skanerze tabelowym (ADR 0001, punkt 5; ADR 0002, punkt 6).
type table struct {
	words         map[string]class // słowa kluczowe, typy i nazwy wbudowane
	lineComments  []string         // przedrostki komentarzy do końca wiersza
	blockComments []blockComment
	quotes        []quote
	escape        func(src []byte, i int) int // długość sekwencji ucieczki od src[i] == '\\'; co najmniej 1
	operators     string                      // znaki, które dostają hl-op
	calls         bool                        // identyfikator przed ( dostaje hl-fn
}

type blockComment struct {
	open, close string
	nested      bool
}

// quote to rodzaj napisu.
type quote struct {
	open, close string
	escapes     bool // \ rozpoczyna sekwencję ucieczki
	multiline   bool // napis może przejść przez koniec wiersza
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

// scan koloruje src w jednym przebiegu. Każda gałąź zużywa co najmniej jeden bajt, a każde spojrzenie
// w przód kończy się na bajtach, które ta sama gałąź zużywa (ADR 0002, punkt 7).
func (t *table) scan(w *writer, src []byte) {
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
		if j := t.comment(src, i); j > i {
			w.emit(comment, src[i:j])
			i = j
			continue
		}
		if q := t.quoteAt(src, i); q != nil {
			i = t.scanQuote(w, src, i, q)
			continue
		}
		if isDigit(b) || b == '.' && i+1 < len(src) && isDigit(src[i+1]) {
			j := scanNumber(src, i)
			w.emit(number, src[i:j])
			i = j
			continue
		}
		r, size := utf8.DecodeRune(src[i:])
		if isIdentStart(r) {
			j := i + size
			for j < len(src) {
				r, size := utf8.DecodeRune(src[j:])
				if !isIdentPart(r) {
					break
				}
				j += size
			}
			c := t.words[string(src[i:j])]
			if c == none && t.calls && parenAfter(src, j) {
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
		w.emit(none, src[i:i+size])
		i += size
	}
}

// comment zwraca koniec komentarza zaczynającego się w src[i] albo i, gdy go tam nie ma.
func (t *table) comment(src []byte, i int) int {
	rest := src[i:]
	for _, p := range t.lineComments {
		if hasPrefix(rest, p) {
			if nl := bytes.IndexByte(rest, '\n'); nl >= 0 {
				return i + nl
			}
			return len(src)
		}
	}
	for _, c := range t.blockComments {
		if !hasPrefix(rest, c.open) {
			continue
		}
		j := i + len(c.open)
		if !c.nested {
			if end := bytes.Index(src[j:], []byte(c.close)); end >= 0 {
				return j + end + len(c.close)
			}
			return len(src)
		}
		for depth := 1; j < len(src); {
			switch {
			case hasPrefix(src[j:], c.close):
				j += len(c.close)
				if depth--; depth == 0 {
					return j
				}
			case hasPrefix(src[j:], c.open):
				j += len(c.open)
				depth++
			default:
				j++
			}
		}
		return len(src)
	}
	return i
}

func (t *table) quoteAt(src []byte, i int) *quote {
	for k := range t.quotes {
		if hasPrefix(src[i:], t.quotes[k].open) {
			return &t.quotes[k]
		}
	}
	return nil
}

// scanQuote zapisuje napis od src[i] i zwraca jego koniec. Niezamknięty napis jednowierszowy kończy się
// przed \n, wielowierszowy na końcu danych.
func (t *table) scanQuote(w *writer, src []byte, i int, q *quote) int {
	start := i
	j := i + len(q.open)
	for j < len(src) {
		if hasPrefix(src[j:], q.close) {
			j += len(q.close)
			w.emit(str, src[start:j])
			return j
		}
		switch src[j] {
		case '\n':
			if !q.multiline {
				w.emit(str, src[start:j])
				return j
			}
		case '\\':
			if q.escapes {
				n := t.escape(src, j)
				w.emit(str, src[start:j])
				w.emit(escape, src[j:j+n])
				j += n
				start = j
				continue
			}
		}
		j++
	}
	w.emit(str, src[start:j])
	return j
}

// scanNumber zwraca koniec liczby od src[i] (cyfra albo kropka przed cyfrą): przedrostki 0x, 0b, 0o,
// podkreślenia między cyframi, część ułamkowa, wykładnik e albo p (szesnastkowy) i przyklejony przyrostek
// z liter i cyfr (i w Go, u i L w C, f32 w Ruście).
func scanNumber(src []byte, i int) int {
	digits := func(j int, ok func(byte) bool) int {
		for j < len(src) && (ok(src[j]) || src[j] == '_') {
			j++
		}
		return j
	}
	j := i
	prefix := byte(0)
	if src[j] == '0' && j+1 < len(src) {
		prefix = src[j+1] | 0x20
	}
	switch prefix {
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
		if j < len(src) && src[j] == '.' && (j+1 >= len(src) || src[j+1] != '.') {
			j = digits(j+1, isDigit)
		}
		j = exponent(src, j, 'e')
	}
	for j < len(src) && (isDigit(src[j]) || src[j]|0x20 >= 'a' && src[j]|0x20 <= 'z' || src[j] == '_') {
		j++
	}
	return j
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

// backslashEscape to sekwencja ucieczki wspólna dla wielu języków: \ i jeden znak.
func backslashEscape(src []byte, i int) int {
	if i+1 >= len(src) || src[i+1] == '\n' {
		return 1
	}
	_, size := utf8.DecodeRune(src[i+1:])
	return 1 + size
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
