# Wymagania: C

Źródło: ISO/IEC 9899:2024 (C23); w nawiasach rozdział normy. Typy z nagłówków biblioteki standardowej
(`<stddef.h>`, `<stdint.h>`, `<uchar.h>`, `<wchar.h>`) są kolorowane jak typy wbudowane.

| Id | Wymaganie |
|---|---|
| LANG-011 | Nazwy języka: `c`, `h`. |
| LANG-012 | Słowa kluczowe C23 dostają `hl-kw` (6.4.1), z wyjątkiem słów nazywających typy (`bool char double float int long short signed unsigned void _Bool _Complex _Imaginary _BitInt _Decimal32 _Decimal64 _Decimal128`), które razem z `size_t ptrdiff_t intptr_t uintptr_t intmax_t uintmax_t wchar_t char8_t char16_t char32_t` i `intN_t`, `uintN_t` dla N = 8, 16, 32, 64 dostają `hl-type`; `true false nullptr NULL` dostają `hl-builtin`. |
| LANG-013 | Komentarze `//` i `/* … */` (bez zagnieżdżania) dostają `hl-com` (6.4.9). |
| LANG-014 | Napisy w cudzysłowach i stałe znakowe w apostrofach, także z przedrostkami `u8`, `u`, `U`, `L`, dostają `hl-str`; sekwencje ucieczki (`\` z jednym znakiem, `\x` z dowolną liczbą cyfr szesnastkowych, do trzech cyfr ósemkowych, `\u` z czterema i `\U` z ośmioma cyframi) dostają `hl-esc`. `\` przed końcem wiersza przedłuża napis na następny wiersz (6.4.4.4, 6.4.5, 5.1.1.2). |
| LANG-015 | Liczby jak w Go (`LANG-007`), z apostrofem jako separatorem cyfr i przyrostkami `u`, `l`, `f` i podobnymi (6.4.4.1, 6.4.4.2). |
| LANG-016 | Dyrektywa preprocesora: `#` na początku wiersza (po ewentualnych odstępach) z nazwą dyrektywy dostaje `hl-meta`; nazwa pliku w `<…>` po `#include`, `#include_next`, `#import` i `#embed` dostaje `hl-str`; reszta wiersza to zwykłe tokeny (6.10). |
| LANG-017 | Operatory `+ - * / % & \| ^ ! < > = : ? ~` dostają `hl-op`; identyfikator przed `(` dostaje `hl-fn` jak w Go (`LANG-009`). |
