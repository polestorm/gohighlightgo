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
- **Linear time on hostile input**: one pass over the bytes, no regular expressions; 142 pathological cases are
  timed in the test suite, and the library is fuzzed.
- **Deterministic**: the same input gives the same output, byte for byte, documented and tested.
- **Made to be built on**: a single function, a stylesheet per theme, a command, and a few lines to connect it to
  a Markdown renderer (for example [gomdgo](https://github.com/polestorm/gomdgo)) behind a build tag.

Where it is weaker: two themes instead of dozens, many languages described more shallowly than in the
older tools (keywords, comments, strings and numbers; few context rules), no detection of the language from the content, and lexical highlighting only (a conversion `T(x)` looks like a call). It is new.

## Contents

[Languages](#languages) · [Usage](#usage) · [Comparison](#comparison) · [Status](#status) ·
[Documentation](#documentation) · [How the comparison was measured](#how-the-comparison-was-measured) ·
[License](#license)

## Languages

| Language | Names |
|---|---|
| Go | `go`, `golang` |
| C, C++ | `c`, `h`; `cpp`, `c++`, `cxx`, `cc`, `hpp`, `arduino`, `ino` |
| Java | `java` |
| JavaScript, TypeScript | `javascript`, `js`, `mjs`, `cjs`, `jsx`; `typescript`, `ts`, `mts`, `cts`, `tsx` |
| C# | `csharp`, `cs`, `c#` |
| Kotlin | `kotlin`, `kt`, `kts` |
| Swift | `swift` |
| Scala | `scala`, `sc` |
| Dart | `dart` |
| Python | `python`, `py`, `python3` |
| Rust | `rust`, `rs` |
| Lua | `lua` |
| R | `r` |
| PHP | `php` (with HTML outside `<?php … ?>`) |
| Perl | `perl`, `pl`, `pm` |
| Objective-C | `objectivec`, `objc`, `objective-c`, `m`, `mm` |
| Elixir, Erlang | `elixir`, `ex`, `exs`; `erlang`, `erl` |
| Clojure | `clojure`, `clj`, `cljs`, `cljc`, `edn` |
| F#, OCaml | `fsharp`, `fs`, `f#`; `ocaml`, `ml`, `mli` |
| Julia | `julia`, `jl` |
| MATLAB | `matlab`, `octave` |
| Groovy | `groovy`, `gradle`, `gvy` |
| PowerShell | `powershell`, `ps1`, `pwsh`, `posh` |
| Visual Basic | `vb`, `vbnet`, `visualbasic`, `vba`, `vbs`, `vbscript` |
| Fortran | `fortran`, `f90`, `f95`, `f03`, `f08`, `for` |
| Makefile, CMake | `make`, `makefile`, `mk`, `mak`; `cmake` |
| Dockerfile | `dockerfile`, `docker`, `containerfile` |
| INI, properties | `ini`, `cfg`, `conf`, `desktop`, `gitconfig`, `editorconfig`, `systemd`, `reg`; `properties`, `props` |
| Protocol Buffers, GraphQL | `protobuf`, `proto`, `proto3`; `graphql`, `gql` |
| HCL (Terraform, Packer), Nix | `hcl`, `terraform`, `tf`, `tfvars`, `packer`; `nix` |
| LaTeX | `latex`, `tex`, `sty`, `cls` |
| Vim script | `vim`, `vimscript`, `viml` |
| AsciiDoc | `asciidoc`, `adoc`, `asc` (source blocks highlighted in their own language) |
| Elm, Nim, Crystal, D | `elm`; `nim`, `nims`, `nimble`; `crystal`, `cr`; `d`, `dlang` |
| Ada, Prolog | `ada`, `adb`, `ads`; `prolog` |
| Scheme, Racket, Common Lisp | `scheme`, `scm`, `ss`; `racket`, `rkt`; `commonlisp`, `lisp`, `cl`, `el`, `elisp`, `emacs-lisp` |
| Solidity | `solidity`, `sol` |
| Verilog, VHDL | `verilog`, `systemverilog`, `sv`, `svh`, `vh`; `vhdl`, `vhd` |
| GLSL | `glsl`, `vert`, `frag` |
| Tcl, AWK | `tcl`, `tk`; `awk`, `gawk`, `mawk`, `nawk` |
| Assembly | `asm`, `nasm`, `x86asm`, `gas`, `s`; `armasm`, `arm`, `aarch64`, `arm64` |
| COBOL | `cobol`, `cbl`, `cob`, `cpy` (also the fixed format) |
| Smalltalk | `smalltalk`, `st`, `squeak`, `pharo` |
| Modula-2, Oberon | `modula2`, `modula-2`, `mod`, `def`; `oberon`, `ob`, `obn` |
| MoonScript | `moonscript`, `moon` |
| fish, csh | `fish`; `csh`, `tcsh` |
| Windows batch | `batch`, `bat`, `cmd` |
| XQuery, SPARQL, Cypher | `xquery`, `xq`, `xqy`; `sparql`, `rq`; `cypher`, `cql` |
| PL/SQL, T-SQL | `plsql`, `pls`; `tsql`, `mssql`, `t-sql` |
| Cython, Mojo | `cython`, `pyx`, `pxd`; `mojo` |
| V, Odin, Hare, Gleam | `v`, `vlang`; `odin`; `hare`, `ha`; `gleam` |
| Templates | `jinja` (also `django`, `twig`, `nunjucks`); `handlebars`, `mustache`; `liquid`; `erb`; `gotemplate`, `helm` (HTML outside the tags) |
| CSV, TSV | `csv`; `tsv`, `tab` |
| JSON5, Hjson, KDL, RON | `json5`; `hjson`; `kdl`; `ron` |
| Jsonnet, Starlark | `jsonnet`, `libsonnet`; `starlark`, `bzl`, `bazel`, `star` |
| Bicep, Rego, Puppet | `bicep`; `rego`; `puppet`, `pp` |
| Org, reStructuredText, Textile | `org`, `orgmode`; `rst`, `restructuredtext`, `rest`; `textile` (code blocks in their own language) |
| Dhall, Haxe, ActionScript | `dhall`; `haxe`, `hx`; `actionscript`, `as3` |
| Delphi forms | `dfm`, `lfm`, `fmx` |
| Luau, Teal | `luau`; `teal`, `tl` |
| F*, Lean, Agda, Idris, Coq | `fstar`, `fst`; `lean`, `lean4`; `agda`; `idris`, `idr`; `coq`, `rocq` |
| PureScript, ReScript, Reason | `purescript`, `purs`; `rescript`, `res`; `reason`, `re`, `reasonml` |
| Ruby | `ruby`, `rb` |
| Haskell | `haskell`, `hs` |
| Zig | `zig` |
| SQL | `sql` |
| JSON | `json`, `jsonc` |
| Pascal, Delphi | `pascal`, `delphi`, `pas`, `objectpascal` |
| XML, HTML | `xml`, `xsd`, `xsl`, `xslt`, `svg`, `plist`; `html`, `htm`, `xhtml`, `vue`, `svelte` (with JavaScript in `<script>` and CSS in `<style>`) |
| CSS | `css`, `scss`, `less` |
| YAML, TOML | `yaml`, `yml`, `ansible`; `toml` |
| Shell | `sh`, `bash`, `shell`, `zsh`, `ksh` |
| diff | `diff`, `patch` |
| Markdown | `markdown`, `md` (fenced code highlighted in its own language) |
| Apex, Hack, Nemerle | `apex`; `hack`, `hh`; `nemerle` |
| AutoHotkey, AutoIt | `autohotkey`, `ahk`; `autoit`, `au3` |
| Ballerina, Chapel, Pony, Vala | `ballerina`, `bal`; `chapel`, `chpl`; `pony`; `vala` |
| Clean, Roc, Mercury, Logtalk | `clean`; `roc`; `mercury`; `logtalk`, `lgt` |
| CoffeeScript, LiveScript, QML | `coffeescript`, `coffee`; `livescript`, `ls`; `qml` |
| Factor, Forth, J, Janet, Red | `factor`; `forth`, `fth`; `j`; `janet`; `red`, `rebol` |
| OpenSCAD, PostScript, Pike | `openscad`, `scad`; `postscript`, `ps`, `eps`; `pike` |
| Raku, REXX, Wolfram Language | `raku`, `perl6`; `rexx`; `wolfram`, `mathematica`, `wl` |
| SAS, Stata, Zeek | `sas`; `stata`; `zeek`, `bro` |
| ABAP, ANTLR, AppleScript | `abap`; `antlr`, `g4`; `applescript` |
| BibTeX, Brainfuck, Cap'n Proto | `bibtex`, `bib`; `brainfuck`, `bf`; `capnp` |
| Ceylon, Eiffel, Io, Modelica | `ceylon`; `eiffel`, `e`; `io`; `modelica`, `mo` |
| Fennel, Hy | `fennel`, `fnl`; `hy` |
| GDScript, Gherkin, gnuplot | `gdscript`, `gd`; `gherkin`, `cucumber`, `feature`; `gnuplot`, `gp` |
| LLVM IR, WGSL, Smali | `llvm`, `ll`; `wgsl`; `smali` |
| nginx, NSIS, RPM spec, DNS zone files | `nginx`; `nsis`, `nsi`; `rpmspec`, `spec`; `dnszone`, `zone` |
| PL/pgSQL, PromQL | `plpgsql`, `postgresql`, `psql`; `promql` |
| Standard ML, Thrift, Turtle, YANG | `sml`; `thrift`; `turtle`, `ttl`; `yang` |
| SourcePawn, POV-Ray, HTTP | `sourcepawn`, `sp`; `povray`, `pov`; `http` |
| Apache, Caddyfile, lighttpd | `apache`, `htaccess`; `caddyfile`; `lighttpd` |
| BNF, EBNF | `bnf`, `abnf`; `ebnf` |
| Berry, Boo, ChaiScript, Dylan, Wren, Squirrel, AngelScript | `berry`; `boo`; `chaiscript`; `dylan`; `wren`; `squirrel`; `angelscript` |
| FreeBASIC, QBasic | `freebasic`, `fb`; `qbasic`, `qb` |
| Isabelle, Whiley, Koka, Carbon | `isabelle`, `thy`; `whiley`; `koka`; `carbon` |
| JSONata, Kusto, DAX, Power Query M, PRQL, Pig Latin | `jsonata`; `kusto`, `kql`; `dax`; `powerquery`, `pq`; `prql`; `pig` |
| MiniZinc, Scilab, q (kdb+), APL, BQN | `minizinc`, `mzn`; `scilab`; `q`, `kdb`; `apl`; `bqn` |
| MLIR, TableGen, Z80 assembly, OpenEdge ABL | `mlir`; `tablegen`, `td`; `z80`; `abl`, `progress` |
| pkg-config, just, Meson, Ninja, GN, Kconfig | `pkgconfig`, `pc`; `justfile`, `just`; `meson`; `ninja`; `gn`; `kconfig` |
| Devicetree, GNU ld scripts, Sieve, SNOBOL4 | `dts`, `dtsi`; `ld`, `lds`; `sieve`; `snobol` |
| TypoScript, Pine Script, Logo, ink, EJS | `typoscript`; `pinescript`, `pine`; `logo`; `ink`; `ejs` |
| HLSL, CUDA, OpenCL C, Metal | `hlsl`, `fx`; `cuda`, `cu`; `opencl`, `ocl`; `metal`, `msl` |
| Graphviz DOT, PlantUML, Mermaid | `dot`, `graphviz`, `gv`; `plantuml`, `puml`; `mermaid`, `mmd` |
| ASN.1, CUE, Pkl, Nickel, Smithy | `asn1`; `cue`; `pkl`; `nickel`, `ncl`; `smithy` |
| Typst, roff (man pages), jq | `typst`, `typ`; `groff`, `roff`, `man`; `jq` |
| Varnish VCL, HAProxy | `vcl`, `varnish`; `haproxy` |
| Nushell, Elvish | `nushell`, `nu`; `elvish`, `elv` |
| Maxima, AMPL, newLISP, MoonBit | `maxima`; `ampl`; `newlisp`; `moonbit`, `mbt` |
| Cairo, Move, Vyper, Clarity | `cairo`; `move`; `vyper`, `vy`; `clarity` |
| G-code, IEC 61131-3 Structured Text | `gcode`, `nc`; `structuredtext`, `iecst`, `scl` |
| OpenQASM, Q# | `openqasm`, `qasm`; `qsharp`, `qs` |
| PureBasic, Xojo | `purebasic`, `pb`; `xojo` |
| 6502, 68000 and AVR assembly | `asm6502`, `6502`, `ca65`; `m68k`; `avrasm`, `avr` |
| Inno Setup, LilyPond, Maple, MEL | `innosetup`, `iss`; `lilypond`, `ly`; `maple`; `mel` |
| Icon, PL/I, TTCN-3, Oz, Frink | `icon`, `unicon`; `pli`, `pl1`; `ttcn3`; `oz`; `frink` |
| Gambas, BlitzBasic, Clipper / xBase | `gambas`; `blitzbasic`, `blitzmax`; `clipper`, `harbour`, `foxpro` |
| CFML, CLIPS, Web IDL | `cfml`, `cfscript`; `clips`; `webidl`, `idl` |
| Cassandra CQL, ArangoDB AQL | `cassandra`; `aql` |
| mcfunction, sed, Xorg, ISC dhcpd, CFEngine 3 | `mcfunction`; `sed`; `xorg`; `dhcpd`; `cfengine3`, `cf3` |
| Turing, ALGOL 60/68, BCPL, Igor Pro | `turing`; `algol`, `algol68`; `bcpl`; `igor`, `ipf` |
| UnrealScript, Papyrus, LSL | `unrealscript`, `uc`; `papyrus`, `psc`; `lsl` |

Names are matched without regard to case. Each language is described from its specification; what is
highlighted, and where the highlighting is only an approximation, is listed in
[requirements/languages/](requirements/languages/). 133 of them are definitions in
[highlight/lang/](highlight/lang/); markup, CSS, YAML, TOML, shell, diff, Markdown, INI, properties, AsciiDoc,
CSV, TSV, Org, reStructuredText, Textile and the template languages have scanners in Go.

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
| Languages | 298 🥇 | 297 | 263 |
| New language without a new release | a text file (`Register`, `-def`) | an XML file (`NewXMLLexer`, `Register`) | a Lua file |
| Themes | 2, and `auto` | 74 | 97 🥇 |
| Dependencies[^deps] | 0 🥇 | 1 | a Lua interpreter |
| Lines of code[^size] | 5,397 and 3,753 lines of definitions 🥇 | 31,199 and 279 lexer files (2.4 MB) | — |
| Added to a program | 536 kB 🥇 | 4,774 kB | 3,416 kB and 2.7 MB of definitions and themes |
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
