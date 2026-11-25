# Wymagania: MEL

Źródło: Autodesk Maya MEL command reference. Definicja: `highlight/lang/mel.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2721 | Nazwy języka: `mel`. |
| LANG-2722 | Słowa z listy keyword (`global`, `proc`, `if`, `else`, `for`, `in` …) dostają `hl-kw`; słowa z listy type (`int`, `float`, `string`, `vector`, `matrix`) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `print`, `polySphere`, `move`, `rotate` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-2723 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2724 | Liczby dostają `hl-num`; znaki `` +-*/%&|!<>=:?.` `` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
