# Wymagania: Hjson

Źródło: Hjson syntax (hjson.github.io).

| Id | Wymaganie |
|---|---|
| LANG-991 | Nazwa języka: `hjson`. |
| LANG-992 | Klucz w cudzysłowie albo bez niego, po którym stoi `:`, dostaje `hl-attr`; `true false null` `hl-builtin`. |
| LANG-993 | Komentarze `#`, `//` i `/* … */` dostają `hl-com`. |
| LANG-994 | Napisy `"…"`, `'…'` i wielowierszowe `'''…'''` dostają `hl-str`; wartość bez cudzysłowu nie dostaje klasy; liczby `hl-num`. |
