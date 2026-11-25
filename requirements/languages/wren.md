# Wymagania: Wren

Źródło: Wren documentation (wren.io): syntax. Definicja: `highlight/lang/wren.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2231 | Nazwy języka: `wren`. |
| LANG-2232 | Słowa z listy keyword (`as`, `break`, `class`, `construct`, `continue`, `else` …) dostają `hl-kw`; słowa z listy builtin (`System`, `Fn`, `Fiber`, `List`, `Map`, `Num` …) dostają `hl-builtin`. |
| LANG-2233 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"""…"""` wielowierszowy; `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2234 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
