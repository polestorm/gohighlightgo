package highlight

import (
	"bytes"
	"unicode/utf8"
)

// Haki dla Rusta (definicja w lang/rust.hl; wymagania: requirements/languages/rust.md) i hak apostrophe.

// rustCharQuote: literał znakowy z sekwencjami ucieczki według tabeli języka.
var rustCharQuote = quote{open: "'", close: "'", escapes: true}

// apostropheSpecial: apostrof jako literał znakowy 'a' albo nazwa 'nazwa (hl-meta), jak w Ruście, bez reszty
// haka rust; dla języków z symbolami i etykietami ('sym w Scali).
func apostropheSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '\'' {
		return i
	}
	return rustChar(t, w, src, i, i)
}

// rustSpecial: atrybuty #[…] i #![…] (znacznik i nazwa dostają hl-meta), literały znakowe 'a' i b'a',
// czasy życia i etykiety 'a (hl-meta) oraz napisy surowe r"…", r#"…"#, br…, cr… z dowolną liczbą #.
func rustSpecial(t *table, w *writer, src []byte, i int) int {
	switch b := src[i]; {
	case b == '#':
		k := i + 1
		if k < len(src) && src[k] == '!' {
			k++
		}
		if k >= len(src) || src[k] != '[' {
			return i
		}
		end := t.identEnd(src, k+1)
		w.emit(meta, src[i:end])
		return end
	case b == '\'':
		return rustChar(t, w, src, i, i)
	case b == 'b' && i+1 < len(src) && src[i+1] == '\'':
		return rustChar(t, w, src, i, i+1)
	case b == 'r' || b == 'b' || b == 'c':
		return rustRaw(w, src, i)
	}
	return i
}

// rustChar rozstrzyga apostrof w src[q] (z przedrostkiem od src[start]): literał znakowy, gdy po jednym
// znaku albo sekwencji ucieczki stoi apostrof; czas życia, gdy stoi identyfikator; inaczej nic.
func rustChar(t *table, w *writer, src []byte, start, q int) int {
	if q+1 >= len(src) {
		return start
	}
	if src[q+1] == '\\' {
		return t.scanQuote(w, src, start, q, &rustCharQuote, false, str)
	}
	r, size := utf8.DecodeRune(src[q+1:])
	if r != '\n' && q+1+size < len(src) && src[q+1+size] == '\'' {
		end := q + 2 + size
		w.emit(str, src[start:end])
		return end
	}
	if start == q {
		if j := t.identEnd(src, q+1); j > q+1 {
			w.emit(meta, src[q:j])
			return j
		}
	}
	return start
}

// rustRaw zapisuje napis surowy od src[i], jeśli tam jest: [b|c]r, n znaków #, cudzysłów; koniec to cudzysłów
// i n znaków #. Każdy cudzysłów w treści jest sprawdzany raz, więc koszt jest liniowy.
func rustRaw(w *writer, src []byte, i int) int {
	k := i
	if src[k] == 'b' || src[k] == 'c' {
		k++
	}
	if k >= len(src) || src[k] != 'r' {
		return i
	}
	k++
	hashes := 0
	for k+hashes < len(src) && src[k+hashes] == '#' {
		hashes++
	}
	if k+hashes >= len(src) || src[k+hashes] != '"' {
		return i
	}
	end := len(src)
	for p := k + hashes + 1; p < len(src); {
		n := bytes.IndexByte(src[p:], '"')
		if n < 0 {
			break
		}
		q := p + n + 1
		if run(src, q, hashes, func(b byte) bool { return b == '#' }) == hashes {
			end = q + hashes
			break
		}
		p = q
	}
	w.emit(str, src[i:end])
	return end
}
