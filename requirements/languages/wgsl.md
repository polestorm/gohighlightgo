# Wymagania: WGSL

Źródło: WebGPU Shading Language (W3C Working Draft), 15 Keyword and Token Summary. Definicja: `highlight/lang/wgsl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1781 | Nazwy języka: `wgsl`. |
| LANG-1782 | Słowa z listy keyword (`fn`, `let`, `var`, `const`, `override`, `struct` …) dostają `hl-kw`; słowa z listy type (`bool`, `i32`, `u32`, `f32`, `f16`, `vec2` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `dot`, `cross`, `normalize`, `length` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-meta`. |
| LANG-1783 | Komentarze `//` do końca wiersza, `/* … */` (zagnieżdżany) dostają `hl-com`. |
| LANG-1784 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
