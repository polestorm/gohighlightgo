# Wymagania: PromQL

Źródło: Prometheus documentation: Querying basics, operators and functions. Definicja: `highlight/lang/promql.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1731 | Nazwy języka: `promql`. |
| LANG-1732 | Słowa z listy keyword (`by`, `without`, `on`, `ignoring`, `group_left`, `group_right` …) dostają `hl-kw`; słowa z listy builtin (`rate`, `irate`, `increase`, `delta`, `deriv`, `histogram_quantile` …) dostają `hl-builtin`; nazwa przed `=` albo `~` albo `!` dostaje `hl-attr`. |
| LANG-1733 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc`; `` `…` `` dostają `hl-str`. |
| LANG-1734 | Liczby dostają `hl-num`; znaki `+-*/%^=!<>~{}[]` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
