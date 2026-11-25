# Wymagania: properties

Źródło: dokumentacja Javy, `java.util.Properties.load(Reader)`.

| Id | Wymaganie |
|---|---|
| LANG-471 | Nazwy języka: `properties`, `props`. |
| LANG-472 | Komentarz `#` albo `!` na początku wiersza dostaje `hl-com`. |
| LANG-473 | Klucz przed `=`, `:` albo pierwszym odstępem dostaje `hl-attr`, separator `hl-op`. |
| LANG-474 | W wartości `\` z następnym znakiem dostaje `hl-esc`, a `\` na końcu wiersza (przy nieparzystej liczbie ukośników) przedłuża wartość na następny wiersz; wartości jak w INI (`LANG-464`). |
