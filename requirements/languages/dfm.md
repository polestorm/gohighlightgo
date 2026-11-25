# Wymagania: formularze Delphi

Źródło: dokumentacja RAD Studio, pliki formularzy w postaci tekstowej (DFM, LFM, FMX).

| Id | Wymaganie |
|---|---|
| LANG-1131 | Nazwy języka: `dfm`, `lfm`, `fmx`. |
| LANG-1132 | `object`, `inherited`, `inline`, `end`, `item` dostają `hl-kw`; nazwa właściwości przed `=` `hl-attr`; nazwa klasy z wielkiej litery (`TForm1`) `hl-type`; `True False nil` `hl-builtin`. |
| LANG-1133 | Napis `'…'` z `''` dostaje `hl-str`; kody znaków `#13` `hl-esc`; liczby (także `$FF`) `hl-num`. |
