# Wymagania: Varnish VCL

Źródło: Varnish Cache documentation: VCL syntax. Definicja: `highlight/lang/vcl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2471 | Nazwy języka: `vcl`, `varnish`. |
| LANG-2472 | Słowa z listy keyword (`vcl`, `backend`, `sub`, `if`, `else`, `elsif` …) dostają `hl-kw`; słowa z listy builtin (`req`, `resp`, `bereq`, `beresp`, `obj`, `client` …) dostają `hl-builtin`. |
| LANG-2473 | Komentarze `#` do końca wiersza, `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `{"…"}` wielowierszowy; `"…"` dostają `hl-str`. |
| LANG-2474 | Liczby dostają `hl-num`; znaki `=~!&|<>+-*/{}();,` dostają `hl-op`. |
