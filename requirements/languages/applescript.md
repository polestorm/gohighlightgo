# Wymagania: AppleScript

Źródło: AppleScript Language Guide (Apple), Lexical Conventions. Definicja: `highlight/lang/applescript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1561 | Nazwy języka: `applescript`, `osascript`. |
| LANG-1562 | Słowa z listy keyword (`tell`, `end`, `if`, `then`, `else`, `repeat` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `missing`, `value`, `me`, `it` …) dostają `hl-builtin`. |
| LANG-1563 | Komentarze `--` do końca wiersza, `#` do końca wiersza, `(* … *)` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1564 | Liczby dostają `hl-num`; znaki `&=<>+-*/^` dostają `hl-op`. |
