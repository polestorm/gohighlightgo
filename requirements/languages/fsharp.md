# Wymagania: F#

Źródło: F# Language Specification i „F# language reference” (F# 8). Definicja: `highlight/lang/fsharp.hl`.

| Id | Wymaganie |
|---|---|
| LANG-351 | Nazwy języka: `fsharp`, `fs`, `f#`. |
| LANG-352 | Słowa kluczowe dostają `hl-kw`; typy z listy `hl-type`, a inne nazwy z wielkiej litery (typy, przypadki unii, moduły) też `hl-type`; `true false null`, `Some None` i funkcje z listy `hl-builtin`. |
| LANG-353 | Komentarze `//` i `(* … *)`, które się zagnieżdżają, dostają `hl-com`. |
| LANG-354 | Napisy `"…"` z ucieczkami, dosłowne `@"…"`, potrójne `"""…"""` i z wstawkami `$"…{x}…"` oraz znaki `'…'` dostają `hl-str` (wstawki i ucieczki `hl-esc`). |
| LANG-355 | Liczby `hl-num`; operatory `hl-op`. Wywołania nie są rozpoznawane, bo funkcja nie wymaga nawiasów. |
