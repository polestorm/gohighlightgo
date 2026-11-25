# Wymagania: Vim script

Źródło: Vim reference manual (9.1): „:help eval”, „:help expression-syntax”, „:help vim9”.

| Id | Wymaganie |
|---|---|
| LANG-531 | Nazwy języka: `vim`, `vimscript`, `viml`. |
| LANG-532 | Polecenia z listy dostają `hl-kw`; funkcje z listy i `v:true` `hl-builtin`; nazwy z zakresem i autoload (`g:nazwa`, `s:f`, `pakiet#f`) są jedną nazwą; `&opcja`, `@r`, `$ZMIENNA` dostają `hl-attr`. |
| LANG-533 | Komentarz `"` na początku polecenia (po odstępach albo po `|`) i komentarz `#` (Vim9) dostają `hl-com`. |
| LANG-534 | Napis `"…"` (w innym miejscu niż komentarz) z ucieczkami i `'…'` z `''` dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
