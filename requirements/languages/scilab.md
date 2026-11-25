# Wymagania: Scilab

Źródło: Scilab 2024 help: language syntax. Definicja: `highlight/lang/scilab.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2051 | Nazwy języka: `scilab`, `sce`, `sci`. |
| LANG-2052 | Słowa z listy keyword (`function`, `endfunction`, `if`, `then`, `else`, `elseif` …) dostają `hl-kw`; słowa z listy builtin (`disp`, `printf`, `mprintf`, `zeros`, `ones`, `size` …) dostają `hl-builtin`. |
| LANG-2053 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką; `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2054 | Liczby dostają `hl-num`; znaki `+-*/^\=<>~&|:.'` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
