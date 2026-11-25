package highlight

import (
	"embed"
	"path"
	"slices"
	"strconv"
	"strings"
)

// Definicje języków z tabeli w formacie opisanym w docs/definitions.md (ADR 0008). Wbudowane leżą
// w highlight/lang/*.hl; Register przyjmuje taki sam tekst od wywołującego.

//go:embed lang/*.hl
var builtinDefinitions embed.FS

// maxDefinition ogranicza rozmiar definicji; ogranicza to też liczbę reguł, które skaner sprawdza na każdej
// pozycji, więc stały czynnik czasu liniowego.
const maxDefinition = 64 << 10

// hooks: konstrukcje spoza tabeli, napisane w Go; definicja może je włączyć nazwą, ale nie może ich zmienić.
var hooks = map[string]hookFunc{
	"c-preprocessor": cPreprocessor,
	"cpp":            cppSpecial,
	"rust":           rustSpecial,
	"json-minus":     jsonMinus,
	"pascal":         pascalSpecial,
	"lua":            luaSpecial,
	"swift":          swiftSpecial,
	"apostrophe":     apostropheSpecial,
	"haskell":        haskellPragma,
	"ruby":           rubySpecial,
	"php":            phpSpecial,
	"perl":           perlSpecial,
	"elixir":         elixirSpecial,
	"matlab":         matlabSpecial,
	"dotted":         dottedSpecial,
	"variables":      variablesSpecial,
	"latex":          latexSpecial,
	"vim":            vimSpecial,
	"hcl":            hclSpecial,
}

// DefinitionError reports an invalid language definition: the line (from 1) and what is wrong with it.
type DefinitionError struct {
	Line int
	Msg  string
}

func (e *DefinitionError) Error() string {
	if e.Line == 0 {
		return "highlight: definition: " + e.Msg
	}
	return sprintf("highlight: definition line %d: %s", e.Line, e.Msg)
}

// parseDefinition czyta definicję. lookup daje język do dyrektywy extends (tylko języki z tabeli).
func parseDefinition(src []byte, lookup func(string) *language) (*language, error) {
	if len(src) > maxDefinition {
		return nil, &DefinitionError{Msg: sprintf("longer than %d bytes", maxDefinition)}
	}
	t := &table{words: map[string]class{}, escape: escapeRule{}.length}
	var names []string
	seen := map[string]bool{}
	for n, raw := range strings.Split(string(src), "\n") {
		fail := func(format string, a ...any) (*language, error) {
			return nil, &DefinitionError{Line: n + 1, Msg: sprintf(format, a...)}
		}
		line := strings.TrimSpace(raw)
		if line == "" || line[0] == '#' {
			continue
		}
		fields := strings.Fields(line)
		key := fields[0]
		var values, opts []string
		// Separatorem opcji jest ostatni samotny „:”, więc „:” może być też wartością (sigil : : attr).
		for k := len(fields) - 1; k >= 1; k-- {
			if fields[k] == ":" {
				values, opts = fields[1:k], fields[k+1:]
				break
			}
		}
		if values == nil && opts == nil {
			values = fields[1:]
		}
		for _, v := range values {
			if !isPrintableASCII(v) && key != "name" && key != "keyword" && key != "type" && key != "builtin" {
				return fail("%s: %q is not printable ASCII", key, v)
			}
		}
		if key != "keyword" && key != "type" && key != "builtin" && key != "string" && key != "prefix" && key != "block-comment" {
			if seen[key] && key != "name" && key != "line-comment" && key != "hook" && key != "sigil" && key != "line-string" {
				return fail("%s given twice", key)
			}
		}
		seen[key] = true
		switch key {
		case "extends":
			if len(names) > 0 || len(t.words) > 0 || len(values) != 1 || len(opts) > 0 {
				return fail("extends takes one name and must come before everything else")
			}
			base := lookup(strings.ToLower(values[0]))
			if base == nil || base.tbl() == nil {
				return fail("extends: no table language %q", values[0])
			}
			t = base.tbl().clone()
		case "name":
			if len(values) == 0 || len(opts) > 0 {
				return fail("name needs at least one value")
			}
			for _, v := range values {
				if v != strings.ToLower(v) || strings.ContainsAny(v, " \t\x00") {
					return fail("name %q must be lower case", v)
				}
				names = append(names, v)
			}
		case "keyword", "type", "builtin":
			c := map[string]class{"keyword": keyword, "type": typeName, "builtin": builtin}[key]
			if len(values) == 0 || len(opts) > 0 {
				return fail("%s needs words", key)
			}
			for _, v := range values {
				if len(v) > 32 {
					return fail("%s %q is longer than 32 bytes", key, v)
				}
				t.words[v] = c
			}
		case "fold-case", "calls", "macros", "decorators", "templates", "keys", "regex", "upper-type", "call-after-paren":
			if len(values) > 0 || len(opts) > 0 {
				return fail("%s takes no values", key)
			}
			switch key {
			case "fold-case":
				t.foldCase = true
			case "calls":
				t.calls = true
			case "macros":
				t.macros = true
			case "decorators":
				t.sigils = append(t.sigils, sigil{char: '@', class: meta})
			case "upper-type":
				t.upperType = true
			case "call-after-paren":
				t.lispCalls = true
			case "templates":
				t.templates = true
			case "keys":
				t.keys = true
			case "regex":
				t.regex = true
			}
		case "ident-extra":
			if len(values) != 1 || len(opts) > 0 || strings.ContainsAny(values[0], "\"'`") {
				return fail("ident-extra takes one group of punctuation characters, without quotes")
			}
			t.identExtra = values[0]
		case "ident-part":
			if len(values) != 1 || len(opts) > 0 || strings.ContainsAny(values[0], "\"`") {
				return fail("ident-part takes one group of punctuation characters")
			}
			t.identPart = values[0]
		case "sigil":
			if len(values) != 1 || len(opts) != 1 {
				return fail("sigil takes characters and a class after a colon: sigil $ : attr")
			}
			c, ok := classByName[opts[0]]
			if !ok {
				return fail("sigil: unknown class %q", opts[0])
			}
			for k := 0; k < len(values[0]); k++ {
				t.sigils = append(t.sigils, sigil{char: values[0][k], class: c})
			}
		case "line-string":
			if len(values) == 0 || len(opts) > 0 {
				return fail("line-string needs at least one prefix")
			}
			t.lineStrings = append(t.lineStrings, values...)
		case "line-comment":
			if len(values) == 0 || len(opts) > 0 {
				return fail("line-comment needs at least one prefix")
			}
			t.lineComments = append(t.lineComments, values...)
		case "block-comment":
			if len(values) != 2 {
				return fail("block-comment needs an opening and a closing delimiter")
			}
			bc := blockComment{open: values[0], close: values[1]}
			for _, o := range opts {
				if o != "nested" {
					return fail("block-comment: unknown option %q", o)
				}
				bc.nested = true
			}
			t.blockComments = append(t.blockComments, bc)
		case "string":
			if len(values) != 2 {
				return fail("string needs an opening and a closing delimiter")
			}
			q := quote{open: values[0], close: values[1]}
			for _, o := range opts {
				switch o {
				case "escapes":
					q.escapes = true
				case "multiline":
					q.multiline = true
				case "doubled":
					q.doubled = true
				case "name":
					q.ident = true
				case "interp=dollar":
					q.interp = '$'
				case "interp=braces":
					q.interp = '{'
				case "interp=hash":
					q.interp = '#'
				case "interp=dollar-braces":
					q.interp = 'B'
				default:
					return fail("string: unknown option %q", o)
				}
			}
			t.quotes = append(t.quotes, q)
		case "prefix":
			if len(values) == 0 {
				return fail("prefix needs at least one prefix")
			}
			mode, interp, anyCase := prefixPlain, byte(0), false
			for _, o := range opts {
				switch o {
				case "raw":
					mode = prefixRaw
				case "literal":
					mode = prefixLiteral
				case "verbatim":
					mode = prefixVerbatim
				case "interp=dollar":
					interp = '$'
				case "interp=braces":
					interp = '{'
				case "any-case":
					anyCase = true
				default:
					return fail("prefix: unknown option %q", o)
				}
			}
			for _, v := range values {
				if len(v) > 8 {
					return fail("prefix %q is longer than 8 bytes", v)
				}
				variants := []string{v}
				if anyCase {
					variants = caseVariants(v)
				}
				for _, text := range variants {
					t.prefixes = append(t.prefixes, prefix{text: text, mode: mode, interp: interp})
				}
			}
		case "escape":
			if len(values) > 0 {
				return fail("escape takes only options, after a colon")
			}
			r := escapeRule{}
			for _, o := range opts {
				k, v, hasValue := strings.Cut(o, "=")
				switch {
				case k == "continuation" && !hasValue:
					r.continuation = true
				case k == "x" && hasValue:
					if v == "any" {
						r.hexX = unbounded
					} else if r.hexX = atoiRange(v, 1, 8); r.hexX == 0 {
						return fail("escape: x must be 1 to 8 or any")
					}
				case k == "octal" && hasValue:
					if r.octal = atoiRange(v, 1, 3); r.octal == 0 {
						return fail("escape: octal must be 1 to 3")
					}
				case k == "u4" && !hasValue:
					r.u4 = true
				case k == "U8" && !hasValue:
					r.u8 = true
				case k == "u-braces" && !hasValue:
					r.uBraces = true
				case k == "java-u" && !hasValue:
					r.javaU = true
				case k == "named" && !hasValue:
					r.named = true
				case k == "paren" && !hasValue:
					r.parenGroup = true
				case k == "char" && hasValue && len(v) == 1 && v[0] > ' ' && v[0] < 0x7f && v != "\n":
					t.escapeChar = v[0]
				default:
					return fail("escape: unknown option %q", o)
				}
			}
			t.escape = r.length
		case "number":
			if len(values) > 0 {
				return fail("number takes only options, after a colon")
			}
			for _, o := range opts {
				switch o {
				case "no-leading-dot":
					t.numbers.noLeadingDot = true
				case "dot-ident":
					t.numbers.dotIdent = true
				case "quote-sep":
					t.numbers.quoteSep = true
				default:
					return fail("number: unknown option %q", o)
				}
			}
		case "operators":
			if len(values) != 1 || len(opts) > 0 {
				return fail("operators takes one group of characters")
			}
			t.operators = values[0]
		case "hook":
			if len(values) != 1 || len(opts) > 0 || hooks[values[0]] == nil {
				return fail("hook: unknown hook %q (known: %s)", strings.Join(values, " "), strings.Join(hookNames(), ", "))
			}
			t.specials = append(t.specials, hooks[values[0]])
		default:
			return fail("unknown directive %q", key)
		}
	}
	if len(names) == 0 {
		return nil, &DefinitionError{Msg: "no name"}
	}
	for _, q := range t.quotes {
		if q.open == "" || q.close == "" {
			return nil, &DefinitionError{Msg: "empty string delimiter"}
		}
	}
	if t.foldCase {
		t.words = foldWords(t.words)
	}
	t.finish()
	return &language{names: names, scan: t.scan, table: t}, nil // Register: bez leniwego parsowania
}

// clone kopiuje tabelę do dyrektywy extends: mapy i wycinki są nowe, więc zmiany nie dotykają bazy.
func (t *table) clone() *table {
	c := *t
	c.words = make(map[string]class, len(t.words))
	for k, v := range t.words {
		c.words[k] = v
	}
	c.lineComments = slices.Clone(t.lineComments)
	c.blockComments = slices.Clone(t.blockComments)
	c.quotes = slices.Clone(t.quotes)
	c.prefixes = slices.Clone(t.prefixes)
	c.sigils = slices.Clone(t.sigils)
	c.lineStrings = slices.Clone(t.lineStrings)
	c.specials = slices.Clone(t.specials)
	return &c
}

// caseVariants zwraca wszystkie zapisy wielkości liter słowa (dla przedrostków najwyżej 8 znaków).
func caseVariants(s string) []string {
	out := []string{""}
	for _, r := range s {
		lower, upper := strings.ToLower(string(r)), strings.ToUpper(string(r))
		var next []string
		for _, prefix := range out {
			next = append(next, prefix+lower)
			if upper != lower {
				next = append(next, prefix+upper)
			}
		}
		out = next
	}
	return out
}

func atoiRange(s string, lo, hi int) int {
	n, err := strconv.Atoi(s)
	if err != nil || n < lo || n > hi {
		return 0
	}
	return n
}

func isPrintableASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] <= ' ' || s[i] >= 0x7f {
			return false
		}
	}
	return true
}

func hookNames() []string {
	var names []string
	for n := range hooks {
		names = append(names, n)
	}
	slices.Sort(names)
	return names
}

// loadBuiltins przygotowuje wbudowane definicje bez parsowania: z każdego pliku czyta tylko nazwy (wiersze
// name), a resztę parsuje przy pierwszym użyciu języka (ADR 0009). Poprawność wszystkich sprawdza
// TestBuiltinDefinitions.
func loadBuiltins() []*language {
	files, err := builtinDefinitions.ReadDir("lang")
	if err != nil {
		panic(err)
	}
	var out []*language
	byPrimary := map[string]*language{}
	base := func(name string) *language { return byPrimary[name] }
	for _, f := range files {
		src, err := builtinDefinitions.ReadFile(path.Join("lang", f.Name()))
		if err != nil {
			panic(err)
		}
		l := &language{names: definitionNames(src), definition: src, base: base}
		if len(l.names) == 0 {
			panic("highlight: built-in definition " + f.Name() + " has no name")
		}
		l.scan = func(w *writer, code []byte) { l.tbl().scan(w, code) }
		byPrimary[l.names[0]] = l
		out = append(out, l)
	}
	return out
}

// definitionNames zwraca wartości wierszy name bez parsowania reszty definicji.
func definitionNames(src []byte) []string {
	var names []string
	for _, line := range strings.Split(string(src), "\n") {
		fields := strings.Fields(line)
		if len(fields) > 1 && fields[0] == "name" {
			names = append(names, fields[1:]...)
		}
	}
	return names
}

// sprintf formatuje komunikat błędu: %q (napis w cudzysłowie), %s (napis albo błąd), %d (liczba). Zastępuje
// fmt, który dokładał do programu około 250 KB (pakiety reflect i os), a biblioteka formatuje tylko te
// komunikaty.
func sprintf(format string, a ...any) string {
	var b strings.Builder
	k := 0
	for i := 0; i < len(format); i++ {
		if format[i] != '%' || i+1 >= len(format) {
			b.WriteByte(format[i])
			continue
		}
		i++
		var arg any
		if k < len(a) {
			arg = a[k]
			k++
		}
		text := ""
		switch v := arg.(type) {
		case string:
			text = v
		case error:
			text = v.Error()
		case int:
			text = strconv.Itoa(v)
		}
		if format[i] == 'q' {
			text = strconv.Quote(text)
		}
		b.WriteString(text)
	}
	return b.String()
}

// classByName: nazwy klas w definicjach (bez przedrostka hl-), dla dyrektywy sigil.
var classByName = map[string]class{"kw": keyword, "type": typeName, "builtin": builtin, "str": str, "esc": escape,
	"num": number, "com": comment, "op": operator, "fn": call, "tag": tag, "attr": attribute, "meta": meta}
