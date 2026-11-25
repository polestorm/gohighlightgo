# Wymagania: QBasic

Źródło: QBasic/QuickBASIC 4.5 language reference. Definicja: `highlight/lang/qbasic.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1941 | Nazwy języka: `qbasic`, `quickbasic`, `qb`. |
| LANG-1942 | Słowa z listy keyword (`DIM`, `AS`, `INTEGER`, `LONG`, `SINGLE`, `DOUBLE` …) dostają `hl-kw`; słowa z listy builtin (`LEN`, `MID$`, `LEFT$`, `RIGHT$`, `STR$`, `VAL` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-1943 | Komentarze `'` do końca wiersza, `REM` do końca wiersza dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-1944 | Liczby dostają `hl-num`; znaki `+-*/\^=<>` dostają `hl-op`. |
