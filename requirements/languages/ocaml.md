# Wymagania: OCaml

Źródło: The OCaml Manual (5.1), „Lexical conventions”. Definicja: `highlight/lang/ocaml.hl`.

| Id | Wymaganie |
|---|---|
| LANG-361 | Nazwy języka: `ocaml`, `ml`, `mli`. |
| LANG-362 | Słowa kluczowe dostają `hl-kw`; typy podstawowe `hl-type`, a nazwy z wielkiej litery (konstruktory, moduły) też `hl-type`; `true false` i funkcje z listy `hl-builtin`; apostrof wewnątrz nazwy (`x'`) należy do nazwy. |
| LANG-363 | Komentarz `(* … *)`, który się zagnieżdża, dostaje `hl-com`. |
| LANG-364 | Napis `"…"` (może obejmować wiele wierszy) z ucieczkami i napis cytowany `{|…|}` dostają `hl-str`; znak `'a'` dostaje `hl-str`, a zmienna typu `'a` `hl-meta`. |
| LANG-365 | Liczby (bez kropki na początku) `hl-num`; operatory `hl-op`. |
