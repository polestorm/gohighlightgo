# Wymagania: Nushell

Źródło: The Nushell book and command reference (nushell.sh). Definicja: `highlight/lang/nushell.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2501 | Nazwy języka: `nushell`, `nu`. |
| LANG-2502 | Słowa z listy keyword (`def`, `let`, `mut`, `const`, `if`, `else` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `null`, `echo`, `print`, `ls` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-2503 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'`; `` `…` `` dostają `hl-str`. |
| LANG-2504 | Liczby dostają `hl-num`; znaki `|=<>!+-*/.:` dostają `hl-op`. |
