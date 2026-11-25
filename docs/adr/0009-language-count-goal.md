# 0009 Cel: liczba języków na poziomie konkurencji

Status: przyjęta · 2026-11-25 (decyzja właściciela; zmienia punkt 2 ADR 0001)

## Kontekst

ADR 0001 zakładał „kilkanaście do kilkudziesięciu języków, nie setki”. Po wprowadzeniu definicji w plikach
(ADR 0008) koszt języka z tabeli spadł do 15–40 wierszy tekstu, a właściciel postawił cel: zrównać się liczbą
języków z konkurencją (chroma 297, highlight 263; stan 2026-11-25).

## Decyzja

1. Celem jest około 260–300 języków. Kolejność: najpierw języki najczęściej spotykane w blokach kodu
   (dokumentacja, README, Stack Overflow), potem formaty konfiguracji i języki niszowe.
2. Zasady się nie zmieniają: każdy język ma definicję napisaną ze specyfikacji albo oficjalnej dokumentacji
   (zasada rodowodu), plik wymagań z odwołaniami do źródła, plik korpusu i przechodzi H2–H4. Konstrukcje spoza
   tabeli są hakami w Go z przypadkami w `TestPathological`.
3. Definicje są wczytywane leniwie (od 2026-11-25, przy 44 plikach): przy starcie pakietu czytane są tylko nazwy
   (62 µs dla 35 plików), a całą definicję parsuje pierwsze użycie języka, raz (`sync.Once`). Poprawność wszystkich
   sprawdza `TestBuiltinDefinitions`.
4. Rozmiar: definicja dokłada do programu kilka KB; przy 300 językach spodziewany rozmiar to około 1–1,5 MB
   (szacunek), wciąż mniej niż chroma (4,8 MB zmierzone) i highlight (3,4 MB i 2,7 MB danych).

## Skutki

+ porównanie w README przestanie mieć wiersz, w którym konkurencja wyraźnie wygrywa;
− praca rozłożona na wiele sesji; korpus i wymagania rosną liniowo z liczbą języków;
− języki z nietypową składnią wymagają nowych haków, czyli nowych wersji biblioteki.
