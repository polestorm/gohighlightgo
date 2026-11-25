# Wymagania: Jsonnet

Źródło: Jsonnet language reference i specyfikacja (jsonnet.org).

| Id | Wymaganie |
|---|---|
| LANG-1021 | Nazwy języka: `jsonnet`, `libsonnet`. |
| LANG-1022 | Słowa kluczowe dostają `hl-kw`; `true false null std $` `hl-builtin`; nazwa pola przed `:` (także `::`, `:::`) `hl-attr`; nazwa przed `(` `hl-fn`. |
| LANG-1023 | Komentarze `//`, `#` i `/* … */` dostają `hl-com`. |
| LANG-1024 | Napisy `"…"` i `'…'` z ucieczkami, dosłowne `@"…"` i bloki tekstu `|||…|||` dostają `hl-str`; liczby `hl-num`. |
