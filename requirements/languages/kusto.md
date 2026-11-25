# Wymagania: Kusto Query Language

Źródło: Kusto Query Language (KQL) reference (Microsoft). Definicja: `highlight/lang/kusto.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2121 | Nazwy języka: `kusto`, `kql`. |
| LANG-2122 | Słowa z listy keyword (`where`, `project`, `extend`, `summarize`, `by`, `order` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `now`, `count`, `dcount`, `sum` …) dostają `hl-builtin`. |
| LANG-2123 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2124 | Liczby dostają `hl-num`; znaki `|=<>!~+-*/%.,()[]` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
