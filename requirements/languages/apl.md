# Wymagania: APL

Źródło: ISO/IEC 13751 (Extended APL) and the Dyalog APL Language Reference. Definicja: `highlight/lang/apl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2281 | Nazwy języka: `apl`, `dyalog`. |
| LANG-2282 | Słowa z listy keyword (`:If`, `:Else`, `:ElseIf`, `:EndIf`, `:For`, `:In` …) dostają `hl-kw`. |
| LANG-2283 | Komentarze `⍝` do końca wiersza dostają `hl-com`; napisy `'…'` z podwojonym ogranicznikiem jako ucieczką; `"…"` dostają `hl-str`. |
| LANG-2284 | Liczby dostają `hl-num`; znaki `+-*/<>=,.:;()[]{}` dostają `hl-op`. |
