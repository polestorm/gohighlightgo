# How the comparison was measured

[← README](../README.md)

Every row of the comparison table can be checked by anyone, including the authors of the other highlighters.
The inputs are frozen files in [`tools/compare/inputs`](../tools/compare/inputs) (described in its
[README](../tools/compare/inputs/README.md)), the commands that measure them are in `tools/compare`, and the
results they printed are kept in [`tools/compare/results`](../tools/compare/results).

```sh
cd tools/compare
go run ./cmd/rivalbench                     # size, time, memory, output size (about 4 minutes)
go run ./cmd/rivalprobe -page /tmp/p.html   # hostile input, text preservation, determinism (about 15 minutes)
go test -run Rivals .                       # XSS-safe output, output form
```

chroma's version is fixed by `tools/compare/go.mod`. highlight needs QLMarkdown 1.5.7 in `/Applications` (macOS);
without it the highlight columns stay empty. `tools/compare/inputs/README.md` also shows how to reproduce single
findings by hand, for example chroma's crash on YANG input or highlight's hang on `${` in Dart, with the probe
programs in `tools/compare/probes`, which highlight standard input with one library and nothing else.

gohighlightgo and chroma are run through their public API. highlight is not available as a Go library, so it is
measured as a whole program, the way it is used in practice, and its share is the difference between runs with
and without highlighting.

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

How much bigger a program gets when it can highlight code. The programs in `tools/compare/probes` read code from
standard input and write HTML: `ours` with gohighlightgo, `chroma` with chroma, and `none` copies the input
unchanged. Each is built with `-trimpath -ldflags="-s -w"`, and the row gives the difference to `none`.
Everything gohighlightgo needs, all 298 languages included, is inside its 536 kB. highlight is not a Go library;
its row is the part of the library QLMarkdown loads for this processor (3,416 kB, with a Lua interpreter), plus
the language definition and theme files it reads from disk at run time (1,548 kB, the sum of the file sizes).
The first two have their definitions and themes inside the program, so the three numbers cover the same things.

### Time and memory on the corpus

The 21 files of `tools/compare/inputs/corpus` (19 languages, 14.5 kB), each highlighted by a separate call, like
the code blocks of one document, and the Go files among them repeated to 1 MB. `rivalbench` runs each case with
Go's `testing.Benchmark` five times and reports the median. chroma runs its lexer and its HTML formatter with
classes, without `<pre>`, which gives output of the same form.

### A document full of code

The 21 files as fenced code blocks of one Markdown document, repeated to 1.3 MB (1,848 blocks). For chroma and
gohighlightgo, the time to highlight every block, median of five benchmark runs. For highlight, the QLMarkdown
command on that document with `--syntax-highlight on` minus the same with `off`, median of 15 runs of each; this
includes what highlight adds around each block, but not the Markdown itself.

### Memory

"Allocated" is the sum of all allocations during one run (`AllocedBytesPerOp` of the benchmark); most of it is
garbage collected during the run, so it measures the work for the garbage collector, not the memory in use.
highlight is written in C++ and runs as a separate program, so its allocations cannot be counted this way.
"Peak memory added" is the maximum resident size (`getrusage`) of a probe program that highlights 1 MB of Go once,
minus the `none` program on the same input, median of 15 runs; for highlight, the QLMarkdown command on a
document with that code, with and without highlighting.

### Hostile input

`go run ./cmd/rivalprobe` in `tools/compare`, 2026-11-25 at commit `bc1f263`, chroma 2.27.0 and highlight 4.21,
Apple M4, four processes at a time. The inputs are the 22 files of `tools/compare/inputs/hostile`: eleven patterns
(unclosed `"`, `'`, `` ` ``, `/*`, `(*`, `<!--`, `<`, `${`, `#`, `\`, `[[`, each followed by `a `), repeated to
about 4 kB and 64 kB, given to every language each library knows: 298 for gohighlightgo, 200 for chroma, 134 for
highlight. Every call of chroma and highlight runs in its own process, which is killed after 15 s; "hang or over
15 s" does not tell a hang from a very slow run. highlight is timed as the whole QLMarkdown command on a Markdown
document with one code block, so its times include starting the program; that makes a slow-down harder to detect,
not easier. Besides hangs, the probe records two kinds of slow-down, which the README leaves out: input ×16
costing time ×96 or more (gohighlightgo 0, chroma 11 cases in 7 languages, highlight 15 in 6), and linear time but
over 1 s for 64 kB (gohighlightgo 0, chroma 119 in 38 languages, highlight 1). Each finding is listed in
[`tools/compare/results/rivalprobe-2026-11-25.json`](../tools/compare/results/rivalprobe-2026-11-25.json). The
inputs are written to be hostile: ordinary code is highlighted by all three without trouble, and the difference
matters for untrusted input such as previews and pasted code.

### Round-trip fidelity

The output with its tags removed and its escapes undone must equal the input, byte for byte. Inputs: the library's
test corpus (one or more files in each of the 298 languages, at commit `bc1f263`) and the seven files of
`tools/compare/inputs/text` (CRLF, a lone CR, tabs, no final newline, invalid UTF-8, NUL, non-ASCII text) in each
language; for invalid UTF-8 and NUL a U+FFFD in their place counts as kept, for both libraries. chroma runs
without its default conversion of line ends to LF (`EnsureLF: false`), so the check is about its lexers, not that
option. Where it changes the text: a lone CR is lost or moved (for example C, C++, JavaScript, Rust and 11 more),
and in diff, make and just also tabs and the last line. highlight's output passes through Markdown, so it is not
compared.

### XSS-safe output

Output for `tools/compare/inputs/xss/attack.txt` (`</code></pre><script>alert(1)</script>`, an `<img onerror>`,
`&` and a comment) in each of the 200 languages that chroma also knows (`TestRivalsSafety`); for highlight, the
same text in Go, HTML, XML, shell and Markdown blocks, checked by hand.
