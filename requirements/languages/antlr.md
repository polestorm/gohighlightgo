# Wymagania: ANTLR 4

Źródło: ANTLR 4 documentation: grammar structure, lexer and parser rules. Definicja: `highlight/lang/antlr.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1551 | Nazwy języka: `antlr`, `antlr4`, `g4`. |
| LANG-1552 | Słowa z listy keyword (`grammar`, `lexer`, `parser`, `fragment`, `import`, `options` …) dostają `hl-kw`; słowa z listy builtin (`EOF`, `skip`, `more`, `pushMode`, `popMode`, `channel` …) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`; nazwa po `@` dostaje `hl-meta`. |
| LANG-1553 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1554 | Liczby dostają `hl-num`; znaki `:;|*+?~=>()-.#` dostają `hl-op`. |
