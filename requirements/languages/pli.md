# Wymagania: PL/I

Źródło: IBM Enterprise PL/I for z/OS Language Reference. Definicja: `highlight/lang/pli.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2741 | Nazwy języka: `pli`, `pl1`. |
| LANG-2742 | Słowa z listy keyword (`PROCEDURE`, `PROC`, `END`, `DECLARE`, `DCL`, `IF` …) dostają `hl-kw`; słowa z listy type (`FIXED`, `FLOAT`, `BINARY`, `DECIMAL`, `CHARACTER`, `CHAR` …) dostają `hl-type`; słowa z listy builtin (`SUBSTR`, `LENGTH`, `INDEX`, `TRIM`, `MOD`, `ABS` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2743 | Komentarze `/* … */` dostają `hl-com`; napisy `'…'` z podwojonym ogranicznikiem jako ucieczką; `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2744 | Liczby dostają `hl-num`; znaki `+-*/=<>|&^;:(),.` dostają `hl-op`. |
