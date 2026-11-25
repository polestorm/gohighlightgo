# Wymagania: ink

Źródło: ink writing language documentation (inkle, WritingWithInk.md). Definicja: `highlight/lang/ink.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2311 | Nazwy języka: `ink`. |
| LANG-2312 | Słowa z listy keyword (`VAR`, `CONST`, `INCLUDE`, `EXTERNAL`, `LIST`, `function` …) dostają `hl-kw`; słowa z listy builtin (`END`, `DONE`, `RANDOM`, `CHOICE_COUNT`, `TURNS_SINCE`) dostają `hl-builtin`; nazwa po `~` dostaje `hl-kw`. |
| LANG-2313 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-2314 | Liczby dostają `hl-num`; znaki `->=+*{}|-:<>` dostają `hl-op`. |
