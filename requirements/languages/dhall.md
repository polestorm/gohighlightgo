# Wymagania: Dhall

Źródło: Dhall language standard (gramatyka dhall.abnf).

| Id | Wymaganie |
|---|---|
| LANG-1101 | Nazwa języka: `dhall`. |
| LANG-1102 | Słowa kluczowe dostają `hl-kw`; typy wbudowane i nazwy z wielkiej litery `hl-type`; `True False None Some` `hl-builtin`; nazwy wbudowane z ukośnikiem (`Natural/fold`) są jedną nazwą. |
| LANG-1103 | Komentarz `--` i zagnieżdżany `{- … -}` dostają `hl-com`. |
| LANG-1104 | Napisy `"…"` i wielowierszowe `''…''` ze wstawkami `${…}` dostają `hl-str`; liczby `hl-num`; operatory `hl-op`. |
