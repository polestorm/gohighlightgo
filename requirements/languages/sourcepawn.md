# Wymagania: SourcePawn

Źródło: SourcePawn language reference (AlliedModders wiki). Definicja: `highlight/lang/sourcepawn.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1801 | Nazwy języka: `sourcepawn`, `sp`, `pawn`. |
| LANG-1802 | Słowa z listy keyword (`public`, `stock`, `static`, `const`, `native`, `forward` …) dostają `hl-kw`; słowa z listy type (`int`, `float`, `bool`, `char`, `void`, `any` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `INVALID_HANDLE`, `PrintToServer`, `PrintToChat`) dostają `hl-builtin`. |
| LANG-1803 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1804 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`; konstrukcje haka `c-preprocessor` jak w opisie haka w `docs/definitions.md`. |
