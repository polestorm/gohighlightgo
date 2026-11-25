# Wymagania: Typst

Źródło: Typst reference (typst.app/docs/reference): syntax. Definicja: `highlight/lang/typst.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2451 | Nazwy języka: `typst`, `typ`. |
| LANG-2452 | Słowa z listy keyword (`let`, `set`, `show`, `if`, `else`, `for` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`) dostają `hl-builtin`; nazwa po `#` dostaje `hl-kw`; nazwa po `@` dostaje `hl-attr`. |
| LANG-2453 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2454 | Liczby dostają `hl-num`; znaki `=*_$<>:+-/` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
