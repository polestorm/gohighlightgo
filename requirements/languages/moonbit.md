# Wymagania: MoonBit

Źródło: MoonBit language documentation (docs.moonbitlang.com). Definicja: `highlight/lang/moonbit.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2541 | Nazwy języka: `moonbit`, `mbt`. |
| LANG-2542 | Słowa z listy keyword (`fn`, `let`, `mut`, `if`, `else`, `match` …) dostają `hl-kw`; słowa z listy type (`Int`, `Int64`, `UInt`, `Double`, `Float`, `Bool` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `None`, `Some`, `Ok`, `Err` …) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-2543 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2544 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
