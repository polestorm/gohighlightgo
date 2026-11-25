# Wymagania: CFML (CFScript)

Źródło: Adobe ColdFusion and Lucee documentation: CFScript. Definicja: `highlight/lang/cfml.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2801 | Nazwy języka: `cfml`, `cfscript`, `cfc`, `coldfusion`. |
| LANG-2802 | Słowa z listy keyword (`component`, `function`, `property`, `return`, `if`, `else` …) dostają `hl-kw`; słowa z listy type (`string`, `numeric`, `boolean`, `struct`, `array`, `query` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `writeOutput`, `arrayLen`, `structKeyExists` …) dostają `hl-builtin`; nazwa po `#` dostaje `hl-attr` (bez rozróżniania wielkości liter). |
| LANG-2803 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką; `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2804 | Liczby dostają `hl-num`; znaki `+-*/%&|!<>=:?.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
