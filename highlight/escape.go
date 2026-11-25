package highlight

import "unicode/utf8"

// unbounded to „bez ograniczenia długości” w escapeRule.hexX i run.
const unbounded = 1 << 30

// escapeRule opisuje sekwencje ucieczki po \ w napisach danego języka. Sekwencja nieznana obejmuje \ i jeden
// znak; niepełna kończy się na ostatniej pasującej cyfrze. Każda sekwencja ma ograniczoną długość albo zużywa
// tylko bajty, których nie trzeba czytać drugi raz (ADR 0002, punkt 7).
type escapeRule struct {
	continuation bool // \ przed końcem wiersza to jedna sekwencja, napis ciągnie się dalej (C, JS, Python)
	hexX         int  // najwięcej cyfr po \x; 0: \x nie jest szczególne
	octal        int  // najwięcej cyfr ósemkowych po \; 0: brak
	u4           bool // \u i cztery cyfry szesnastkowe
	u8           bool // \U i osiem cyfr szesnastkowych
	uBraces      bool // \u{…} (JS, Rust)
	javaU        bool // \u, \uu, \uuu… i cztery cyfry (Java)
	named        bool // \N{nazwa} (Python)
}

// length zwraca długość sekwencji ucieczki od src[i] == '\\'.
func (r escapeRule) length(src []byte, i int) int {
	if i+1 >= len(src) {
		return 1
	}
	switch c := src[i+1]; {
	case c == '\n':
		if r.continuation {
			return 2
		}
		return 1
	case c == '\r' && r.continuation:
		if i+2 < len(src) && src[i+2] == '\n' {
			return 3
		}
		return 2
	case c == 'x' && r.hexX > 0:
		return 2 + run(src, i+2, r.hexX, isHex)
	case c == 'u' && r.uBraces && i+2 < len(src) && src[i+2] == '{':
		return braced(src, i, i+2)
	case c == 'u' && r.javaU:
		k := i + 1
		for k < len(src) && src[k] == 'u' {
			k++
		}
		return k - i + run(src, k, 4, isHex)
	case c == 'u' && r.u4:
		return 2 + run(src, i+2, 4, isHex)
	case c == 'U' && r.u8:
		return 2 + run(src, i+2, 8, isHex)
	case c == 'N' && r.named && i+2 < len(src) && src[i+2] == '{':
		return braced(src, i, i+2)
	case isOctal(c) && r.octal > 0:
		return 1 + run(src, i+1, r.octal, isOctal)
	}
	_, size := utf8.DecodeRune(src[i+1:])
	return 1 + size
}

// braced zwraca długość sekwencji od src[i] z nawiasem { w src[k]: do } w tym samym wierszu, najwyżej 64 bajty
// dalej; bez } sekwencją są tylko dwa pierwsze znaki.
func braced(src []byte, i, k int) int {
	for j := k + 1; j < len(src) && j <= k+64 && src[j] != '\n'; j++ {
		if src[j] == '}' {
			return j + 1 - i
		}
	}
	return 2
}
