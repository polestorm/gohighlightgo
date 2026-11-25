# Wymagania: ALGOL 60/68

Źródło: Revised Report on the Algorithmic Language ALGOL 60 and ALGOL 68 Revised Report. Definicja: `highlight/lang/algol.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2901 | Nazwy języka: `algol`, `algol60`, `algol68`, `a68`. |
| LANG-2902 | Słowa z listy keyword (`begin`, `end`, `if`, `then`, `else`, `elif` …) dostają `hl-kw`; słowa z listy type (`integer`, `real`, `boolean`, `string`, `array`, `int` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `print`, `read`) dostają `hl-builtin`. |
| LANG-2903 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2904 | Liczby dostają `hl-num`; znaki `+-*/:=<>^&|~(),;` dostają `hl-op`. |
