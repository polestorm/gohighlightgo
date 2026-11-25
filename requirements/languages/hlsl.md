# Wymagania: HLSL

Źródło: Microsoft HLSL reference (Shader Model 6): keywords and data types. Definicja: `highlight/lang/hlsl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2331 | Nazwy języka: `hlsl`, `fx`. |
| LANG-2332 | Słowa z listy keyword (`if`, `else`, `for`, `while`, `do`, `switch` …) dostają `hl-kw`; słowa z listy type (`void`, `bool`, `int`, `uint`, `float`, `double` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `mul`, `dot`, `cross`, `normalize` …) dostają `hl-builtin`. |
| LANG-2333 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2334 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`; konstrukcje haka `c-preprocessor` jak w opisie haka w `docs/definitions.md`. |
