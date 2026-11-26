# How the comparison was measured

[← README](../README.md)

gohighlightgo and chroma are run through their public API on the same inputs by `TestRivalsOutput`,
`TestRivalsSafety` and `BenchmarkRivals` in `tools/compare` (a separate module, so the library itself has no
dependencies). highlight is not available as a Go library, so it is measured as a whole program, the way it is
used in practice, and its share is the difference between runs with and without highlighting.

The hostile-input rows come from `tools/compare/cmd/rivalprobe`, which runs every call of chroma and highlight
in a separate process with a time limit, so a hang is recorded and the measurement goes on. It writes a progress
page and a JSON summary; the numbers and the method are in [testing.md](testing.md) (in Polish).

The corpus is small and written to cover each language's specification, not to imitate typical code; the 1 MB
and 1.3 MB inputs are the corpus repeated. The three highlighters split code into tokens differently (chroma
marks whitespace and punctuation, gohighlightgo does not), so the output sizes compare the work each one does,
not identical results.

## Notes on the rows

### highlight

chroma is the usual highlighter in Go. highlight, by André Simon, is an independent C++ program; there is no
command-line build of it here, so it is measured through the copy that QLMarkdown 1.5.7 (a Markdown preview for
macOS) bundles as a library, running QLMarkdown's command. gohighlightgo is written from the language
specifications, not from the code or the language definitions of other highlighters
([provenance.md](provenance.md), in Polish).

### Size added to a program

How much bigger a program gets when it can highlight code. A small program that reads code from standard input and
writes HTML is built twice, with the library and without it (`-trimpath -ldflags="-s -w"`, macOS arm64), and the
row gives the difference: 536 kB for gohighlightgo, about 4.8 MB for chroma. Everything gohighlightgo needs, all
298 languages included, is inside those 536 kB. highlight is not a Go library; its row is the arm64 part of the
library QLMarkdown loads, which includes a Lua interpreter (3,416 kB), plus the language definition and theme
files it reads from disk at run time (2,740 kB). The first two have their definitions and themes inside the
program, so the three numbers cover the same things.

### Time and memory on the corpus

The 21 files of this project's corpus one by one, in the 19 languages it had on 2026-11-25 before the definition
files, and the Go files of the corpus joined to 1 MB. Median of five runs of `BenchmarkRivals`. chroma runs its
lexer and its HTML formatter with classes, without `<pre>`, which gives output of the same form. Measured on
2026-11-25 on one Apple-silicon laptop. 🥇 marks the best value in a row.

### A document full of code

The corpus as a Markdown document with one fenced block per file, joined to 1.3 MB, converted to an HTML page with
and without highlighting; the difference is the cost of highlighting. gohighlightgo: the `gomdgo-mermaid` command
built with and without highlighting; highlight: QLMarkdown's command with `--syntax-highlight on` and `off`.
Median of 15 runs of each process.

### Hostile input

`go run ./cmd/rivalprobe` in `tools/compare`, 2026-11-25, chroma 2.27.0 and highlight 4.21 on one Apple-silicon
laptop, four processes at a time. Eleven hostile patterns (unclosed `"`, `'`, `` ` ``, `/*`, `(*`, `<!--`, `<`,
`${`, `#`, `\`, `[[`, each followed by `a `), repeated to about 4 kB and 64 kB, in every language the library
knows: 298 for gohighlightgo, 200 for chroma, 134 for highlight. Every call of chroma and highlight runs in its
own process, which is killed after 15 s; "hang or over 15 s" does not tell a hang from a very slow run. highlight
is timed as the whole QLMarkdown command on a Markdown document with one code block, so its times include starting
the program; that makes a slow-down harder to detect, not easier. The inputs are written to be hostile: ordinary
code is highlighted by all three without trouble, and the difference matters for untrusted input such as previews
and pasted code.

### Removing the tags gives back the input

The corpus and seven awkward inputs (CRLF, a lone CR, tabs, no final newline, invalid UTF-8, NUL, non-ASCII text)
in each language; for invalid UTF-8 and NUL a U+FFFD in their place counts as kept, for both libraries. chroma
runs without its default conversion of line ends to LF (`EnsureLF: false`). Where it changes the text: a lone CR
is lost or moved (for example C, C++, JavaScript, Rust and 11 more), and in diff, make and just also tabs and the
last line. highlight's output passes through Markdown, so it is not compared.

### Memory

"Allocated" is the sum of all allocations during one run (`B/op` of `BenchmarkRivals`); most of it is garbage
collected during the run, so it measures the work for the garbage collector, not the memory in use. highlight is
written in C++ and runs as a separate program, so its allocations cannot be counted this way. "Peak memory added"
is the maximum resident size of a process that highlights 1 MB of Go once minus the same process without
highlighting (`/usr/bin/time -l`): for the two Go libraries a test process (22 MB and 93 MB against 18 MB), for
highlight the QLMarkdown command on a document with that code, with `--syntax-highlight on` and `off` (36 MB
against 25 MB). Measured on 2026-11-25 and 2026-11-26.

### XSS-safe output

Output for `</code></pre><script>alert(1)</script>` and an `<img onerror>` in each of the 200 languages that
chroma also knows (`TestRivalsSafety`); for highlight, the same text in Go, HTML, XML, shell and Markdown blocks.
