# Wymagania: Wolfram Language

Źródło: Wolfram Language documentation (reference.wolfram.com): składnia i operatory. Definicja: `highlight/lang/wolfram.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1521 | Nazwy języka: `wolfram`, `mathematica`, `wl`, `nb`. |
| LANG-1522 | Stałe z listy (`True`, `Pi`, `Infinity` …) dostają `hl-builtin`; inne nazwy z wielkiej litery (funkcje wbudowane, `Plot`, `Module`) `hl-type`. |
| LANG-1523 | Komentarz `(* … *)`, który się zagnieżdża, dostaje `hl-com`; napis `"…"` `hl-str`; liczby `hl-num`; operatory (także `[ ]`, `_`, `#`, `&`) `hl-op`. |
