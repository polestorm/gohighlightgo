# Wymagania: Isabelle

Źródło: The Isabelle/Isar Reference Manual (Isabelle2024). Definicja: `highlight/lang/isabelle.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1951 | Nazwy języka: `isabelle`, `thy`. |
| LANG-1952 | Słowa z listy keyword (`theory`, `imports`, `begin`, `end`, `definition`, `fun` …) dostają `hl-kw`; słowa z listy builtin (`auto`, `simp`, `blast`, `force`, `induction`, `cases` …) dostają `hl-builtin`. |
| LANG-1953 | Komentarze `(* … *)` (zagnieżdżany) dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `‹…›`; `` `…` `` dostają `hl-str`. |
| LANG-1954 | Liczby dostają `hl-num`; znaki `+-*/=<>:|!&?.` dostają `hl-op`. |
