# Wymagania: RON

Źródło: Rusty Object Notation, gramatyka (ron-rs/ron, docs/grammar.md).

| Id | Wymaganie |
|---|---|
| LANG-1011 | Nazwa języka: `ron`. |
| LANG-1012 | Nazwa pola przed `:` dostaje `hl-attr`; nazwa struktury z wielkiej litery `hl-type`; `true false Some None` `hl-builtin`. |
| LANG-1013 | Komentarze `//` i zagnieżdżany `/* … */` dostają `hl-com`. |
| LANG-1014 | Napisy `"…"` (także surowe `r#"…"#`) i znaki `'a'` dostają `hl-str`; liczby `hl-num`. |
