# Wymagania: fish

Źródło: dokumentacja fish shell (3.7): „The fish language” i polecenia wbudowane. Definicja: `highlight/lang/fish.hl`.

| Id | Wymaganie |
|---|---|
| LANG-771 | Nazwa języka: `fish`. |
| LANG-772 | Słowa kluczowe dostają `hl-kw`; polecenia wbudowane z listy `hl-builtin`; zmienna `$nazwa` `hl-attr`. |
| LANG-773 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-774 | Napis `"…"` ze wstawkami `$x` i `'…'` dostają `hl-str`; operatory `| & ; < > = !` `hl-op`. |
