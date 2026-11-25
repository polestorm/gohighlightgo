package highlight

import "bytes"

// Haki dla C i C++ (definicje w lang/c.hl i lang/cpp.hl; wymagania: requirements/languages/c.md i cpp.md).

// cPreprocessor koloruje dyrektywę preprocesora: # na początku wiersza i nazwa dyrektywy dostają hl-meta,
// a nazwa pliku w <…> po #include, #include_next, #import i #embed dostaje hl-str. Reszta wiersza to zwykłe
// tokeny.
func cPreprocessor(t *table, w *writer, src []byte, i int) int {
	if src[i] != '#' || !atLineStart(src, i) {
		return i
	}
	j := i + 1
	for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
		j++
	}
	k := t.identEnd(src, j)
	w.emit(meta, src[i:k])
	switch string(src[j:k]) {
	case "include", "include_next", "import", "embed":
	default:
		return k
	}
	m := k
	for m < len(src) && (src[m] == ' ' || src[m] == '\t') {
		m++
	}
	if m < len(src) && src[m] == '<' {
		if end := bytes.IndexByte(src[m:lineEnd(src, m)], '>'); end >= 0 {
			w.emit(none, src[k:m])
			w.emit(str, src[m:m+end+1])
			return m + end + 1
		}
	}
	return k
}

// cppSpecial: preprocesor jak w C oraz napisy surowe R"ogr(…)ogr" z przedrostkami u8, u, U, L. Ogranicznik ma
// najwyżej 16 znaków bez spacji, nawiasów, \ i cudzysłowu; bez ( po nim to nie jest napis surowy.
func cppSpecial(t *table, w *writer, src []byte, i int) int {
	if j := cPreprocessor(t, w, src, i); j > i {
		return j
	}
	if b := src[i]; b != 'R' && b != 'u' && b != 'U' && b != 'L' {
		return i
	}
	k := 0
	for _, p := range []string{`R"`, `u8R"`, `uR"`, `UR"`, `LR"`} {
		if hasPrefix(src[i:], p) {
			k = i + len(p)
			break
		}
	}
	if k == 0 {
		return i
	}
	d := k
	for d < len(src) && d-k < 16 && isRawDelimiter(src[d]) {
		d++
	}
	if d >= len(src) || src[d] != '(' {
		return i
	}
	// Koniec to ), ogranicznik i cudzysłów. Każdy nawias ) w treści jest sprawdzany raz, kosztem najwyżej
	// 17 bajtów, więc szukanie jest liniowe.
	delim := src[k:d]
	end := len(src)
	for p := d + 1; p < len(src); {
		n := bytes.IndexByte(src[p:], ')')
		if n < 0 {
			break
		}
		q := p + n + 1
		if bytes.HasPrefix(src[q:], delim) && q+len(delim) < len(src) && src[q+len(delim)] == '"' {
			end = q + len(delim) + 1
			break
		}
		p = q
	}
	w.emit(str, src[i:end])
	return end
}

// isRawDelimiter: znak dozwolony w ograniczniku napisu surowego C++.
func isRawDelimiter(b byte) bool {
	return b != ' ' && b != '(' && b != ')' && b != '\\' && b != '"' && b != '\t' && b != '\n' && b != '\v' && b != '\f'
}
