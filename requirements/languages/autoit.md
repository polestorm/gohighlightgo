# Wymagania: AutoIt

Źródło: AutoIt v3 documentation, Language Reference. Definicja: `highlight/lang/autoit.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1261 | Nazwy języka: `autoit`, `au3`. |
| LANG-1262 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; funkcje i stałe z listy `hl-builtin`; zmienna `$nazwa` `hl-attr`; makro `@Nazwa` `hl-builtin`; dyrektywa `#include` `hl-meta`. |
| LANG-1263 | Komentarz `;` i blok `#cs` … `#ce` (także `#comments-start` … `#comments-end`) dostają `hl-com`; napisy `"…"` i `'…'` z podwojonym ogranicznikiem `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
