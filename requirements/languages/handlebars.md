# Wymagania: Handlebars i Mustache

Źródło: Handlebars Language Guide i specyfikacja Mustache (mustache(5)).

| Id | Wymaganie |
|---|---|
| LANG-921 | Nazwy języka: `handlebars`, `hbs`, `mustache`. |
| LANG-922 | Wstawki `{{ … }}` i `{{{ … }}}` są kolorowane językiem wyrażeń: pomocniki blokowe (`if`, `unless`, `each`, `with`, `else`) `hl-kw`, `this` `hl-builtin`, `@index` i inne `@nazwa` `hl-attr`, `#`, `/`, `^`, `>` `hl-op`, napisy `hl-str`. |
| LANG-923 | Komentarze `{{!-- … --}}` i `{{! … }}` dostają `hl-com`. |
| LANG-924 | Tekst poza wstawkami jest kolorowany jak HTML (`LANG-122`); ograniczniki wstawek dostają `hl-meta`; niezamknięta wstawka sięga końca danych. |
