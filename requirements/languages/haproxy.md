# Wymagania: HAProxy configuration

Źródło: HAProxy Configuration Manual (2.8). Definicja: `highlight/lang/haproxy.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2481 | Nazwy języka: `haproxy`, `haproxycfg`. |
| LANG-2482 | Słowa z listy keyword (`global`, `defaults`, `frontend`, `backend`, `listen`, `bind` …) dostają `hl-kw`; słowa z listy builtin (`http`, `tcp`, `roundrobin`, `leastconn`, `check`, `ssl` …) dostają `hl-builtin`. |
| LANG-2483 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2484 | Liczby dostają `hl-num`; znaki `{}!` dostają `hl-op`. |
