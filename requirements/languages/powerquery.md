# Wymagania: Power Query M

Źródło: Power Query M formula language specification (Microsoft). Definicja: `highlight/lang/powerquery.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2031 | Nazwy języka: `powerquery`, `pq`. |
| LANG-2032 | Słowa z listy keyword (`let`, `in`, `if`, `then`, `else`, `each` …) dostają `hl-kw`; słowa z listy type (`number`, `text`, `logical`, `date`, `datetime`, `duration` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `#table`, `#date`, `#datetime` …) dostają `hl-builtin`. |
| LANG-2033 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2034 | Liczby dostają `hl-num`; znaki `+-*/&=<>?.@` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
