# Wymagania: nginx

Źródło: nginx documentation: configuration file structure and core directives. Definicja: `highlight/lang/nginx.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1701 | Nazwy języka: `nginx`, `nginxconf`. |
| LANG-1702 | Słowa z listy keyword (`http`, `server`, `location`, `upstream`, `events`, `stream` …) dostają `hl-kw`; słowa z listy builtin (`on`, `off`) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-1703 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1704 | Liczby dostają `hl-num`; znaki `=~^*;{}` dostają `hl-op`. |
