# Wymagania: Xorg configuration

Źródło: xorg.conf(5) manual page. Definicja: `highlight/lang/xorg.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2881 | Nazwy języka: `xorg`, `xorgconf`. |
| LANG-2882 | Słowa z listy keyword (`Section`, `EndSection`, `SubSection`, `EndSubSection`, `Identifier`, `Driver` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `on`, `off`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2883 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2884 | Liczby dostają `hl-num`. |
