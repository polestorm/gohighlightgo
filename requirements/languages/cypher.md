# Wymagania: Cypher

Źródło: openCypher 9 i Neo4j Cypher Manual (5). Definicja: `highlight/lang/cypher.hl`.

| Id | Wymaganie |
|---|---|
| LANG-821 | Nazwy języka: `cypher`, `cql`. |
| LANG-822 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; `true false null` i funkcje z listy `hl-builtin`. |
| LANG-823 | Komentarze `//` i `/* … */` dostają `hl-com`. |
| LANG-824 | Napisy `"…"` i `'…'` dostają `hl-str`; etykieta i typ relacji `:Nazwa` `hl-attr`; parametr `$nazwa` `hl-meta`; nazwa w `` `…` `` bez klasy; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
