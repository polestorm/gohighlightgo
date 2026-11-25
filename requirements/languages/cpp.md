# Wymagania: C++

Źródło: ISO/IEC 14882:2024 (C++23); w nawiasach rozdział normy. Reszta jak w C (`c.md`).

| Id | Wymaganie |
|---|---|
| LANG-021 | Nazwy języka: `cpp`, `c++`, `cxx`, `cc`, `hpp`. |
| LANG-022 | Słowa kluczowe i zastępcze zapisy operatorów C++23 (`and`, `bitor`, `not_eq` …) oraz identyfikatory o szczególnym znaczeniu `final`, `override`, `import`, `module` dostają `hl-kw` ([lex.key], [lex.name]); typy jak w C (`LANG-012`); `true false nullptr NULL` dostają `hl-builtin`. |
| LANG-023 | Komentarze, napisy, stałe znakowe, liczby, preprocesor, operatory i wywołania jak w C (`LANG-013`–`LANG-017`). |
| LANG-024 | Napis surowy `R"ogr(…)ogr"` (także z przedrostkami `u8`, `u`, `U`, `L`) dostaje `hl-str` w całości, bez sekwencji ucieczki, i może obejmować wiele wierszy; ogranicznik ma najwyżej 16 znaków bez spacji, nawiasów, `\` i cudzysłowu; niezamknięty sięga końca danych ([lex.string]). |
