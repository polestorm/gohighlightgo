# Wymagania: Elm

Źródło: Elm syntax guide (elm-lang.org/docs/syntax) i biblioteka core Elm 0.19. Definicja: `highlight/lang/elm.hl`.

| Id | Wymaganie |
|---|---|
| LANG-551 | Nazwa języka: `elm`. |
| LANG-552 | Słowa kluczowe dostają `hl-kw`; nazwy z wielkiej litery (typy, konstruktory, moduły) `hl-type`, a z listy (`True`, `Just`, `Maybe` …) `hl-builtin`. |
| LANG-553 | Komentarz `--` i zagnieżdżany `{- … -}` dostają `hl-com`. |
| LANG-554 | Napisy `"…"` i `"""…"""` z ucieczkami (`\u{…}`) oraz znak `'a'` dostają `hl-str`; liczby (bez kropki na początku) `hl-num`; operatory `hl-op`. |
