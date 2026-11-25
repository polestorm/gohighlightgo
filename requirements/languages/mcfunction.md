# Wymagania: mcfunction

Źródło: Minecraft Java Edition command syntax (Minecraft Wiki: Commands). Definicja: `highlight/lang/mcfunction.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2861 | Nazwy języka: `mcfunction`, `mcf`. |
| LANG-2862 | Słowa z listy keyword (`execute`, `as`, `at`, `if`, `unless`, `run` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `survival`, `creative`) dostają `hl-builtin`; nazwa po `@` dostaje `hl-attr`. |
| LANG-2863 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2864 | Liczby dostają `hl-num`; znaki `~^=[]{},:.` dostają `hl-op`. |
