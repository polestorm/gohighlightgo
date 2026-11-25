# Wymagania: Agda

Źródło: Agda User Manual (2.6): Lexical Structure.

| Id | Wymaganie |
|---|---|
| LANG-1181 | Nazwa języka: `agda`. |
| LANG-1182 | Słowa kluczowe dostają `hl-kw`; `Set Prop Setω` i nazwy z wielkiej litery `hl-type`; nazwa może zawierać `'` i `-`; pragma `{-# … #-}` `hl-meta`. |
| LANG-1183 | Komentarz `--` i zagnieżdżany `{- … -}` dostają `hl-com`. |
| LANG-1184 | Napis `"…"` dostaje `hl-str`; liczby `hl-num`; `: = | . \` `hl-op`. |
