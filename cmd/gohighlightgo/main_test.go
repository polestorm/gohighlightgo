package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/polestorm/gohighlightgo/highlight"
)

func runCLI(args []string, stdin string) (code int, out, errOut string) {
	var o, e bytes.Buffer
	code = run(args, strings.NewReader(stdin), &o, &e)
	return code, o.String(), e.String()
}

func TestRun(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "main.GO")
	if err := os.WriteFile(file, []byte("func f() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		args       []string
		stdin      string
		code       int
		out, inErr string
	}{
		{[]string{"-lang", "go"}, "x := 1", 0, `x <span class="hl-op">:=</span> <span class="hl-num">1</span>`, ""},
		{[]string{file}, "", 0, `<span class="hl-kw">func</span> <span class="hl-fn">f</span>() {}` + "\n", ""},
		{[]string{"-lang", "python", file}, "", 0, `func <span class="hl-fn">f</span>`, ""}, // -lang przed rozszerzeniem
		{[]string{}, "x", 2, "", "no -lang given"},
		{[]string{"-lang", "klingon"}, "x", 2, "", "unknown language"},
		{[]string{"-lang", "go", "-theme", "x"}, "x", 2, "", "unknown -theme"},
		{[]string{"a.go", "b.go"}, "", 2, "", "at most one"},
		{[]string{filepath.Join(dir, "missing.go")}, "", 1, "", "no such file"},
		{[]string{"-bogus"}, "", 2, "", "flag provided but not defined"},
	}
	for _, c := range cases {
		code, out, errOut := runCLI(c.args, c.stdin)
		if code != c.code || !strings.Contains(out, c.out) || !strings.Contains(errOut, c.inErr) {
			t.Errorf("%v: code %d, out %q, err %q", c.args, code, out, errOut)
		}
	}
}

func TestStandaloneAndOutput(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "out.html")
	code, _, errOut := runCLI([]string{"-lang", "go", "-standalone", "-theme", "dark", "-title", "<a>", "-o", target}, "nil")
	if code != 0 {
		t.Fatalf("code %d: %s", code, errOut)
	}
	page, _ := os.ReadFile(target)
	css, _ := highlight.ThemeCSS("dark")
	for _, want := range []string{"<title>&lt;a&gt;</title>", css, `<pre class="hl"><code class="language-go"><span class="hl-builtin">nil</span></code></pre>`} {
		if !strings.Contains(string(page), want) {
			t.Errorf("page has no %q", want)
		}
	}
}

func TestDefinitionFlag(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "demo.hl")
	bad := filepath.Join(dir, "bad.hl")
	os.WriteFile(good, []byte("name clidemo\nkeyword hello\n"), 0o644)
	os.WriteFile(bad, []byte("name clibad\nbogus\n"), 0o644)
	if code, out, errOut := runCLI([]string{"-def", good, "-lang", "clidemo"}, "hello world"); code != 0 || out != `<span class="hl-kw">hello</span> world` {
		t.Errorf("code %d, out %q, err %q", code, out, errOut)
	}
	if code, _, errOut := runCLI([]string{"-def", bad, "-lang", "go"}, "x"); code != 2 || !strings.Contains(errOut, "line 2") {
		t.Errorf("bad definition: code %d, err %q", code, errOut)
	}
	if code, _, _ := runCLI([]string{"-def", filepath.Join(dir, "missing.hl"), "-lang", "go"}, "x"); code != 1 {
		t.Errorf("missing definition: code %d", code)
	}
}

func TestList(t *testing.T) {
	code, out, _ := runCLI([]string{"-list"}, "")
	if code != 0 || strings.Join(highlight.Languages(), "\n")+"\n" != out {
		t.Errorf("code %d, out %q", code, out)
	}
}

// TestExtensions: każde rozszerzenie wskazuje język, który biblioteka zna (ADR 0005, „Skutki”), a każdy język
// ma co najmniej jedno rozszerzenie.
func TestExtensions(t *testing.T) {
	covered := map[string]bool{}
	for ext, lang := range extensions {
		if _, ok := highlight.HTML(lang, nil); !ok {
			t.Errorf("%s: unknown language %q", ext, lang)
		}
		if ext != strings.ToLower(ext) || !strings.HasPrefix(ext, ".") {
			t.Errorf("bad extension %q", ext)
		}
		covered[lang] = true
	}
	for _, name := range []string{"go", "c", "cpp", "java", "javascript", "typescript", "python", "rust", "sql",
		"json", "pascal", "xml", "html", "css", "yaml", "toml", "sh", "diff", "markdown"} {
		if !covered[name] {
			t.Errorf("language %s has no extension", name)
		}
	}
}
