# Wymagania: pliki strefy DNS

Źródło: RFC 1035, 5 Master Files (zone file format). Definicja: `highlight/lang/dnszone.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1831 | Nazwy języka: `dnszone`, `zone`, `bind`. |
| LANG-1832 | Słowa z listy keyword (`$ORIGIN`, `$TTL`, `$INCLUDE`) dostają `hl-kw`; słowa z listy builtin (`IN`, `CH`, `HS`, `A`, `AAAA`, `CNAME` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-attr`. |
| LANG-1833 | Komentarze `;` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1834 | Liczby dostają `hl-num`; znaki `()` dostają `hl-op`. |
