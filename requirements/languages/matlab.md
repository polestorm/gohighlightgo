# Wymagania: MATLAB

Źródło: dokumentacja MATLAB R2024a, „Language Fundamentals” i lista z `iskeyword`. Definicja: `highlight/lang/matlab.hl`.

| Id | Wymaganie |
|---|---|
| LANG-381 | Nazwy języka: `matlab`, `octave`. |
| LANG-382 | Słowa kluczowe dostają `hl-kw`; funkcje z listy i `true false pi` `hl-builtin`. |
| LANG-383 | Komentarz `%` do końca wiersza i komentarz blokowy `%{` … `%}` (znaczniki w osobnych wierszach) dostają `hl-com`. |
| LANG-384 | Napis `"…"` z `""` jako ucieczką dostaje `hl-str`; apostrof po nazwie, liczbie, `)`, `]`, `}`, `.` albo apostrofie to transpozycja (`hl-op`), a w innym miejscu zaczyna napis `'…'` z `''` w środku, do końca wiersza. |
| LANG-385 | Liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
