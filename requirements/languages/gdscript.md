# Wymagania: GDScript

Źródło: Godot Engine documentation (4.x): GDScript reference. Definicja: `highlight/lang/gdscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1631 | Nazwy języka: `gdscript`, `gd`. |
| LANG-1632 | Słowa z listy keyword (`if`, `elif`, `else`, `for`, `while`, `match` …) dostają `hl-kw`; słowa z listy type (`int`, `float`, `bool`, `String`, `Vector2`, `Vector3` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `PI`, `TAU`, `INF` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-meta`; nazwa po `$` dostaje `hl-attr`. |
| LANG-1633 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"""…"""` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1634 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
