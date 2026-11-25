# Wymagania: Visual Basic

Źródło: Visual Basic language reference: Keywords, Comments in Code, String Data Type. Definicja: `highlight/lang/vb.hl`.

| Id | Wymaganie |
|---|---|
| LANG-411 | Nazwy języka: `vb`, `vbnet`, `visualbasic`, `vba`, `vbs`, `vbscript`. |
| LANG-412 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; typy z listy `hl-type`; `True False Nothing` i funkcje z listy `hl-builtin`. |
| LANG-413 | Komentarz `'` albo `REM` do końca wiersza dostaje `hl-com`. |
| LANG-414 | Napis `"…"` z `""` jako ucieczką dostaje `hl-str`. |
| LANG-415 | Liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
