# Wymagania: AMPL

Źródło: AMPL: A Modeling Language for Mathematical Programming (Fourer, Gay, Kernighan), appendix A. Definicja: `highlight/lang/ampl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2651 | Nazwy języka: `ampl`. |
| LANG-2652 | Słowa z listy keyword (`set`, `param`, `var`, `maximize`, `minimize`, `subject` …) dostają `hl-kw`; słowa z listy builtin (`card`, `max`, `min`, `abs`, `Infinity`) dostają `hl-builtin`. |
| LANG-2653 | Komentarze `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką; `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2654 | Liczby dostają `hl-num`; znaki `+-*/^=<>:;,{}[]().` dostają `hl-op`. |
