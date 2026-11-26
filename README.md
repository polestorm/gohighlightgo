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
  timed in the test suite, and the library is fuzzed. On 64 kB of hostile input chroma hangs or needs more than
  15 s in 101 of 2,200 cases and highlight in 8 of 1,474; gohighlightgo in none of 3,278.
- **Deterministic**: the same input gives the same output, byte for byte, documented and tested.
- **Made to be built on**: a single function, a stylesheet per theme, a command, and a few lines to connect it to
  a Markdown renderer (for example [gomdgo](https://github.com/polestorm/gomdgo)) behind a build tag.

Where it is weaker: two themes instead of dozens, many languages described more shallowly than in the
older tools (keywords, comments, strings and numbers; few context rules), no detection of the language from the content, and lexical highlighting only (a conversion `T(x)` looks like a call). It is new.

## Contents

- [Why gohighlightgo](#why-gohighlightgo) · [Comparison](#comparison) · [Status](#status) ·
  [Documentation](#documentation) · [How the comparison was measured](#how-the-comparison-was-measured) ·
  [License](#license)
- [Usage](docs/usage.md): the API, classes and stylesheets, command line, connecting to a Markdown renderer
- [Languages](docs/languages.md): the 298 languages and the names each one accepts

## Comparison

| | gohighlightgo 0.1.0 | chroma 2.27.0 | highlight (QLMarkdown 1.5.7)[^hl] |
|---|:---:|:---:|:---:|
| Languages | 298 🥇 | 297 | 263 |
| New language without a new release | a text file (`Register`, `-def`) | an XML file (`NewXMLLexer`, `Register`) | a Lua file |
| Themes | 2, and `auto` | 74 | 97 🥇 |
| Dependencies[^deps] | 0 🥇 | 1 | a Lua interpreter |
| Added to a program[^size] | 536 kB 🥇 | 4,774 kB | 3,416 kB and 2.7 MB of definitions and themes |
| Time, the corpus as 21 blocks (14.5 kB)[^bench] | 0.16 ms 🥇 | 13.3 ms | — |
| Time, 1 MB of Go | 10.0 ms 🥇 | 690 ms | — |
| Time, a document with 1.3 MB of code[^doc] | 17 ms 🥇 | — | 2,079 ms |
| Memory allocated, the corpus as 21 blocks[^mem] | 0.14 MB 🥇 | 17.5 MB | — |
| Memory allocated, 1 MB of Go[^mem] | 4.2 MB in 45 allocations 🥇 | 820 MB in 12 million | — |
| Peak memory of the process, 1 MB of Go[^mem] | 22 MB 🥇 | 93 MB | — |
| Output for the corpus | 65.9 kB, 1,817 spans 🥇 | 117.4 kB, 4,317 spans | — |
| HTML in the code is escaped[^safe] | yes | yes | yes |
| Guarantees hold for added languages | yes, by construction 🥇 | not documented | not documented |
| Elements span line breaks | never 🥇 | yes (whitespace) | — |
| Hostile input, 64 kB: hang or over 15 s[^hostile] | 0 of 3,278 🥇 | 101 of 2,200 (59 languages) | 8 of 1,474 (8 languages) |
| Hostile input: input ×16 costs time ×96 or more | 0 🥇 | 11 (7 languages) | 15 (6 languages) |
| Hostile input: linear, but over 1 s for 64 kB | 0 🥇 | 119 (38 languages) | 1 |
| Hangs on ordinary non-ASCII text | no 🥇 | yes, JSONata (`zażółć`, `日本`) | no |
| Crashes on input | no 🥇 | yes, YANG (`/*` repeated) | no |
| Removing the tags gives back the input[^text] | yes, 298 languages 🥇 | not in 15 of 200 (a lone CR, for example C, C++, JavaScript, Rust; also tabs and the last line in diff, make, just) | — |
| Deterministic output | yes, documented and tested | yes, in the same test | — |
| Language detection from the content | no | yes | — |
| Command line | yes | yes | yes |
| License | Apache 2.0 | MIT | GPL 3.0 |

[^hl]: highlight is measured inside QLMarkdown, the Markdown preview for macOS that bundles it as a library;
    it is only run as a program. gohighlightgo is written from the language specifications, not from the code or
    the language definitions of other highlighters ([docs/provenance.md](docs/provenance.md)).
[^deps]: Modules compiled into a program. chroma requires two more modules for its tests.
[^size]: The size of that program minus the same program without highlighting (`-trimpath -ldflags="-s
    -w"`, macOS arm64); for highlight, the arm64 part of its library in QLMarkdown, which also holds the Lua
    interpreter, and its language definitions and themes.
[^bench]: The 21 files of this project's corpus one by one, in the 19 languages it had on 2026-11-25 before
    the definition files, and the Go files of the corpus joined to 1 MB. Median of five runs of `BenchmarkRivals`.
    chroma runs its lexer and its HTML formatter with classes, without `<pre>`, which gives output of the same
    form. Measured on 2026-11-25 on one Apple-silicon laptop. 🥇 marks the best value in a row.
[^doc]: The corpus as a Markdown document with one fenced block per file, joined to 1.3 MB, converted to an
    HTML page with and without highlighting; the difference is the cost of highlighting. gohighlightgo: the
    `gomdgo-mermaid` command built with and without highlighting; highlight: QLMarkdown's command with
    `--syntax-highlight on` and `off`. Median of 15 runs of each process.
[^hostile]: `go run ./cmd/rivalprobe` in `tools/compare`, 2026-11-25, chroma 2.27.0 and QLMarkdown 1.5.7 on one
    Apple-silicon laptop, four processes at a time. Eleven hostile patterns (unclosed `"`, `'`, `` ` ``, `/*`, `(*`,
    `<!--`, `<`, `${`, `#`, `\`, `[[`, each followed by `a `), repeated to about 4 kB and 64 kB, in every language
    the library knows: 298 for gohighlightgo, 200 for chroma, 134 for highlight. Every call of chroma and highlight
    runs in its own process, which is killed after 15 s; "hang or over 15 s" does not tell a hang from a very slow
    run. highlight is timed as the whole QLMarkdown command on a Markdown document with one code block, so its times
    include starting the program; that makes a slow-down harder to detect, not easier. The inputs are written to be
    hostile: ordinary code is highlighted by all three without trouble, and the difference matters for untrusted
    input such as previews and pasted code.
[^text]: The corpus and seven awkward inputs (CRLF, a lone CR, tabs, no final newline, invalid UTF-8, NUL, non-ASCII
    text) in each language; for invalid UTF-8 and NUL a U+FFFD in their place counts as kept, for both libraries.
    chroma runs without its default conversion of line ends to LF (`EnsureLF: false`). highlight's output passes
    through Markdown, so it is not compared.
[^mem]: "Allocated" is the sum of all allocations during one run (`B/op` of `BenchmarkRivals`); most of it is
    garbage collected during the run, so it measures the work for the garbage collector, not the memory in use.
    "Peak" is the maximum resident size of a test process that highlights 1 MB of Go once (`/usr/bin/time -l`);
    the same process without highlighting takes 18 MB. Measured on 2026-11-25.
[^safe]: Output for `</code></pre><script>alert(1)</script>` and an `<img onerror>` in each of the 200
    languages that chroma also knows (`TestRivalsSafety`); for highlight, the same text in Go, HTML, XML, shell and Markdown blocks.

## Status

Version 0.1.0: the planned languages, themes and command are done; hardening in progress. The API may still
change before 1.0. Requires Go 1.22 or newer.

- 298 languages, each with a requirements file and a corpus file whose output is checked byte for byte.
- Language definitions: the 11 table languages moved to definition files without a single byte of output
  changing; any definition the parser accepts is fuzzed together with random code.
- Safety and round trip (removing the tags gives back the input) are checked on every test input and in fuzzing.
- Fuzzing: short runs during development (about 13 minutes in total, about 18 million inputs, each in every
  language supported at the time) without a failure after one fix; the 24-hour run is still to be done.
- Not yet done: highlighting in the gomdgo command and in the gopeekgo preview (the adapters are ready).

## Documentation

- This README, [`docs/usage.md`](docs/usage.md), [`docs/languages.md`](docs/languages.md) and the package documentation (`go doc github.com/polestorm/gohighlightgo/highlight`).
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

The hostile-input rows come from `tools/compare/cmd/rivalprobe`, which runs every call of chroma and highlight
in a separate process with a time limit, so a hang is recorded and the measurement goes on. It writes a progress
page and a JSON summary; the numbers and the method are in `docs/testing.md`.

The corpus is small and written to cover each language's specification, not to imitate typical code; the 1 MB
and 1.3 MB inputs are the corpus repeated. The three highlighters split code into tokens differently (chroma
marks whitespace and punctuation, gohighlightgo does not), so the output sizes compare the work each one does,
not identical results.

## License

Apache License 2.0, see `LICENSE` and `NOTICE`. gohighlightgo contains no code by other authors.
