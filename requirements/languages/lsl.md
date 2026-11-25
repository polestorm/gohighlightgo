# Wymagania: Linden Scripting Language

Źródło: Second Life wiki: LSL Portal. Definicja: `highlight/lang/lsl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2971 | Nazwy języka: `lsl`. |
| LANG-2972 | Słowa z listy keyword (`default`, `state`, `if`, `else`, `for`, `while` …) dostają `hl-kw`; słowa z listy type (`integer`, `float`, `string`, `key`, `vector`, `rotation` …) dostają `hl-type`; słowa z listy builtin (`TRUE`, `FALSE`, `NULL_KEY`, `ZERO_VECTOR`, `PI`, `llSay` …) dostają `hl-builtin`. |
| LANG-2973 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2974 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=~<>@` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
