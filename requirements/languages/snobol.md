# Wymagania: SNOBOL4

Źródło: The SNOBOL4 Programming Language (Griswold, Poage, Polonsky), 2nd edition. Definicja: `highlight/lang/snobol.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2071 | Nazwy języka: `snobol`, `snobol4`, `sno`. |
| LANG-2072 | Słowa z listy keyword (`END`, `DEFINE`, `RETURN`, `FRETURN`, `NRETURN`) dostają `hl-kw`; słowa z listy builtin (`OUTPUT`, `INPUT`, `SIZE`, `DUPL`, `TRIM`, `SPAN` …) dostają `hl-builtin`. |
| LANG-2073 | napisy `"…"`; `'…'` dostają `hl-str`. |
| LANG-2074 | Liczby dostają `hl-num`; znaki `=:|$.*()+-/` dostają `hl-op`; konstrukcje haka `cobol` jak w opisie haka w `docs/definitions.md`. |
