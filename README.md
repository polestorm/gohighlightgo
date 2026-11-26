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

Compared with the established highlighters (details in [Comparison](#comparison)):

- **Fast**: about 100 MB/s, the fastest of the highlighters compared.
- **Small**: 5,397 lines of Go and 3,753 lines of language definitions, 536 kB added to a program, no
  dependencies, nothing loaded at run time.
- **Extensible without a release**: a language made of words, comments, strings and numbers is a short text
  file ([docs/definitions.md](docs/definitions.md)); a program adds its own with `highlight.Register`, the
  command with `-def`. Definitions cannot contain regular expressions or code, so every guarantee below holds for
  them too.
- **Safe by construction**: the output contains only escaped text and `<span>` elements with classes from a
  fixed list of 14, never nested and never spanning a line break. An independent checker verifies this in the
  tests and in fuzzing, and removing the tags gives back the input.
- **Linear time on hostile input**: one pass over the bytes, no regular expressions; 283 pathological cases are
  timed in the test suite, and the library is fuzzed. In the comparison it is the only one that never hangs or
  slows down on hostile input.
- **Deterministic**: the same input gives the same output, byte for byte, documented and tested.
- **Made to be built on**: a single function, a stylesheet per theme, a command, and a few lines to connect it to
  a Markdown renderer (for example [gomdgo](https://github.com/polestorm/gomdgo)) behind a build tag.

Where it is weaker: two themes instead of dozens, many languages described more shallowly than in the
older tools (keywords, comments, strings and numbers; few context rules), no detection of the language from the content, and lexical highlighting only (a conversion `T(x)` looks like a call). It is new.

## Contents

- [Why gohighlightgo](#why-gohighlightgo) · [Comparison](#comparison) ·
  [Documentation](#documentation) · [License](#license)
- [Usage](docs/usage.md): the API, classes and stylesheets, command line, connecting to a Markdown renderer
- [Languages](docs/languages.md): the 298 languages and the names each one accepts
- [How the comparison was measured](docs/comparison.md): the tests and tools behind the comparison table

## Comparison

| | gohighlightgo 0.1.0 | [chroma](https://github.com/alecthomas/chroma) 2.27.0 | [highlight](http://www.andre-simon.de/) 4.21[^hl] |
|---|:---:|:---:|:---:|
| Languages | 298 🥇 | 297 | 263 |
| New language without a new release | a text file (`Register`, `-def`) | an XML file (`NewXMLLexer`, `Register`) | a Lua file |
| Dependencies[^deps] | 0 🥇 | 1 | a Lua interpreter |
| Size added to a program[^size] | 536 kB 🥇 | 4 774 kB | 6 156 kB |
| Time, the corpus as 21 blocks (14.5 kB)[^bench] | 0.16 ms 🥇 | 13.3 ms | *not measured* |
| Time, 1 MB of Go | 10.0 ms 🥇 | 690 ms | *not measured* |
| Time, a document with 1.3 MB of code[^doc] | 17 ms 🥇 | *not measured* | 2 079 ms |
| Memory allocated, the corpus as 21 blocks[^mem] | 0.14 MB 🥇 | 17.5 MB | *not measured* |
| Memory allocated, 1 MB of Go[^mem] | 4.2 MB 🥇 | 820 MB | *not measured* |
| Peak memory added, 1 MB of Go[^mem] | 4 MB 🥇 | 75 MB | 11 MB |
| Output for the corpus | 65.9 kB 🥇 | 117.4 kB | *not measured* |
| XSS-safe output[^safe] | yes | yes | yes |
| Guarantees hold for added languages | yes, by construction 🥇 | not documented | not documented |
| Elements span line breaks | never 🥇 | yes (whitespace) | *not checked* |
| Hostile input, 64 kB: hang or over 15 s[^hostile] | 0 of 3 278 🥇 | 101 of 2 200 (59 languages) | 8 of 1 474 (8 languages) |
| Hostile input: input ×16 costs time ×96 or more | 0 🥇 | 11 (7 languages) | 15 (6 languages) |
| Hostile input: linear, but over 1 s for 64 kB | 0 🥇 | 119 (38 languages) | 1 |
| Hangs on ordinary non-ASCII text | no 🥇 | yes, JSONata (`zażółć`, `日本`) | no |
| Crashes on input | no 🥇 | yes, YANG (`/*` repeated) | no |
| Removing the tags gives back the input[^text] | yes, 298 languages 🥇 | not in 15 of 200 | *not measured* |
| Deterministic output | yes, documented and tested | yes, in the same test | *not checked* |
| Language detection from the content | no | yes | *not checked* |
| Themes | 2, and `auto` | 74 | 97 🥇 |
| Command line | yes | yes | yes |
| License | Apache 2.0 | MIT | GPL 3.0 |

How the numbers were measured: [docs/comparison.md](docs/comparison.md) and the notes below.

[^hl]: An independent C++ program by André Simon, measured through the copy that QLMarkdown 1.5.7 bundles.
    [More](docs/comparison.md#highlight)
[^deps]: Modules compiled into a program. chroma requires two more modules for its tests.
[^size]: A small program that turns standard input into HTML, built with and without highlighting; language
    definitions and themes are counted for all three. [More](docs/comparison.md#size-added-to-a-program)
[^bench]: Median of five runs of `BenchmarkRivals`, 2026-11-25, one Apple-silicon laptop. 🥇 marks the best value
    in a row. [More](docs/comparison.md#time-and-memory-on-the-corpus)
[^doc]: A Markdown document turned into HTML with and without highlighting; median of 15 runs.
    [More](docs/comparison.md#a-document-full-of-code)
[^hostile]: Eleven hostile patterns at about 4 kB and 64 kB in every language; each call of chroma and highlight
    in its own process, stopped after 15 s. [More](docs/comparison.md#hostile-input)
[^text]: The corpus and seven awkward inputs (CRLF, a lone CR, tabs, invalid UTF-8 and others) in each language.
    [More](docs/comparison.md#removing-the-tags-gives-back-the-input)
[^mem]: Allocated: the sum of all allocations in one run, mostly freed during it. Peak added: the largest memory
    use of a process minus the same process without highlighting. [More](docs/comparison.md#memory)
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
