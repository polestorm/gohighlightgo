# Wymagania: FreeBASIC

Źródło: FreeBASIC manual (keywords and language reference). Definicja: `highlight/lang/freebasic.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1931 | Nazwy języka: `freebasic`, `fb`, `bas`. |
| LANG-1932 | Słowa z listy keyword (`Dim`, `As`, `Integer`, `Long`, `Single`, `Double` …) dostają `hl-kw`; słowa z listy builtin (`True`, `False`, `Len`, `Mid`, `Left`, `Right` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-1933 | Komentarze `'` do końca wiersza, `REM` do końca wiersza, `/' … '/` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-1934 | Liczby dostają `hl-num`; znaki `+-*/\^=<>&` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
