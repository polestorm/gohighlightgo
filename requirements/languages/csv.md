# Wymagania: CSV

Źródło: RFC 4180.

| Id | Wymaganie |
|---|---|
| LANG-961 | Nazwa języka: `csv`. |
| LANG-962 | Pierwszy wiersz (nagłówek) dostaje `hl-attr`; separatory `,` i `;` dostają `hl-op`. |
| LANG-963 | Pole w cudzysłowie (z `""` w środku, także przez koniec wiersza) dostaje `hl-str`; pole z liczbą `hl-num`. |
