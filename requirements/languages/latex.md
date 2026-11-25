# Wymagania: LaTeX

Źródło: The TeXbook (konwencje leksykalne) i LaTeX2e: An unofficial reference manual.

| Id | Wymaganie |
|---|---|
| LANG-521 | Nazwy języka: `latex`, `tex`, `sty`, `cls`. |
| LANG-522 | Polecenie `\nazwa` dostaje `hl-kw`; `\` z jednym znakiem spoza liter (`\%`, `\\`, `\{`) dostaje `hl-esc` i nie zaczyna komentarza. |
| LANG-523 | Komentarz `%` do końca wiersza dostaje `hl-com`. |
| LANG-524 | Matematyka `$…$` i `$$…$$` dostaje `hl-str`; `& ^ _ ~` dostają `hl-op`. |
