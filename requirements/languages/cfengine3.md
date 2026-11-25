# Wymagania: CFEngine 3

Źródło: CFEngine 3 Reference: language concepts. Definicja: `highlight/lang/cfengine3.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2921 | Nazwy języka: `cfengine3`, `cf3`, `cfengine`. |
| LANG-2922 | Słowa z listy keyword (`bundle`, `body`, `agent`, `common`, `edit_line`, `files` …) dostają `hl-kw`; słowa z listy builtin (`string`, `slist`, `int`, `real`, `data`, `ifvarclass` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-attr`. |
| LANG-2923 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc`; `'…'` z ucieczkami `hl-esc`; `` `…` `` dostają `hl-str`. |
| LANG-2924 | Liczby dostają `hl-num`; znaki `=>:;,(){}.!&|` dostają `hl-op`; konstrukcje haka `variables` jak w opisie haka w `docs/definitions.md`. |
