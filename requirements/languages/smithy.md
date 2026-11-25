# Wymagania: Smithy

Źródło: Smithy 2.0 specification (smithy.io): IDL. Definicja: `highlight/lang/smithy.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2441 | Nazwy języka: `smithy`. |
| LANG-2442 | Słowa z listy keyword (`namespace`, `use`, `service`, `resource`, `operation`, `structure` …) dostają `hl-kw`; słowa z listy type (`String`, `Integer`, `Long`, `Boolean`, `Timestamp`, `Blob` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`; nazwa po `@` dostaje `hl-meta`; nazwa po `$` dostaje `hl-meta`. |
| LANG-2443 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"""…"""` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2444 | Liczby dostają `hl-num`; znaki `:={}[](),.#` dostają `hl-op`. |
