# Wymagania: Berry

Źródło: Berry language reference manual (berry-lang.github.io). Definicja: `highlight/lang/berry.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1871 | Nazwy języka: `berry`, `be`. |
| LANG-1872 | Słowa z listy keyword (`if`, `elif`, `else`, `while`, `for`, `def` …) dostają `hl-kw`; słowa z listy builtin (`print`, `type`, `classname`, `size`, `super`, `self` …) dostają `hl-builtin`. |
| LANG-1873 | Komentarze `#` do końca wiersza, `#- … -#` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1874 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
