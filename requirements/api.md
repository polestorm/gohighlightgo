# Wymagania: publiczne API

Decyzje: `docs/adr/0001-scope-and-acceptance.md` (punkt „API”).

| Id | Wymaganie |
|---|---|
| API-001 | `HTML(lang, code)` dla znanego języka zwraca `ok == true` i kod zapisany jako fragment HTML z elementami `<span class="…">` (postać: `output.md`). Pusty kod daje pusty wynik. |
| API-002 | Dla nieznanego języka (także pustej nazwy) `HTML` zwraca `nil, false`; wywołujący zapisuje blok bez kolorowania. |
| API-003 | Nazwa języka jest porównywana bez rozróżniania wielkości liter; każdy język ma nazwę główną i może mieć aliasy (np. `go` i `golang`). |
| API-004 | `Languages()` zwraca wszystkie nazwy i aliasy, posortowane, bez powtórzeń, jako nową kopię: zmiana zwróconego wycinka nie wpływa na bibliotekę. |
| API-005 | Funkcje pakietu można wołać równolegle z wielu gorutyn; pakiet nie ma stanu zmiennego. |
| API-006 | `ThemeCSS(nazwa)` zwraca arkusz dla motywów `light` (domyślny, `DefaultTheme`), `dark` i `auto` (jasny, a przy ustawieniu systemu na ciemny wygląd ciemny); reguły działają wewnątrz elementu z klasą `hl`, który dostaje kolor tekstu i tło motywu; każda klasa z `OUT-002` ma regułę, a jej kolor ma z tłem kontrast co najmniej 4,5 : 1 (WCAG 2.2, 1.4.3). Arkusz jest powtarzalny co do bajta. `Themes()` zwraca nazwy posortowane. Nieznana nazwa: `ok == false`. |
| API-007 | `Register(definicja)` dodaje język opisany w formacie definicji (`DEF-001`): od tej chwili przyjmuje go `HTML`, wymienia `Languages` i koloruje nim blok kodu w Markdown. Nazwa zajęta daje błąd. `Register` można wołać równolegle z `HTML`. |
