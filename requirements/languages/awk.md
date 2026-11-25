# Wymagania: AWK

Źródło: POSIX.1-2024, opis programu awk, i The GNU Awk User's Guide. Definicja: `highlight/lang/awk.hl`.

| Id | Wymaganie |
|---|---|
| LANG-691 | Nazwy języka: `awk`, `gawk`, `mawk`, `nawk`. |
| LANG-692 | Słowa kluczowe i wzorce `BEGIN`, `END` dostają `hl-kw`; funkcje i zmienne wbudowane (`length`, `gsub`, `NR`, `NF` …) `hl-builtin`. |
| LANG-693 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-694 | Napis `"…"` z ucieczkami i wyrażenie regularne `/…/` (rozpoznane jak w JavaScripcie) dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
