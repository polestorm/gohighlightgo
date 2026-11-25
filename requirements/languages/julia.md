# Wymagania: Julia

Źródło: Julia Documentation (1.10): „Strings”, „Variables”, lista słów kluczowych. Definicja: `highlight/lang/julia.hl`.

| Id | Wymaganie |
|---|---|
| LANG-371 | Nazwy języka: `julia`, `jl`. |
| LANG-372 | Słowa kluczowe dostają `hl-kw`; typy z listy i nazwy z wielkiej litery `hl-type`; `true false nothing missing` i funkcje z listy `hl-builtin`; makro `@nazwa` `hl-meta`; `!` na końcu nazwy należy do nazwy (`push!`). |
| LANG-373 | Komentarz `#` do końca wiersza i `#= … =#`, który się zagnieżdża, dostają `hl-com`. |
| LANG-374 | Napisy `"…"` i `"""…"""` z ucieczkami i wstawkami `$x`, `$(…)` dostają `hl-str`; znak `'a'` dostaje `hl-str`, a apostrof po nazwie albo nawiasie (transpozycja) nie zaczyna znaku. |
| LANG-375 | Liczby (także `1im`) `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
