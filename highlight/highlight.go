package highlight

import (
	"bytes"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"unicode/utf8"
)

// language łączy nazwy języka z funkcją, która go koloruje.
type language struct {
	names []string // nazwa główna, potem aliasy; małymi literami
	scan  func(w *writer, src []byte)
	table *table // dla języków z definicji (do extends); nil dla skanerów w Go

	// Definicja wbudowana jest parsowana przy pierwszym użyciu (ADR 0009): definition to jej tekst, base daje
	// język do extends, once pilnuje jednego parsowania także przy wywołaniach równoległych.
	definition []byte
	base       func(string) *language
	once       sync.Once
}

// tbl zwraca tabelę języka z definicji, parsując ją przy pierwszym użyciu; nil dla skanerów w Go.
func (l *language) tbl() *table {
	if l.definition != nil {
		l.once.Do(func() {
			parsed, err := parseDefinition(l.definition, l.base)
			if err != nil {
				panic("highlight: built-in definition " + l.names[0] + ": " + err.Error())
			}
			l.table = parsed.table
		})
	}
	return l.table
}

// languages to języki wbudowane: z definicji w lang/*.hl i ze skanerów w Go. Nazwa główna wyznacza plik
// wymagań i katalog korpusu.
var languages []*language

// registry to niezmienna migawka języków; Register wymienia ją w całości, więc HTML czyta bez blokady.
type registry struct {
	byName map[string]*language
	names  []string // wszystkie nazwy i aliasy, posortowane
	langs  []*language
}

var (
	current    atomic.Pointer[registry]
	registerMu sync.Mutex // porządkuje równoległe wywołania Register
)

func init() {
	languages = append(loadBuiltins(), &langXML, &langHTML, &langCSS, &langYAML, &langTOML, &langShell,
		&langDiff, &langMarkdown, &langINI, &langProperties, &langAsciiDoc, &langJinja, &langHandlebars, &langLiquid, &langERB, &langGoTemplate,
		&langCSV, &langTSV, &langOrg, &langRST, &langTextile)
	reg, err := newRegistry(languages)
	if err != nil {
		panic(err)
	}
	current.Store(reg)
}

func newRegistry(langs []*language) (*registry, error) {
	r := &registry{byName: map[string]*language{}, langs: langs}
	for _, l := range langs {
		for _, n := range l.names {
			if r.byName[n] != nil {
				return nil, &DefinitionError{Msg: sprintf("name %q is already in use", n)}
			}
			r.byName[n] = l
			r.names = append(r.names, n)
		}
	}
	slices.Sort(r.names)
	return r, nil
}

// lookup zwraca język o nazwie (bez rozróżniania wielkości liter) albo nil.
func lookup(name string) *language {
	return current.Load().byName[strings.ToLower(name)]
}

// Register adds a language described in the definition format (see docs/definitions.md in the repository).
// The language is then accepted by HTML and listed by Languages, also as the language of a fenced block in
// Markdown. The guarantees of HTML hold for it too: the format cannot express anything that would make the
// output unsafe or the time worse than linear. Register returns a *DefinitionError for an invalid definition
// and when one of its names is already in use; it is safe to call while other goroutines call HTML.
func Register(definition []byte) error {
	registerMu.Lock()
	defer registerMu.Unlock()
	old := current.Load()
	l, err := parseDefinition(definition, func(name string) *language { return old.byName[name] })
	if err != nil {
		return err
	}
	reg, err := newRegistry(append(slices.Clone(old.langs), l))
	if err != nil {
		return err
	}
	current.Store(reg)
	return nil
}

// HTML returns code escaped for HTML, with <span class="hl-…"> elements around tokens. The result has no
// <pre> or <code> wrapper, its elements never nest and never contain a line break, and only &, < and > are
// escaped. Invalid UTF-8 and NUL bytes are replaced with U+FFFD. The language name is matched without regard
// to case; ok is false (and out is nil) when the language is not known.
func HTML(lang string, code []byte) (out []byte, ok bool) {
	l := lookup(lang)
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
	return slices.Clone(current.Load().names)
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
