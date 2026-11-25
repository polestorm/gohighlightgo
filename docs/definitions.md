# Language definitions

A language whose tokens are words, comments, strings, numbers and operators is described by a short text file.
The built-in ones are in [`highlight/lang/`](../highlight/lang/); a program adds its own with
`highlight.Register(definition)`, and the command with `-def file.hl`.

A definition cannot contain regular expressions or code, so every language it describes keeps the library's
guarantees: the output is escaped, its elements come from the fixed list of classes, and the time is linear in
the length of the input. Constructs that a table cannot describe are written in Go and switched on by name
(`hook`).

## Example

```
# Go: The Go Programming Language Specification (Go 1.22)
name     go golang

keyword  break case chan const continue default defer else fallthrough for func go goto
keyword  if import interface map package range return select struct switch type var
type     any bool byte comparable error float32 float64 int int8 int16 int32 int64 rune string uint uintptr
builtin  true false iota nil append cap clear close copy delete len make max min new panic print println

line-comment   //
block-comment  /* */
string   " "  : escapes
string   ' '  : escapes
string   ` `  : multiline
escape        : x=2 octal=3 u4 U8
operators  +-*/%&|^<>=!:~
calls
```

## Syntax

One directive per line: a name, values separated by spaces, and optionally options after a lone `:` (the last
one on the line, so `:` can also be a value: `sigil : : attr`). Lines that
start with `#` and empty lines are ignored. Values are printable ASCII without spaces, except words, which may
contain any letters.

| Directive | Values | Options | Meaning |
|---|---|---|---|
| `name` | names | | the primary name, then aliases; lower case; may be repeated |
| `extends` | a language | | starts from a copy of an earlier table language; must be the first directive |
| `keyword`, `type`, `builtin` | words | | words coloured `hl-kw`, `hl-type`, `hl-builtin`; at most 32 bytes each; may be repeated |
| `fold-case` | | | words match without regard to case (SQL, Pascal) |
| `ident-extra` | characters | | punctuation allowed in identifiers besides letters, digits and `_` (for example `$`) |
| `ident-part` | characters | | punctuation allowed inside identifiers but not at their start (`'` in Haskell: `foldl'`) |
| `upper-type` | | | a name that starts with an upper-case letter and is not a listed word is coloured `hl-type` |
| `sigil` | characters | a class: `attr`, `meta`, `builtin`, … | a character (once or twice) before a name gives the name that class: `sigil $ : attr` for PHP variables |
| `line-string` | prefixes | | a string to the end of the line (`\\` in Zig) |
| `line-comment` | prefixes | | a comment to the end of the line |
| `block-comment` | open close | `nested` | a comment between two delimiters; an unclosed one runs to the end |
| `string` | open close | `escapes`, `multiline`, `doubled`, `name`, `interp=dollar`, `interp=dollar-braces`, `interp=braces`, `interp=hash` | see below |
| `prefix` | prefixes | `raw`, `literal`, `verbatim`, `interp=dollar`, `interp=braces`, `any-case` | see below |
| `escape` | | `continuation`, `x=N` or `x=any`, `octal=N`, `u4`, `U8`, `u-braces`, `java-u`, `named`, `paren` | escape sequences after `\` in strings with `escapes` |
| `number` | | `no-leading-dot`, `dot-ident`, `quote-sep` | differences in number syntax |
| `operators` | characters | | characters coloured `hl-op` |
| `calls` | | | a name followed by `(` is coloured `hl-fn` |
| `call-after-paren` | | | a name right after `(` is coloured `hl-fn` (Lisp: `(println x)`) |
| `macros` | | | a name followed by `!` (not `!=`) is coloured `hl-fn` |
| `decorators` | | | `@name` is coloured `hl-meta` (the same as `sigil @ : meta`) |
| `templates` | | | JavaScript template literals with `${…}` |
| `regex` | | | JavaScript regular expression literals, decided by the previous token |
| `keys` | | | a string followed by `:` is coloured `hl-attr` (JSON keys) |
| `hook` | a hook | | a construct written in Go: `c-preprocessor`, `cpp`, `rust`, `json-minus`, `pascal`, `lua`, `swift`, `apostrophe`, `haskell`, `ruby`, `php`, `perl`, `elixir`, `matlab`, `dotted`, `variables`, `latex`, `vim`, `hcl`, `lisp`, `verilog-number`, `tcl`, `asm-label`; may be repeated, tried in order |

**Strings.** `escapes`: `\` starts an escape sequence (shaped by `escape`); `multiline`: the string may cross line
ends, otherwise an unclosed string ends before the line break; `doubled`: the closing delimiter written twice is an
escape (`''` in SQL); `name`: a quoted name, not a string, so it gets no class (`"column"` in SQL);
`interp=dollar`: `$name` and `${…}` are coloured `hl-esc`; `interp=braces`: `{…}` is coloured `hl-esc` and `{{` is
an escape; `interp=hash`: `#{…}` is coloured `hl-esc` (Ruby); `interp=dollar-braces`: only `${…}` (Nix, HCL, CMake). The first matching delimiter wins, so longer ones (`"""`) go first.

**Prefixes** are glued to a string: `u8"…"`, `r"…"`, `@"…"`. `raw`: no escape sequences, but `\` still protects
the next character from closing the string (Python); `literal`: `\` is an ordinary character (Dart); `verbatim`:
like `literal`, the doubled delimiter is an escape and the string may span lines (C#); `any-case`: every
spelling of upper and lower case is accepted. `raw`, `literal` and `verbatim` also turn interpolation off,
unless the prefix gives its own `interp=` (`raw"…$x"` in Scala). With `interp=braces`, `{{` and `}}` are escapes. The longest matching prefix wins.

**Escapes.** `continuation`: `\` before a line end continues the string; `x=N`: `\x` and up to N hex digits;
`octal=N`: up to N octal digits; `u4`, `U8`: `\u` with four and `\U` with eight hex digits; `u-braces`: `\u{…}`;
`java-u`: `\u`, `\uu`, … and four hex digits; `named`: `\N{…}`; `paren`: `\( … )` (Swift interpolation);
`char=X`: the escape character is X instead of `\` (`` char=` `` in PowerShell). Any other escape character with
one character after it is an escape too. `interp=dollar` also takes `$(…)`.

**Numbers.** Decimal, `0x`, `0o`, `0b`, `_` between digits, fractions, exponents (`e`, or `p` in hex) and a
suffix of letters (`10u32`, `1.5f`) are built in. `no-leading-dot`: a number cannot start with `.`; `dot-ident`:
a dot before a name ends the number (`1.max(2)`); `quote-sep`: `'` separates digits (C++).

## Limits

A definition is at most 64 kB, a word at most 32 bytes and a prefix at most 8. The time per byte of input grows
with the number of rules (strings, comments, prefixes), so a definition with thousands of rules is slow, but still
linear. An invalid definition is rejected with a `*highlight.DefinitionError` that gives the line.
