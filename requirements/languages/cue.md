# Wymagania: CUE

Źródło: The CUE Language Specification (cuelang.org/docs/reference/spec). Definicja: `highlight/lang/cue.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2411 | Nazwy języka: `cue`. |
| LANG-2412 | Słowa z listy keyword (`package`, `import`, `let`, `if`, `for`, `in`) dostają `hl-kw`; słowa z listy type (`string`, `int`, `float`, `number`, `bool`, `bytes` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `_`, `close`, `len`) dostają `hl-builtin`. |
| LANG-2413 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"""…"""` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2414 | Liczby dostają `hl-num`; znaki `+-*/&|!<>=:?.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
