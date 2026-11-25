# Wymagania: lighttpd configuration

Źródło: lighttpd documentation: Configuration file syntax. Definicja: `highlight/lang/lighttpd.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1971 | Nazwy języka: `lighttpd`. |
| LANG-1972 | Słowa z listy keyword (`server.modules`, `server.document-root`, `server.port`, `server.username`, `server.errorlog`, `index-file.names` …) dostają `hl-kw`; słowa z listy builtin (`enable`, `disable`) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`. |
| LANG-1973 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1974 | Liczby dostają `hl-num`; znaki `=+(),>` dostają `hl-op`. |
