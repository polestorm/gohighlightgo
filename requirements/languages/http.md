# Wymagania: HTTP

Źródło: RFC 9110 (HTTP Semantics) and RFC 9112 (HTTP/1.1): message syntax. Definicja: `highlight/lang/http.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1841 | Nazwy języka: `http`, `https`. |
| LANG-1842 | Słowa z listy keyword (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`, `HEAD` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `null`) dostają `hl-builtin`; nazwa przed `:` dostaje `hl-attr`. |
| LANG-1843 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1844 | Liczby dostają `hl-num`; znaki `/:=&?` dostają `hl-op`. |
