package highlight

import (
	"bytes"
	"slices"
	"strings"
	"unicode/utf8"
)

// language łączy nazwy języka z funkcją, która go koloruje.
type language struct {
	names []string // nazwa główna, potem aliasy; małymi literami
	scan  func(w *writer, src []byte)
}

// languages to wszystkie obsługiwane języki; nazwa główna pierwszego wpisu wyznacza plik wymagań
// i katalog korpusu.
var languages = []*language{
	&langGo, &langC, &langCpp, &langJava, &langJavaScript, &langTypeScript, &langPython, &langRust,
	&langSQL, &langJSON, &langPascal,
}

var (
	byName   = map[string]*language{}
	allNames []string
)

func init() {
	for _, l := range languages {
		for _, n := range l.names {
			if byName[n] != nil {
				panic("highlight: duplicate language name " + n)
			}
			byName[n] = l
			allNames = append(allNames, n)
		}
	}
	slices.Sort(allNames)
}

// HTML returns code escaped for HTML, with <span class="hl-…"> elements around tokens. The result has no
// <pre> or <code> wrapper, its elements never nest and never contain a line break, and only &, < and > are
// escaped. Invalid UTF-8 and NUL bytes are replaced with U+FFFD. The language name is matched without regard
// to case; ok is false (and out is nil) when the language is not known.
func HTML(lang string, code []byte) (out []byte, ok bool) {
	l := byName[strings.ToLower(lang)]
	if l == nil {
		return nil, false
	}
	code = sanitize(code)
	// Gęsty kod daje wynik 4–5 razy dłuższy od wejścia (korpus Go, 2026-11-25): przy tym zapasie bufor
	// zwykle nie rośnie wcale albo raz.
	w := writer{out: make([]byte, 0, 4*len(code)+16)}
	l.scan(&w, code)
	w.close()
	return w.out, true
}

// Languages returns every accepted language name and alias, sorted, in a new slice.
func Languages() []string {
	return slices.Clone(allNames)
}

// sanitize zamienia każdą serię bajtów niebędących poprawnym UTF-8 i każdy znak U+0000 na U+FFFD (OUT-006).
// Poprawne źródło wraca bez kopiowania.
func sanitize(src []byte) []byte {
	if !utf8.Valid(src) {
		src = bytes.ToValidUTF8(src, []byte("�"))
	}
	if bytes.IndexByte(src, 0) >= 0 {
		src = bytes.ReplaceAll(src, []byte{0}, []byte("�"))
	}
	return src
}
