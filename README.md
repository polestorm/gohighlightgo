# gohighlightgo

Syntax highlighting for source code in pure Go: code in, HTML `<span>` elements out.

Status: **early development**. Nineteen languages work: Go, C, C++, Java, JavaScript, TypeScript, Python, Rust,
SQL, JSON, Pascal/Delphi, XML, HTML, CSS, YAML, TOML, shell, diff and Markdown; see `docs/roadmap.md` for the remaining stages and `docs/adr/0001-scope-and-acceptance.md` for the scope.

## Goals

- **No dependencies** outside the Go standard library.
- **Safe by construction**: every byte of the input is HTML-escaped; the output contains only `<span>`
  elements with a fixed, small set of classes.
- **Linear time**, also for hostile input: one pass over the bytes, no regular expressions, no backtracking.
- **Deterministic**: the same input gives the same output, byte for byte.
- **Small**: a table-driven scanner and a few dozen languages, not hundreds. A language that is not known
  is left alone.
- **Independent of any Markdown parser.** The library does not import
  [gomdgo](https://github.com/polestorm/gomdgo); a caller connects the two in a few lines.

## API

```go
import "github.com/polestorm/gohighlightgo/highlight"

out, ok := highlight.HTML("go", src)      // escaped code with <span class="hl-…">; ok is false for an unknown language
names := highlight.Languages()            // names and aliases: "go", "golang", "sh", "bash", …
css, _ := highlight.ThemeCSS("light")     // stylesheet for the classes: "light" (default), "dark" or "auto"
```

The stylesheet applies inside an element with the class `hl`, so a block is written as
`<pre class="hl"><code class="language-go">…</code></pre>`.

## Command line

```sh
go install github.com/polestorm/gohighlightgo/cmd/gohighlightgo@latest
gohighlightgo main.go > main.html                  # fragment; the language comes from the extension
gohighlightgo -lang sql -standalone -theme dark < q.sql > q.html
gohighlightgo -list
```

## Use with gomdgo

gomdgo hands fenced code blocks to the caller through `Config.CodeBlocks` (language names match without regard
to case). The adapter lives in the caller's module, so neither library depends on the other:

```go
blocks := map[string]markdown.CodeBlockFunc{}
for _, name := range highlight.Languages() {
	blocks[name] = func(b markdown.CodeBlock) ([]byte, error) {
		code, _ := highlight.HTML(b.Lang, b.Code)
		return wrap(b.Lang, code), nil // <pre class="hl"><code class="language-…">…</code></pre>
	}
}
```

A program that wants highlighting to be optional puts the adapter in a file with a build tag
(`//go:build highlight`), so `go build -tags highlight` includes it and a plain `go build` leaves it out.

## Family

| Project | Role |
|---|---|
| [gomdgo](https://github.com/polestorm/gomdgo) | Markdown → HTML |
| [gomermaidgo](https://github.com/polestorm/gomermaidgo) | Mermaid diagrams → SVG |
| [gopdfgo](https://github.com/polestorm/gopdfgo) | HTML + CSS + SVG → PDF |
| [gopressgo](https://github.com/polestorm/gopressgo) | the pipeline Markdown → PDF |
| [gopeekgo](https://github.com/polestorm/gopeekgo) | Quick Look preview of Markdown on macOS |
| gohighlightgo | syntax highlighting for code blocks |

## Licence

Apache License 2.0 — see `LICENSE` and `NOTICE`.
