# Wymagania: Oz

Źródło: The Mozart Programming System documentation: The Oz Base Environment. Definicja: `highlight/lang/oz.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2761 | Nazwy języka: `oz`, `mozart`. |
| LANG-2762 | Słowa z listy keyword (`local`, `in`, `end`, `proc`, `fun`, `if` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `unit`, `nil`, `Browse`, `Show`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-2763 | Komentarze `%` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2764 | Liczby dostają `hl-num`; znaki `+-*/=<>:|#!.~@$` dostają `hl-op`. |
