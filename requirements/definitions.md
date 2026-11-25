# Wymagania: definicje języków

Decyzje: `docs/adr/0008-language-definitions.md`; opis formatu: `docs/definitions.md`.

| Id | Wymaganie |
|---|---|
| DEF-001 | Język z tabeli opisuje plik tekstowy: jedna dyrektywa w wierszu (`name`, `keyword`, `type`, `builtin`, `line-comment`, `block-comment`, `string`, `prefix`, `escape`, `number`, `operators`, `ident-extra`, `fold-case`, `calls`, `macros`, `decorators`, `templates`, `keys`, `regex`, `hook`, `extends`), wartości oddzielone odstępami, opcje po samotnym `:`, wiersze od `#` to komentarze. Wbudowane języki z tabeli są zapisane w tym formacie w `highlight/lang/*.hl`. |
| DEF-002 | Definicja błędna (nieznana dyrektywa albo opcja, zła liczba wartości, pusty ogranicznik, nazwa nie małymi literami, brak nazwy, słowo dłuższe niż 32 bajty, przedrostek dłuższy niż 8, plik dłuższy niż 64 KB) jest odrzucana błędem `*DefinitionError` z numerem wiersza. |
| DEF-003 | `hook` włącza tylko konstrukcje napisane w Go i nazwane w bibliotece; `extends` kopiuje tabelę wcześniej znanego języka z tabeli i nie zmienia jej. |
| DEF-004 | Dla każdej definicji przyjętej przez parser wynik spełnia `SAFE-001`, `SAFE-002` i `DET-003` (fuzzing definicji i kodu); format nie zawiera wyrażeń regularnych, więc czas pozostaje liniowy (`DET-002`). |
