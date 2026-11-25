# Wymagania: D

Źródło: D Language Specification (dlang.org/spec): „Lexical”, „Attributes”. Definicja: `highlight/lang/d.hl`.

| Id | Wymaganie |
|---|---|
| LANG-581 | Nazwy języka: `d`, `dlang`. |
| LANG-582 | Słowa kluczowe dostają `hl-kw`; typy wbudowane `hl-type`; `true false null` i `writeln` `hl-builtin`; atrybut `@nazwa` `hl-meta`. |
| LANG-583 | Komentarze `//`, `/* … */` i zagnieżdżany `/+ … +/` dostają `hl-com`. |
| LANG-584 | Napisy `"…"` z ucieczkami (mogą obejmować wiele wierszy), WYSIWYG `` `…` `` i `r"…"` oraz znak `'a'` dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
