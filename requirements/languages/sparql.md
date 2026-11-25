# Wymagania: SPARQL

Źródło: SPARQL 1.1 Query Language i SPARQL 1.1 Update (W3C 2013), 19 „SPARQL Grammar”. Definicja: `highlight/lang/sparql.hl`.

| Id | Wymaganie |
|---|---|
| LANG-811 | Nazwy języka: `sparql`, `rq`. |
| LANG-812 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; `a`, `true false` i funkcje z listy `hl-builtin`; nazwa z przedrostkiem (`foaf:name`) jest jedną nazwą. |
| LANG-813 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-814 | Napisy `"…"`, `'…'`, `"""…"""`, `'''…'''` i IRI `<http://…>` (bez odstępów) dostają `hl-str`; zmienne `?x` i `$x` `hl-attr`; `<` z odstępem to operator; liczby `hl-num`. |
