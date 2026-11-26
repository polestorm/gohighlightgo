# Inputs of the comparison

Every number in the comparison table of the [README](../../../README.md) comes from the files in this directory
and the commands in `tools/compare`. The files are frozen: they do not change when the library's own test corpus
grows, so a measurement can be repeated later and by anyone, including the authors of the other highlighters.

| Directory | What it is | Used by |
|---|---|---|
| `corpus/` | 21 files in 19 languages (14.5 kB without their first lines), copied from `testdata/corpus` at commit `c5ac60f` | `rivalbench`: time, memory, output size, the 1 MB and 1.3 MB inputs |
| `text/` | seven awkward inputs: CRLF, a lone CR, tabs, no final newline, invalid UTF-8, NUL, non-ASCII text | `rivalprobe`, stage `text` |
| `hostile/` | eleven hostile patterns, each followed by `a ` and repeated 2,000 times (`-small`) and 32,000 times (`-large`) | `rivalprobe`, stages `*-time` |
| `xss/` | `</code></pre><script>…`, an `<img onerror>`, `&` and a comment | `TestRivalsSafety` |

Rules for the files:

- The first line of each file in `corpus/` refers to this project's requirements and is not highlighted; the rest
  is. Its directory name is the language.
- The same file in `text/` and `hostile/` is given to every language each highlighter knows.
- `go test ./cmd/rivalprobe` checks that the files are byte for byte the inputs the probe uses.

## Repeating the measurements

Requirements: Go 1.25 or newer. For highlight, [QLMarkdown](https://github.com/sbarex/QLMarkdown) 1.5.7, which
bundles highlight 4.21, installed in `/Applications` (macOS only). Without it, the highlight columns stay empty.

```sh
cd tools/compare
go run ./cmd/rivalbench                     # size, time, memory, output size: a Markdown table (about 4 minutes)
go run ./cmd/rivalprobe -page /tmp/p.html   # hostile input and text preservation, JSON on stdout (about 15 minutes)
go test -run Rivals .                       # XSS and output checks
```

chroma's version is fixed by `go.mod` (2.27.0). The results of 2026-11-25 and 2026-11-26 are in
[`../results/`](../results/).

## Checking one claim by hand

The probe programs in `tools/compare/probes` highlight standard input with one library and nothing else:

```sh
cd tools/compare
go build -o /tmp/chroma ./probes/chroma && go build -o /tmp/ours ./probes/ours

/tmp/chroma yang < inputs/hostile/block-comment-small.txt        # chroma: panic: unknown state comment
/tmp/ours   yang < inputs/hostile/block-comment-small.txt        # gohighlightgo: output at once

/tmp/chroma jsonata < inputs/text/non-ascii.txt                  # chroma: does not finish, memory keeps growing
                                                                 # (3.6 GB after a minute); stop it with Ctrl-C

time /tmp/chroma make < inputs/hostile/dquote-small.txt > /dev/null   # 4 kB: about 35 ms, with starting the process
time /tmp/chroma make < inputs/hostile/dquote-large.txt > /dev/null   # 64 kB (×16): about 2.7 s (×80)
time /tmp/ours   make < inputs/hostile/dquote-large.txt > /dev/null   # gohighlightgo: a few ms

printf 'a\rb\r' | /tmp/chroma c | od -c   # chroma: the CR becomes a line feed (its default; rivalprobe turns
                                         # that off with EnsureLF: false, and the CR is still lost or moved)
printf 'a\rb\r' | /tmp/ours c | od -c     # gohighlightgo: a \r b \r
```

For highlight, a code block in a Markdown file:

```sh
Q=/Applications/QLMarkdown.app/Contents/Resources/qlmarkdown_cli
{ echo '~~~dart'; cat inputs/hostile/dollar-brace-large.txt; echo; echo '~~~'; } > /tmp/dart.md
time $Q --syntax-highlight on /tmp/dart.md > /dev/null          # highlight: still running after 20 s
```

The full list of cases where a highlighter hangs, crashes, slows down or changes the text is in
[`../results/rivalprobe-2026-11-25.json`](../results/rivalprobe-2026-11-25.json), as `language "pattern"`;
the pattern names map to the files in `hostile/` in the order `"` dquote, `'` squote, `` ` `` backtick,
`/*` block-comment, `(*` paren-star, `<!--` html-comment, `<` lt, `${` dollar-brace, `#` hash, `\` backslash,
`[[` double-bracket.
