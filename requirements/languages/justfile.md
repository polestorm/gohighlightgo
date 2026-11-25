# Wymagania: just

Źródło: just programmer's manual (just.systems/man). Definicja: `highlight/lang/justfile.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2141 | Nazwy języka: `justfile`, `just`. |
| LANG-2142 | Słowa z listy keyword (`set`, `export`, `alias`, `import`, `mod`, `if` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `os`, `arch`, `env_var`, `justfile_directory` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-meta`. |
| LANG-2143 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"""…"""` z ucieczkami `hl-esc`, wielowierszowy; `"…"` z ucieczkami `hl-esc`; `'…'` dostają `hl-str`. |
| LANG-2144 | Liczby dostają `hl-num`; znaki `:=+/&|!` dostają `hl-op`; konstrukcje haka `variables` jak w opisie haka w `docs/definitions.md`. |
