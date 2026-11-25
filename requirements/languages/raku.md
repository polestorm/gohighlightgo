# Wymagania: Raku

Źródło: dokumentacja Raku (docs.raku.org): składnia, operatory, cytowanie. Definicja: `highlight/lang/raku.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1461 | Nazwy języka: `raku`, `perl6`, `p6`, `rakumod`. |
| LANG-1462 | Słowa kluczowe i deklaratory dostają `hl-kw`; `True False Nil Any Mu self` `hl-builtin`; zmienne `$x`, `@x`, `%x`, `&x` `hl-attr`; nazwa może zawierać `-` i `'`. |
| LANG-1463 | Komentarz `#` dostaje `hl-com`; napisy `"…"` (ze wstawkami `$x`) i `'…'` `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
