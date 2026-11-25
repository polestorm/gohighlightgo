# Wymagania: Ceylon

Źródło: Ceylon language specification 1.3, lexical structure. Definicja: `highlight/lang/ceylon.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1601 | Nazwy języka: `ceylon`. |
| LANG-1602 | Słowa z listy keyword (`assembly`, `module`, `package`, `import`, `alias`, `class` …) dostają `hl-kw`; słowa z listy type (`String`, `Integer`, `Float`, `Boolean`, `Character`, `Anything` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `print`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-1603 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"""…"""` wielowierszowy; `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1604 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
