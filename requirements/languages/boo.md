# Wymagania: Boo

Źródło: Boo language guide (boo-lang). Definicja: `highlight/lang/boo.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1901 | Nazwy języka: `boo`. |
| LANG-1902 | Słowa z listy keyword (`import`, `namespace`, `class`, `interface`, `struct`, `enum` …) dostają `hl-kw`; słowa z listy type (`int`, `long`, `double`, `string`, `bool`, `object` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `print`, `len`, `range` …) dostają `hl-builtin`. |
| LANG-1903 | Komentarze `#` do końca wiersza, `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"""…"""` wielowierszowy; `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1904 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
