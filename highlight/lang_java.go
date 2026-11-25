package highlight

// Java według The Java Language Specification, Java SE 21; wymagania: requirements/languages/java.md.
var langJava = language{
	names: []string{"java"},
	scan: (&table{
		words: words(
			wordGroup{keyword, `abstract assert break case catch class const continue default do else enum
				extends final finally for goto if implements import instanceof interface native new package
				private protected public return static strictfp super switch synchronized this throw throws
				transient try volatile while var record sealed permits yield when`},
			wordGroup{typeName, `boolean byte char double float int long short void`},
			wordGroup{builtin, `true false null`},
		),
		identExtra:    "$",
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/"}},
		quotes: []quote{
			{open: `"""`, close: `"""`, escapes: true, multiline: true},
			{open: `"`, close: `"`, escapes: true},
			{open: `'`, close: `'`, escapes: true},
		},
		escape:     escapeRule{octal: 3, javaU: true}.length,
		operators:  "+-*/%&|^!<>=:?~",
		calls:      true,
		decorators: true,
	}).scan,
}
