# Wymagania: PRQL

Źródło: PRQL language book (prql-lang.org). Definicja: `highlight/lang/prql.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2041 | Nazwy języka: `prql`. |
| LANG-2042 | Słowa z listy keyword (`from`, `select`, `filter`, `derive`, `group`, `aggregate` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `null`, `sum`, `count`, `average` …) dostają `hl-builtin`. |
| LANG-2043 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2044 | Liczby dostają `hl-num`; znaki `+-*/%=<>!&|~.:` dostają `hl-op`. |
