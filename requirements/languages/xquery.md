# Wymagania: XQuery

Źródło: XQuery 3.1 (W3C Recommendation 2017), A.2 „Lexical structure”. Definicja: `highlight/lang/xquery.hl`.

| Id | Wymaganie |
|---|---|
| LANG-801 | Nazwy języka: `xquery`, `xq`, `xqy`. |
| LANG-802 | Słowa kluczowe dostają `hl-kw`; `true false` i funkcje z listy (także z przedrostkiem `fn:`) `hl-builtin`; nazwa może zawierać `-`, `.`, `:`. |
| LANG-803 | Komentarz `(: … :)`, który się zagnieżdża, dostaje `hl-com`. |
| LANG-804 | Napisy `"…"` i `'…'` z podwojonym ogranicznikiem dostają `hl-str`; zmienna `$nazwa` `hl-attr`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
