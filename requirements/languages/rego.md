# Wymagania: Rego

Źródło: Open Policy Agent, Policy Language reference.

| Id | Wymaganie |
|---|---|
| LANG-1051 | Nazwa języka: `rego`. |
| LANG-1052 | Słowa kluczowe dostają `hl-kw`; `true false null input data` i funkcje wbudowane z listy `hl-builtin`. |
| LANG-1053 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-1054 | Napis `"…"` z ucieczkami i surowy `` `…` `` dostają `hl-str`; liczby `hl-num`; operatory (`:=`, `==`) `hl-op`; nazwa przed `(` `hl-fn`. |
