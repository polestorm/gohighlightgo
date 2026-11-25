# Wymagania: Lean 4

Źródło: Lean language reference (lean-lang.org/doc/reference): struktura leksykalna, polecenia, taktyki.

| Id | Wymaganie |
|---|---|
| LANG-1171 | Nazwy języka: `lean`, `lean4`. |
| LANG-1172 | Polecenia i słowa kluczowe dostają `hl-kw`; taktyki z listy i `true false sorry rfl` `hl-builtin`; nazwy z wielkiej litery `hl-type`; atrybut `@nazwa` `hl-meta`; nazwa może zawierać `' . ! ?`. |
| LANG-1173 | Komentarz `--` i zagnieżdżany `/- … -/` (także `/-- … -/`) dostają `hl-com`. |
| LANG-1174 | Napis `"…"` i znak `'a'` dostają `hl-str`; liczby `hl-num`; operatory `hl-op`. |
