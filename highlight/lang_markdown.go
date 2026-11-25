package highlight

import (
	"bytes"
	"strings"
)

// Markdown według CommonMark Spec 0.31.2 i rozszerzeń GFM 0.29 (listy zadań); koloruje źródło dokumentu, bez
// parsowania. Wymagania: requirements/languages/markdown.md.
var langMarkdown = language{
	names: []string{"markdown", "md"},
	scan:  scanMarkdown,
}

// scanMarkdown koloruje wiersz po wierszu: nagłówki ATX (hl-kw), linie poziome i podkreślenia nagłówków
// (hl-meta), ogrodzenia bloków kodu (hl-meta) z treścią kolorowaną językiem z napisu informacyjnego albo jako
// hl-str, znaczniki cytatu i list (hl-op), definicje odsyłaczy, kod w wierszu, adresy odnośników i ucieczki.
func scanMarkdown(w *writer, src []byte) {
	var scratch backtickRuns
	for i := 0; i < len(src); {
		e := lineEnd(src, i)
		line := src[i:e]
		indent := 0
		for indent < len(line) && indent < 4 && line[indent] == ' ' {
			indent++
		}
		rest := line[indent:]
		switch {
		case indent < 4 && fenceLen(rest) > 0:
			i = markdownFence(w, src, i, e, indent)
			continue
		case indent < 4 && atxHeading(rest):
			w.emit(none, line[:indent])
			w.emit(keyword, rest)
		case indent < 4 && thematicBreak(rest):
			w.emit(none, line[:indent])
			w.emit(meta, rest)
		case indent < 4 && len(rest) > 1 && rest[0] == '<' && (rest[1] == '/' || rest[1] == '!' || isNameStart(rest, 1)):
			w.emit(none, line[:indent])
			markup{html: false}.scan(w, rest)
		default:
			markdownLine(w, line, &scratch)
		}
		i = e
		if i < len(src) {
			w.emit(none, src[i:i+1])
			i++
		}
	}
}

// fenceLen zwraca długość ogrodzenia ``` albo ~~~ (co najmniej 3) na początku s albo 0. Napis informacyjny po
// ogrodzeniu z odwrotnych apostrofów nie może ich zawierać.
func fenceLen(s []byte) int {
	if len(s) < 3 || s[0] != '`' && s[0] != '~' {
		return 0
	}
	n := run(s, 0, len(s), func(b byte) bool { return b == s[0] })
	if n < 3 || s[0] == '`' && bytes.IndexByte(s[n:], '`') >= 0 {
		return 0
	}
	return n
}

// markdownFence zapisuje blok kodu w ogrodzeniu, którego wiersz otwierający to src[i:e] z wcięciem indent,
// i zwraca początek wiersza po ogrodzeniu zamykającym (albo koniec danych). Szuka zamknięcia raz, wiersz po
// wierszu; treść koloruje język z pierwszego słowa napisu informacyjnego, a gdy go nie ma (albo to Markdown,
// co dałoby kolorowanie zagnieżdżone dowolnie głęboko), dostaje hl-str.
func markdownFence(w *writer, src []byte, i, e, indent int) int {
	open := src[i+indent : e]
	char, n := open[0], fenceLen(open)
	w.emit(none, src[i:i+indent])
	w.emit(meta, open)
	info := firstWord(open[n:])
	body := e
	if body < len(src) {
		body++
	}
	end, closeStart, closeEnd := len(src), len(src), len(src)
	for p := body; p < len(src); {
		q := lineEnd(src, p)
		line := src[p:q]
		k := 0
		for k < len(line) && k < 4 && line[k] == ' ' {
			k++
		}
		if k < 4 {
			m := run(line, k, len(line), func(b byte) bool { return b == char })
			if m >= n && len(bytes.Trim(line[k+m:], " \t\r")) == 0 {
				closeStart, closeEnd, end = p, q, q
				break
			}
		}
		if q >= len(src) {
			break
		}
		p = q + 1
	}
	if e < len(src) {
		w.emit(none, src[e:e+1])
	}
	content := src[min(body, closeStart):closeStart]
	var inner *language
	if len(info) > 0 && len(info) <= 32 {
		var buf [32]byte
		for k, c := range info {
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			buf[k] = c
		}
		inner = byName[string(buf[:len(info)])]
	}
	if inner != nil && inner.names[0] != "markdown" {
		inner.scan(w, content)
		w.close()
	} else {
		w.emit(str, content)
	}
	if closeStart < len(src) {
		w.emit(meta, src[closeStart:closeEnd])
	}
	if end < len(src) {
		w.emit(none, src[end:end+1])
		end++
	}
	return end
}

// firstWord zwraca pierwsze słowo napisu informacyjnego (bez odstępów wokół).
func firstWord(s []byte) []byte {
	k := 0
	for k < len(s) && (s[k] == ' ' || s[k] == '\t') {
		k++
	}
	e := k
	for e < len(s) && !isSpace(s[e]) {
		e++
	}
	return s[k:e]
}

// atxHeading: 1–6 znaków # i odstęp albo koniec wiersza.
func atxHeading(s []byte) bool {
	n := run(s, 0, 7, func(b byte) bool { return b == '#' })
	return n >= 1 && n <= 6 && (n == len(s) || s[n] == ' ' || s[n] == '\t' || s[n] == '\r')
}

// thematicBreak: linia pozioma (co najmniej trzy -, * albo _, ze spacjami) albo podkreślenie nagłówka (= albo -).
func thematicBreak(s []byte) bool {
	s = bytes.TrimRight(s, " \t\r")
	if len(s) == 0 {
		return false
	}
	c := s[0]
	if c == '=' {
		return len(bytes.Trim(s, "=")) == 0
	}
	if c != '-' && c != '*' && c != '_' {
		return false
	}
	count := 0
	for _, b := range s {
		switch b {
		case c:
			count++
		case ' ', '\t':
		default:
			return false
		}
	}
	return count >= 3
}

// markdownLine: znaczniki cytatu, jeden znacznik listy (z polem zadania), definicja odsyłacza, potem treść.
func markdownLine(w *writer, line []byte, scratch *backtickRuns) {
	k := 0
	for {
		s := k
		for k < len(line) && (line[k] == ' ' || line[k] == '\t') {
			k++
		}
		w.emit(none, line[s:k])
		if k < len(line) && line[k] == '>' {
			w.emit(operator, line[k:k+1])
			k++
			continue
		}
		break
	}
	if n := listMarker(line[k:]); n > 0 {
		w.emit(operator, line[k:k+n])
		k += n
		s := k
		for k < len(line) && (line[k] == ' ' || line[k] == '\t') {
			k++
		}
		w.emit(none, line[s:k])
		if len(line)-k >= 3 && line[k] == '[' && line[k+2] == ']' && strings.IndexByte(" xX", line[k+1]) >= 0 {
			w.emit(meta, line[k:k+3])
			k += 3
		}
	}
	if k < len(line) && line[k] == '[' {
		if n := bytes.Index(line[k:], []byte("]:")); n > 1 {
			w.emit(attribute, line[k:k+n+2])
			k += n + 2
			s := k
			for k < len(line) && (line[k] == ' ' || line[k] == '\t') {
				k++
			}
			w.emit(none, line[s:k])
			w.emit(str, line[k:])
			return
		}
	}
	markdownInline(w, line[k:], scratch)
}

// listMarker zwraca długość znacznika listy na początku s (-, *, +, albo 1–9 cyfr z . lub )) z odstępem po nim.
func listMarker(s []byte) int {
	n := 0
	switch {
	case len(s) > 0 && (s[0] == '-' || s[0] == '*' || s[0] == '+'):
		n = 1
	default:
		d := run(s, 0, 9, isDigit)
		if d == 0 || d >= len(s) || s[d] != '.' && s[d] != ')' {
			return 0
		}
		n = d + 1
	}
	if n < len(s) && s[n] != ' ' && s[n] != '\t' && s[n] != '\r' {
		return 0
	}
	return n
}

// markdownInline: kod w wierszu `…` (hl-str), adres odnośnika ](…) i autolink <…> (hl-str), ucieczki \* i
// odwołania znakowe (hl-esc). Kod w wierszu szuka zamknięcia o tej samej długości w obrębie wiersza; listę
// następnych serii tej samej długości liczy się raz, więc koszt jest liniowy.
// backtickRuns: bufory dla markdownInline, wspólne dla wszystkich wierszy jednego wywołania.
type backtickRuns struct {
	starts, lengths, closer, last []int
}

// find wypełnia starts (początki serii `), lengths i closer (koniec następnej serii tej samej długości albo -1).
// last[n] to indeks ostatnio widzianej serii długości n; po użyciu jest czyszczony, więc bufor zostaje gotowy.
func (b *backtickRuns) find(s []byte) {
	b.starts, b.lengths, b.closer = b.starts[:0], b.lengths[:0], b.closer[:0]
	for k := 0; k < len(s); {
		if s[k] != '`' {
			k++
			continue
		}
		n := run(s, k, len(s), func(c byte) bool { return c == '`' })
		b.starts, b.lengths = append(b.starts, k), append(b.lengths, n)
		k += n
	}
	for len(b.last) <= len(s) {
		b.last = append(b.last, -1)
	}
	for r := len(b.starts) - 1; r >= 0; r-- {
		c := b.last[b.lengths[r]]
		end := -1
		if c >= 0 {
			end = b.starts[c] + b.lengths[c]
		}
		b.closer = append(b.closer, end)
		b.last[b.lengths[r]] = r
	}
	for _, n := range b.lengths {
		b.last[n] = -1
	}
	for l, r := 0, len(b.closer)-1; l < r; l, r = l+1, r-1 {
		b.closer[l], b.closer[r] = b.closer[r], b.closer[l]
	}
}

func markdownInline(w *writer, s []byte, scratch *backtickRuns) {
	var runs, closer []int // początki serii odwrotnych apostrofów i końce serii zamykających
	if bytes.IndexByte(s, '`') >= 0 {
		scratch.find(s)
		runs, closer = scratch.starts, scratch.closer
	}
	runIndex := 0
	noParen := false // po nieudanym szukaniu ) nie ma go już w reszcie wiersza
	for k := 0; k < len(s); {
		b := s[k]
		switch {
		case b == '\\' && k+1 < len(s) && s[k+1] < 0x80 && strings.IndexByte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", s[k+1]) >= 0:
			w.emit(escape, s[k:k+2])
			k += 2
		case b == '&' && entityLen(s, k) > 0:
			n := entityLen(s, k)
			w.emit(escape, s[k:k+n])
			k += n
		case b == '`':
			for runIndex < len(runs) && runs[runIndex] < k {
				runIndex++
			}
			n := run(s, k, len(s), func(b byte) bool { return b == '`' })
			if runIndex < len(runs) && runs[runIndex] == k && closer[runIndex] > 0 {
				end := closer[runIndex]
				w.emit(str, s[k:end])
				k = end
			} else {
				w.emit(none, s[k:k+n])
				k += n
			}
		case b == ']' && k+1 < len(s) && s[k+1] == '(' && !noParen:
			w.emit(none, s[k:k+2])
			k += 2
			n := bytes.IndexByte(s[k:], ')')
			if n < 0 {
				noParen = true
				continue
			}
			w.emit(str, s[k:k+n])
			k += n
		case b == '<' && autolinkLen(s, k) > 0:
			n := autolinkLen(s, k)
			w.emit(str, s[k:k+n])
			k += n
		default:
			j := k + 1
			for j < len(s) && strings.IndexByte("\\&`]<", s[j]) < 0 {
				j++
			}
			w.emit(none, s[k:j])
			k = j
		}
	}
}

// autolinkLen zwraca długość autolinku <schemat:…> albo <adres@poczty> od s[k] albo 0. Szukanie kończy się
// na odstępie albo < lub >, więc każdy bajt jest sprawdzany przez najwyżej jedną próbę.
func autolinkLen(s []byte, k int) int {
	j := k + 1
	colon, at := false, false
	for j < len(s) && s[j] != '>' && s[j] != '<' && !isSpace(s[j]) {
		colon = colon || s[j] == ':'
		at = at || s[j] == '@'
		j++
	}
	if j >= len(s) || s[j] != '>' || j == k+1 || !colon && !at {
		return 0
	}
	return j + 1 - k
}
