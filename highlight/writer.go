package highlight

import "bytes"

// writer składa wynik z tokenów. Pilnuje postaci z ADR 0002: elementy się nie zagnieżdżają, sąsiednie
// tokeny tej samej klasy dzielą jeden element, \n zamyka element, pusty token nie otwiera elementu.
type writer struct {
	out  []byte
	open class // klasa otwartego elementu albo none
}

// emit zapisuje tekst tokenu klasy c.
func (w *writer) emit(c class, text []byte) {
	for len(text) > 0 {
		line := text
		nl := bytes.IndexByte(text, '\n')
		if nl >= 0 {
			line = text[:nl]
		}
		if len(line) > 0 {
			if w.open != c {
				w.close()
				if c != none {
					w.out = append(w.out, `<span class="`...)
					w.out = append(w.out, classNames[c]...)
					w.out = append(w.out, `">`...)
					w.open = c
				}
			}
			w.out = appendEscaped(w.out, line)
		}
		if nl < 0 {
			return
		}
		w.close()
		w.out = append(w.out, '\n')
		text = text[nl+1:]
	}
}

// close zamyka otwarty element.
func (w *writer) close() {
	if w.open != none {
		w.out = append(w.out, "</span>"...)
		w.open = none
	}
}

// appendEscaped dopisuje tekst z zamianą &, < i > (OUT-001).
func appendEscaped(out, text []byte) []byte {
	start := 0
	for i, c := range text {
		var rep string
		switch c {
		case '&':
			rep = "&amp;"
		case '<':
			rep = "&lt;"
		case '>':
			rep = "&gt;"
		default:
			continue
		}
		out = append(out, text[start:i]...)
		out = append(out, rep...)
		start = i + 1
	}
	return append(out, text[start:]...)
}
