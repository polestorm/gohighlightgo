# Wymagania: SQL

Źródło: ISO/IEC 9075-2:2016 (słowa zastrzeżone, literały, komentarze) oraz słowa powszechne w dokumentacji
PostgreSQL, SQLite i MySQL (`LIMIT`, `INDEX`, `IF`, `AUTO_INCREMENT`, `TEXT`). Dialekty nie są rozróżniane.

| Id | Wymaganie |
|---|---|
| LANG-081 | Nazwa języka: `sql`. |
| LANG-082 | Słowa kluczowe dostają `hl-kw`, nazwy typów `hl-type`, a `TRUE FALSE NULL UNKNOWN` i funkcje agregujące i tekstowe z listy w `highlight/lang_sql.go` `hl-builtin`; wielkość liter nie ma znaczenia. |
| LANG-083 | Komentarz `--` sięga końca wiersza; komentarz `/* … */` się zagnieżdża; oba dostają `hl-com`. |
| LANG-084 | Napis `'…'` dostaje `hl-str` i może obejmować wiele wierszy; podwojony apostrof `''` w środku dostaje `hl-esc`; przedrostki `N`, `X`, `B`, `E`, `U&` (w dowolnej wielkości liter) należą do napisu. Nazwy w `"…"` i `` `…` `` nie dostają klasy, a słowa kluczowe w nich nie są kolorowane. |
| LANG-085 | Liczby jak w Go (`LANG-007`); operatory `+ - * / % < > = ! \| & ^ ~ :` dostają `hl-op`. Identyfikator przed `(` nie dostaje `hl-fn`: w SQL to najczęściej nazwa tabeli (`INSERT INTO t (a, b)`, `CREATE TABLE t (…)`), a funkcje wbudowane mają `hl-builtin`. |
