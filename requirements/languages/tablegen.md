# Wymagania: TableGen

Źródło: LLVM TableGen Programmer's Reference. Definicja: `highlight/lang/tablegen.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2081 | Nazwy języka: `tablegen`, `td`. |
| LANG-2082 | Słowa z listy keyword (`class`, `def`, `defm`, `defset`, `defvar`, `multiclass` …) dostają `hl-kw`; słowa z listy type (`bit`, `bits`, `int`, `string`, `list`, `code` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`; nazwa po `!` dostaje `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-2083 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `[{…}]` wielowierszowy dostają `hl-str`. |
| LANG-2084 | Liczby dostają `hl-num`; znaki `=:<>,;#.?` dostają `hl-op`. |
