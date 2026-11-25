package highlight

import "bytes"

// Haki dla COBOL-a, plików wsadowych Windows i SPARQL (definicje w lang/cobol.hl, batch.hl, sparql.hl).

// cobolFixed: w formacie stałym (ISO/IEC 1989:2023, 6.1) znak * albo / w kolumnie 7 robi z wiersza komentarz;
// sprawdzane tylko na początku wiersza, gdy kolumny 1–6 to cyfry albo odstępy (numer sekwencyjny).
func cobolFixed(t *table, w *writer, src []byte, i int) int {
	if !(i == 0 || src[i-1] == '\n') {
		return i
	}
	e := lineEnd(src, i)
	line := src[i:e]
	if len(line) < 7 || line[6] != '*' && line[6] != '/' {
		return i
	}
	for _, b := range line[:6] {
		if !isDigit(b) && b != ' ' {
			return i
		}
	}
	w.emit(comment, line)
	return e
}

// batchSpecial: zmienne %NAZWA%, %~dp0, %1, %%i i !NAZWA! (z opóźnionym rozwijaniem) dostają hl-esc; etykieta
// :nazwa na początku wiersza (po odstępach) dostaje hl-fn; komentarz :: na początku wiersza dostaje hl-com
// (Windows Commands reference: set, call, for, goto).
func batchSpecial(t *table, w *writer, src []byte, i int) int {
	b := src[i]
	switch {
	case b == ':' && atLineStart(src, i):
		if i+1 < len(src) && src[i+1] == ':' {
			e := lineEnd(src, i)
			w.emit(comment, src[i:e])
			return e
		}
		if j := t.identEnd(src, i+1); j > i+1 {
			w.emit(call, src[i:j])
			return j
		}
	case b == '%' && i+1 < len(src):
		next := src[i+1]
		switch {
		case next == '%':
			j := i + 2
			if j < len(src) && (isWordByte(src[j]) || src[j] == '~') {
				for j < len(src) && (isWordByte(src[j]) || src[j] == '~') {
					j++
				}
				w.emit(escape, src[i:j])
				return j
			}
		case isDigit(next) || next == '*':
			w.emit(escape, src[i:i+2])
			return i + 2
		case next == '~':
			j := i + 2
			for j < len(src) && j < i+16 && isWordByte(src[j]) {
				j++
			}
			w.emit(escape, src[i:j])
			return j
		default:
			if n := closingWithin(src, i, '%', 64); n > 0 {
				w.emit(escape, src[i:i+n])
				return i + n
			}
		}
	case b == '!' && i+1 < len(src) && isWordByte(src[i+1]):
		if n := closingWithin(src, i, '!', 64); n > 0 {
			w.emit(escape, src[i:i+n])
			return i + n
		}
	}
	return i
}

// closingWithin zwraca długość od src[i] do znaku c zamykającego nazwę (bez odstępów i końca wiersza), najwyżej
// limit bajtów dalej, albo 0.
func closingWithin(src []byte, i int, c byte, limit int) int {
	for j := i + 1; j < len(src) && j < i+limit; j++ {
		switch src[j] {
		case c:
			if j > i+1 {
				return j + 1 - i
			}
			return 0
		case ' ', '\t', '\n', '\r', '"':
			return 0
		}
	}
	return 0
}

// iriSpecial: IRI w nawiasach ostrych <http://…> (SPARQL 1.1, 19.5 IRIREF): < i znaki bez odstępów, <, >, " do >
// w tym samym wierszu dostają hl-str; < z odstępem po nim jest operatorem.
func iriSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '<' || i+1 >= len(src) || isSpace(src[i+1]) || src[i+1] == '=' {
		return i
	}
	// Szukanie kończy się na pierwszym znaku, którego w IRI nie ma (także \n), więc każdy bajt jest czytany przez
	// najwyżej jedną próbę.
	n := bytes.IndexAny(src[i+1:], "<> \t\r\n\"{}|^`")
	if n < 0 || src[i+1+n] != '>' {
		return i
	}
	w.emit(str, src[i:i+n+2])
	return i + n + 2
}

// remComment: słowo REM (bez względu na wielkość liter) z odstępem albo końcem wiersza po nim zaczyna komentarz
// do końca wiersza (plik wsadowy, Visual Basic); słowo, które tylko zaczyna się od rem (remove), nie.
func remComment(t *table, w *writer, src []byte, i int) int {
	if i+3 > len(src) || src[i]|0x20 != 'r' || src[i+1]|0x20 != 'e' || src[i+2]|0x20 != 'm' {
		return i
	}
	if i > 0 && isWordByte(src[i-1]) || i+3 < len(src) && !isSpace(src[i+3]) {
		return i
	}
	e := lineEnd(src, i)
	w.emit(comment, src[i:e])
	return e
}
