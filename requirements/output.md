# Wymagania: postać wyniku

Decyzje: `docs/adr/0002-engine-and-output.md`.

| Id | Wymaganie |
|---|---|
| OUT-001 | Znaki `&`, `<`, `>` są zapisywane jako `&amp;`, `&lt;`, `&gt;`; każdy inny bajt (także `"` i `'`) trafia do wyniku bez zmian. |
| OUT-002 | Klasa elementu pochodzi z listy: `hl-kw`, `hl-type`, `hl-builtin`, `hl-str`, `hl-esc`, `hl-num`, `hl-com`, `hl-op`, `hl-fn`, `hl-tag`, `hl-attr`, `hl-ins`, `hl-del`, `hl-meta`. |
| OUT-003 | Elementy się nie zagnieżdżają i nie są puste. Sąsiednie tokeny tej samej klasy tworzą jeden element; odstępy między tokenami i tekst bez klasy stoją poza elementami. |
| OUT-004 | Znak `\n` stoi zawsze poza elementem: token przez kilka wierszy dostaje osobny element w każdym wierszu. |
| OUT-005 | Wynik nie zawiera `<pre>` ani `<code>`; owinięcie należy do wywołującego. |
| OUT-006 | Każda nieprzerwana seria bajtów niebędących poprawnym UTF-8 i każdy znak U+0000 są zastępowane jednym znakiem U+FFFD; wynik jest poprawnym UTF-8. |
