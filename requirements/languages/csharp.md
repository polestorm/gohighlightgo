# Wymagania: C#

Źródło: C# language specification (ECMA-334) i dokumentacja języka C# 12 („Keywords”, „Strings”, „Raw string
literals”, „Preprocessor directives”). Definicja: `highlight/lang/csharp.hl`.

| Id | Wymaganie |
|---|---|
| LANG-191 | Nazwy języka: `csharp`, `cs`, `c#`. |
| LANG-192 | Słowa kluczowe i słowa kontekstowe `async await var yield record init when where nameof required partial with` dostają `hl-kw`; typy wbudowane `hl-type`; `true false null` dostają `hl-builtin`. |
| LANG-193 | Komentarze `//` i `/* … */`; dyrektywa preprocesora `#` na początku wiersza z nazwą dostaje `hl-meta`, jak w C (`LANG-016`). |
| LANG-194 | Napisy `"…"` i znaki `'…'` z ucieczkami (`\x` z czterema cyframi, `\u`, `\U`) dostają `hl-str` i `hl-esc`; napis dosłowny `@"…"` (bez ucieczek, `""` jako `hl-esc`, wiele wierszy), napis z wstawkami `$"…{x}…"` (wstawka i `{{` jako `hl-esc`), oba razem `$@"…"`, oraz napis surowy `"""…"""` (wiele wierszy). |
| LANG-195 | Liczby z przyrostkami (`m`, `f`, `d`, `L`, `u`) dostają `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
