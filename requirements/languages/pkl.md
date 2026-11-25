# Wymagania: Pkl

Źródło: Pkl Language Reference (pkl-lang.org). Definicja: `highlight/lang/pkl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2421 | Nazwy języka: `pkl`. |
| LANG-2422 | Słowa z listy keyword (`module`, `amends`, `extends`, `import`, `class`, `typealias` …) dostają `hl-kw`; słowa z listy type (`String`, `Int`, `Float`, `Boolean`, `Duration`, `DataSize` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`; nazwa po `@` dostaje `hl-meta`. |
| LANG-2423 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"""…"""` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2424 | Liczby dostają `hl-num`; znaki `+-*/%&|!<>=:?.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
