# Wymagania: gnuplot

Źródło: gnuplot 5.4 documentation: commands and expressions. Definicja: `highlight/lang/gnuplot.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1651 | Nazwy języka: `gnuplot`, `gp`, `gnu`. |
| LANG-1652 | Słowa z listy keyword (`plot`, `splot`, `replot`, `set`, `unset`, `show` …) dostają `hl-kw`; słowa z listy builtin (`sin`, `cos`, `tan`, `exp`, `log`, `sqrt` …) dostają `hl-builtin`. |
| LANG-1653 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-1654 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
