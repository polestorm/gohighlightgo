# Wymagania: MiniZinc

Źródło: The MiniZinc Handbook (2.8): language reference. Definicja: `highlight/lang/minizinc.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1981 | Nazwy języka: `minizinc`, `mzn`, `dzn`. |
| LANG-1982 | Słowa z listy keyword (`var`, `par`, `int`, `float`, `bool`, `set` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `alldifferent`, `show`, `abs`, `max` …) dostają `hl-builtin`. |
| LANG-1983 | Komentarze `%` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1984 | Liczby dostają `hl-num`; znaki `+-*/<>=!:\/.|` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
