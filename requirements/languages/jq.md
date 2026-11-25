# Wymagania: jq

Źródło: jq manual (jqlang.github.io/jq/manual). Definicja: `highlight/lang/jq.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2461 | Nazwy języka: `jq`. |
| LANG-2462 | Słowa z listy keyword (`def`, `if`, `then`, `elif`, `else`, `end` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `null`, `length`, `keys`, `map` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`; nazwa po `@` dostaje `hl-builtin`. |
| LANG-2463 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2464 | Liczby dostają `hl-num`; znaki `|.,[]{}:=+-*/%<>!?` dostają `hl-op`. |
