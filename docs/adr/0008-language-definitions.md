# 0008 Definicje języków w plikach tekstowych

Status: przyjęta · 2026-11-25 (decyzja właściciela po rozważeniu interpretera definicji chroma i highlight)

## Kontekst

Właściciel chce rozszerzać bibliotekę o kolejne składnie, także bez nowego wydania (plik z definicją wskazany
przez użytkownika). Rozważone drogi:

| Droga | Dodane do programu | Prędkość | Czas liniowy | Uwagi |
|---|--:|--:|:-:|---|
| interpreter XML chroma | 4,8 MB (biblioteka chroma, zmierzone) albo szac. 1–1,5 MB (własny) | ok. 1 MB/s (zmierzone) | nie | wyrażenia regularne z nawracaniem (regexp2) |
| interpreter Lua i definicje highlight | szac. 3–4 MB | ok. 1 MB/s (highlight, zmierzone) | nie | definicje na GPL 3.0; Lua i wyrażenia w składni Boost |
| natywny Go na każdy język | ok. 10 KB na język | ok. 100 MB/s | tak | nowa składnia tylko z nowym wydaniem |
| własny format czytany przez silnik tabelowy | ok. 10 KB na język | jak natywny | tak, z konstrukcji | wybrany |

Go nie ładuje kodu w czasie działania w praktyce (pakiet `plugin` działa w wąskich przypadkach), więc rozszerzanie
bez wydania wymaga danych.

## Decyzja

1. **Własny format tekstowy** (`docs/definitions.md`): jedna dyrektywa w wierszu, odpowiadająca polom tabeli z
   ADR 0003 (słowa, komentarze, napisy, przedrostki, sekwencje ucieczki, liczby, operatory, flagi). Bez wyrażeń
   regularnych i bez kodu, więc bezpieczny wynik i czas liniowy wynikają z konstrukcji, także dla definicji
   z zewnątrz (`DEF-004`, `FuzzDefinition`).
2. **Haki w Go** dla konstrukcji spoza tabeli, włączane nazwą (`hook rust`): preprocesor C, napisy surowe C++,
   Rusta i Swifta, długie nawiasy Lua, apostrof jako znak albo symbol, minus w JSON, Pascal. Nowy hak to nowa
   wersja biblioteki.
3. **Wbudowane języki z tabeli w plikach** `highlight/lang/*.hl`, wkompilowanych przez `go:embed`; jedna binarka
   zostaje. Odczyt 11 definicji przy starcie pakietu: 83 µs i 166 KB (pomiar z 2026-11-25). Przeniesienie nie
   zmieniło żadnego wzorca korpusu ani prędkości.
4. **API**: `Register(definicja) error` z błędem `*DefinitionError` (wiersz i opis); `extends` kopiuje
   wcześniejszy język z tabeli. Rejestr to niezmienna migawka wymieniana atomowo, więc `HTML` nie bierze blokady.
   Polecenie dostaje `-def plik` (można powtarzać).
5. **Języki kontekstowe** (dziś XML, HTML, CSS, YAML, TOML, powłoka, diff, Markdown) zostają skanerami w Go.
6. **Nie interpretujemy definicji chroma ani highlight**; nowe języki piszemy ze specyfikacji (zasada rodowodu bez
   zmian).
7. Przy okazji: przedrostek napisu jest sprawdzany na początku tokenu, więc działają przedrostki spoza znaków
   identyfikatora (`@"…"` w C#) i `U&'…'` w SQL, który dotąd nie był rozpoznawany.

## Skutki

+ nowy język z tabeli to 15–40 wierszy czytelnego tekstu; autor spoza projektu może go dodać bez Go;
+ te same gwarancje dla definicji wbudowanych i zewnętrznych;
− koszt na bajt rośnie z liczbą reguł definicji; limit 64 KB ogranicza go do stałej;
− składnie wymagające nowego haka nadal wymagają wydania biblioteki.
