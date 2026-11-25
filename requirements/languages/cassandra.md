# Wymagania: Cassandra CQL

Źródło: Apache Cassandra documentation: The Cassandra Query Language (CQL 3). Definicja: `highlight/lang/cassandra.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2841 | Nazwy języka: `cassandra`, `cassandra-cql` (`cql` to Cypher). |
| LANG-2842 | Słowa z listy keyword (`SELECT`, `FROM`, `WHERE`, `AND`, `INSERT`, `INTO` …) dostają `hl-kw`; słowa z listy type (`ascii`, `bigint`, `blob`, `boolean`, `counter`, `date` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `now`, `uuid`, `toTimestamp`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2843 | Komentarze `--` do końca wiersza, `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `'…'` z podwojonym ogranicznikiem jako ucieczką; `$$…$$` wielowierszowy; `"…"` (nazwa, bez klasy) dostają `hl-str`. |
| LANG-2844 | Liczby dostają `hl-num`; znaki `=<>+-*/;,.(){}[]?:` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
