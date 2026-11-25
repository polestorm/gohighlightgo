# Wymagania: T-SQL

Źródło: Transact-SQL reference (SQL Server 2022): słowa zastrzeżone, typy danych, zmienne. Definicja: `highlight/lang/tsql.hl`.

| Id | Wymaganie |
|---|---|
| LANG-841 | Nazwy języka: `tsql`, `mssql`, `t-sql`. |
| LANG-842 | Wszystko jak w SQL (`LANG-082`–`LANG-085`), a dodatkowo słowa T-SQL (`GO`, `TOP`, `EXEC`, `TRY`, `CATCH` …) dostają `hl-kw`, typy (`NVARCHAR`, `DATETIME2`, `UNIQUEIDENTIFIER` …) `hl-type`, funkcje z listy `hl-builtin`, zmienna `@nazwa` `hl-attr`, a nazwa w `[…]` nie dostaje klasy. |
