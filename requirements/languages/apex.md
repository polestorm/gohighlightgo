# Wymagania: Apex

Źródło: Apex Developer Guide (Salesforce): słowa zastrzeżone, typy danych, adnotacje. Definicja: `highlight/lang/apex.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1241 | Nazwy języka: `apex`, `cls-apex`, `trigger`. |
| LANG-1242 | Słowa kluczowe (bez rozróżniania wielkości liter, także słowa DML i SOQL) dostają `hl-kw`; typy `hl-type`; `true false null System` `hl-builtin`; adnotacja `@Nazwa` `hl-meta`. |
| LANG-1243 | Komentarze `//` i `/* … */` dostają `hl-com`; napis `'…'` z ucieczkami `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
