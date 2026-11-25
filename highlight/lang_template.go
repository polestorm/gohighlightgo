package highlight

import (
	"bytes"
	"sync"
)

// Języki szablonów: Jinja, Handlebars, Liquid, ERB, szablony Go (wymagania: requirements/languages/jinja.md,
// handlebars.md, liquid.md, erb.md, gotemplate.md) oraz CSV i TSV (csv.md, tsv.md).

// tagPair to ogranicznik wstawki szablonu; comment: treść to komentarz.
type tagPair struct {
	open, close string
	comment     bool
}

// templateScanner koloruje szablon: tekst poza wstawkami jak HTML, ograniczniki wstawek hl-meta, komentarze
// szablonu hl-com, a treść wstawki językiem wyrażeń inner (wbudowana tabela albo język, np. Ruby w ERB).
type templateScanner struct {
	tags  []tagPair // dłuższe ograniczniki przed krótszymi z tym samym początkiem
	inner func() func(*writer, []byte)
}

func (s templateScanner) scan(w *writer, src []byte) {
	inner := s.inner()
	for i := 0; i < len(src); {
		// Najbliższy ogranicznik otwierający: szukamy pierwszego bajtu { albo < i sprawdzamy przedrostki.
		p := i
		var tp *tagPair
		for p < len(src) {
			k := bytes.IndexAny(src[p:], "{<")
			if k < 0 {
				p = len(src)
				break
			}
			p += k
			for n := range s.tags {
				if hasPrefix(src[p:], s.tags[n].open) {
					tp = &s.tags[n]
					break
				}
			}
			if tp != nil {
				break
			}
			p++
		}
		if p > i {
			markup{html: true}.scan(w, src[i:p])
			w.close()
		}
		if tp == nil {
			return
		}
		body := p + len(tp.open)
		end, closeEnd := len(src), len(src)
		if n := bytes.Index(src[body:], []byte(tp.close)); n >= 0 {
			end, closeEnd = body+n, body+n+len(tp.close)
		}
		if tp.comment {
			w.emit(comment, src[p:closeEnd])
		} else {
			w.emit(meta, src[p:body])
			inner(w, src[body:end])
			w.close()
			w.emit(meta, src[end:closeEnd])
		}
		i = closeEnd
	}
}

// innerTable zwraca skaner wbudowanej tabeli wyrażeń szablonu (definicja w Go, nie jako osobny język).
func innerTable(def string) func() func(*writer, []byte) {
	var t *table
	var once sync.Once
	return func() func(*writer, []byte) {
		once.Do(func() {
			l, err := parseDefinition([]byte(def), lookup)
			if err != nil {
				panic("highlight: template expressions: " + err.Error())
			}
			t = l.table
		})
		return t.scan
	}
}

var jinjaInner = innerTable(`name -jinja-expr
keyword  if elif else endif for endfor in block endblock extends include import from macro endmacro set
keyword  endset with endwith filter endfilter call endcall raw endraw not and or is as recursive
keyword  autoescape endautoescape trans endtrans
builtin  true false none True False None loop super self range
string   " "  : escapes
string   ' '  : escapes
operators  +-*/%=<>!|~.:
calls
`)

var handlebarsInner = innerTable(`name -handlebars-expr
keyword  if unless each with else lookup log let
builtin  this true false null undefined
string   " "  : escapes
string   ' '  : escapes
sigil    @    : attr
operators  #/^>&=~.
`)

var liquidInner = innerTable(`name -liquid-expr
keyword  if elsif else endif unless endunless case when endcase for endfor in break continue assign
keyword  capture endcapture include render raw endraw comment endcomment tablerow endtablerow cycle
keyword  increment decrement liquid echo and or contains
builtin  true false nil empty blank forloop
string   " "  : escapes
string   ' '  : escapes
operators  =<>!|:.-
`)

var goTemplateInner = innerTable(`name -go-template-expr
keyword  if else end range with define template block break continue
builtin  and or not len index slice print printf println html js urlquery eq ne lt le gt ge call true
builtin  false nil
string   " "  : escapes
string   ` + "`" + ` ` + "`" + `  : multiline
sigil    $    : attr
sigil    .    : attr
operators  |=:-
`)

// ejsInner: wnętrze <% … %> w EJS to JavaScript.
func ejsInner() func(*writer, []byte) {
	return func(w *writer, src []byte) { lookup("javascript").scan(w, src) }
}

var langEJS = language{
	names: []string{"ejs"},
	scan: templateScanner{tags: []tagPair{
		{open: "<%#", close: "%>", comment: true}, {open: "<%=", close: "%>"}, {open: "<%-", close: "%>"},
		{open: "<%_", close: "%>"}, {open: "<%", close: "%>"},
	}, inner: ejsInner}.scan,
}

// erbInner: wnętrze <% … %> w ERB to Ruby.
func erbInner() func(*writer, []byte) {
	return func(w *writer, src []byte) { lookup("ruby").scan(w, src) }
}

var langJinja = language{
	names: []string{"jinja", "jinja2", "j2", "django", "twig", "nunjucks", "njk"},
	scan: templateScanner{tags: []tagPair{
		{open: "{#", close: "#}", comment: true}, {open: "{{", close: "}}"}, {open: "{%", close: "%}"},
	}, inner: jinjaInner}.scan,
}

var langHandlebars = language{
	names: []string{"handlebars", "hbs", "mustache"},
	scan: templateScanner{tags: []tagPair{
		{open: "{{!--", close: "--}}", comment: true}, {open: "{{!", close: "}}", comment: true},
		{open: "{{{", close: "}}}"}, {open: "{{", close: "}}"},
	}, inner: handlebarsInner}.scan,
}

var langLiquid = language{
	names: []string{"liquid"},
	scan: templateScanner{tags: []tagPair{
		{open: "{{", close: "}}"}, {open: "{%", close: "%}"},
	}, inner: liquidInner}.scan,
}

var langERB = language{
	names: []string{"erb", "rhtml", "eruby"},
	scan: templateScanner{tags: []tagPair{
		{open: "<%#", close: "%>", comment: true}, {open: "<%=", close: "%>"}, {open: "<%-", close: "%>"},
		{open: "<%", close: "%>"},
	}, inner: erbInner}.scan,
}

var langGoTemplate = language{
	names: []string{"gotemplate", "gotmpl", "go-template", "tmpl", "helm"},
	scan: templateScanner{tags: []tagPair{
		{open: "{{/*", close: "*/}}", comment: true}, {open: "{{- /*", close: "*/ -}}", comment: true},
		{open: "{{", close: "}}"},
	}, inner: goTemplateInner}.scan,
}

// csvScanner koloruje dane rozdzielane: pierwszy wiersz (nagłówek) hl-attr, separatory hl-op, pola
// w cudzysłowie (z "" w środku, także przez koniec wiersza) hl-str, liczby hl-num (RFC 4180; TSV: IANA
// text/tab-separated-values).
type csvScanner struct{ sep byte }

var langCSV = language{names: []string{"csv"}, scan: csvScanner{sep: ','}.scan}
var langTSV = language{names: []string{"tsv", "tab"}, scan: csvScanner{sep: '\t'}.scan}

func (c csvScanner) scan(w *writer, src []byte) {
	header := true
	for i := 0; i < len(src); {
		b := src[i]
		switch {
		case b == '\n':
			w.emit(none, src[i:i+1])
			header = false
			i++
		case b == c.sep || b == ';' && c.sep == ',':
			w.emit(operator, src[i:i+1])
			i++
		case b == '"':
			j := i + 1
			for j < len(src) {
				if src[j] == '"' {
					if j+1 < len(src) && src[j+1] == '"' {
						j += 2
						continue
					}
					j++
					break
				}
				j++
			}
			cl := str
			if header {
				cl = attribute
			}
			w.emit(cl, src[i:j])
			i = j
		default:
			j := i
			for j < len(src) && src[j] != '\n' && src[j] != c.sep && !(c.sep == ',' && src[j] == ';') && src[j] != '"' {
				j++
			}
			field := src[i:j]
			cl := none
			switch {
			case header:
				cl = attribute
			case isConfigNumber(string(bytes.TrimSpace(field))):
				cl = number
			}
			w.emit(cl, field)
			i = j
		}
	}
}
