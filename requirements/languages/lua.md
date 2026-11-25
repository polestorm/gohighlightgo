# Wymagania: Lua

Źródło: Lua 5.4 Reference Manual, 3.1 „Lexical Conventions” i 6 „Standard Libraries”. Definicja:
`highlight/lang/lua.hl`.

| Id | Wymaganie |
|---|---|
| LANG-221 | Nazwa języka: `lua`. |
| LANG-222 | 22 słowa zastrzeżone dostają `hl-kw`; `true false nil`, funkcje podstawowe i nazwy bibliotek standardowych (`string`, `table`, `math` …) dostają `hl-builtin`. |
| LANG-223 | Komentarz `--` do końca wiersza i komentarz długi `--[[ … ]]`, `--[==[ … ]==]` (koniec z tą samą liczbą `=`) dostają `hl-com`. |
| LANG-224 | Napisy `"…"` i `'…'` z ucieczkami (`\x` z dwiema cyframi, `\u{…}`, `\` przed końcem wiersza) oraz napis długi `[[ … ]]`, `[==[ … ]==]` (wiele wierszy, bez ucieczek) dostają `hl-str`. |
| LANG-225 | Liczby (także szesnastkowe z wykładnikiem `p`) dostają `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
