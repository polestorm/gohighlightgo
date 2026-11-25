package highlight

import (
	"bytes"
	"errors"
	"strings"
	"sync"
	"testing"
)

// REQ: DEF-002 DEF-003
func TestDefinitionErrors(t *testing.T) {
	cases := []struct{ src, want string }{
		{"", "no name"},
		{"# only a comment\n", "no name"},
		{"name Go", "lower case"},
		{"name x\nbogus 1", "line 2: unknown directive"},
		{"name x\nstring \"", "line 2: string needs"},
		{"name x\nstring \" \" : sometimes", "unknown option"},
		{"name x\nblock-comment /*", "line 2: block-comment needs"},
		{"name x\nescape : x=9", "x must be"},
		{"name x\nescape x=2", "only options"},
		{"name x\nhook nothing", "unknown hook"},
		{"extends nothing\nname x", "no table language"},
		{"name x\nextends go", "must come before"},
		{"name x\ncalls yes", "takes no values"},
		{"name x\noperators + -", "one group"},
		{"name x\nprefix toolongprefix", "longer than 8"},
		{"name x\nkeyword " + strings.Repeat("a", 33), "longer than 32"},
		{"name x\nnumber : sideways", "unknown option"},
		{"name x\noperators ab\noperators cd", "given twice"},
		{"name x\nident-extra \"", "without quotes"},
		{strings.Repeat("#", maxDefinition+1), "longer than"},
	}
	for _, c := range cases {
		_, err := parseDefinition([]byte(c.src), lookup)
		var de *DefinitionError
		if err == nil || !errors.As(err, &de) || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%q: error %v, want %q", c.src, err, c.want)
		}
	}
}

// REQ: API-007 DEF-001 DEF-004
func TestRegister(t *testing.T) {
	def := []byte(`# a test language
name     zzlang zzl
keyword  when then
type     Num
builtin  yes no
line-comment  ;;
block-comment  {- -}  : nested
string   " "  : escapes interp=dollar
prefix   raw  : literal
escape   : x=2
operators  +-*=
calls
`)
	if err := Register(def); err != nil {
		t.Fatal(err)
	}
	out, ok := HTML("ZZL", []byte(`when f(x) then "a\x41 $v ${w}" raw"\" ;; c {- {- -} -} 1`))
	want := `<span class="hl-kw">when</span> <span class="hl-fn">f</span>(x) <span class="hl-kw">then</span> ` +
		`<span class="hl-str">"a</span><span class="hl-esc">\x41</span><span class="hl-str"> </span><span class="hl-esc">$v</span>` +
		`<span class="hl-str"> </span><span class="hl-esc">${w}</span><span class="hl-str">"</span> <span class="hl-str">raw"\"</span> ` +
		`<span class="hl-com">;; c {- {- -} -} 1</span>`
	if !ok || string(out) != want {
		t.Errorf("got  %s\nwant %s", out, want)
	}
	found := false
	for _, n := range Languages() {
		found = found || n == "zzl"
	}
	if !found {
		t.Error("registered alias missing from Languages")
	}
	if err := Register([]byte("name zzlang")); err == nil || !strings.Contains(err.Error(), "already in use") {
		t.Errorf("duplicate name: %v", err)
	}
	if err := Register([]byte("extends zzlang\nname zzlang2\nkeyword also")); err != nil {
		t.Errorf("extends a registered language: %v", err)
	}
	if out, _ := HTML("zzlang2", []byte("when also")); string(out) != `<span class="hl-kw">when</span> <span class="hl-kw">also</span>` {
		t.Errorf("extended language: %s", out)
	}
	if out, _ := HTML("zzlang", []byte("also")); string(out) != "also" {
		t.Errorf("extends changed its base: %s", out)
	}
	// Język zarejestrowany działa też w bloku Markdown.
	if out, _ := HTML("markdown", []byte("```zzl\nwhen\n```\n")); !strings.Contains(string(out), `<span class="hl-kw">when</span>`) {
		t.Errorf("registered language in a Markdown fence: %s", out)
	}
}

// REQ: API-007
func TestRegisterConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				if g == 0 && i < 10 {
					Register([]byte("name zzconc" + string(rune('a'+i)) + "\nkeyword k"))
				}
				if _, ok := HTML("go", []byte("func f() {}")); !ok {
					t.Error("go missing during Register")
				}
				_ = Languages()
			}
		}(g)
	}
	wg.Wait()
}

// REQ: DEF-004
// FuzzDefinition: dowolna definicja przyjęta przez parser daje wynik bezpieczny, odwracalny i bez paniki dla
// dowolnego kodu (gwarancje HTML z konstrukcji formatu).
func FuzzDefinition(f *testing.F) {
	files, _ := builtinDefinitions.ReadDir("lang")
	for _, file := range files {
		src, _ := builtinDefinitions.ReadFile("lang/" + file.Name())
		f.Add(src, []byte("x = \"a\\n\" // c\n/* b */ 1.5 'c' @d"))
	}
	f.Add([]byte("name q\nstring a a\nline-comment a\nblock-comment a a : nested\nprefix a : verbatim interp=braces"), []byte("aaaa{a}aa"))
	f.Fuzz(func(t *testing.T, def, code []byte) {
		l, err := parseDefinition(def, lookup)
		if err != nil {
			return
		}
		w := writer{}
		l.scan(&w, sanitize(code))
		w.close()
		if err := checkSafe(w.out); err != nil {
			t.Fatalf("%v\ndefinition: %q\ncode: %q\noutput: %q", err, def, code, w.out)
		}
		if got := plainText(w.out); !bytes.Equal(got, expectedText(code)) {
			t.Fatalf("round trip differs\ndefinition: %q\ncode: %q", def, code)
		}
	})
}
