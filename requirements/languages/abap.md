# Wymagania: ABAP

Źródło: ABAP Keyword Documentation (SAP, 7.5x). Definicja: `highlight/lang/abap.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1541 | Nazwy języka: `abap`. |
| LANG-1542 | Słowa z listy keyword (`REPORT`, `DATA`, `TYPES`, `TYPE`, `LIKE`, `BEGIN` …) dostają `hl-kw`; słowa z listy type (`i`, `f`, `p`, `c`, `n`, `d` …) dostają `hl-type`; słowa z listy builtin (`abap_true`, `abap_false`, `sy`, `me`, `space`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-1543 | Komentarze `"` do końca wiersza dostają `hl-com`; napisy `'…'` z podwojonym ogranicznikiem jako ucieczką; `` `…` `` z podwojonym ogranicznikiem jako ucieczką; `|…|` ze wstawkami `hl-esc` dostają `hl-str`. |
| LANG-1544 | Liczby dostają `hl-num`; znaki `+-*/=<>:.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`; konstrukcje haka `cobol` jak w opisie haka w `docs/definitions.md`. |
