# Wymagania: Koka

Źródło: The Koka Programming Language (koka-lang.github.io). Definicja: `highlight/lang/koka.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2221 | Nazwy języka: `koka`, `kk`. |
| LANG-2222 | Słowa z listy keyword (`fun`, `fn`, `val`, `var`, `if`, `then` …) dostają `hl-kw`; słowa z listy type (`int`, `string`, `bool`, `char`, `double`, `list` …) dostają `hl-type`; słowa z listy builtin (`True`, `False`, `Nothing`, `Just`, `println`, `show`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-2223 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2224 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
