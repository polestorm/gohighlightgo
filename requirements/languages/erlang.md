# Wymagania: Erlang

Źródło: Erlang Reference Manual (OTP 26): „Data Types”, „Expressions”, „Preprocessor”. Definicja: `highlight/lang/erlang.hl`.

| Id | Wymaganie |
|---|---|
| LANG-331 | Nazwy języka: `erlang`, `erl`. |
| LANG-332 | Słowa zastrzeżone dostają `hl-kw`; atomy i funkcje z listy `hl-builtin`; zmienna (nazwa z wielkiej litery) `hl-type`; makro `?NAZWA` `hl-meta`. |
| LANG-333 | Komentarz `%` do końca wiersza dostaje `hl-com`. |
| LANG-334 | Napis `"…"` z ucieczkami dostaje `hl-str`; atom w apostrofach `'…'` nie dostaje klasy. |
| LANG-335 | Liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
