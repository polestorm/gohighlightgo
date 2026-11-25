# Wymagania: UnrealScript

Źródło: UnrealScript Language Reference (Epic Games, UE3). Definicja: `highlight/lang/unrealscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2951 | Nazwy języka: `unrealscript`, `uc`. |
| LANG-2952 | Słowa z listy keyword (`class`, `extends`, `var`, `local`, `function`, `event` …) dostają `hl-kw`; słowa z listy type (`int`, `float`, `bool`, `byte`, `string`, `name` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `None`, `Log`, `Spawn`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2953 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` (nazwa, bez klasy) dostają `hl-str`. |
| LANG-2954 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.$@` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
