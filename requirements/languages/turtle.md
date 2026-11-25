# Wymagania: Turtle

Źródło: RDF 1.1 Turtle (W3C Recommendation 2014). Definicja: `highlight/lang/turtle.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1771 | Nazwy języka: `turtle`, `ttl`, `n3`, `trig`. |
| LANG-1772 | Słowa z listy keyword (`@prefix`, `@base`, `PREFIX`, `BASE`) dostają `hl-kw`; słowa z listy builtin (`a`, `true`, `false`) dostają `hl-builtin`; nazwa po `@` dostaje `hl-kw`. |
| LANG-1773 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"""…"""` z ucieczkami `hl-esc`, wielowierszowy; `'''…'''` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1774 | Liczby dostają `hl-num`; znaki `;,.^[]` dostają `hl-op`; konstrukcje haka `iri` jak w opisie haka w `docs/definitions.md`. |
