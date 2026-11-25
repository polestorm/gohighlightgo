# Wymagania: Io

Źródło: Io Programming Guide (iolanguage.org). Definicja: `highlight/lang/io.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1671 | Nazwy języka: `io`. |
| LANG-1672 | Słowa z listy keyword (`method`, `block`, `if`, `then`, `else`, `while` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `nil`, `Object`, `Lobby`, `writeln` …) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-1673 | Komentarze `//` do końca wiersza, `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"""…"""` wielowierszowy; `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1674 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`. |
