# Wymagania: Cairo

Źródło: The Cairo Programming Language book (Starknet, 2.x). Definicja: `highlight/lang/cairo.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2551 | Nazwy języka: `cairo`. |
| LANG-2552 | Słowa z listy keyword (`fn`, `let`, `mut`, `const`, `if`, `else` …) dostają `hl-kw`; słowa z listy type (`felt252`, `u8`, `u16`, `u32`, `u64`, `u128` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `None`, `Some`, `Ok`, `Err` …) dostają `hl-builtin`; nazwa po `#` dostaje `hl-meta`. |
| LANG-2553 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2554 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`; konstrukcje haka `rust` jak w opisie haka w `docs/definitions.md`. |
