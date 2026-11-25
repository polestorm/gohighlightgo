# Wymagania: PostScript

Źródło: PostScript Language Reference, third edition (Adobe), 3.2 „Syntax”. Definicja: `highlight/lang/postscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1441 | Nazwy języka: `postscript`, `ps`, `eps`. |
| LANG-1442 | Operatory sterujące (`def`, `if`, `ifelse`, `for` …) dostają `hl-kw`; operatory z listy (`moveto`, `stroke`, `showpage` …) `hl-builtin`; nazwa dosłowna `/nazwa` `hl-attr`. |
| LANG-1443 | Komentarz `%` dostaje `hl-com`; napis `( … )` (może obejmować wiele wierszy, z ucieczkami) i ASCII85 `<~ … ~>` `hl-str`; liczby `hl-num`; `[ ] { } < >` `hl-op`. |
