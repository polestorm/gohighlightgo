# Wymagania: Pike

Źródło: Pike reference manual (8.0). Definicja: `highlight/lang/pike.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1431 | Nazwy języka: `pike`, `pmod`. |
| LANG-1432 | Słowa kluczowe dostają `hl-kw`; typy `hl-type`; `true false UNDEFINED write werror` `hl-builtin`; dyrektywy preprocesora jak w C. |
| LANG-1433 | Komentarze `//` i `/* … */` dostają `hl-com`; napisy `"…"` i znaki `'…'` z ucieczkami `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
