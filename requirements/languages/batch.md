# Wymagania: pliki wsadowe Windows

Źródło: Windows Commands reference: cmd, set, if, for, call, goto, rem. Definicja: `highlight/lang/batch.hl`.

| Id | Wymaganie |
|---|---|
| LANG-791 | Nazwy języka: `batch`, `bat`, `cmd`. |
| LANG-792 | Polecenia i słowa z listy (bez rozróżniania wielkości liter) dostają `hl-kw` albo `hl-builtin`. |
| LANG-793 | Komentarz `REM` (słowo, nie przedrostek: `remove` nim nie jest) i `::` na początku wiersza dostają `hl-com`; etykieta `:nazwa` na początku wiersza `hl-fn`. |
| LANG-794 | Zmienne `%NAZWA%`, `%~dp0`, `%1`, `%%i` i `!NAZWA!` dostają `hl-esc`; napis `"…"` `hl-str`; `@ = | & < > ( )` `hl-op`. |
