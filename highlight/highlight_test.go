package highlight

import (
	"bytes"
	"flag"
	"os"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"
)

var update = flag.Bool("update", false, "przepisz wzorce korpusu (testdata/corpus/*/*.html)")

// TestCorpus porównuje wynik dla plików korpusu z wzorcami; język to nazwa katalogu.
// go test -update przepisuje wzorce, różnicę ogląda się w git diff.
func TestCorpus(t *testing.T) {
	for _, path := range corpusFiles(t) {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		lang := path[len("../testdata/corpus/"):]
		lang = lang[:strings.IndexByte(lang, '/')]
		got, ok := HTML(lang, src)
		if !ok {
			t.Fatalf("%s: unknown language %q", path, lang)
		}
		golden := strings.TrimSuffix(path, ".txt") + ".html"
		if *update {
			if err := os.WriteFile(golden, got, 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		want, err := os.ReadFile(golden)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, want) {
			t.Errorf("%s: output differs from %s\ngot:\n%s", path, golden, got)
		}
	}
}

// REQ: API-001 API-002 API-003
func TestHTML(t *testing.T) {
	for _, name := range []string{"go", "golang", "Go", "GO", "GoLang"} {
		out, ok := HTML(name, []byte("func"))
		if !ok || string(out) != `<span class="hl-kw">func</span>` {
			t.Errorf("HTML(%q) = %q, %v", name, out, ok)
		}
	}
	if out, ok := HTML("go", nil); !ok || len(out) != 0 {
		t.Errorf("empty code: %q, %v", out, ok)
	}
	for _, name := range []string{"", "gopher", "go ", "c++?"} {
		if out, ok := HTML(name, []byte("func")); ok || out != nil {
			t.Errorf("HTML(%q) = %q, %v; want nil, false", name, out, ok)
		}
	}
}

// REQ: API-004
func TestLanguages(t *testing.T) {
	names := Languages()
	if !slices.IsSorted(names) {
		t.Errorf("not sorted: %v", names)
	}
	if len(slices.Compact(slices.Clone(names))) != len(names) {
		t.Errorf("duplicates: %v", names)
	}
	for _, l := range languages {
		for _, n := range l.names {
			if !slices.Contains(names, n) {
				t.Errorf("%s missing from Languages()", n)
			}
			if n != strings.ToLower(n) {
				t.Errorf("name %q is not lower case", n)
			}
			if _, ok := HTML(n, nil); !ok {
				t.Errorf("HTML does not accept %q", n)
			}
		}
	}
	names[0] = "changed"
	if Languages()[0] == "changed" {
		t.Error("Languages returns the internal slice")
	}
}

// REQ: OUT-001
func TestEscapeOnly(t *testing.T) {
	out, _ := HTML("go", []byte(`x := "<a href='&'>"`))
	want := `x <span class="hl-op">:=</span> <span class="hl-str">"&lt;a href='&amp;'&gt;"</span>`
	if string(out) != want {
		t.Errorf("got  %s\nwant %s", out, want)
	}
}

// REQ: OUT-006
func TestInvalidUTF8(t *testing.T) {
	out, _ := HTML("go", []byte("a\xff\xfeb \x00 \"\xc3\""))
	want := "a�b � <span class=\"hl-str\">\"�\"</span>"
	if string(out) != want {
		t.Errorf("got  %q\nwant %q", out, want)
	}
}

// REQ: DET-001 API-005
func TestDeterministic(t *testing.T) {
	var src []byte
	for _, path := range corpusFiles(t) {
		b, _ := os.ReadFile(path)
		src = append(src, b...)
	}
	want, _ := HTML("go", src)
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				if got, _ := HTML("go", src); !bytes.Equal(got, want) {
					t.Error("output differs between calls")
					return
				}
				_ = Languages()
			}
		}()
	}
	wg.Wait()
}

// pathological: dane złośliwe dla skanera; gen(n) daje wejście długości rzędu n.
var pathological = []struct {
	lang, name string
	gen        func(n int) string
}{
	{"go", "unclosed block comment openers", func(n int) string { return strings.Repeat("/*", n/2) }},
	{"go", "block comment closers", func(n int) string { return strings.Repeat("*/", n/2) }},
	{"go", "half comment markers", func(n int) string { return strings.Repeat("/*/", n/3) }},
	{"go", "line comments", func(n int) string { return strings.Repeat("//\n", n/3) }},
	{"go", "long line comment", func(n int) string { return "//" + strings.Repeat("a", n) }},
	{"go", "unclosed string with escapes", func(n int) string { return `"` + strings.Repeat(`\"`, n/2) }},
	{"go", "backslashes", func(n int) string { return `"` + strings.Repeat(`\`, n) }},
	{"go", "short unicode escapes", func(n int) string { return `"` + strings.Repeat(`\U0010`, n/6) }},
	{"go", "quotes", func(n int) string { return strings.Repeat(`"`, n) }},
	{"go", "apostrophes", func(n int) string { return strings.Repeat(`'`, n) }},
	{"go", "unclosed raw string", func(n int) string { return "`" + strings.Repeat("a\n", n/2) }},
	{"go", "backticks", func(n int) string { return strings.Repeat("`", n) }},
	{"go", "unclosed strings on lines", func(n int) string { return strings.Repeat("\"a\n", n/3) }},
	{"go", "long identifier", func(n int) string { return strings.Repeat("a", n) }},
	{"go", "long unicode identifier", func(n int) string { return strings.Repeat("ż", n/2) }},
	{"go", "identifiers and spaces before a paren", func(n int) string { return "a" + strings.Repeat(" ", n) + "(" }},
	{"go", "calls", func(n int) string { return strings.Repeat("a(", n/2) }},
	{"go", "identifiers without parens", func(n int) string { return strings.Repeat("a ", n/2) }},
	{"go", "keywords", func(n int) string { return strings.Repeat("func ", n/5) }},
	{"go", "hex prefixes", func(n int) string { return strings.Repeat("0x", n/2) }},
	{"go", "long number", func(n int) string { return "0x" + strings.Repeat("_", n) + "p" }},
	{"go", "exponents", func(n int) string { return strings.Repeat("1e+", n/3) }},
	{"go", "dots", func(n int) string { return strings.Repeat(".", n) }},
	{"go", "operators", func(n int) string { return strings.Repeat("+=", n/2) }},
	{"go", "html characters", func(n int) string { return strings.Repeat("<&>", n/3) }},
	{"go", "line breaks", func(n int) string { return strings.Repeat("\n", n) }},
	{"go", "invalid utf-8", func(n int) string { return strings.Repeat("\xff", n) }},
	{"go", "nul bytes", func(n int) string { return strings.Repeat("\x00", n) }},
	{"c", "directive lines", func(n int) string { return strings.Repeat("#\n", n/2) }},
	{"c", "hashes after spaces", func(n int) string { return strings.Repeat("  #", n/3) }},
	{"c", "unclosed include", func(n int) string { return "#include <" + strings.Repeat("a", n) }},
	{"c", "many unclosed includes", func(n int) string { return strings.Repeat("#include <a\n", n/12) }},
	{"c", "long hex escape", func(n int) string { return `"\x` + strings.Repeat("f", n) }},
	{"c", "digit separators", func(n int) string { return "1" + strings.Repeat("'1", n/2) }},
	{"cpp", "raw string starts", func(n int) string { return strings.Repeat(`R"`, n/2) }},
	{"cpp", "raw string long delimiters", func(n int) string { return strings.Repeat(`R"abcdefghijklmnop`, n/18) }},
	{"cpp", "unclosed raw string", func(n int) string { return `R"x(` + strings.Repeat(")x", n/2) }},
	{"cpp", "raw string openers", func(n int) string { return strings.Repeat(`u8R"(`, n/5) }},
	{"java", "text block openers", func(n int) string { return strings.Repeat(`"""`, n/3) }},
	{"java", "long unicode escape", func(n int) string { return `"\` + strings.Repeat("u", n) }},
	{"java", "annotations", func(n int) string { return strings.Repeat("@a", n/2) }},
	{"java", "at signs", func(n int) string { return strings.Repeat("@", n) }},
	{"javascript", "nested template expressions", func(n int) string { return strings.Repeat("`${", n/3) }},
	{"javascript", "closing braces", func(n int) string { return strings.Repeat("}", n) }},
	{"javascript", "braces in a template expression", func(n int) string { return "`${" + strings.Repeat("{", n) }},
	{"javascript", "empty expressions", func(n int) string { return "`" + strings.Repeat("${}", n/3) }},
	{"javascript", "template escapes", func(n int) string { return "`" + strings.Repeat(`\u{`, n/3) }},
	{"javascript", "unclosed regular expressions", func(n int) string { return "=" + strings.Repeat("/[", n/2) }},
	{"javascript", "regular expressions", func(n int) string { return strings.Repeat("(/a/", n/4) }},
	{"javascript", "divisions", func(n int) string { return strings.Repeat("a/", n/2) }},
	{"javascript", "regular expression classes", func(n int) string { return "(/" + strings.Repeat("[/]", n/3) }},
	{"python", "triple quote openers", func(n int) string { return strings.Repeat(`'''`, n/3) }},
	{"python", "prefixed strings", func(n int) string { return strings.Repeat(`rb"`, n/3) }},
	{"python", "raw backslashes", func(n int) string { return `r"` + strings.Repeat(`\`, n) }},
	{"python", "named escapes", func(n int) string { return `"` + strings.Repeat(`\N{`, n/3) }},
	{"python", "decorators", func(n int) string { return strings.Repeat("@a\n", n/3) }},
	{"rust", "apostrophes", func(n int) string { return strings.Repeat("'", n) }},
	{"rust", "lifetimes", func(n int) string { return strings.Repeat("'a ", n/3) }},
	{"rust", "raw prefixes", func(n int) string { return strings.Repeat("r#", n/2) }},
	{"rust", "raw string with near closers", func(n int) string {
		return "r" + strings.Repeat("#", 100) + `"` + strings.Repeat(`"`+strings.Repeat("#", 99), n/100)
	}},
	{"rust", "nested comment openers", func(n int) string { return strings.Repeat("/*", n/2) }},
	{"rust", "byte chars", func(n int) string { return strings.Repeat("b'", n/2) }},
	{"rust", "attributes", func(n int) string { return strings.Repeat("#[", n/2) }},
	{"rust", "macros", func(n int) string { return strings.Repeat("a!", n/2) }},
	{"sql", "apostrophes", func(n int) string { return strings.Repeat("'", n) }},
	{"sql", "doubled apostrophes", func(n int) string { return "'" + strings.Repeat("''", n/2) }},
	{"sql", "nested comment openers", func(n int) string { return strings.Repeat("/*", n/2) }},
	{"sql", "line comments", func(n int) string { return strings.Repeat("--\n", n/3) }},
	{"sql", "prefixed strings", func(n int) string { return strings.Repeat("N'", n/2) }},
	{"sql", "long keyword-like word", func(n int) string { return strings.Repeat("SELECT", n/6) }},
	{"json", "keys", func(n int) string { return strings.Repeat(`"a":`, n/4) }},
	{"json", "quotes", func(n int) string { return strings.Repeat(`"`, n) }},
	{"json", "long string before a colon", func(n int) string { return `"` + strings.Repeat(`\"`, n/2) + `":` }},
	{"json", "minus signs", func(n int) string { return strings.Repeat("-", n) }},
	{"json", "negative numbers", func(n int) string { return strings.Repeat("-1,", n/3) }},
	{"pascal", "directive openers", func(n int) string { return strings.Repeat("{$", n/2) }},
	{"pascal", "star directive openers", func(n int) string { return strings.Repeat("(*$", n/3) }},
	{"pascal", "comment openers", func(n int) string { return strings.Repeat("{", n) }},
	{"pascal", "hashes", func(n int) string { return strings.Repeat("#", n) }},
	{"pascal", "char codes", func(n int) string { return strings.Repeat("#$", n/2) }},
	{"pascal", "number prefixes", func(n int) string { return strings.Repeat("$%&", n/3) }},
	{"pascal", "doubled apostrophes", func(n int) string { return "'" + strings.Repeat("''", n/2) }},
	{"xml", "unclosed tags", func(n int) string { return strings.Repeat("<a", n/2) }},
	{"xml", "attributes without values", func(n int) string { return "<a" + strings.Repeat(" b=", n/3) }},
	{"xml", "unclosed attribute value", func(n int) string { return `<a b="` + strings.Repeat("x", n) }},
	{"xml", "comment openers", func(n int) string { return strings.Repeat("<!--", n/4) }},
	{"xml", "cdata openers", func(n int) string { return strings.Repeat("<![CDATA[", n/9) }},
	{"xml", "processing instruction openers", func(n int) string { return strings.Repeat("<?", n/2) }},
	{"xml", "ampersands", func(n int) string { return strings.Repeat("&", n) }},
	{"xml", "almost entities", func(n int) string { return strings.Repeat("&aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", n/37) }},
	{"html", "less-than signs in a script", func(n int) string { return "<script>" + strings.Repeat("<", n) }},
	{"html", "script openers", func(n int) string { return strings.Repeat("<script>", n/8) }},
	{"html", "near script closers", func(n int) string { return "<script>" + strings.Repeat("</scrip", n/7) }},
	{"html", "closing tag starts", func(n int) string { return strings.Repeat("</", n/2) }},
	{"css", "opening braces", func(n int) string { return strings.Repeat("{", n) }},
	{"css", "closing braces", func(n int) string { return strings.Repeat("}", n) }},
	{"css", "properties", func(n int) string { return "{" + strings.Repeat("a:", n/2) }},
	{"css", "at rules", func(n int) string { return strings.Repeat("@a", n/2) }},
	{"css", "hashes", func(n int) string { return strings.Repeat("#", n) }},
	{"css", "url openers", func(n int) string { return strings.Repeat("url(", n/4) }},
	{"css", "pseudo classes", func(n int) string { return strings.Repeat("::", n/2) }},
	{"css", "string escapes", func(n int) string { return `"` + strings.Repeat(`\`, n) }},
	{"yaml", "list markers", func(n int) string { return strings.Repeat("- ", n/2) }},
	{"yaml", "keys on one line", func(n int) string { return strings.Repeat("a: ", n/3) }},
	{"yaml", "block scalar", func(n int) string { return "a: |\n" + strings.Repeat(" a\n", n/3) }},
	{"yaml", "block scalar indicators", func(n int) string { return strings.Repeat("|\n", n/2) }},
	{"yaml", "anchors", func(n int) string { return strings.Repeat("&", n) }},
	{"yaml", "quotes", func(n int) string { return strings.Repeat(`"`, n) }},
	{"yaml", "colons without spaces", func(n int) string { return strings.Repeat("a:b", n/3) }},
	{"yaml", "flow openers", func(n int) string { return strings.Repeat("[{", n/2) }},
	{"yaml", "spaces before a colon", func(n int) string { return "a" + strings.Repeat(" ", n) + ":" }},
	{"toml", "header openers", func(n int) string { return strings.Repeat("[", n) }},
	{"toml", "header lines", func(n int) string { return strings.Repeat("[a\n", n/3) }},
	{"toml", "quoted keys", func(n int) string { return strings.Repeat(`"a" `, n/4) }},
	{"toml", "numbers with dots", func(n int) string { return strings.Repeat("1.", n/2) }},
	{"toml", "dotted keys", func(n int) string { return strings.Repeat("a.", n/2) }},
	{"sh", "command substitution openers", func(n int) string { return strings.Repeat("$(", n/2) }},
	{"sh", "closing parens", func(n int) string { return strings.Repeat(")", n) }},
	{"sh", "substitutions in strings", func(n int) string { return strings.Repeat(`"$(`, n/3) }},
	{"sh", "arithmetic openers", func(n int) string { return strings.Repeat("$((", n/3) }},
	{"sh", "many here documents", func(n int) string { return strings.Repeat("<<a ", n/8) + "\n" + strings.Repeat("x\n", n/4) }},
	{"sh", "here document operators", func(n int) string { return strings.Repeat("<<", n/2) }},
	{"sh", "brace expansions", func(n int) string { return strings.Repeat("${", n/2) }},
	{"sh", "backslashes", func(n int) string { return strings.Repeat(`\`, n) }},
	{"sh", "apostrophes", func(n int) string { return strings.Repeat("'", n) }},
	{"sh", "backticks", func(n int) string { return strings.Repeat("`", n) }},
	{"sh", "assignments", func(n int) string { return strings.Repeat("a=b ", n/4) }},
	{"sh", "for loops", func(n int) string { return strings.Repeat("for ", n/4) }},
	{"diff", "huge hunk", func(n int) string { return "@@ -1,1000000000 +1 @@\n" + strings.Repeat("+\n", n/2) }},
	{"diff", "hunk headers", func(n int) string { return strings.Repeat("@@", n/2) }},
	{"diff", "header lines", func(n int) string { return strings.Repeat("+++ \n", n/5) }},
	{"markdown", "fences", func(n int) string { return strings.Repeat("```\n", n/4) }},
	{"markdown", "unclosed long fence", func(n int) string { return "````\n" + strings.Repeat("```\n", n/4) }},
	{"markdown", "backtick runs of growing length", func(n int) string {
		var b strings.Builder
		for k := 1; b.Len() < n; k++ {
			b.WriteString(strings.Repeat("`", k) + " ")
		}
		return b.String()
	}},
	{"markdown", "backticks", func(n int) string { return strings.Repeat("` ", n/2) }},
	{"markdown", "link destinations", func(n int) string { return strings.Repeat("](", n/2) }},
	{"markdown", "autolink starts", func(n int) string { return strings.Repeat("<a:", n/3) }},
	{"markdown", "nested quotes", func(n int) string { return strings.Repeat("> ", n/2) }},
	{"markdown", "escapes", func(n int) string { return strings.Repeat(`\*`, n/2) }},
	{"markdown", "headings", func(n int) string { return strings.Repeat("# a\n", n/4) }},
	{"markdown", "html lines", func(n int) string { return strings.Repeat("<div", n/4) }},
	{"markdown", "reference definitions", func(n int) string { return strings.Repeat("[a]:\n", n/5) }},
}

// REQ: DET-002
func TestPathological(t *testing.T) {
	if testing.Short() {
		t.Skip("timing test")
	}
	// Rozmiary różnią się 16 razy: wzrost liniowy to czynnik 16, kwadratowy 256. Próg 96 jak w gomdgo.
	// Mierzymy czas procesora przy wyłączonym odśmiecaniu, najlepszy z trzech przebiegów (docs/testing.md).
	const small, factor, slack = 20_000, 16, 6
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	measure := func(lang string, src []byte, runs int) time.Duration {
		best := time.Duration(1 << 62)
		for i := 0; i < runs; i++ {
			runtime.GC()
			start := cpuTime()
			HTML(lang, src)
			best = min(best, cpuTime()-start)
		}
		return best
	}
	for _, c := range pathological {
		var base, big time.Duration
		for attempt := 0; attempt < 2; attempt++ {
			base = max(measure(c.lang, []byte(c.gen(small)), 3), 100*time.Microsecond)
			big = measure(c.lang, []byte(c.gen(small*factor)), 3)
			if big <= base*factor*slack {
				break
			}
		}
		if big > base*factor*slack {
			t.Errorf("%s/%s: %v for size %d, %v for size %d", c.lang, c.name, base, small, big, small*factor)
		}
	}
}

// REQ: DET-003 SAFE-001 SAFE-002 DET-001
func FuzzHTML(f *testing.F) {
	for _, in := range trickyInputs {
		f.Add(in)
	}
	for _, path := range corpusFiles(f) {
		src, _ := os.ReadFile(path)
		f.Add(string(src))
	}
	for _, c := range pathological {
		f.Add(c.gen(30))
	}
	names := primaryNames()
	f.Fuzz(func(t *testing.T, src string) {
		for _, name := range names {
			out, ok := HTML(name, []byte(src))
			if !ok {
				t.Fatalf("%s: language not found", name)
			}
			if err := checkSafe(out); err != nil {
				t.Fatalf("%s: %v\ninput:  %q\noutput: %q", name, err, src, out)
			}
			if got, want := plainText(out), expectedText([]byte(src)); !bytes.Equal(got, want) {
				t.Fatalf("%s: round trip differs\ninput: %q\ngot:   %q", name, src, got)
			}
			if again, _ := HTML(name, []byte(src)); !bytes.Equal(again, out) {
				t.Fatalf("%s: output differs between calls\ninput: %q", name, src)
			}
		}
	})
}

// BenchmarkCorpus: korpus danego języka powtórzony do około 1 MB.
func BenchmarkCorpus(b *testing.B) {
	for _, name := range primaryNames() {
		var src []byte
		for _, path := range corpusFiles(b) {
			if strings.Contains(path, "/"+name+"/") {
				part, _ := os.ReadFile(path)
				src = append(src, part...)
			}
		}
		src = bytes.Repeat(src, 1<<20/len(src)+1)
		b.Run(name, func(b *testing.B) {
			b.SetBytes(int64(len(src)))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				HTML(name, src)
			}
		})
	}
}
