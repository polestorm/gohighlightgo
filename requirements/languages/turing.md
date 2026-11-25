# Wymagania: Turing

Źródło: Turing Reference Manual (Holt Software). Definicja: `highlight/lang/turing.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2891 | Nazwy języka: `turing`, `tu`. |
| LANG-2892 | Słowa z listy keyword (`var`, `const`, `if`, `then`, `elsif`, `else` …) dostają `hl-kw`; słowa z listy type (`int`, `real`, `string`, `boolean`, `char`, `nat`) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `length`, `upper`, `index`) dostają `hl-builtin`. |
| LANG-2893 | Komentarze `%` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2894 | Liczby dostają `hl-num`; znaki `+-*/=<>:.,()` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
