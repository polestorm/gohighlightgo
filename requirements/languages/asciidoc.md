# Wymagania: AsciiDoc

Źródło: AsciiDoc Language Documentation (Eclipse AsciiDoc).

| Id | Wymaganie |
|---|---|
| LANG-541 | Nazwy języka: `asciidoc`, `adoc`, `asc`. |
| LANG-542 | Tytuł (1–6 znaków `=` i odstęp) dostaje `hl-kw`; wpis atrybutu `:nazwa:` `hl-attr`; atrybuty bloku `[…]` `hl-meta`; znaczniki list `*`, `-`, `.` `hl-op`. |
| LANG-543 | Komentarz `//` i blok `////` dostają `hl-com`. |
| LANG-544 | Ograniczniki bloków (`----`, `....`, `====`, `****`, `____`, `++++`, `|===`) dostają `hl-meta`; treść bloku `----` po `[source,język]` jest kolorowana tym językiem, a bez niego (i w bloku `....`) dostaje `hl-str`; kod `` `…` `` dostaje `hl-str`, odwołanie `{nazwa}` `hl-esc`. |
