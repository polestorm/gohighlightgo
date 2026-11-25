# Wymagania: Carbon

Źródło: Carbon language design documents (carbon-lang). Definicja: `highlight/lang/carbon.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2211 | Nazwy języka: `carbon`. |
| LANG-2212 | Słowa z listy keyword (`package`, `library`, `import`, `api`, `fn`, `var` …) dostają `hl-kw`; słowa z listy type (`i8`, `i16`, `i32`, `i64`, `u8`, `u16` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `Print`) dostają `hl-builtin`. |
| LANG-2213 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `'''…'''` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2214 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
