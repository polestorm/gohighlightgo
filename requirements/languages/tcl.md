# Wymagania: Tcl

Źródło: Tcl(n) (dwanaście reguł składni) i polecenia wbudowane Tcl 8.6. Definicja: `highlight/lang/tcl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-681 | Nazwy języka: `tcl`, `tk`. |
| LANG-682 | Polecenia sterujące i wbudowane z listy dostają `hl-kw` albo `hl-builtin`; zmienna `$nazwa` dostaje `hl-attr`. |
| LANG-683 | `#` na początku polecenia (po odstępach albo po `;`) zaczyna komentarz do końca wiersza (`hl-com`); w innym miejscu jest zwykłym znakiem. |
| LANG-684 | Napis `"…"` z ucieczkami i wstawkami `$x`, `${…}` dostaje `hl-str`; liczby `hl-num`; operatory `hl-op`. |
