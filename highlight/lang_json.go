package highlight

// JSON według RFC 8259; komentarze // i /* */ są tolerowane (JSONC, np. pliki ustawień edytorów).
// Wymagania: requirements/languages/json.md.
var langJSON = language{
	names: []string{"json", "jsonc"},
	scan: (&table{
		words:         words(wordGroup{builtin, `true false null`}),
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/"}},
		quotes:        []quote{{open: `"`, close: `"`, escapes: true}},
		escape:        escapeRule{u4: true}.length,
		keys:          true,
		special:       jsonMinus,
	}).scan,
}

// jsonMinus: minus przed cyfrą należy do liczby.
func jsonMinus(t *table, w *writer, src []byte, i int) int {
	if src[i] != '-' || i+1 >= len(src) || !isDigit(src[i+1]) {
		return i
	}
	end := scanNumber(src, i+1, t.numbers)
	w.emit(number, src[i:end])
	return end
}
