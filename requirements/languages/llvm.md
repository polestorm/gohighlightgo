# Wymagania: LLVM IR

Źródło: LLVM Language Reference Manual (llvm.org/docs/LangRef.html). Definicja: `highlight/lang/llvm.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1681 | Nazwy języka: `llvm`, `ll`, `llvm-ir`. |
| LANG-1682 | Słowa z listy keyword (`define`, `declare`, `global`, `constant`, `private`, `internal` …) dostają `hl-kw`; słowa z listy type (`void`, `i1`, `i8`, `i16`, `i32`, `i64` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `undef`, `poison`, `zeroinitializer`) dostają `hl-builtin`; nazwa po `%` dostaje `hl-attr`; nazwa po `@` dostaje `hl-fn`; nazwa po `!` dostaje `hl-meta`; nazwa po `#` dostaje `hl-meta`. |
| LANG-1683 | Komentarze `;` do końca wiersza dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-1684 | Liczby dostają `hl-num`; znaki `=*,` dostają `hl-op`. |
