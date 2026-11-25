# Wymagania: Graphviz DOT

Źródło: The DOT Language (graphviz.org/doc/info/lang.html). Definicja: `highlight/lang/dot.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2371 | Nazwy języka: `dot`, `graphviz`, `gv`. |
| LANG-2372 | Słowa z listy keyword (`graph`, `digraph`, `subgraph`, `node`, `edge`, `strict`) dostają `hl-kw`; słowa z listy builtin (`label`, `shape`, `color`, `style`, `fillcolor`, `rankdir` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2373 | Komentarze `//` do końca wiersza, `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `<…>` wielowierszowy dostają `hl-str`. |
| LANG-2374 | Liczby dostają `hl-num`; znaki `->=;,[]{}` dostają `hl-op`. |
