package highlight

import "bytes"

// C według ISO/IEC 9899:2024 (C23), C++ według ISO/IEC 14882:2024 (C++23); wymagania:
// requirements/languages/c.md i cpp.md.

const cTypes = `bool char double float int long short signed unsigned void _Bool _Complex _Imaginary _BitInt
	_Decimal32 _Decimal64 _Decimal128 size_t ptrdiff_t intptr_t uintptr_t intmax_t uintmax_t wchar_t
	char8_t char16_t char32_t int8_t int16_t int32_t int64_t uint8_t uint16_t uint32_t uint64_t`

var langC = language{
	names: []string{"c", "h"},
	scan: (&table{
		words: words(
			wordGroup{keyword, `alignas alignof auto break case const constexpr continue default do else enum
				extern for goto if inline register restrict return sizeof static static_assert struct switch
				thread_local typedef typeof typeof_unqual union volatile while _Alignas _Alignof _Atomic
				_Generic _Noreturn _Static_assert _Thread_local`},
			wordGroup{typeName, cTypes},
			wordGroup{builtin, `true false nullptr NULL`},
		),
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/"}},
		quotes:        cQuotes,
		prefixes:      cPrefixes,
		escape:        cEscape,
		numbers:       numberRule{quoteSep: true},
		operators:     "+-*/%&|^!<>=:?~",
		calls:         true,
		special:       cPreprocessor,
	}).scan,
}

var langCpp = language{
	names: []string{"cpp", "c++", "cxx", "cc", "hpp"},
	scan: (&table{
		words: words(
			wordGroup{keyword, `alignas alignof and and_eq asm auto bitand bitor break case catch class compl
				concept const consteval constexpr constinit const_cast continue co_await co_return co_yield
				decltype default delete do dynamic_cast else enum explicit export extern for friend goto if
				inline mutable namespace new noexcept not not_eq operator or or_eq private protected public
				register reinterpret_cast requires return sizeof static static_assert static_cast struct
				switch template this thread_local throw try typedef typeid typename union using virtual
				volatile while xor xor_eq final override import module`},
			wordGroup{typeName, cTypes},
			wordGroup{builtin, `true false nullptr NULL`},
		),
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/"}},
		quotes:        cQuotes,
		prefixes:      cPrefixes,
		escape:        cEscape,
		numbers:       numberRule{quoteSep: true},
		operators:     "+-*/%&|^!<>=:?~",
		calls:         true,
		special:       cppSpecial,
	}).scan,
}

var (
	cQuotes = []quote{
		{open: `"`, close: `"`, escapes: true},
		{open: `'`, close: `'`, escapes: true},
	}
	cPrefixes = []prefix{{text: "u8"}, {text: "u"}, {text: "U"}, {text: "L"}}
	cEscape   = escapeRule{continuation: true, hexX: unbounded, octal: 3, u4: true, u8: true}.length
)

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
