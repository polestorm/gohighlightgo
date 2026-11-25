# Wymagania: Liquid

Źródło: Liquid template language reference (Shopify).

| Id | Wymaganie |
|---|---|
| LANG-931 | Nazwa języka: `liquid`. |
| LANG-932 | Wstawki `{{ … }}` i `{% … %}` są kolorowane językiem wyrażeń: tagi (`if`, `for`, `assign`, `capture`, `render`, `comment` …) `hl-kw`, `true false nil empty blank forloop` `hl-builtin`, napisy `hl-str`, liczby `hl-num`, filtry po `|` `hl-op`. |
| LANG-933 | Tekst poza wstawkami jest kolorowany jak HTML (`LANG-122`); ograniczniki wstawek dostają `hl-meta`; niezamknięta wstawka sięga końca danych. |
