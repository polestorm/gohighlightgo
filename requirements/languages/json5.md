# Wymagania: JSON5

Źródło: The JSON5 Data Interchange Format (1.0.0).

| Id | Wymaganie |
|---|---|
| LANG-981 | Nazwa języka: `json5`. |
| LANG-982 | Klucz w cudzysłowie albo bez niego, po którym stoi `:`, dostaje `hl-attr`; `true false null Infinity NaN` `hl-builtin`. |
| LANG-983 | Komentarze `//` i `/* … */` dostają `hl-com`. |
| LANG-984 | Napisy `"…"` i `'…'` z ucieczkami (także `\` przed końcem wiersza) dostają `hl-str`; liczby (także szesnastkowe i ze znakiem) `hl-num`. |
