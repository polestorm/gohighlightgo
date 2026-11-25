# Wymagania: Puppet

Źródło: Puppet language reference (8.x): słowa zastrzeżone, typy zasobów, zmienne, napisy.

| Id | Wymaganie |
|---|---|
| LANG-1061 | Nazwy języka: `puppet`, `pp`. |
| LANG-1062 | Słowa zastrzeżone dostają `hl-kw`; typy zasobów i typy danych z listy `hl-type`; `true false undef` i funkcje z listy `hl-builtin`; zmienna `$nazwa` `hl-attr`. |
| LANG-1063 | Komentarze `#` i `/* … */` dostają `hl-com`. |
| LANG-1064 | Napis `"…"` ze wstawkami `$x`, `${…}` i `'…'` dostają `hl-str`; liczby `hl-num`; `=>` i inne operatory `hl-op`; nazwa przed `(` `hl-fn`. |
