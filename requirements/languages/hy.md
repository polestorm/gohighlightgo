# Wymagania: Hy

Źródło: Hy documentation (1.0): syntax and macros. Definicja: `highlight/lang/hy.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1661 | Nazwy języka: `hy`, `hylang`. |
| LANG-1662 | Słowa z listy keyword (`defn`, `defmacro`, `defclass`, `fn`, `let`, `setv` …) dostają `hl-kw`; słowa z listy builtin (`True`, `False`, `None`, `print`, `len`, `range` …) dostają `hl-builtin`; nazwa po `:` dostaje `hl-attr`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-1663 | Komentarze `;` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, wielowierszowy dostają `hl-str`. |
| LANG-1664 | Liczby dostają `hl-num`. |
