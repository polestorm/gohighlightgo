# Wymagania: SAS

Źródło: SAS 9.4 Programming Documentation: DATA step, PROC, macro language. Definicja: `highlight/lang/sas.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1491 | Nazwy języka: `sas`. |
| LANG-1492 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; funkcje z listy `hl-builtin`; instrukcja makra `%nazwa` `hl-meta`; zmienna makra `&nazwa` `hl-attr`. |
| LANG-1493 | Komentarz `/* … */` dostaje `hl-com`; napisy `"…"` i `'…'` z podwojonym ogranicznikiem `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
