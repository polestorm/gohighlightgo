# Wymagania: CLIPS

Źródło: CLIPS Reference Manual, Volume I: Basic Programming Guide (6.4). Definicja: `highlight/lang/clips.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2811 | Nazwy języka: `clips`, `clp`. |
| LANG-2812 | Słowa z listy keyword (`defrule`, `deftemplate`, `deffacts`, `deffunction`, `defglobal`, `defclass` …) dostają `hl-kw`; słowa z listy builtin (`TRUE`, `FALSE`, `nil`, `crlf`, `t`) dostają `hl-builtin`; nazwa po `?` dostaje `hl-attr`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-2813 | Komentarze `;` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2814 | Liczby dostają `hl-num`; znaki `()<=>&|~:` dostają `hl-op`. |
