# Wymagania: Apache HTTP Server configuration

Źródło: Apache HTTP Server 2.4 documentation: configuration files and directive index. Definicja: `highlight/lang/apache.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1851 | Nazwy języka: `apache`, `apacheconf`, `htaccess`. |
| LANG-1852 | Słowa z listy keyword (`ServerRoot`, `ServerName`, `ServerAdmin`, `Listen`, `DocumentRoot`, `DirectoryIndex` …) dostają `hl-kw`; słowa z listy builtin (`On`, `Off`, `all`, `granted`, `denied`, `None` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-1853 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1854 | Liczby dostają `hl-num`; znaki `<>/` dostają `hl-op`; konstrukcje haka `variables` jak w opisie haka w `docs/definitions.md`. |
