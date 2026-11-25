# Wymagania: Gleam

Źródło: Gleam language tour i dokumentacja (1.x). Definicja: `highlight/lang/gleam.hl`.

| Id | Wymaganie |
|---|---|
| LANG-901 | Nazwa języka: `gleam`. |
| LANG-902 | Słowa kluczowe dostają `hl-kw`; `True False Nil Ok Error` `hl-builtin`; inne nazwy z wielkiej litery (typy, konstruktory) `hl-type`; atrybut `@nazwa` `hl-meta`. |
| LANG-903 | Komentarze `//`, `///` i `////` dostają `hl-com`. |
| LANG-904 | Napis `"…"` z ucieczkami (`\u{…}`) dostaje `hl-str`; liczby `hl-num`; operatory (`|>`, `<>`) `hl-op`; nazwa przed `(` `hl-fn`. |
