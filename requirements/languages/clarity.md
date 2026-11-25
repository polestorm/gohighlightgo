# Wymagania: Clarity

Źródło: Clarity language reference (docs.stacks.co). Definicja: `highlight/lang/clarity.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2581 | Nazwy języka: `clarity`, `clar`. |
| LANG-2582 | Słowa z listy keyword (`define-public`, `define-private`, `define-read-only`, `define-constant`, `define-data-var`, `define-map` …) dostają `hl-kw`; słowa z listy type (`uint`, `int`, `bool`, `principal`, `buff`, `string-ascii` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `none`, `some`, `ok`, `err` …) dostają `hl-builtin`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-2583 | Komentarze `;;` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2584 | Liczby dostają `hl-num`. |
