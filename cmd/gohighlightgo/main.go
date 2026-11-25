// Polecenie gohighlightgo koloruje kod źródłowy biblioteką gohighlightgo.
//
//	gohighlightgo [flagi] [plik]
//
// Bez nazwy pliku (albo z nazwą "-") czyta standardowe wejście; wynik trafia na standardowe wyjście albo do
// pliku wskazanego flagą -o. Język podaje flaga -lang albo rozszerzenie pliku (docs/adr/0005-command-line.md).
package main

import (
	"flag"
	"fmt"
	"html"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/polestorm/gohighlightgo/highlight"
)

// extensions: rozszerzenie pliku → nazwa języka. To nie jest wykrywanie z treści (ADR 0001, punkt 4).
var extensions = map[string]string{
	".go": "go", ".c": "c", ".h": "c", ".cc": "cpp", ".cpp": "cpp", ".cxx": "cpp", ".hpp": "cpp", ".hh": "cpp",
	".java": "java", ".js": "javascript", ".mjs": "javascript", ".cjs": "javascript", ".ts": "typescript",
	".mts": "typescript", ".cts": "typescript", ".py": "python", ".rs": "rust", ".sql": "sql", ".json": "json",
	".pas": "pascal", ".dpr": "pascal", ".dpk": "pascal", ".pp": "pascal", ".inc": "pascal", ".xml": "xml",
	".xsd": "xml", ".xsl": "xml", ".xslt": "xml", ".svg": "xml", ".plist": "xml", ".html": "html", ".htm": "html",
	".xhtml": "html", ".css": "css", ".yaml": "yaml", ".yml": "yaml", ".toml": "toml", ".sh": "sh", ".bash": "sh",
	".zsh": "sh", ".ksh": "sh", ".diff": "diff", ".patch": "diff", ".md": "markdown", ".markdown": "markdown",
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

// run wykonuje polecenie i zwraca kod wyjścia: 0 przy powodzeniu, 1 przy błędzie odczytu lub zapisu, 2 przy
// błędnych argumentach albo nieznanym języku.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("gohighlightgo", flag.ContinueOnError)
	fs.SetOutput(stderr)
	var (
		output     = fs.String("o", "", "write the result to this `file` instead of standard output")
		lang       = fs.String("lang", "", "`language` of the code (default: from the file extension)")
		list       = fs.Bool("list", false, "print the accepted language names and exit")
		standalone = fs.Bool("standalone", false, "write a complete HTML document with a stylesheet")
		title      = fs.String("title", "", "document `title` for -standalone (default: the file name)")
		themeName  = fs.String("theme", highlight.DefaultTheme, "`theme` for -standalone: "+strings.Join(highlight.Themes(), ", "))
		defs       []string
	)
	fs.Func("def", "add a language from a definition `file` (docs/definitions.md); may be repeated", func(s string) error {
		defs = append(defs, s)
		return nil
	})
	fs.Usage = func() {
		fmt.Fprintln(stderr, "usage: gohighlightgo [flags] [file]")
		fmt.Fprintln(stderr, "Highlights source code as HTML. Reads standard input when no file is given.")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}
	fail := func(code int, format string, a ...any) int {
		fmt.Fprintf(stderr, "gohighlightgo: %s\n", fmt.Sprintf(format, a...))
		return code
	}
	for _, d := range defs {
		src, err := os.ReadFile(d)
		if err != nil {
			return fail(1, "%v", err)
		}
		if err := highlight.Register(src); err != nil {
			return fail(2, "%s: %v", d, err)
		}
	}
	if *list {
		fmt.Fprintln(stdout, strings.Join(highlight.Languages(), "\n"))
		return 0
	}
	if fs.NArg() > 1 {
		return fail(2, "expected at most one input file, got %d", fs.NArg())
	}
	name := fs.Arg(0)
	if *lang == "" {
		*lang = extensions[strings.ToLower(filepath.Ext(name))]
		if *lang == "" {
			return fail(2, "no -lang given and no language known for %q", name)
		}
	}
	css, ok := highlight.ThemeCSS(*themeName)
	if !ok {
		return fail(2, "unknown -theme %q (%s)", *themeName, strings.Join(highlight.Themes(), ", "))
	}

	var src []byte
	var err error
	if name == "" || name == "-" {
		src, err = io.ReadAll(stdin)
	} else {
		src, err = os.ReadFile(name)
	}
	if err != nil {
		return fail(1, "%v", err)
	}
	body, ok := highlight.HTML(*lang, src)
	if !ok {
		return fail(2, "unknown language %q (see -list)", *lang)
	}
	out := body
	if *standalone {
		if *title == "" {
			*title = filepath.Base(name)
		}
		out = []byte("<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"utf-8\">\n<title>" + html.EscapeString(*title) +
			"</title>\n<style>\nbody { margin: 0; }\npre.hl { margin: 0; padding: 1em; min-height: 100vh; box-sizing: border-box; }\n" +
			css + "</style>\n</head>\n<body>\n<pre class=\"hl\"><code class=\"language-" + html.EscapeString(*lang) + "\">")
		out = append(out, body...)
		out = append(out, "</code></pre>\n</body>\n</html>\n"...)
	}
	if *output == "" {
		_, err = stdout.Write(out)
	} else {
		err = os.WriteFile(*output, out, 0o644)
	}
	if err != nil {
		return fail(1, "%v", err)
	}
	return 0
}
