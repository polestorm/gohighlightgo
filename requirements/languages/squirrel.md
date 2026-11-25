# Wymagania: Squirrel

Źródło: Squirrel 3.2 Reference Manual: The language. Definicja: `highlight/lang/squirrel.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2241 | Nazwy języka: `squirrel`, `nut`. |
| LANG-2242 | Słowa z listy keyword (`base`, `break`, `case`, `catch`, `class`, `clone` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `print`, `array`) dostają `hl-builtin`. |
| LANG-2243 | Komentarze `//` do końca wiersza, `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2244 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
