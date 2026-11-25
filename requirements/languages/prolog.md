# Wymagania: Prolog

Źródło: ISO/IEC 13211-1 (Prolog: General core) i SWI-Prolog Reference Manual. Definicja: `highlight/lang/prolog.hl`.

| Id | Wymaganie |
|---|---|
| LANG-601 | Nazwa języka: `prolog`. |
| LANG-602 | Operatory słowne (`is`, `mod`, `rem` …) dostają `hl-kw`; predykaty wbudowane z listy `hl-builtin`; zmienna (nazwa z wielkiej litery) `hl-type`. |
| LANG-603 | Komentarz `%` i `/* … */` dostają `hl-com`. |
| LANG-604 | Napis `"…"` i `` `…` `` dostają `hl-str`; atom w apostrofach `'…'` nie dostaje klasy; operatory (`:-`, `->`, `=..`) `hl-op`; nazwa przed `(` `hl-fn`. |
