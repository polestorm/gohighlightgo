# Wymagania: diff

Źródło: GNU Diffutils Manual („Unified Format”, „Normal Format”) i dokumentacja `git diff` („Generating patch
text with -p”).

| Id | Wymaganie |
|---|---|
| LANG-171 | Nazwy języka: `diff`, `patch`. |
| LANG-172 | Nagłówek fragmentu `@@ -a,b +c,d @@` dostaje `hl-meta` (do drugiego `@@`; reszta wiersza bez klasy); wewnątrz fragmentu, dopóki nie wyczerpią się liczby wierszy z nagłówka, wiersz `+` dostaje `hl-ins`, `-` dostaje `hl-del`, a wiersz zaczynający się od `+++` albo `---` jest treścią, nie nagłówkiem. |
| LANG-173 | Poza fragmentem wiersze `+++ `, `--- `, `diff `, `index `, `new file`, `deleted file`, `rename `, `similarity `, `old mode`, `new mode`, `Binary files`, `Only in ` dostają `hl-meta`; wiersze `>` i `<` formatu zwykłego dostają `hl-ins` i `hl-del`; `\ No newline at end of file` dostaje `hl-com`. |
