# Wymagania: Fortran

Źródło: ISO/IEC 1539-1:2018 (Fortran 2018), rozdział 6 „Lexical tokens and source form”. Definicja: `highlight/lang/fortran.hl`.

| Id | Wymaganie |
|---|---|
| LANG-421 | Nazwy języka: `fortran`, `f90`, `f95`, `f03`, `f08`, `for`. |
| LANG-422 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; typy wbudowane `hl-type`; funkcje wbudowane z listy `hl-builtin`. |
| LANG-423 | Komentarz `!` do końca wiersza dostaje `hl-com`. |
| LANG-424 | Napisy `"…"` i `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-425 | Operatory w kropkach (`.and.`, `.eq.`, `.not.`) dostają `hl-op`, a `.true.` i `.false.` `hl-builtin`; liczby `hl-num`; nazwa przed `(` dostaje `hl-fn`. |
