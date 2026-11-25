# Wymagania: Pine Script

Źródło: TradingView Pine Script v5 language reference. Definicja: `highlight/lang/pinescript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2261 | Nazwy języka: `pinescript`, `pine`. |
| LANG-2262 | Słowa z listy keyword (`if`, `else`, `for`, `to`, `by`, `while` …) dostają `hl-kw`; słowa z listy type (`int`, `float`, `bool`, `color`, `string`, `line` …) dostają `hl-type`; słowa z listy builtin (`close`, `open`, `high`, `low`, `volume`, `ta` …) dostają `hl-builtin`. |
| LANG-2263 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2264 | Liczby dostają `hl-num`; znaki `+-*/%=<>!?:.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
