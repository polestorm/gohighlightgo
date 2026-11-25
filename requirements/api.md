# Wymagania: publiczne API

Decyzje: `docs/adr/0001-scope-and-acceptance.md` (punkt „API”).

| Id | Wymaganie |
|---|---|
| API-001 | `HTML(lang, code)` dla znanego języka zwraca `ok == true` i kod zapisany jako fragment HTML z elementami `<span class="…">` (postać: `output.md`). Pusty kod daje pusty wynik. |
| API-002 | Dla nieznanego języka (także pustej nazwy) `HTML` zwraca `nil, false`; wywołujący zapisuje blok bez kolorowania. |
| API-003 | Nazwa języka jest porównywana bez rozróżniania wielkości liter; każdy język ma nazwę główną i może mieć aliasy (np. `go` i `golang`). |
| API-004 | `Languages()` zwraca wszystkie nazwy i aliasy, posortowane, bez powtórzeń, jako nową kopię: zmiana zwróconego wycinka nie wpływa na bibliotekę. |
| API-005 | Funkcje pakietu można wołać równolegle z wielu gorutyn; pakiet nie ma stanu zmiennego. |
