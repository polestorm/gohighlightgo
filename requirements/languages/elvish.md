# Wymagania: Elvish

Źródło: Elvish language reference (elv.sh/ref/language.html). Definicja: `highlight/lang/elvish.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2511 | Nazwy języka: `elvish`, `elv`. |
| LANG-2512 | Słowa z listy keyword (`var`, `set`, `tmp`, `del`, `fn`, `if` …) dostają `hl-kw`; słowa z listy builtin (`echo`, `put`, `print`, `each`, `peach`, `range` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-2513 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2514 | Liczby dostają `hl-num`; znaki `|=<>!&{}[]()` dostają `hl-op`. |
