# 0007 Pierwsze wydanie v0.1.0 i publikacja

Status: przyjęta · 2026-11-25 (decyzja właściciela: przygotować do publikacji i dołączyć do `gomdgo` i `gopeekgo`
przed końcem fuzzingu; zmienia kolejność z `docs/roadmap.md`, gdzie wydanie było po 24 h fuzzingu)

## Kontekst

Odbiorcy wskazują w `go.mod` tylko wydane wersje (`gopeekgo`, ADR 0001 i 0003 tamże), a `gomdgo` jest publiczne,
więc jego moduł `cmd/gomdgo-mermaid` nie może wymagać prywatnego modułu. Właściciel chce dołączyć kolorowanie do
obu odbiorców teraz; wynik fuzzingu 24 h nie zmieni tej kolejności.

## Decyzja

1. **Wydanie to znacznik `v0.1.0`** na commicie z pełnym README i porównaniem; znaczników nie przesuwamy (jak
   ADR 0017 w `gomdgo`). Poprawki po fuzzingu wychodzą jako `v0.1.1` i dalej.
2. **Repozytorium staje się publiczne** przed dołączeniem do `gomdgo` (na wyraźną zgodę właściciela przy samej
   zmianie widoczności).
3. **Fuzzing 24 h nie jest warunkiem v0.1.0**; jest warunkiem dalszych wydań i wersji 1.0. Stan fuzzingu README
   podaje wprost.
4. **README po angielsku** na wzór `gomdgo`: zalety, słabe strony, tabela porównania z chroma i z highlight
   (w QLMarkdown) z przypisami o sposobie pomiaru; dokumentacja wewnętrzna zostaje po polsku.

## Skutki

+ odbiorcy mogą wskazać wydaną wersję w `go.mod`; łatki z `docs/adapters/` stosuje się bez `go.work`;
− v0.1.0 nie spełnia jeszcze kryterium H4 w części „24 h fuzzingu”; ryzyko ogranicza 142 przypadków
  patologicznych i dotychczasowe przebiegi bez błędu.

## Uzupełnienie 2026-11-26: v0.1.1 przed 24 h fuzzingu

Decyzja właściciela: `v0.1.1` wychodzi przed 24 h fuzzingu, bo od `v0.1.0` nie zmienił się kod biblioteki
(`highlight/`, `cmd/`): to wydanie README, `docs/comparison.md` i powtarzalnego porównania w `tools/compare`
(zamrożone wejścia, `rivalbench`, zapisane wyniki). 24 h fuzzingu (`FuzzHTML`, `FuzzDefinition`) zlecone tego samego
dnia na Synology przez wspólną kolejkę; warunek z punktu 3 obowiązuje dalej dla pierwszego wydania ze zmianą
kodu biblioteki i dla 1.0. Sekcję „Status” usunięto z README (decyzja
właściciela z 2026-11-26, README ma być krótkie), więc stan fuzzingu podają `docs/testing.md` i `docs/roadmap.md`,
a nie README, jak chciał punkt 3.

`v0.1.2` (ten sam dzień, przed upublicznieniem): komentarze dokumentacyjne poleceń (`cmd/gohighlightgo`,
`tools/review`) i modułu `tools/compare` przetłumaczone na angielski, bo pokazuje je pkg.go.dev; binarki
zbudowane z `-trimpath -buildvcs=false -ldflags="-s -w -buildid="` mają tę samą sumę przed i po. Kod biblioteki bez
zmian, więc fuzzing zlecony na `b3e7860` dotyczy też tego wydania.
