# Wymagania: JSONata

Źródło: JSONata documentation (docs.jsonata.org). Definicja: `highlight/lang/jsonata.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1961 | Nazwy języka: `jsonata`. |
| LANG-1962 | Słowa z listy keyword (`function`, `and`, `or`, `in`) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `null`, `$sum`, `$count`, `$string` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-1963 | Komentarze `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1964 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.@#` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
