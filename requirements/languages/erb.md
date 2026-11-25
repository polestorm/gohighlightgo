# Wymagania: ERB

Źródło: dokumentacja ERB (Ruby standard library).

| Id | Wymaganie |
|---|---|
| LANG-941 | Nazwy języka: `erb`, `rhtml`, `eruby`. |
| LANG-942 | Wstawki `<% … %>`, `<%= … %>` i `<%- … %>` są kolorowane jak Ruby (`LANG-272`–`LANG-275`). |
| LANG-943 | Komentarz `<%# … %>` dostaje `hl-com`. |
| LANG-944 | Tekst poza wstawkami jest kolorowany jak HTML (`LANG-122`); ograniczniki wstawek dostają `hl-meta`; niezamknięta wstawka sięga końca danych. |
