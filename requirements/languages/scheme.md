# Wymagania: Scheme

Źródło: R7RS small, 2 „Lexical conventions” i 4 „Expressions”. Definicja: `highlight/lang/scheme.hl`.

| Id | Wymaganie |
|---|---|
| LANG-611 | Nazwy języka: `scheme`, `scm`, `ss`. |
| LANG-612 | Formy składniowe z listy (`define`, `lambda`, `let`, `if`, `cond` …) dostają `hl-kw`; procedury z listy `hl-builtin`; `#t`, `#f`, `#true`, `#false` `hl-builtin`; nazwa może zawierać `! ? * < > = / + - . :`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-613 | Komentarz `;` i zagnieżdżany `#| … |#` dostają `hl-com`. |
| LANG-614 | Napis `"…"` i znak `#\a`, `#\space` dostają `hl-str`; cytat `'nazwa` `hl-meta`; liczby `hl-num`. |
