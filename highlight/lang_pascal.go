package highlight

// Hak dla Pascala (definicja w lang/pascal.hl; wymagania: requirements/languages/pascal.md).

// pascalSpecial: dyrektywy kompilatora {$…} i (*$…*) (hl-meta), kody znaków #13 i #$0D (hl-esc) oraz liczby
// szesnastkowe $FF, dwójkowe %1010 i ósemkowe &777 (hl-num).
func pascalSpecial(t *table, w *writer, src []byte, i int) int {
	next := byte(0)
	if i+1 < len(src) {
		next = src[i+1]
	}
	switch src[i] {
	case '{':
		if next == '$' {
			end := blockEnd(src, i+2, "{", "}", false)
			w.emit(meta, src[i:end])
			return end
		}
	case '(':
		if hasPrefix(src[i:], "(*$") {
			end := blockEnd(src, i+3, "(*", "*)", false)
			w.emit(meta, src[i:end])
			return end
		}
	case '#':
		n := run(src, i+1, unbounded, isDigit)
		if next == '$' {
			n = run(src, i+2, unbounded, isHex)
			if n > 0 {
				n++
			}
		}
		if n > 0 {
			w.emit(escape, src[i:i+1+n])
			return i + 1 + n
		}
	case '$':
		return pascalNumber(w, src, i, isHex)
	case '%':
		return pascalNumber(w, src, i, isBinary)
	case '&':
		return pascalNumber(w, src, i, isOctal)
	}
	return i
}

// pascalNumber zapisuje liczbę z przedrostkiem w src[i] i cyframi spełniającymi ok (z _ między nimi).
func pascalNumber(w *writer, src []byte, i int, ok func(byte) bool) int {
	n := run(src, i+1, unbounded, func(b byte) bool { return ok(b) || b == '_' })
	if n == 0 || !ok(src[i+1]) {
		return i
	}
	w.emit(number, src[i:i+1+n])
	return i + 1 + n
}
