# Wymagania: Dart

Źródło: Dart Programming Language Specification (Dart 3) i „Language tour” (Keywords, Strings). Definicja:
`highlight/lang/dart.hl`.

| Id | Wymaganie |
|---|---|
| LANG-241 | Nazwa języka: `dart`. |
| LANG-242 | Słowa zastrzeżone, wbudowane i kontekstowe z listy (bez `get`, `set`) dostają `hl-kw`; typy podstawowe `hl-type`; `true false null print` dostają `hl-builtin`; `$` może być częścią nazwy. |
| LANG-243 | Komentarze `//` i `/* … */`, które się zagnieżdżają, dostają `hl-com`. |
| LANG-244 | Napisy `'…'`, `"…"` i wielowierszowe `'''…'''`, `"""…"""` z ucieczkami i wstawkami `$nazwa`, `${…}` (`hl-esc`) dostają `hl-str`; napis surowy `r'…'` nie ma ucieczek ani wstawek. |
| LANG-245 | Adnotacja `@nazwa` dostaje `hl-meta`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
