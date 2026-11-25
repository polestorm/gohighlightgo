# Wymagania: ChaiScript

Źródło: ChaiScript language reference (chaiscript.com). Definicja: `highlight/lang/chaiscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1911 | Nazwy języka: `chaiscript`, `chai`. |
| LANG-1912 | Słowa z listy keyword (`def`, `var`, `auto`, `class`, `attr`, `if` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `print`, `puts`, `size`) dostają `hl-builtin`. |
| LANG-1913 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1914 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
