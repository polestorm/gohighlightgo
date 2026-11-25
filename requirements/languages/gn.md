# Wymagania: GN

Źródło: GN reference (gn.googlesource.com): language and functions. Definicja: `highlight/lang/gn.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2201 | Nazwy języka: `gn`, `gni`. |
| LANG-2202 | Słowa z listy keyword (`if`, `else`, `foreach`, `import`, `template`, `declare_args`) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `executable`, `static_library`, `shared_library`, `source_set` …) dostają `hl-builtin`. |
| LANG-2203 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc` dostają `hl-str`. |
| LANG-2204 | Liczby dostają `hl-num`; znaki `+-=<>!&|.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
