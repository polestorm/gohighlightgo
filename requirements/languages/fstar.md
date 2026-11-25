# Wymagania: F*

Źródło: Proof-oriented Programming in F* (podręcznik języka): struktura leksykalna i słowa kluczowe.

| Id | Wymaganie |
|---|---|
| LANG-1161 | Nazwy języka: `fstar`, `fst`, `fsti`. |
| LANG-1162 | Słowa kluczowe i kwalifikatory dostają `hl-kw`; typy z listy i nazwy z wielkiej litery `hl-type`; `True False true false Some None` `hl-builtin`; apostrof wewnątrz nazwy należy do nazwy. |
| LANG-1163 | Komentarze `//` i zagnieżdżany `(* … *)` dostają `hl-com`. |
| LANG-1164 | Napis `"…"` i znak `'a'` dostają `hl-str`; zmienna typu `'a` `hl-meta`; liczby `hl-num`. |
