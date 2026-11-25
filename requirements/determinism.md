# Wymagania: powtarzalność i odporność

Decyzje: `docs/adr/0001-scope-and-acceptance.md` (H4), `docs/adr/0002-engine-and-output.md` (punkt 7).

| Id | Wymaganie |
|---|---|
| DET-001 | To samo wejście daje ten sam wynik, bajt w bajt, przy każdym wywołaniu, także równoległym. |
| DET-002 | Czas kolorowania rośnie liniowo z długością wejścia, także dla danych złośliwych (niezamknięte komentarze i napisy, długie serie ograniczników, znaków ucieczki, identyfikatorów i nawiasów). |
| DET-003 | Żadne wejście nie powoduje paniki (fuzzing). |
