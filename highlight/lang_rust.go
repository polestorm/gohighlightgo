package highlight

import (
	"bytes"
	"unicode/utf8"
)

// Rust według The Rust Reference (wydanie 2021) i dokumentacji preludium biblioteki standardowej;
// wymagania: requirements/languages/rust.md.
var langRust = language{
	names: []string{"rust", "rs"},
	scan: (&table{
		words: words(
			wordGroup{keyword, `as async await break const continue crate dyn else enum extern fn for if impl
				in let loop match mod move mut pub ref return self Self static struct super trait type
				unsafe use where while union macro_rules abstract become box do final macro override priv
				try typeof unsized virtual yield`},
			wordGroup{typeName, `bool char str i8 i16 i32 i64 i128 isize u8 u16 u32 u64 u128 usize f32 f64
				String Vec Option Result Box`},
			wordGroup{builtin, `true false Some None Ok Err`},
		),
		lineComments:  []string{"//"},
		blockComments: []blockComment{{open: "/*", close: "*/", nested: true}},
		quotes:        []quote{{open: `"`, close: `"`, escapes: true, multiline: true}},
		prefixes:      []prefix{{text: "b"}, {text: "c"}},
		escape:        rustEscape,
		numbers:       numberRule{noLeadingDot: true, dotIdent: true},
		operators:     "+-*/%&|^!<>=:?~@",
		calls:         true,
		macros:        true,
		special:       rustSpecial,
	}).scan,
}

var (
	rustEscape    = escapeRule{continuation: true, hexX: 2, uBraces: true}.length
	rustCharQuote = quote{open: "'", close: "'", escapes: true}
)

// rustSpecial: atrybuty #[…] i #![…] (znacznik i nazwa dostają hl-meta), literały znakowe 'a' i b'a',
// czasy życia i etykiety 'a (hl-meta) oraz napisy surowe r"…", r#"…"#, br…, cr… z dowolną liczbą #.
func rustSpecial(t *table, w *writer, src []byte, i int) int {
	switch b := src[i]; {
	case b == '#':
		k := i + 1
		if k < len(src) && src[k] == '!' {
			k++
		}
		if k >= len(src) || src[k] != '[' {
			return i
		}
		end := t.identEnd(src, k+1)
		w.emit(meta, src[i:end])
		return end
	case b == '\'':
		return rustChar(t, w, src, i, i)
	case b == 'b' && i+1 < len(src) && src[i+1] == '\'':
		return rustChar(t, w, src, i, i+1)
	case b == 'r' || b == 'b' || b == 'c':
		return rustRaw(w, src, i)
	}
	return i
}

// rustChar rozstrzyga apostrof w src[q] (z przedrostkiem od src[start]): literał znakowy, gdy po jednym
// znaku albo sekwencji ucieczki stoi apostrof; czas życia, gdy stoi identyfikator; inaczej nic.
func rustChar(t *table, w *writer, src []byte, start, q int) int {
	if q+1 >= len(src) {
		return start
	}
	if src[q+1] == '\\' {
		return t.scanQuote(w, src, start, q, &rustCharQuote, false, str)
	}
	r, size := utf8.DecodeRune(src[q+1:])
	if r != '\n' && q+1+size < len(src) && src[q+1+size] == '\'' {
		end := q + 2 + size
		w.emit(str, src[start:end])
		return end
	}
	if start == q {
		if j := t.identEnd(src, q+1); j > q+1 {
			w.emit(meta, src[q:j])
			return j
		}
	}
	return start
}

// rustRaw zapisuje napis surowy od src[i], jeśli tam jest: [b|c]r, n znaków #, cudzysłów; koniec to cudzysłów
// i n znaków #. Każdy cudzysłów w treści jest sprawdzany raz, więc koszt jest liniowy.
func rustRaw(w *writer, src []byte, i int) int {
	k := i
	if src[k] == 'b' || src[k] == 'c' {
		k++
	}
	if k >= len(src) || src[k] != 'r' {
		return i
	}
	k++
	hashes := 0
	for k+hashes < len(src) && src[k+hashes] == '#' {
		hashes++
	}
	if k+hashes >= len(src) || src[k+hashes] != '"' {
		return i
	}
	end := len(src)
	for p := k + hashes + 1; p < len(src); {
		n := bytes.IndexByte(src[p:], '"')
		if n < 0 {
			break
		}
		q := p + n + 1
		if run(src, q, hashes, func(b byte) bool { return b == '#' }) == hashes {
			end = q + hashes
			break
		}
		p = q
	}
	w.emit(str, src[i:end])
	return end
}
