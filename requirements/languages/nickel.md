# Wymagania: Nickel

Źródło: Nickel user manual (nickel-lang.org). Definicja: `highlight/lang/nickel.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2431 | Nazwy języka: `nickel`, `ncl`. |
| LANG-2432 | Słowa z listy keyword (`let`, `rec`, `in`, `fun`, `if`, `then` …) dostają `hl-kw`; słowa z listy type (`Number`, `String`, `Bool`, `Array`, `Dyn`) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `std`) dostają `hl-builtin`. |
| LANG-2433 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `m%"…"%` wielowierszowy, ze wstawkami `hl-esc`; `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc` dostają `hl-str`. |
| LANG-2434 | Liczby dostają `hl-num`; znaki `+-*/%&|!<>=:?.@` dostają `hl-op`. |
