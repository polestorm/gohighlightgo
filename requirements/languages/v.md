# Wymagania: V

Źródło: dokumentacja V (vlang.io/docs, 0.4). Definicja: `highlight/lang/v.hl`.

| Id | Wymaganie |
|---|---|
| LANG-871 | Nazwy języka: `v`, `vlang`. |
| LANG-872 | Słowa kluczowe dostają `hl-kw`; typy wbudowane `hl-type`; `true false nil` i funkcje z listy `hl-builtin`. |
| LANG-873 | Komentarze `//` i zagnieżdżany `/* … */` dostają `hl-com`. |
| LANG-874 | Napisy `"…"` i `'…'` z ucieczkami i wstawkami `$x`, `${…}`, napis surowy `r'…'` i znak `` `a` `` dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
