# Wymagania: Fennel

Źródło: Fennel reference (fennel-lang.org/reference). Definicja: `highlight/lang/fennel.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1621 | Nazwy języka: `fennel`, `fnl`. |
| LANG-1622 | Słowa z listy keyword (`fn`, `lambda`, `λ`, `let`, `local`, `var` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `nil`, `print`, `require`, `tostring` …) dostają `hl-builtin`; nazwa po `:` dostaje `hl-attr`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-1623 | Komentarze `;` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, wielowierszowy dostają `hl-str`. |
| LANG-1624 | Liczby dostają `hl-num`. |
