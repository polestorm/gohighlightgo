# Wymagania: YAML

Źródło: YAML Ain't Markup Language, wersja 1.2.2; w nawiasach rozdział.

| Id | Wymaganie |
|---|---|
| LANG-141 | Nazwy języka: `yaml`, `yml`. |
| LANG-142 | Klucz (skalar bez cudzysłowu albo napis), po którym stoi `:` z odstępem albo końcem wiersza, dostaje `hl-attr` (7.4, 8.2.2). |
| LANG-143 | Napisy `"…"` (z ucieczkami jak w JSON oraz `\x`, `\U`; `hl-esc`) i `'…'` (z `''` jako `hl-esc`) dostają `hl-str` i mogą obejmować wiele wierszy (7.3). |
| LANG-144 | Skalar bez cudzysłowu, który jest liczbą według schematu podstawowego (także `0x`, `0o`, `.inf`, `.nan`), dostaje `hl-num`; `null`, `~`, `true`, `false` (w zapisach ze schematu podstawowego) dostają `hl-builtin`; inne zostają bez klasy (10.3.2). |
| LANG-145 | Komentarz `#` na początku wiersza albo po odstępie dostaje `hl-com`; wskaźniki `-`, `?`, `:` (z odstępem po nich) i `\|`, `>` skalarów blokowych dostają `hl-op`; kotwice `&a`, aliasy `*a` i znaczniki `!tag`, `!!str` dostają `hl-meta`; `---` i `...` na początku wiersza oraz dyrektywy `%` dostają `hl-meta` (6.6, 6.8, 9.2). |
| LANG-146 | Treść skalara blokowego (`\|`, `>`, z modyfikatorami) w kolejnych wierszach wciętych głębiej niż wiersz ze wskaźnikiem (i w pustych wierszach między nimi) dostaje `hl-str` (8.1). |
