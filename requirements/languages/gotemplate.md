# Wymagania: szablony Go

Źródło: dokumentacja pakietu text/template (Go 1.22).

| Id | Wymaganie |
|---|---|
| LANG-951 | Nazwy języka: `gotemplate`, `gotmpl`, `go-template`, `tmpl`, `helm`. |
| LANG-952 | Wstawka `{{ … }}` (także `{{-` i `-}}`) jest kolorowana językiem akcji: słowa (`if`, `range`, `with`, `define`, `template`, `block`, `end`) `hl-kw`, funkcje (`and`, `len`, `printf`, `eq` …) `hl-builtin`, zmienne `$x` i pola `.Nazwa` `hl-attr`, napisy `hl-str`, `|` i `:=` `hl-op`. |
| LANG-953 | Komentarz `{{/* … */}}` dostaje `hl-com`. |
| LANG-954 | Tekst poza wstawkami jest kolorowany jak HTML (`LANG-122`); ograniczniki wstawek dostają `hl-meta`; niezamknięta wstawka sięga końca danych. |
