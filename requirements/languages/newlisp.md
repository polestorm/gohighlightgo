# Wymagania: newLISP

Źródło: newLISP Users Manual and Reference (10.7). Definicja: `highlight/lang/newlisp.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2531 | Nazwy języka: `newlisp`, `nl`. |
| LANG-2532 | Słowa z listy keyword (`define`, `fn`, `lambda`, `let`, `letn`, `local` …) dostają `hl-kw`; słowa z listy builtin (`true`, `nil`, `println`, `print`, `append`, `list` …) dostają `hl-builtin`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-2533 | Komentarze `;` do końca wiersza, `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `{…}` wielowierszowy; `[text]…[/text]` wielowierszowy dostają `hl-str`. |
| LANG-2534 | Liczby dostają `hl-num`. |
