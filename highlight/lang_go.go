package highlight

// Go według The Go Programming Language Specification (Go 1.22); wymagania: requirements/languages/go.md.
var langGo = language{
	names: []string{"go", "golang"},
	scan: (&table{
		words: words(
			wordGroup{keyword, `break case chan const continue default defer else fallthrough for func go goto
				if import interface map package range return select struct switch type var`},
			wordGroup{typeName, `any bool byte comparable complex64 complex128 error float32 float64
				int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr`},
			wordGroup{builtin, `true false iota nil append cap clear close complex copy delete imag len make
				max min new panic print println real recover`},
		),
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/"}},
		quotes: []quote{
			{open: `"`, close: `"`, escapes: true},
			{open: `'`, close: `'`, escapes: true},
			{open: "`", close: "`", multiline: true},
		},
		escape:    goEscape,
		operators: "+-*/%&|^<>=!:~",
		calls:     true,
	}).scan,
}

// goEscape: sekwencje ucieczki z rozdziału „Rune literals”. Niepełna sekwencja kończy się na ostatniej
// pasującej cyfrze; nieznana (np. \q) obejmuje \ i jeden znak.
func goEscape(src []byte, i int) int {
	if i+1 >= len(src) {
		return 1
	}
	switch c := src[i+1]; {
	case c == 'x':
		return 2 + run(src, i+2, 2, isHex)
	case c == 'u':
		return 2 + run(src, i+2, 4, isHex)
	case c == 'U':
		return 2 + run(src, i+2, 8, isHex)
	case isOctal(c):
		return 1 + run(src, i+1, 3, isOctal)
	}
	return backslashEscape(src, i)
}
