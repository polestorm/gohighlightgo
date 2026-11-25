# Wymagania: Logo

Źródło: UCBLogo (Berkeley Logo) user manual. Definicja: `highlight/lang/logo.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2301 | Nazwy języka: `logo`, `ucblogo`. |
| LANG-2302 | Słowa z listy keyword (`to`, `end`, `repeat`, `if`, `ifelse`, `make` …) dostają `hl-kw`; słowa z listy builtin (`forward`, `fd`, `back`, `bk`, `left`, `lt` …) dostają `hl-builtin`; nazwa po `:` dostaje `hl-attr`; nazwa po `"` dostaje `hl-str` (bez rozróżniania wielkości liter). |
| LANG-2303 | Komentarze `;` do końca wiersza dostają `hl-com`. |
| LANG-2304 | Liczby dostają `hl-num`; znaki `+-*/=<>[]()` dostają `hl-op`. |
