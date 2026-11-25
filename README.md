# gohighlightgo

Syntax highlighting for source code in pure Go: code in, HTML `<span class="hl-…">` elements out. Written from
the language specifications, with no dependencies outside the Go standard library. The output is safe by
construction, the same input always gives the same HTML, byte for byte, and the time is linear in the length of
the input, also for hostile input.

```go
import "github.com/polestorm/gohighlightgo/highlight"

out, ok := highlight.HTML("go", src) // escaped code with <span class="hl-…">; ok is false for an unknown language
```

## Why gohighlightgo

Compared with [chroma](https://github.com/alecthomas/chroma), the usual highlighter in Go, and with
[highlight](http://www.andre-simon.de/) by André Simon, the engine behind the QLMarkdown preview on macOS
(details in [Comparison](#comparison)):

- **Fast**: about 100 MB/s; 98.5–98.8% less time than chroma on the test inputs and 99% less than highlight on
  a document full of code.
- **Small**: 3,429 lines of Go and 365 lines of language definitions, 268 kB added to a program, no
  dependencies, nothing loaded at run time.
- **Extensible without a release**: a language made of words, comments, strings and numbers is a short text
  file ([docs/definitions.md](docs/definitions.md)); a program adds its own with `highlight.Register`, the
  command with `-def`. Definitions cannot contain regular expressions or code, so every guarantee below holds for
  them too.
- **Safe by construction**: the output contains only escaped text and `<span>` elements with classes from a
  fixed list of 14, never nested and never spanning a line break. An independent checker verifies this in the
  tests and in fuzzing, and removing the tags gives back the input.
- **Linear time on hostile input**: one pass over the bytes, no regular expressions; 142 pathological cases are
  timed in the test suite, and the library is fuzzed.
- **Deterministic**: the same input gives the same output, byte for byte, documented and tested.
- **Made to be built on**: a single function, a stylesheet per theme, a command, and a few lines to connect it to
  a Markdown renderer (for example [gomdgo](https://github.com/polestorm/gomdgo)) behind a build tag.

Where it is weaker: 26 languages instead of about 300, two themes instead of dozens, no detection of the language
from the content, and lexical highlighting only (a conversion `T(x)` looks like a call). It is new.

## Contents

[Languages](#languages) · [Usage](#usage) · [Comparison](#comparison) · [Status](#status) ·
[Documentation](#documentation) · [How the comparison was measured](#how-the-comparison-was-measured) ·
[License](#license)

## Languages

| Language | Names |
|---|---|
| Go | `go`, `golang` |
| C, C++ | `c`, `h`; `cpp`, `c++`, `cxx`, `cc`, `hpp` |
| Java | `java` |
| JavaScript, TypeScript | `javascript`, `js`, `mjs`, `cjs`; `typescript`, `ts`, `mts`, `cts` |
| C# | `csharp`, `cs`, `c#` |
| Kotlin | `kotlin`, `kt`, `kts` |
| Swift | `swift` |
| Scala | `scala`, `sc` |
| Dart | `dart` |
| Python | `python`, `py`, `python3` |
| Rust | `rust`, `rs` |
| Lua | `lua` |
| R | `r` |
| SQL | `sql` |
| JSON | `json`, `jsonc` |
| Pascal, Delphi | `pascal`, `delphi`, `pas`, `objectpascal` |
| XML, HTML | `xml`, `xsd`, `xsl`, `xslt`, `svg`, `plist`; `html`, `htm`, `xhtml` (with JavaScript in `<script>` and CSS in `<style>`) |
| CSS | `css` |
| YAML, TOML | `yaml`, `yml`; `toml` |
| Shell | `sh`, `bash`, `shell`, `zsh`, `ksh` |
| diff | `diff`, `patch` |
| Markdown | `markdown`, `md` (fenced code highlighted in its own language) |

Names are matched without regard to case. Each language is described from its specification; what is
highlighted, and where the highlighting is only an approximation, is listed in
[requirements/languages/](requirements/languages/). Eighteen of them are definitions in
[highlight/lang/](highlight/lang/); markup, CSS, YAML, TOML, shell, diff and Markdown have scanners in Go.

## Usage

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

### Command line

```sh
go install github.com/polestorm/gohighlightgo/cmd/gohighlightgo@latest
gohighlightgo main.go > main.html                  # fragment; the language comes from the extension
gohighlightgo -lang sql -standalone -theme dark < q.sql > q.html
gohighlightgo -list
gohighlightgo -def mylang.hl -lang mylang code.txt  # a language from a definition file
```

### With gomdgo

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
Complete adapters are in [docs/adapters/](docs/adapters/).

## Comparison

| | gohighlightgo 0.1.0 | chroma 2.27.0 | highlight (QLMarkdown 1.5.7)[^hl] |
|---|:---:|:---:|:---:|
| Languages | 26 | 297 🥇 | 263 |
| New language without a new release | a text file (`Register`, `-def`) | an XML file (`NewXMLLexer`, `Register`) | a Lua file |
| Themes | 2, and `auto` | 74 | 97 🥇 |
| Dependencies[^deps] | 0 🥇 | 1 | a Lua interpreter |
| Lines of code[^size] | 3,429 and 365 lines of definitions 🥇 | 31,199 and 279 lexer files (2.4 MB) | — |
| Added to a program | 268 kB 🥇 | 4,774 kB | 3,416 kB and 2.7 MB of definitions and themes |
| Time, the corpus as 21 blocks (14.5 kB)[^bench] | 0.16 ms 🥇 | 13.3 ms | — |
| Time, 1 MB of Go | 10.0 ms 🥇 | 690 ms | — |
| Time, a document with 1.3 MB of code[^doc] | 17 ms 🥇 | — | 2,079 ms |
| Memory, the corpus as 21 blocks | 0.14 MB 🥇 | 17.5 MB | — |
| Memory, 1 MB of Go | 4.2 MB 🥇 | 819 MB | — |
| Output for the corpus | 65.9 kB, 1,817 spans 🥇 | 117.4 kB, 4,317 spans | — |
| HTML in the code is escaped[^safe] | yes | yes | yes |
| Guarantees hold for added languages | yes, by construction 🥇 | not documented | not documented |
| Elements span line breaks | never 🥇 | yes (whitespace) | — |
| Time on hostile input | linear, tested | not documented | not documented |
| Deterministic output | yes, documented and tested | not documented | not documented |
| Language detection from the content | no | yes | — |
| Command line | yes | yes | yes |
| License | Apache 2.0 | MIT | GPL 3.0 |

[^hl]: highlight is measured inside QLMarkdown, the Markdown preview for macOS that bundles it as a library;
    it is only run as a program. gohighlightgo is written from the language specifications, not from the code or
    the language definitions of other highlighters ([docs/provenance.md](docs/provenance.md)).
[^deps]: Modules compiled into a program. chroma requires two more modules for its tests.
[^size]: Lines of Go in the packages a "standard input to HTML" program pulls in, without tests. "Added to a
    program" is the size of that program minus the same program without highlighting (`-trimpath -ldflags="-s
    -w"`, macOS arm64); for highlight, the arm64 part of its library in QLMarkdown, which also holds the Lua
    interpreter, and its language definitions and themes.
[^bench]: The 21 files of this project's corpus one by one, in the 19 languages it had on 2026-11-25 before
    the definition files, and the Go files of the corpus joined to 1 MB. Median of five runs of `BenchmarkRivals`; memory is the number of bytes allocated per run.
    chroma runs its lexer and its HTML formatter with classes, without `<pre>`, which gives output of the same
    form. Measured on 2026-11-25 on one Apple-silicon laptop. 🥇 marks the best value in a row.
[^doc]: The corpus as a Markdown document with one fenced block per file, joined to 1.3 MB, converted to an
    HTML page with and without highlighting; the difference is the cost of highlighting. gohighlightgo: the
    `gomdgo-mermaid` command built with and without highlighting; highlight: QLMarkdown's command with
    `--syntax-highlight on` and `off`. Median of 15 runs of each process.
[^safe]: Output for `</code></pre><script>alert(1)</script>` and an `<img onerror>` in each of the 26
    languages (`TestRivalsSafety`); for highlight, the same text in Go, HTML, XML, shell and Markdown blocks.

## Status

Version 0.1.0: the planned languages, themes and command are done; hardening in progress. The API may still
change before 1.0. Requires Go 1.22 or newer.

- 26 languages, each with a requirements file and a corpus file whose output is checked byte for byte.
- Language definitions: the 11 table languages moved to definition files without a single byte of output
  changing; any definition the parser accepts is fuzzed together with random code.
- Safety and round trip (removing the tags gives back the input) are checked on every test input and in fuzzing.
- Fuzzing: short runs during development (about 13 minutes in total, about 18 million inputs, each in every
  language supported at the time) without a failure after one fix; the 24-hour run is still to be done.
- Not yet done: highlighting in the gomdgo command and in the gopeekgo preview (the adapters are ready).

## Documentation

- This README and the package documentation (`go doc github.com/polestorm/gohighlightgo/highlight`).
- [`docs/definitions.md`](docs/definitions.md): the format of language definitions.
- In Polish: [`docs/provenance.md`](docs/provenance.md) (sources and clean-room rules),
  [`docs/testing.md`](docs/testing.md) (tests, measurements, comparison), [`requirements/`](requirements/)
  (numbered requirements, one file per language), [`docs/adr/`](docs/adr/) (decisions),
  [`docs/roadmap.md`](docs/roadmap.md).
- `go run ./tools/review` builds one HTML page with every corpus file in both themes next to the output of
  QLMarkdown, for review by a person.

## How the comparison was measured

gohighlightgo and chroma are run through their public API on the same inputs by `TestRivalsOutput`,
`TestRivalsSafety` and `BenchmarkRivals` in `tools/compare` (a separate module, so the library itself has no
dependencies). highlight is not available as a Go library, so it is measured as a whole program, the way it is
used in practice, and its share is the difference between runs with and without highlighting.

The corpus is small and written to cover each language's specification, not to imitate typical code; the 1 MB
and 1.3 MB inputs are the corpus repeated. The three highlighters split code into tokens differently (chroma
marks whitespace and punctuation, gohighlightgo does not), so the output sizes compare the work each one does,
not identical results.

## License

Apache License 2.0, see `LICENSE` and `NOTICE`. gohighlightgo contains no code by other authors.
