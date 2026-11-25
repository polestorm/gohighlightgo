# Wymagania: Groovy

Źródło: Groovy Language Documentation (4.0), „Syntax”. Definicja: `highlight/lang/groovy.hl`.

| Id | Wymaganie |
|---|---|
| LANG-391 | Nazwy języka: `groovy`, `gradle`, `gvy`. |
| LANG-392 | Słowa kluczowe (z `def`, `as`, `in`, `trait`) dostają `hl-kw`; typy z listy `hl-type`; `true false null println print` `hl-builtin`; adnotacja `@Nazwa` `hl-meta`. |
| LANG-393 | Komentarze `//` i `/* … */` dostają `hl-com`. |
| LANG-394 | Napisy `"…"` i `"""…"""` z ucieczkami i wstawkami `$x`, `${…}` oraz `'…'` i `'''…'''` bez wstawek dostają `hl-str`; wyrażenie `/…/` jak w JavaScripcie. |
| LANG-395 | Liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
