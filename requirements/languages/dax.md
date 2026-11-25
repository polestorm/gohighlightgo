# Wymagania: DAX

Źródło: DAX function reference and DAX syntax (Microsoft). Definicja: `highlight/lang/dax.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2131 | Nazwy języka: `dax`. |
| LANG-2132 | Słowa z listy keyword (`VAR`, `RETURN`, `EVALUATE`, `DEFINE`, `MEASURE`, `ORDER` …) dostają `hl-kw`; słowa z listy builtin (`CALCULATE`, `SUM`, `SUMX`, `FILTER`, `ALL`, `ALLEXCEPT` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2133 | Komentarze `//` do końca wiersza, `--` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką; `'…'` (nazwa, bez klasy) dostają `hl-str`. |
| LANG-2134 | Liczby dostają `hl-num`; znaki `+-*/^&=<>|,()[]` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
