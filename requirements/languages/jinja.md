# Wymagania: Jinja

Źródło: Jinja Template Designer Documentation (3.1); składnia wspólna z Django, Twig i Nunjucks.

| Id | Wymaganie |
|---|---|
| LANG-911 | Nazwy języka: `jinja`, `jinja2`, `j2`, `django`, `twig`, `nunjucks`, `njk`. |
| LANG-912 | Wstawki `{{ … }}` i `{% … %}` (także z `-` przy ogranicznikach) są kolorowane językiem wyrażeń: słowa kluczowe (`if`, `for`, `block`, `extends`, `macro` …) `hl-kw`, `true false none loop` `hl-builtin`, napisy `hl-str`, liczby `hl-num`, operatory i filtry po `|` `hl-op`, nazwa przed `(` `hl-fn`. |
| LANG-913 | Komentarz `{# … #}` dostaje `hl-com`. |
| LANG-914 | Tekst poza wstawkami jest kolorowany jak HTML (`LANG-122`); ograniczniki wstawek dostają `hl-meta`; niezamknięta wstawka sięga końca danych. |
