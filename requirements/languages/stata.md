# Wymagania: Stata

Źródło: Stata 18 manuals: [P] Programming, [U] Language syntax. Definicja: `highlight/lang/stata.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1501 | Nazwy języka: `stata`, `do`, `ado`. |
| LANG-1502 | Słowa programowania (`foreach`, `program`, `local`, `capture` …) dostają `hl-kw`; polecenia z listy (`generate`, `regress`, `tabulate` …) `hl-builtin`; makro globalne `$nazwa` `hl-attr`. |
| LANG-1503 | Komentarze `//` i `/* … */` dostają `hl-com`; napis `"…"` `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
