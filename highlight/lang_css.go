package highlight

import (
	"strings"
	"unicode/utf8"
)

// CSS według CSS Syntax Module Level 3 i CSS Selectors Level 4; słowa wartości z CSS Values and Units Level 4
// i CSS Cascading Level 4; wymagania: requirements/languages/css.md.
var langCSS = language{
	names: []string{"css"},
	scan:  scanCSS,
}

var (
	cssEscape  = escapeRule{continuation: true}.length
	cssStrings = []quote{{open: `"`, close: `"`, escapes: true}, {open: "'", close: "'", escapes: true}}
)

var cssQuotes = &table{escape: func(src []byte, i int) int {
	// \ i do sześciu cyfr szesnastkowych z jedną spacją po nich, albo \ i jeden znak.
	if n := run(src, i+1, 6, isHex); n > 0 {
		if i+1+n < len(src) && src[i+1+n] == ' ' {
			n++
		}
		return 1 + n
	}
	return cssEscape(src, i)
}}

// cssValueWords: słowa wartości, które dostają hl-builtin; bez rozróżniania wielkości liter.
var cssValueWords = &table{foldCase: true, words: words(wordGroup{builtin, `inherit initial unset revert revert-layer auto none normal
	transparent currentcolor bold italic solid dashed dotted block inline inline-block flex inline-flex grid
	contents absolute relative fixed static sticky hidden visible center left right top bottom both
	important`})}

// scanCSS rozróżnia selektory i deklaracje: identyfikator na początku instrukcji wewnątrz {} z dwukropkiem po
// nim to właściwość (hl-attr), a do ; albo } trwa jej wartość. Poza wartością nazwy to selektory: element
// (hl-tag), .klasa i #id (hl-type), :pseudoklasa i ::pseudoelement (hl-builtin).
func scanCSS(w *writer, src []byte) {
	depth := 0
	inValue := false  // wartość właściwości albo nagłówek reguły @
	stmtStart := true // początek instrukcji: po {, ; albo }
	for i := 0; i < len(src); {
		b := src[i]
		next := byte(0)
		if i+1 < len(src) {
			next = src[i+1]
		}
		switch {
		case isSpace(b):
			j := i + 1
			for j < len(src) && isSpace(src[j]) {
				j++
			}
			w.emit(none, src[i:j])
			i = j
			continue
		case b == '/' && next == '*':
			end := blockEnd(src, i+2, "", "*/", false)
			w.emit(comment, src[i:end])
			i = end
			continue
		case b == '"' || b == '\'':
			q := &cssStrings[0]
			if b == '\'' {
				q = &cssStrings[1]
			}
			i = cssQuotes.scanQuote(w, src, i, i, q, false, str)
		case b == '@' && cssIdentEnd(src, i+1) > i+1:
			j := cssIdentEnd(src, i+1)
			w.emit(keyword, src[i:j])
			i = j
			inValue = true
		case b == '{':
			depth++
			inValue, stmtStart = false, true
			w.emit(none, src[i:i+1])
			i++
			continue
		case b == '}':
			if depth > 0 {
				depth--
			}
			inValue, stmtStart = false, true
			w.emit(none, src[i:i+1])
			i++
			continue
		case b == ';':
			inValue, stmtStart = false, true
			w.emit(none, src[i:i+1])
			i++
			continue
		case b == '!' && cssIdentEnd(src, i+1) > i+1:
			j := cssIdentEnd(src, i+1)
			w.emit(keyword, src[i:j])
			i = j
		case b == '#' && i+1 < len(src) && cssNameEnd(src, i+1) > i+1:
			j := cssNameEnd(src, i+1)
			if inValue {
				w.emit(number, src[i:j]) // kolor #fff
			} else {
				w.emit(typeName, src[i:j])
			}
			i = j
		case isDigit(b) || b == '.' && isDigit(next) || (b == '+' || b == '-') && inValue && (isDigit(next) || next == '.'):
			j := i
			if b == '+' || b == '-' {
				j++
			}
			j = scanNumber(src, j, numberRule{})
			if j < len(src) && src[j] == '%' {
				j++
			}
			w.emit(number, src[i:j])
			i = j
		case b == '.' && !inValue && cssIdentEnd(src, i+1) > i+1:
			j := cssIdentEnd(src, i+1)
			w.emit(typeName, src[i:j])
			i = j
		case b == ':' && !inValue && (cssIdentEnd(src, i+1) > i+1 || next == ':'):
			j := i + 1
			if next == ':' {
				j++
			}
			j = cssIdentEnd(src, j)
			w.emit(builtin, src[i:j])
			i = j
		case cssIdentEnd(src, i) > i:
			j := cssIdentEnd(src, i)
			word := src[i:j]
			switch {
			case j < len(src) && src[j] == '(':
				w.emit(call, word)
				i = j
				if strings.EqualFold(string(word), "url") {
					i = cssURL(w, src, i)
				}
				stmtStart = false
				continue
			case !inValue && depth > 0 && stmtStart && colonAfter(src, j):
				w.emit(attribute, word)
				inValue = true
			case inValue:
				w.emit(cssValueWords.wordClass(word), word)
			default:
				w.emit(tag, word)
			}
			i = j
		case strings.IndexByte(">+~*/=", b) >= 0:
			w.emit(operator, src[i:i+1])
			i++
		default:
			_, size := utf8.DecodeRune(src[i:])
			w.emit(none, src[i:i+size])
			i += size
		}
		stmtStart = false
	}
}

// cssURL: po url( bez cudzysłowu adres do ) w tym samym wierszu dostaje hl-str.
func cssURL(w *writer, src []byte, i int) int {
	w.emit(none, src[i:i+1])
	j := i + 1
	if j < len(src) && (src[j] == '"' || src[j] == '\'') {
		return j
	}
	e := j
	for e < len(src) && src[e] != ')' && src[e] != '\n' {
		e++
	}
	w.emit(str, src[j:e])
	return e
}

// cssIdentEnd zwraca koniec identyfikatora CSS od src[i] albo i: --nazwa, -nazwa albo nazwa; litery, cyfry,
// - i _ oraz znaki spoza ASCII.
func cssIdentEnd(src []byte, i int) int {
	j := i
	if j < len(src) && src[j] == '-' {
		j++
		if j < len(src) && src[j] == '-' {
			return cssNameEnd(src, j+1)
		}
	}
	if j >= len(src) {
		return i
	}
	r, _ := utf8.DecodeRune(src[j:])
	if !isIdentStart(r) {
		return i
	}
	return cssNameEnd(src, j)
}

// cssNameEnd zwraca koniec ciągu znaków nazwy CSS (litery, cyfry, -, _, znaki spoza ASCII) od src[i].
func cssNameEnd(src []byte, i int) int {
	for i < len(src) {
		r, size := utf8.DecodeRune(src[i:])
		if r != '-' && !isIdentPart(r) {
			break
		}
		i += size
	}
	return i
}
