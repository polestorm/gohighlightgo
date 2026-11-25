# Wymagania: Gambas

Źródło: Gambas documentation: language reference. Definicja: `highlight/lang/gambas.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2781 | Nazwy języka: `gambas`. |
| LANG-2782 | Słowa z listy keyword (`Public`, `Private`, `Static`, `Dim`, `As`, `New` …) dostają `hl-kw`; słowa z listy type (`Integer`, `Long`, `Float`, `String`, `Boolean`, `Byte` …) dostają `hl-type`; słowa z listy builtin (`True`, `False`, `Null`, `Len`, `Str`, `Val` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2783 | Komentarze `'` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2784 | Liczby dostają `hl-num`; znaki `+-*/\^=<>&.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
