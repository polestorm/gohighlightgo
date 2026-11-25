# Wymagania: Odin

Źródło: Odin overview i specyfikacja (odin-lang.org/docs). Definicja: `highlight/lang/odin.hl`.

| Id | Wymaganie |
|---|---|
| LANG-881 | Nazwa języka: `odin`. |
| LANG-882 | Słowa kluczowe dostają `hl-kw`; typy wbudowane `hl-type`; `true false nil` i procedury wbudowane `hl-builtin`; dyrektywa `#nazwa` i atrybut `@nazwa` `hl-meta`. |
| LANG-883 | Komentarze `//` i zagnieżdżany `/* … */` dostają `hl-com`. |
| LANG-884 | Napisy `"…"` z ucieczkami, surowe `` `…` `` i znak `'a'` dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
