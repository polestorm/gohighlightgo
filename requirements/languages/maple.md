# Wymagania: Maple

Źródło: Maple Programming Guide (Maplesoft). Definicja: `highlight/lang/maple.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2711 | Nazwy języka: `maple`, `mpl`. |
| LANG-2712 | Słowa z listy keyword (`proc`, `end`, `if`, `then`, `elif`, `else` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `FAIL`, `Pi`, `I`, `infinity` …) dostają `hl-builtin`. |
| LANG-2713 | Komentarze `#` do końca wiersza, `(* … *)` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `` `…` `` dostają `hl-str`. |
| LANG-2714 | Liczby dostają `hl-num`; znaki `+-*/^=<>:;$@!.,|` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
