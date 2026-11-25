# Wymagania: CMake

Źródło: cmake-language(7) i cmake-commands(7) (CMake 3.28).

| Id | Wymaganie |
|---|---|
| LANG-451 | Nazwa języka: `cmake`. |
| LANG-452 | Polecenia sterujące (`if`, `foreach`, `function` …) i operatory słowne (`NOT`, `AND`, `STREQUAL` …) dostają `hl-kw`; polecenia i stałe z listy (`set`, `add_executable`, `ON`, `PUBLIC` …) `hl-builtin`; bez rozróżniania wielkości liter. |
| LANG-453 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-454 | Napis `"…"` (może obejmować wiele wierszy) dostaje `hl-str`; odwołania `${…}` dostają `hl-esc` także poza napisem; nazwa przed `(` dostaje `hl-fn`. |
