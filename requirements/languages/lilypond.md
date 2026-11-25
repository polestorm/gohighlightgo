# Wymagania: LilyPond

Źródło: LilyPond Notation Reference (2.24). Definicja: `highlight/lang/lilypond.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2701 | Nazwy języka: `lilypond`, `ly`. |
| LANG-2702 | Słowa z listy builtin (`major`, `minor`) dostają `hl-builtin`; nazwa po `\` dostaje `hl-kw`; nazwa po `#` dostaje `hl-meta`. |
| LANG-2703 | Komentarze `%` do końca wiersza, `%{ … %}` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2704 | Liczby dostają `hl-num`; znaki `{}<>~|'.,=()[]-_^` dostają `hl-op`. |
