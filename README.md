# gohighlightgo

Syntax highlighting for source code in pure Go. Written from the language specifications, with no dependencies
outside the Go standard library. The output is safe by construction, the same input always gives the same HTML,
byte for byte, and the time is linear in the length of the input, also for hostile input.

```go
import "github.com/polestorm/gohighlightgo/highlight"

out, ok := highlight.HTML("go", src) // escaped code with <span class="hl-…">; ok is false for an unknown language
```

## Why gohighlightgo

Compared with the established highlighters (details in [Comparison](#comparison)):

- **Fast**: about 100 MB/s, the fastest of the highlighters compared.
- **Small**: 5,397 lines of Go and 3,753 lines of language definitions, 536 kB added to a program, no
  dependencies, nothing loaded at run time.
- **Extensible without a release**: a language made of words, comments, strings and numbers is a short text
  file ([docs/definitions.md](docs/definitions.md)).
- **Safe by construction**: only escaped text and `<span>` elements from a fixed list of classes, checked in
  tests and fuzzing.
- **Linear time on hostile input**: one pass, no regular expressions; the only one compared that never hangs or
  slows down on hostile input.
- **Made to be built on**: a single function, a stylesheet per theme, a command, and a few lines to connect it to
  a Markdown renderer (for example [gomdgo](https://github.com/polestorm/gomdgo)) behind a build tag.

Where it is weaker: many languages described more shallowly than in the older tools (keywords, comments, strings
and numbers; few context rules), no detection of the language from the content, and lexical highlighting only (a
conversion `T(x)` looks like a call). It is new.

## Contents

- [Why gohighlightgo](#why-gohighlightgo) · [Comparison](#comparison) ·
  [Documentation](#documentation) · [License](#license)
- [Usage](docs/usage.md): the API, classes and stylesheets, command line, connecting to a Markdown renderer
- [Languages](docs/languages.md): the 298 languages and the names each one accepts
- [How the comparison was measured](docs/comparison.md): the tests and tools behind the comparison table

## Comparison

| | gohighlightgo 0.1.0 | [chroma](https://github.com/alecthomas/chroma) 2.27.0 | [highlight](http://www.andre-simon.de/) 4.21[^hl] |
|---|:---:|:---:|:---:|
| Languages | 298 | 297 | 263 |
| Adding a new language | a `.hl` text file | an XML file | a Lua file |
| Dependencies[^deps] | 0 | 1 | a Lua interpreter |
| Size added to a program[^size] | 536 kB 🥇 | 4 774 kB | 4 964 kB |
| Time, 21 files (14.5 kB)[^bench] | 0.17 ms 🥇 | 13 ms | *not measured* |
| Time, 1 MB of Go | 11 ms 🥇 | 687 ms | *not measured* |
| Time, 1.3 MB of code[^doc] | 15 ms 🥇 | 1 138 ms | 2 256 ms |
| Memory allocated, 21 files[^mem] | 0.14 MB 🥇 | 16.9 MB | *not measured* |
| Memory allocated, 1 MB of Go[^mem] | 4.2 MB 🥇 | 820 MB | *not measured* |
| Peak memory added, 1 MB of Go[^mem] | 7 MB 🥇 | 54 MB | 10 MB |
| Output for the 21 files | 65.9 kB 🥇 | 117.4 kB | *not measured* |
| XSS-safe output[^safe] | yes | yes | yes |
| Elements span line breaks | never 🥇 | yes (whitespace) | *not checked* |
| Hostile input, 64 kB: hang or over 15 s[^hostile] | 0 / 3 278 🥇 | 101 / 2 200 (59 lang) | 8 / 1 474 (8 lang) |
| Hangs on non-ASCII text | no 🥇 | yes, JSONata (`zażółć`, `日本`) | no 🥇 |
| Crashes on input | no 🥇 | yes, YANG (`/*` repeated) | no 🥇 |
| Round-trip fidelity[^text] | 298/298 🥇 | 185/200 | *not measured* |
| Language detection | no | yes 🥇 | *not checked* |
| Themes | 2, and `auto` | 74 | 97 🥇 |
| Output formats | HTML | HTML, SVG, terminal, JSON 🥇 | HTML, LaTeX, RTF, ODT, SVG, terminal and more[^formats] 🥇 |
| Command line | yes | yes | yes |
| GitHub popularity[^pop] | new | 5 088 stars, 568 forks 🥇 | 130 stars, 175 forks (GitLab) 🥇 |
| License | Apache 2.0 | MIT | GPL 3.0 |

To repeat the comparison, see [docs/comparison.md](docs/comparison.md).

[^hl]: An independent C++ program by André Simon, measured through the copy that QLMarkdown 1.5.7 bundles.
    [More](docs/comparison.md#highlight)
[^deps]: Modules compiled into a program. chroma requires two more modules for its tests.
[^size]: A small program that turns standard input into HTML, built with and without highlighting; language
    definitions and themes are counted for all three. [More](docs/comparison.md#size-added-to-a-program)
[^bench]: `go run ./cmd/rivalbench`, median of five runs, 2026-11-26, Apple M4. 🥇 marks the best value in a
    row. [More](docs/comparison.md#time-and-memory-on-the-corpus)
[^doc]: The 21 files as code blocks, repeated to 1.3 MB: for chroma and gohighlightgo the time to highlight every
    block; for highlight the QLMarkdown command with highlighting minus without.
    [More](docs/comparison.md#a-document-full-of-code)
[^hostile]: Eleven hostile patterns at about 4 kB and 64 kB in every language; each call of chroma and highlight
    in its own process, stopped after 15 s. [More](docs/comparison.md#hostile-input)
[^text]: Removing the tags and escapes gives back the input, byte for byte; on the corpus and seven awkward
    inputs in each language. [More](docs/comparison.md#round-trip-fidelity)
[^mem]: Allocated: the sum of all allocations in one run, mostly freed during it. Peak added: the largest memory
    use of a process minus the same process without highlighting. [More](docs/comparison.md#memory)
[^formats]: From its manual: HTML, XHTML, LaTeX, TeX, RTF, ODT, SVG, BBCode, Pango markup and terminal (16, 256
    and true colour). chroma's list is from `formatters.Names()`.
[^pop]: Stars and forks on 2026-11-26; highlight is developed on GitLab.
[^safe]: `</code></pre><script>alert(1)</script>` and an `<img onerror>` in every language chroma also knows; for
    highlight, in five languages. [More](docs/comparison.md#xss-safe-output)

## Documentation

- This README, [`docs/usage.md`](docs/usage.md), [`docs/languages.md`](docs/languages.md), [`docs/comparison.md`](docs/comparison.md) and the package documentation (`go doc github.com/polestorm/gohighlightgo/highlight`).
- [`docs/definitions.md`](docs/definitions.md): the format of language definitions.
- In Polish: [`docs/provenance.md`](docs/provenance.md) (sources and clean-room rules),
  [`docs/testing.md`](docs/testing.md) (tests, measurements, comparison), [`requirements/`](requirements/)
  (numbered requirements, one file per language), [`docs/adr/`](docs/adr/) (decisions),
  [`docs/roadmap.md`](docs/roadmap.md).
- `go run ./tools/review` builds one HTML page with every corpus file in both themes next to the output of
  highlight, for review by a person.

## License

Apache License 2.0, see `LICENSE` and `NOTICE`. gohighlightgo contains no code by other authors.
