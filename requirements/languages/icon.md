# Wymagania: Icon

Źródło: The Icon Programming Language (Griswold & Griswold), 3rd edition. Definicja: `highlight/lang/icon.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2731 | Nazwy języka: `icon`, `icn`, `unicon`. |
| LANG-2732 | Słowa z listy keyword (`procedure`, `end`, `local`, `static`, `global`, `record` …) dostają `hl-kw`; słowa z listy builtin (`write`, `writes`, `read`, `find`, `upto`, `many` …) dostają `hl-builtin`; nazwa po `&` dostaje `hl-builtin`. |
| LANG-2733 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2734 | Liczby dostają `hl-num`; znaki `+-*/%=<>|!:~\?@^.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
