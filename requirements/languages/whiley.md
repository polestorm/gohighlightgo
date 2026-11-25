# Wymagania: Whiley

Źródło: Whiley Language Specification (whiley.org). Definicja: `highlight/lang/whiley.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2101 | Nazwy języka: `whiley`. |
| LANG-2102 | Słowa z listy keyword (`function`, `method`, `type`, `import`, `public`, `private` …) dostają `hl-kw`; słowa z listy type (`int`, `bool`, `byte`, `void`, `null`) dostają `hl-type`; słowa z listy builtin (`true`, `false`) dostają `hl-builtin`. |
| LANG-2103 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2104 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
