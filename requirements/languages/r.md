# Wymagania: R

Źródło: R Language Definition (R 4.x), 10.3 „Tokens”. Definicja: `highlight/lang/r.hl`.

| Id | Wymaganie |
|---|---|
| LANG-251 | Nazwa języka: `r`. |
| LANG-252 | Słowa zastrzeżone `if else repeat while function for in next break` dostają `hl-kw`; stałe `TRUE FALSE NULL Inf NaN NA …` i funkcje z listy dostają `hl-builtin`; nazwa może zawierać kropkę (`data.frame`, `is.na`). |
| LANG-253 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-254 | Napisy `"…"` i `'…'` z ucieczkami (`\x`, cyfry ósemkowe, `\u`, `\U`, `\u{…}`) dostają `hl-str` i `hl-esc`; nazwa w `` `…` `` nie dostaje klasy. |
| LANG-255 | Liczby z przyrostkami `L` i `i` dostają `hl-num`; operatory (także `<-`, `%in%`, `$`, `@`) dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
