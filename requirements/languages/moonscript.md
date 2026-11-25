# Wymagania: MoonScript

Źródło: MoonScript language guide (0.5). Definicja: `highlight/lang/moonscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-761 | Nazwy języka: `moonscript`, `moon`. |
| LANG-762 | Słowa kluczowe dostają `hl-kw`; `true false nil self` i funkcje z listy `hl-builtin`; nazwa z wielkiej litery (klasa) `hl-type`; `@pole` `hl-attr`. |
| LANG-763 | Komentarz `--` i długi `--[[ … ]]` dostają `hl-com`. |
| LANG-764 | Napis `"…"` z wstawkami `#{…}`, `'…'` i napis długi `[[ … ]]` dostają `hl-str`; liczby `hl-num`; operatory (`->`, `=>`, `\`) `hl-op`; nazwa przed `(` `hl-fn`. |
