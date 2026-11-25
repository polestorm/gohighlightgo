# Wymagania: BCPL

Źródło: BCPL: the language and its compiler (Richards & Whitby-Strevens). Definicja: `highlight/lang/bcpl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2911 | Nazwy języka: `bcpl`. |
| LANG-2912 | Słowa z listy keyword (`GET`, `LET`, `AND`, `BE`, `VALOF`, `RESULTIS` …) dostają `hl-kw`; słowa z listy builtin (`TRUE`, `FALSE`, `writef`, `writes`, `newline`) dostają `hl-builtin`. |
| LANG-2913 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2914 | Liczby dostają `hl-num`; znaki `:=+-*/<>~&|!@(){}$` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
