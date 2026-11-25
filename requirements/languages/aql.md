# Wymagania: ArangoDB AQL

Źródło: ArangoDB documentation: AQL fundamentals. Definicja: `highlight/lang/aql.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2851 | Nazwy języka: `aql`, `arangodb`. |
| LANG-2852 | Słowa z listy keyword (`FOR`, `IN`, `RETURN`, `FILTER`, `SORT`, `LIMIT` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `null`, `LENGTH`, `COUNT`, `SUM` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-attr` (bez rozróżniania wielkości liter). |
| LANG-2853 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc`; `` `…` `` (nazwa, bez klasy) dostają `hl-str`. |
| LANG-2854 | Liczby dostają `hl-num`; znaki `=<>!+-*/%.,:?()[]{}` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
