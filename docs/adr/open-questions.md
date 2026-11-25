# Pytania otwarte

Pytania 2, 4, 5 i 7 rozstrzygnął ADR 0002 (2026-11-25).

1. **Lista języków pierwszego wydania** i ich kolejność (ADR 0001, punkt 2). Rozstrzygnięte częściowo
   2026-11-25 (decyzja właściciela): Pascal/Delphi przechodzi do etapu 2, bo mieści się w silniku tabelowym
   (słowa bez rozróżniania wielkości liter, komentarze `{}`, `(* *)` i `//`, napisy z podwojonym apostrofem,
   liczby `$` i `#`). Otwarte: czy XML z JPK ma pierwszeństwo w etapie 3.
3. **Motywy**: jeden jasny i jeden ciemny arkusz w `highlight.CSS` czy więcej?
6. **Wyjście inne niż HTML**: czy `gopdfgo` potrzebuje listy tokenów zamiast HTML?
8. **Wielkość liter w `gomdgo`**: `Config.CodeBlocks` rozróżnia wielkość liter, więc blok ```` ```Go ```` nie trafi
   do adaptera. Funkcja domyślna albo porównanie bez rozróżniania wielkości liter to zmiana w `gomdgo`, która
   zeruje tam licznik fuzzingu K5. Do decyzji przed etapem 4.
