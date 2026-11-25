# Wymagania: Haskell

Źródło: Haskell 2010 Language Report, rozdział 2 „Lexical Structure” i 12 „Compiler Pragmas”, oraz moduł
Prelude. Definicja: `highlight/lang/haskell.hl`.

| Id | Wymaganie |
|---|---|
| LANG-281 | Nazwy języka: `haskell`, `hs`. |
| LANG-282 | Słowa zastrzeżone i częste słowa rozszerzeń (`forall`, `qualified`, `hiding`, `family` …) dostają `hl-kw`; konstruktory i funkcje Prelude z listy `hl-builtin`; inna nazwa z wielkiej litery (typ, konstruktor, moduł) `hl-type`; apostrof wewnątrz nazwy (`foldl'`, `x'`) należy do nazwy. |
| LANG-283 | Komentarz `--` do końca wiersza i komentarz `{- … -}`, który się zagnieżdża, dostają `hl-com`; pragma `{-# … #-}` dostaje `hl-meta`. |
| LANG-284 | Napis `"…"` z ucieczkami dostaje `hl-str` i `hl-esc`; znak `'a'`, `'\n'` dostaje `hl-str`, a `'nazwa` (cytat w Template Haskell) `hl-meta`. |
| LANG-285 | Liczby (bez kropki na początku) dostają `hl-num`; znaki operatorów `!#$%&*+./<=>?@\^|-~:` dostają `hl-op`. Wywołania nie są rozpoznawane, bo funkcja nie wymaga nawiasów. |
