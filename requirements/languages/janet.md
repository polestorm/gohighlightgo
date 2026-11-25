# Wymagania: Janet

Źródło: Janet language documentation (janet-lang.org/docs). Definicja: `highlight/lang/janet.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1371 | Nazwy języka: `janet`. |
| LANG-1372 | Formy specjalne i makra z listy dostają `hl-kw`; funkcje z listy i `true false nil` `hl-builtin`; słowo kluczowe `:nazwa` `hl-attr`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-1373 | Komentarz `#` dostaje `hl-com`; napis `"…"` i długi napis w odwrotnych apostrofach `hl-str`; liczby `hl-num`. |
