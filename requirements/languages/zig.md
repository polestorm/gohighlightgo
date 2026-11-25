# Wymagania: Zig

Źródło: Zig Language Reference (0.13): „Keyword Reference”, „Primitive Types”, „String Literals and Unicode Code
Point Literals”, „Builtin Functions”. Definicja: `highlight/lang/zig.hl`.

| Id | Wymaganie |
|---|---|
| LANG-291 | Nazwa języka: `zig`. |
| LANG-292 | Słowa kluczowe dostają `hl-kw`; typy podstawowe `hl-type`; `true false null undefined` `hl-builtin`; funkcja wbudowana `@nazwa` (`@import`, `@as`) `hl-builtin`. |
| LANG-293 | Komentarze `//`, `///` i `//!` dostają `hl-com`. |
| LANG-294 | Napis `"…"` i znak `'…'` z ucieczkami (`\x` z dwiema cyframi, `\u{…}`) dostają `hl-str` i `hl-esc`; wiersz napisu wielowierszowego od `\\` do końca wiersza dostaje `hl-str`. |
| LANG-295 | Liczby (bez kropki na początku) dostają `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
