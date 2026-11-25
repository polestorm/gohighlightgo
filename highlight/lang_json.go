package highlight

// Hak dla JSON (definicja w lang/json.hl; wymagania: requirements/languages/json.md).

// jsonMinus: minus przed cyfrą należy do liczby.
func jsonMinus(t *table, w *writer, src []byte, i int) int {
	if src[i] != '-' || i+1 >= len(src) || !isDigit(src[i+1]) {
		return i
	}
	end := scanNumber(src, i+1, t.numbers)
	w.emit(number, src[i:end])
	return end
}
