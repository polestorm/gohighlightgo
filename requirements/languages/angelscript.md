# Wymagania: AngelScript

Źródło: AngelScript manual: Script language reference. Definicja: `highlight/lang/angelscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2251 | Nazwy języka: `angelscript`, `as`. |
| LANG-2252 | Słowa z listy keyword (`and`, `abstract`, `auto`, `break`, `case`, `cast` …) dostają `hl-kw`; słowa z listy type (`void`, `bool`, `int`, `int8`, `int16`, `int32` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `print`) dostają `hl-builtin`. |
| LANG-2253 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"""…"""` wielowierszowy; `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2254 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.@` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
