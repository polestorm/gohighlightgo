# Usage

[← README](../README.md)

```go
out, ok := highlight.HTML("python", src)  // fragment: escaped code with <span class="hl-…">
names := highlight.Languages()            // every accepted name and alias, sorted
css, _ := highlight.ThemeCSS("light")     // "light" (the default), "dark" or "auto" (follows the system)
err := highlight.Register(definition)     // add a language from a definition file (docs/definitions.md)
```

The classes are `hl-kw` (keyword), `hl-type`, `hl-builtin`, `hl-str`, `hl-esc` (escape in a string),
`hl-num`, `hl-com` (comment), `hl-op`, `hl-fn` (name before a parenthesis), `hl-tag`, `hl-attr`, `hl-ins`,
`hl-del` and `hl-meta` (preprocessor, annotations). The stylesheet applies inside an element with the class
`hl`, so a block is written as `<pre class="hl"><code class="language-go">…</code></pre>`. Invalid UTF-8 and NUL
bytes are replaced with U+FFFD.

## Command line

```sh
go install github.com/polestorm/gohighlightgo/cmd/gohighlightgo@latest
gohighlightgo main.go > main.html                  # fragment; the language comes from the extension
gohighlightgo -lang sql -standalone -theme dark < q.sql > q.html
gohighlightgo -list
gohighlightgo -def mylang.hl -lang mylang code.txt  # a language from a definition file
```

## With gomdgo

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
Complete adapters are in [adapters/](adapters/).
