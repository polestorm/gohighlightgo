# Wymagania: Caddyfile

Źródło: Caddy 2 documentation: The Caddyfile, directives. Definicja: `highlight/lang/caddyfile.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1861 | Nazwy języka: `caddyfile`, `caddy`. |
| LANG-1862 | Słowa z listy keyword (`reverse_proxy`, `file_server`, `root`, `encode`, `header`, `redir` …) dostają `hl-kw`; słowa z listy builtin (`gzip`, `zstd`, `on`, `off`) dostają `hl-builtin`; nazwa po `@` dostaje `hl-meta`. |
| LANG-1863 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `` `…` `` dostają `hl-str`. |
| LANG-1864 | Liczby dostają `hl-num`; znaki `{}*` dostają `hl-op`. |
