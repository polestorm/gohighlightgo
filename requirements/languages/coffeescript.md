# Wymagania: CoffeeScript

Źródło: CoffeeScript 2 documentation (coffeescript.org). Definicja: `highlight/lang/coffeescript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1301 | Nazwy języka: `coffeescript`, `coffee`, `cson`. |
| LANG-1302 | Słowa kluczowe dostają `hl-kw`; `true false null undefined yes no on off this` `hl-builtin`; `@pole` `hl-attr`. |
| LANG-1303 | Komentarz `#` i blok `### … ###` dostają `hl-com`; napisy `"…"`, `"""…"""` ze wstawkami `#{…}` oraz `'…'`, `'''…'''` `hl-str`; wyrażenie regularne jak w JavaScripcie; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
