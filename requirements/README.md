# Wymagania

Każde wymaganie ma identyfikator `PREFIKS-nnn` (od `LANG-1001` cztery cyfry). Testy i pliki korpusu odwołują się do nich przez
`REQ: PREFIKS-nnn`; test `TestRequirementsCovered` sprawdza pokrycie w obie strony.

| Prefiks | Plik | Zakres |
|---|---|---|
| `API` | `api.md` | publiczne API, nazwy języków i aliasy |
| `OUT` | `output.md` | postać wyniku: klasy, zagnieżdżanie, ucieczki |
| `SAFE` | `safety.md` | bezpieczeństwo wyniku (H2, H3) |
| `DET` | `determinism.md` | powtarzalność, czas liniowy, odporność (H4) |
| `DEF` | `definitions.md` | format definicji języków, `Register` |
| `LANG` | `languages/<język>.md` | co kolorujemy w danym języku, ze wskazaniem rozdziału jego specyfikacji |

Numeracja `LANG` jest wspólna dla wszystkich języków: każdy język dostaje blok dziesięciu numerów, a nowy język
zaczyna od następnej wolnej dziesiątki (zakresy widać w plikach w `languages/`; następny język od `LANG-2981`).
Każdy język z `Languages()` musi mieć plik
w `languages/` nazwany jego nazwą główną (`TestLanguageRequirements`).

Pliki powstają razem z etapami w `../docs/roadmap.md`.
