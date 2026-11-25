# 0005 Polecenie `gohighlightgo`

Status: przyjęta · 2026-11-25 (decyzja właściciela: polecenie w etapie 4; szczegóły to propozycja wykonawcy,
do potwierdzenia przy wykonaniu)

## Kontekst

Biblioteka nie ma programu wiersza poleceń. Porównanie z odniesieniem (H6) jest uczciwe tylko proces do procesu,
bo `highlight` André Simona to program; polecenie przydaje się też do sprawdzania ręcznego i w skryptach.
Wzorem jest `cmd/gomdgo` (ADR 0014 w `gomdgo`).

## Decyzja

1. `cmd/gohighlightgo` w module głównym, jeden plik, tylko biblioteka standardowa: czyta plik albo standardowe
   wejście, pisze na standardowe wyjście albo do pliku (`-o`).
2. Flagi: `-lang` (nazwa języka; wymagana, gdy nie da się jej wziąć z rozszerzenia pliku), `-list` (nazwy
   języków), `-standalone` z `-title` (kompletna strona z `<pre><code class="language-…">` i `highlight.CSS`,
   do otwarcia w przeglądarce). Bez `-standalone` wynik to sam fragment z `HTML`.
3. Język z rozszerzenia pliku (`.go`, `.py`, `.yml` …) według stałej tabeli w poleceniu; to nie jest wykrywanie
   z treści, które ADR 0001 (punkt 4) wyklucza. Rozszerzenie nieznane i brak `-lang`: błąd argumentów.
4. Kody wyjścia jak w `gomdgo`: 0 powodzenie, 1 błąd odczytu lub zapisu, 2 błędne argumenty albo nieznany język.
   Komunikaty po angielsku.
5. Polecenie powstaje w etapie 4, razem z `highlight.CSS`, bo bez arkusza kolorów `-standalone` nie ma sensu.

## Skutki

+ etap 5 (`tools/compare`) może mierzyć oba programy tak samo: czas procesu, pamięć, rozmiar pliku wykonywalnego;
+ zasada zera zależności zostaje;
− tabela rozszerzeń w poleceniu to druga lista nazw obok `Languages()`; test sprawdza, że każde rozszerzenie
  wskazuje język, który biblioteka zna.
