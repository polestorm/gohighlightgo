# Wymagania: Eiffel

Źródło: ECMA-367 (Eiffel: Analysis, Design and Programming Language), 8.32 lexical components. Definicja: `highlight/lang/eiffel.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1611 | Nazwy języka: `eiffel`, `e`. |
| LANG-1612 | Słowa z listy keyword (`across`, `agent`, `alias`, `all`, `and`, `as` …) dostają `hl-kw`; słowa z listy type (`INTEGER`, `REAL`, `DOUBLE`, `BOOLEAN`, `CHARACTER`, `STRING` …) dostają `hl-type`; słowa z listy builtin (`True`, `False`, `Void`, `print`, `io`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type` (bez rozróżniania wielkości liter). |
| LANG-1613 | Komentarze `--` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1614 | Liczby dostają `hl-num`; znaki `+-*/=<>:.~` dostają `hl-op`. |
