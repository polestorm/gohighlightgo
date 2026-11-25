# Wymagania: Standard ML

Źródło: The Definition of Standard ML (Revised), chapter 2 Syntax of the Core. Definicja: `highlight/lang/sml.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1751 | Nazwy języka: `sml`, `standardml`, `ml-sml`. |
| LANG-1752 | Słowa z listy keyword (`abstype`, `and`, `andalso`, `as`, `case`, `datatype` …) dostają `hl-kw`; słowa z listy type (`int`, `real`, `string`, `char`, `bool`, `unit` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `nil`, `NONE`, `SOME`, `print`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-1753 | Komentarze `(* … *)` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1754 | Liczby dostają `hl-num`; znaki `+-*/=<>:^@!~|` dostają `hl-op`; konstrukcje haka `apostrophe` jak w opisie haka w `docs/definitions.md`. |
