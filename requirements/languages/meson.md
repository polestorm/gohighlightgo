# Wymagania: Meson

Źródło: Meson build system reference manual: syntax and functions. Definicja: `highlight/lang/meson.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2151 | Nazwy języka: `meson`. |
| LANG-2152 | Słowa z listy keyword (`if`, `elif`, `else`, `endif`, `foreach`, `endforeach` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `project`, `executable`, `library`, `shared_library` …) dostają `hl-builtin`. |
| LANG-2153 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `'''…'''` z ucieczkami `hl-esc`, wielowierszowy; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2154 | Liczby dostają `hl-num`; znaki `+-*/%=<>!:?.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
