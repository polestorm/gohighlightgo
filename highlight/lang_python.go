package highlight

import "strings"

// Python według The Python Language Reference i The Python Standard Library („Built-in Functions”,
// „Built-in Types”), wersja 3.12; wymagania: requirements/languages/python.md.
var langPython = language{
	names: []string{"python", "py", "python3"},
	scan: (&table{
		words: words(
			wordGroup{keyword, `and as assert async await break class continue def del elif else except
				finally for from global if import in is lambda nonlocal not or pass raise return try while
				with yield`},
			wordGroup{typeName, `bool bytearray bytes complex dict float frozenset int list memoryview object
				range set slice str tuple type`},
			wordGroup{builtin, `True False None NotImplemented Ellipsis abs aiter all anext any ascii bin
				breakpoint callable chr classmethod compile delattr dir divmod enumerate eval exec filter
				format getattr globals hasattr hash help hex id input isinstance issubclass iter len locals map
				max min next oct open ord pow print property repr reversed round setattr sorted staticmethod
				sum super vars zip __import__`},
		),
		lineComments: []string{"#"},
		quotes: []quote{
			{open: `"""`, close: `"""`, escapes: true, multiline: true},
			{open: `'''`, close: `'''`, escapes: true, multiline: true},
			{open: `"`, close: `"`, escapes: true},
			{open: `'`, close: `'`, escapes: true},
		},
		prefixes:   pythonPrefixes(),
		escape:     escapeRule{continuation: true, hexX: 2, octal: 3, u4: true, u8: true, named: true}.length,
		operators:  "+-*/%@&|^~<>=!:",
		calls:      true,
		decorators: true,
	}).scan,
}

// pythonPrefixes zwraca przedrostki napisów we wszystkich zapisach wielkości liter; przedrostek z r jest surowy.
func pythonPrefixes() []prefix {
	var out []prefix
	for _, p := range []string{"r", "u", "f", "b", "br", "rb", "fr", "rf"} {
		for mask := 0; mask < 1<<len(p); mask++ {
			b := []byte(p)
			for k := range b {
				if mask&(1<<k) != 0 {
					b[k] -= 'a' - 'A'
				}
			}
			out = append(out, prefix{text: string(b), raw: strings.Contains(p, "r")})
		}
	}
	return out
}
