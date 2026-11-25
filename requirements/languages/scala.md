# Wymagania: Scala

Źródło: Scala 3 Language Reference i Scala Language Specification, rozdział „Lexical Syntax”. Definicja:
`highlight/lang/scala.hl`.

| Id | Wymaganie |
|---|---|
| LANG-231 | Nazwy języka: `scala`, `sc`. |
| LANG-232 | Słowa zastrzeżone Scali 3 i słowa miękkie z listy dostają `hl-kw`; typy z `scala` i `Predef` z listy `hl-type`; `true false null None Some println print` dostają `hl-builtin`. |
| LANG-233 | Komentarze `//` i `/* … */`, które się zagnieżdżają, dostają `hl-com`. |
| LANG-234 | Napis `"…"` z ucieczkami i `"""…"""` (wiele wierszy, bez ucieczek) dostają `hl-str`; z przedrostkiem `s` albo `f` wstawki `$nazwa`, `${…}` dostają `hl-esc`; z przedrostkiem `raw` `\` jest zwykłym znakiem. Znak `'a'` dostaje `hl-str`, a `'nazwa` (symbol, cytat) dostaje `hl-meta`. |
| LANG-235 | Adnotacja `@nazwa` dostaje `hl-meta`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
