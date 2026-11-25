# Wymagania

Każde wymaganie ma identyfikator `PREFIKS-nnn`. Testy i pliki korpusu odwołują się do nich przez
`REQ: PREFIKS-nnn`; test `TestRequirementsCovered` sprawdza pokrycie w obie strony.

| Prefiks | Plik | Zakres |
|---|---|---|
| `API` | `api.md` | publiczne API, nazwy języków i aliasy |
| `OUT` | `output.md` | postać wyniku: klasy, zagnieżdżanie, ucieczki |
| `SAFE` | `safety.md` | bezpieczeństwo wyniku (H2, H3) |
| `DET` | `determinism.md` | powtarzalność, czas liniowy, odporność (H4) |
| `LANG` | `languages/<język>.md` | co kolorujemy w danym języku, ze wskazaniem rozdziału jego specyfikacji |

Numeracja `LANG` jest wspólna dla wszystkich języków: nowy język zaczyna od następnej wolnej dziesiątki
(Go 001–009, C 011–017, C++ 021–024, Java 031–036, JavaScript 041–048, TypeScript 051–052, Python 061–066,
Rust 071–077, SQL 081–085, JSON 091–094, Pascal 101–107; następny język od `LANG-111`). Każdy język z `Languages()` musi mieć plik
w `languages/` nazwany jego nazwą główną (`TestLanguageRequirements`).

Pliki powstają razem z etapami w `../docs/roadmap.md`.
