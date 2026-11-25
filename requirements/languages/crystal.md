# Wymagania: Crystal

Źródło: Crystal language reference (1.11): „Syntax and semantics”, „Literals”. Definicja: `highlight/lang/crystal.hl`.

| Id | Wymaganie |
|---|---|
| LANG-571 | Nazwy języka: `crystal`, `cr`. |
| LANG-572 | Słowa kluczowe dostają `hl-kw`; nazwy z wielkiej litery (typy jak `Int32`, stałe) `hl-type`; `true false nil` i makra z listy `hl-builtin`; `@zmienna` i `$zmienna` `hl-attr`. |
| LANG-573 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-574 | Napisy, literały `%w[…]`, dokumenty `<<-KONIEC`, symbole `:nazwa`, klucze `nazwa:` i wyrażenia regularne jak w Rubym (`LANG-274`, `LANG-275`). |
