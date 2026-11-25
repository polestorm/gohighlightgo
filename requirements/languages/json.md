# Wymagania: JSON

Źródło: RFC 8259. Komentarze są tolerowane, bo bloki opisane jako `json` bywają w praktyce JSONC
(pliki ustawień edytorów, `tsconfig.json`).

| Id | Wymaganie |
|---|---|
| LANG-091 | Nazwy języka: `json`, `jsonc`. |
| LANG-092 | Napis dostaje `hl-str`, a napis, po którym (po spacjach i tabulatorach) stoi `:`, dostaje `hl-attr` (klucz obiektu); sekwencje ucieczki (`\` z jednym znakiem, `\u` z czterema cyframi) dostają `hl-esc`; napis nie przechodzi przez koniec wiersza (RFC 8259, 7). |
| LANG-093 | Liczba razem z poprzedzającym ją minusem dostaje `hl-num`; `true false null` dostają `hl-builtin`; nawiasy, przecinki i dwukropki zostają bez klasy (RFC 8259, 3, 6). |
| LANG-094 | Komentarze `//` i `/* … */` dostają `hl-com`. |
