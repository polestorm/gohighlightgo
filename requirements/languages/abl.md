# Wymagania: OpenEdge ABL

Źródło: Progress OpenEdge ABL Reference (12.x). Definicja: `highlight/lang/abl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2001 | Nazwy języka: `abl`, `openedge`, `progress`. |
| LANG-2002 | Słowa z listy keyword (`DEFINE`, `VARIABLE`, `AS`, `INITIAL`, `NO-UNDO`, `FOR` …) dostają `hl-kw`; słowa z listy type (`CHARACTER`, `INTEGER`, `DECIMAL`, `LOGICAL`, `DATE`, `DATETIME` …) dostają `hl-type`; słowa z listy builtin (`TRUE`, `FALSE`, `YES`, `NO`, `TODAY`, `NOW`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2003 | Komentarze `/* … */` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2004 | Liczby dostają `hl-num`; znaki `+-*/=<>:.` dostają `hl-op`. |
