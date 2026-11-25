package highlight

import "bytes"

// TOML według TOML v1.0.0; wymagania: requirements/languages/toml.md.
var langTOML = language{
	names: []string{"toml"},
	scan:  scanTOML,
}

var (
	tomlStrings = &table{escape: escapeRule{u4: true, u8: true}.length, quotes: tomlQuotes}
	tomlQuotes  = []quote{
		{open: `"""`, close: `"""`, escapes: true, multiline: true},
		{open: `'''`, close: `'''`, multiline: true},
		{open: `"`, close: `"`, escapes: true},
		{open: `'`, close: `'`},
	}
)

// scanTOML: nagłówek tabeli [a.b] i [[a]] na początku wiersza dostaje hl-tag, klucz przed = (także z kropkami
// i w cudzysłowie) hl-attr, liczby, daty i czasy hl-num, true i false hl-builtin.
func scanTOML(w *writer, src []byte) {
	lineStart := true
	for i := 0; i < len(src); {
		b := src[i]
		switch {
		case b == '\n':
			w.emit(none, src[i:i+1])
			i++
			lineStart = true
			continue
		case b == ' ' || b == '\t' || b == '\r':
			w.emit(none, src[i:i+1])
			i++
			continue
		case b == '#':
			e := lineEnd(src, i)
			w.emit(comment, src[i:e])
			i = e
		case b == '[' && lineStart:
			e := lineEnd(src, i)
			closing := []byte("]")
			if hasPrefix(src[i:], "[[") {
				closing = []byte("]]")
			}
			if n := bytes.Index(src[i:e], closing); n >= 0 {
				e = i + n + len(closing)
			}
			w.emit(tag, src[i:e])
			i = e
		case tomlStrings.quoteAt(src, i) != nil && (b == '"' || b == '\''):
			i = tomlString(w, src, i)
		case b == '=':
			w.emit(operator, src[i:i+1])
			i++
		case isTOMLBare(b):
			// Liczba, data i czas zaczynają się od cyfry albo znaku i mogą zawierać kropkę; klucz bez cudzysłowu
			// kończy się na kropce klucza złożonego.
			numeric := isDigit(b) || b == '+' || b == '-'
			j := i + 1
			for j < len(src) && (isTOMLBare(src[j]) || numeric && src[j] == '.') {
				j++
			}
			word := src[i:j]
			c := none
			switch {
			case numeric && tomlEqualAfter(src, j) || !numeric && tomlKeyAfter(src, j):
				c = attribute
			case string(word) == "true" || string(word) == "false":
				c = builtin
			case isDigit(word[0]) || len(word) > 1 && (word[0] == '+' || word[0] == '-') ||
				string(word) == "inf" || string(word) == "nan":
				c = number
			}
			w.emit(c, word)
			i = j
		default:
			w.emit(none, src[i:i+1])
			i++
		}
		lineStart = false
	}
}

// tomlString zapisuje napis od src[i]; napis przed = albo . to klucz w cudzysłowie.
func tomlString(w *writer, src []byte, i int) int {
	q := tomlStrings.quoteAt(src, i)
	c := str
	if end := tomlStrings.scanQuote(nil, src, i, i, q, false, str); tomlKeyAfter(src, end) {
		c = attribute
	}
	return tomlStrings.scanQuote(w, src, i, i, q, false, c)
}

// tomlKeyAfter mówi, czy po src[j:] (po spacjach i tabulatorach) stoi = albo kropka klucza złożonego.
func tomlKeyAfter(src []byte, j int) bool {
	for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
		j++
	}
	return j < len(src) && (src[j] == '=' || src[j] == '.')
}

// tomlEqualAfter mówi, czy po src[j:] (po spacjach i tabulatorach) stoi =.
func tomlEqualAfter(src []byte, j int) bool {
	for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
		j++
	}
	return j < len(src) && src[j] == '='
}

// isTOMLBare: znaki klucza bez cudzysłowu oraz liczb, dat i czasów.
func isTOMLBare(b byte) bool {
	return isDigit(b) || b|0x20 >= 'a' && b|0x20 <= 'z' || b == '_' || b == '-' || b == '+' || b == ':'
}
