# Wymagania: Haxe

Źródło: Haxe 4 Manual: składnia, słowa kluczowe, napisy, metadane, kompilacja warunkowa.

| Id | Wymaganie |
|---|---|
| LANG-1111 | Nazwy języka: `haxe`, `hx`. |
| LANG-1112 | Słowa kluczowe dostają `hl-kw`; typy podstawowe `hl-type`; `true false null trace` `hl-builtin`; metadane `@nazwa` `hl-meta`; dyrektywa `#if` jak w C. |
| LANG-1113 | Komentarze `//` i `/* … */` dostają `hl-com`. |
| LANG-1114 | Napis `"…"` i `'…'` (ten drugi ze wstawkami `$x`, `${…}`) dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
