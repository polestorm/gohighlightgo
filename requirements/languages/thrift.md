# Wymagania: Apache Thrift

Źródło: Apache Thrift interface description language (thrift.apache.org/docs/idl). Definicja: `highlight/lang/thrift.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1761 | Nazwy języka: `thrift`. |
| LANG-1762 | Słowa z listy keyword (`namespace`, `include`, `struct`, `union`, `exception`, `service` …) dostają `hl-kw`; słowa z listy type (`bool`, `byte`, `i8`, `i16`, `i32`, `i64` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-1763 | Komentarze `//` do końca wiersza, `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1764 | Liczby dostają `hl-num`; znaki `=:;,<>` dostają `hl-op`. |
