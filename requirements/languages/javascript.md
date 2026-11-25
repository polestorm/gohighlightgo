# Wymagania: JavaScript

Źródło: ECMAScript 2024 Language Specification (ECMA-262, 15. wydanie); w nawiasach rozdział.

| Id | Wymaganie |
|---|---|
| LANG-041 | Nazwy języka: `javascript`, `js`, `mjs`, `cjs`, `jsx` (znaczniki JSX bez osobnego koloru). |
| LANG-042 | Słowa zastrzeżone, słowa zastrzeżone w trybie ścisłym i `async`, `await`, `let`, `static`, `yield` dostają `hl-kw` (12.7.2); `true false null undefined NaN Infinity` dostają `hl-builtin`. Identyfikator może zawierać `$` (12.7). |
| LANG-043 | Komentarze `//` i `/* … */` dostają `hl-com` (12.4). |
| LANG-044 | Napisy w cudzysłowach i apostrofach dostają `hl-str`; sekwencje ucieczki (`\` z jednym znakiem, `\x` z dwiema cyframi, `\u` z czterema, `\u{…}`) dostają `hl-esc`; `\` przed końcem wiersza przedłuża napis (12.9.4). |
| LANG-045 | Szablon `` `…` `` dostaje `hl-str` i może obejmować wiele wierszy; `${` i odpowiadające mu `}` dostają `hl-esc`, a wyrażenie między nimi jest kolorowane jak zwykły kod, także z zagnieżdżonymi szablonami i nawiasami `{}` (12.9.6). |
| LANG-046 | Liczby jak w Go (`LANG-007`) z przyrostkiem `n` (BigInt) (12.9.3). |
| LANG-048 | Wyrażenie regularne `/…/flagi` dostaje `hl-str` w całości, gdy `/` stoi na początku danych, po operatorze, po `(`, `[`, `{`, `,`, `;` albo po słowie `return typeof instanceof in new delete void throw case do else yield await`; po identyfikatorze, liczbie, napisie, `)`, `]` i `}` to dzielenie. Wyrażenie kończy się na `/` poza klasą `[…]`, `\\` chroni następny znak; bez zamknięcia w tym wierszu `/` jest operatorem, a w reszcie wiersza nie szukamy już wyrażeń (12.9.5; specyfikacja rozstrzyga to składnią, tu wystarcza poprzedni token). |
| LANG-047 | Dekorator `@nazwa` dostaje `hl-meta`; operatory `+ - * / % & \| ^ ! < > = : ? ~` dostają `hl-op`; identyfikator przed `(` dostaje `hl-fn`. |
