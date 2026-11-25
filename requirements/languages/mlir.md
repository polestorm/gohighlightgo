# Wymagania: MLIR

Źródło: MLIR Language Reference (mlir.llvm.org/docs/LangRef). Definicja: `highlight/lang/mlir.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1991 | Nazwy języka: `mlir`. |
| LANG-1992 | Słowa z listy keyword (`func`, `module`, `return`, `br`, `cond_br`, `constant` …) dostają `hl-kw`; słowa z listy type (`i1`, `i8`, `i16`, `i32`, `i64`, `f16` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `unit`) dostają `hl-builtin`; nazwa po `%` dostaje `hl-attr`; nazwa po `@` dostaje `hl-fn`; nazwa po `^` dostaje `hl-meta`; nazwa po `#` dostaje `hl-meta`; nazwa po `!` dostaje `hl-type`. |
| LANG-1993 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1994 | Liczby dostają `hl-num`; znaki `=:->,<>x` dostają `hl-op`. |
