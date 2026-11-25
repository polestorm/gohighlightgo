# Wymagania: YANG

Źródło: RFC 7950 (The YANG 1.1 Data Modeling Language), 6 Lexical Tokenization. Definicja: `highlight/lang/yang.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1791 | Nazwy języka: `yang`. |
| LANG-1792 | Słowa z listy keyword (`module`, `submodule`, `namespace`, `prefix`, `import`, `include` …) dostają `hl-kw`; słowa z listy type (`string`, `boolean`, `int8`, `int16`, `int32`, `int64` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`) dostają `hl-builtin`. |
| LANG-1793 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, wielowierszowy; `'…'` wielowierszowy dostają `hl-str`. |
| LANG-1794 | Liczby dostają `hl-num`; znaki `;{}+` dostają `hl-op`. |
