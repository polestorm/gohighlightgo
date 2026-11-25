# Wymagania: POV-Ray

Źródło: POV-Ray 3.7 documentation: Scene Description Language reference. Definicja: `highlight/lang/povray.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1821 | Nazwy języka: `povray`, `pov`, `inc-pov`. |
| LANG-1822 | Słowa z listy keyword (`camera`, `light_source`, `sphere`, `box`, `cylinder`, `cone` …) dostają `hl-kw`; słowa z listy builtin (`x`, `y`, `z`, `true`, `false`, `on` …) dostają `hl-builtin`; nazwa po `#` dostaje `hl-meta`. |
| LANG-1823 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1824 | Liczby dostają `hl-num`; znaki `+-*/=<>!&|?:<>.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
