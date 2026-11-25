# Wymagania: TTCN-3

Źródło: ETSI ES 201 873-1 (TTCN-3 Core Language). Definicja: `highlight/lang/ttcn3.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2751 | Nazwy języka: `ttcn3`, `ttcn`. |
| LANG-2752 | Słowa z listy keyword (`module`, `import`, `from`, `all`, `type`, `record` …) dostają `hl-kw`; słowa z listy type (`integer`, `float`, `boolean`, `charstring`, `universal`, `octetstring` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `pass`, `fail`, `inconc`, `none` …) dostają `hl-builtin`. |
| LANG-2753 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2754 | Liczby dostają `hl-num`; znaki `+-*/:=<>&!.?@` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
