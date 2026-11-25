# Wymagania: Move

Źródło: The Move Book (Sui and Aptos dialects): language reference. Definicja: `highlight/lang/move.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2561 | Nazwy języka: `move`. |
| LANG-2562 | Słowa z listy keyword (`module`, `script`, `fun`, `public`, `entry`, `native` …) dostają `hl-kw`; słowa z listy type (`u8`, `u16`, `u32`, `u64`, `u128`, `u256` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `borrow_global`, `borrow_global_mut`, `exists`, `move_to` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-num`. |
| LANG-2563 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2564 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
