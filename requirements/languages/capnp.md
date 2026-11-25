# Wymagania: Cap'n Proto

Źródło: Cap'n Proto schema language documentation. Definicja: `highlight/lang/capnp.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1591 | Nazwy języka: `capnp`, `capnproto`. |
| LANG-1592 | Słowa z listy keyword (`struct`, `enum`, `union`, `interface`, `using`, `import` …) dostają `hl-kw`; słowa z listy type (`Void`, `Bool`, `Int8`, `Int16`, `Int32`, `Int64` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `void`) dostają `hl-builtin`; nazwa po `@` dostaje `hl-meta`; nazwa po `$` dostaje `hl-meta`. |
| LANG-1593 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1594 | Liczby dostają `hl-num`; znaki `=:;-()` dostają `hl-op`. |
