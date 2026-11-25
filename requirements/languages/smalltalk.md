# Wymagania: Smalltalk

Źródło: ANSI INCITS 319-1998 (Smalltalk) i podsumowanie składni z „Pharo by Example”. Definicja: `highlight/lang/smalltalk.hl`.

| Id | Wymaganie |
|---|---|
| LANG-731 | Nazwy języka: `smalltalk`, `st`, `squeak`, `pharo`. |
| LANG-732 | `self super thisContext` dostają `hl-kw`; `true false nil` `hl-builtin`; nazwa z wielkiej litery (klasa) `hl-type`. |
| LANG-733 | Komentarz `"…"` (w cudzysłowach, może obejmować wiele wierszy) dostaje `hl-com`. |
| LANG-734 | Napis `'…'` z `''` jako ucieczką i znak `$a` dostają `hl-str`; symbol `#nazwa` `hl-attr`; liczby `hl-num`; operatory `hl-op`. |
