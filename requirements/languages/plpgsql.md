# Wymagania: PL/pgSQL

Źródło: PostgreSQL 16 documentation, chapter 43 PL/pgSQL. Definicja: `highlight/lang/plpgsql.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1721 | Nazwy języka: `plpgsql`, `pgsql`, `postgresql`, `postgres`, `psql`. |
| LANG-1722 | Wszystko jak w SQL (`LANG-082`–`LANG-085`), a dodatkowo słowa z listy keyword (`DECLARE`, `BEGIN`, `EXCEPTION`, `LOOP`, `EXIT`, `RAISE` …) dostają `hl-kw`; słowa z listy type (`SERIAL`, `BIGSERIAL`, `JSONB`, `JSON`, `UUID`, `BYTEA` …) dostają `hl-type`; słowa z listy builtin (`NOW`, `COALESCE`, `ARRAY_AGG`, `STRING_AGG`) dostają `hl-builtin`. |
