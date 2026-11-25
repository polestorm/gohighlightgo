# Wymagania: Nim

Źródło: Nim Manual (2.0), „Lexical Analysis”. Definicja: `highlight/lang/nim.hl`.

| Id | Wymaganie |
|---|---|
| LANG-561 | Nazwy języka: `nim`, `nims`, `nimble`. |
| LANG-562 | Słowa kluczowe dostają `hl-kw`; typy wbudowane `hl-type`; `true false` i procedury z listy `hl-builtin`. |
| LANG-563 | Komentarz `#` (także `##`) i zagnieżdżany `#[ … ]#` dostają `hl-com`. |
| LANG-564 | Napisy `"…"` z ucieczkami, surowe `r"…"` i potrójne `"""…"""` oraz znak `'a'` dostają `hl-str`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` `hl-fn`. |
