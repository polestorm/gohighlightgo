# Wymagania: Frink

Źródło: Frink documentation (frinklang.org). Definicja: `highlight/lang/frink.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2771 | Nazwy języka: `frink`, `frk`. |
| LANG-2772 | Słowa z listy keyword (`if`, `then`, `else`, `while`, `for`, `to` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `undef`, `meter`, `second`, `kilogram` …) dostają `hl-builtin`. |
| LANG-2773 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"""…"""` wielowierszowy; `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc` dostają `hl-str`. |
| LANG-2774 | Liczby dostają `hl-num`; znaki `+-*/^=<>:|!.->[]` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
