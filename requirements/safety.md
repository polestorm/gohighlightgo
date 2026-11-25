# Wymagania: bezpieczeństwo wyniku

Decyzje: `docs/adr/0001-scope-and-acceptance.md` (H2, H3), `docs/adr/0002-engine-and-output.md`.

| Id | Wymaganie |
|---|---|
| SAFE-001 | Wynik po usunięciu znaczników `<span class="…">` i `</span>` i zamianie `&lt;`, `&gt;`, `&amp;` z powrotem na znaki jest równy wejściu po zamianie z `OUT-006`, bajt w bajt, dla każdego wejścia i każdego języka (kryterium H2). |
| SAFE-002 | Wynik zawiera wyłącznie tekst bez znaków `<` i `>`, odwołania `&amp;`, `&lt;`, `&gt;` oraz znaczniki `<span class="K">` i `</span>` z K z listy `OUT-002`, w parach, bez zagnieżdżeń. Sprawdza to sprawdzian niezależny od biblioteki, także w fuzzingu (kryterium H3). |
