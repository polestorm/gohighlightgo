# Wymagania: Sieve

Źródło: RFC 5228 (Sieve: An Email Filtering Language). Definicja: `highlight/lang/sieve.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2061 | Nazwy języka: `sieve`, `siv`. |
| LANG-2062 | Słowa z listy keyword (`require`, `if`, `elsif`, `else`, `stop`, `keep` …) dostają `hl-kw`; słowa z listy builtin (`:is`, `:contains`, `:matches`, `:over`, `:under`, `:all` …) dostają `hl-builtin`; nazwa po `:` dostaje `hl-attr`. |
| LANG-2063 | Komentarze `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `text:….` dostają `hl-str`. |
| LANG-2064 | Liczby dostają `hl-num`; znaki `[],;{}` dostają `hl-op`. |
