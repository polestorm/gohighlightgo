# Wymagania: Clipper / xBase

Źródło: CA-Clipper 5.3 Reference Guide and Harbour documentation. Definicja: `highlight/lang/clipper.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2821 | Nazwy języka: `clipper`, `harbour`, `xbase`, `prg`, `foxpro`. |
| LANG-2822 | Słowa z listy keyword (`FUNCTION`, `PROCEDURE`, `RETURN`, `LOCAL`, `STATIC`, `PRIVATE` …) dostają `hl-kw`; słowa z listy builtin (`.T.`, `.F.`, `.AND.`, `.OR.`, `.NOT.`, `NIL` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2823 | Komentarze `//` do końca wiersza, `&&` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"`; `'…'`; `[…]` dostają `hl-str`. |
| LANG-2824 | Liczby dostają `hl-num`; znaki `+-*/^%=<>:$!@` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`; konstrukcje haka `c-preprocessor` jak w opisie haka w `docs/definitions.md`. |
