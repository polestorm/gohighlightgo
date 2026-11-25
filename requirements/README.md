# Wymagania

Każde wymaganie ma identyfikator `PREFIKS-nnn`. Testy i pliki korpusu odwołują się do nich przez
`REQ: PREFIKS-nnn`; test `TestRequirementsCovered` sprawdzi pokrycie w obie strony (do napisania w etapie 1).

| Prefiks | Plik | Zakres |
|---|---|---|
| `API` | `api.md` | publiczne API, nazwy języków i aliasy |
| `OUT` | `output.md` | postać wyniku: klasy, zagnieżdżanie, ucieczki |
| `SAFE` | `safety.md` | bezpieczeństwo wyniku (H2, H3) |
| `DET` | `determinism.md` | powtarzalność, czas liniowy, odporność (H4) |
| `LANG` | `languages/<język>.md` | co kolorujemy w danym języku, ze wskazaniem rozdziału jego specyfikacji |

Pliki powstają razem z etapami w `../docs/roadmap.md`.
