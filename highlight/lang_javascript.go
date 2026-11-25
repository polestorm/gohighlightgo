package highlight

// JavaScript według ECMAScript 2024 (ECMA-262, 15. wydanie), TypeScript według dokumentacji języka
// (TypeScript Handbook, wersja 5); wymagania: requirements/languages/javascript.md i typescript.md.

var jsKeywords = wordGroup{keyword, `async await break case catch class const continue debugger default delete
	do else enum export extends finally for function if implements import in instanceof interface let new
	package private protected public return static super switch this throw try typeof var void while with
	yield`}

var jsBuiltins = wordGroup{builtin, `true false null undefined NaN Infinity`}

var langJavaScript = language{
	names: []string{"javascript", "js", "mjs", "cjs"},
	scan:  jsTable(jsKeywords, jsBuiltins).scan,
}

var langTypeScript = language{
	names: []string{"typescript", "ts", "mts", "cts"},
	scan: jsTable(jsKeywords, jsBuiltins,
		wordGroup{keyword, `abstract as asserts declare infer is keyof module namespace override readonly
			satisfies type unique`},
		wordGroup{typeName, `any bigint boolean never number object string symbol unknown`},
	).scan,
}

func jsTable(groups ...wordGroup) *table {
	return &table{
		words:         words(groups...),
		identExtra:    "$",
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/"}},
		quotes: []quote{
			{open: `"`, close: `"`, escapes: true},
			{open: `'`, close: `'`, escapes: true},
		},
		escape:     escapeRule{continuation: true, hexX: 2, u4: true, uBraces: true}.length,
		operators:  "+-*/%&|^!<>=:?~",
		calls:      true,
		decorators: true,
		templates:  true,
	}
}
