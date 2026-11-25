# Wymagania: Dylan

Źródło: Dylan Reference Manual (opendylan.org). Definicja: `highlight/lang/dylan.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1921 | Nazwy języka: `dylan`. |
| LANG-1922 | Słowa z listy keyword (`define`, `end`, `method`, `class`, `function`, `generic` …) dostają `hl-kw`; słowa z listy builtin (`#t`, `#f`, `format-out`, `print`) dostają `hl-builtin`. |
| LANG-1923 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1924 | Liczby dostają `hl-num`; znaki `+-*/:=<>~&|` dostają `hl-op`. |
