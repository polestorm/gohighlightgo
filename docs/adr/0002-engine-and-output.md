# 0002 Silnik i postać wyniku

Status: przyjęta · 2026-11-25 (punkty 1–4: rekomendacje wykonawcy przyjęte przez właściciela tego samego dnia;
zamykają pytania otwarte 2, 4, 5 i 7)

## Kontekst

ADR 0001 zostawił do potwierdzenia nazwy klas, zapis niepoprawnego UTF-8 i zasadę rodowodu. Przed pierwszym
plikiem korpusu trzeba też było ustalić, czy elementy `<span>` się zagnieżdżają i czy przechodzą przez koniec
wiersza: od tego zależy każdy wzorzec w `testdata/corpus/`.

## Decyzja

1. **Niepoprawny UTF-8 i U+0000** są zamieniane na wejściu na U+FFFD, tak jak w `gomdgo` (jego ADR 0004): każda
   nieprzerwana seria niepoprawnych bajtów daje jeden znak U+FFFD. Wynik jest zawsze poprawnym UTF-8.
   Kryterium H2 porównuje wynik z wejściem po tej zamianie. Bloki kodu z `gomdgo` przychodzą już oczyszczone,
   więc dla nich różnicy nie ma.
2. **Klasy**: przedrostek `hl-` i stała lista z ADR 0001 (`hl-kw`, `hl-type`, `hl-builtin`, `hl-str`, `hl-esc`,
   `hl-num`, `hl-com`, `hl-op`, `hl-fn`, `hl-tag`, `hl-attr`, `hl-ins`, `hl-del`, `hl-meta`). Zgodności z cudzymi
   motywami nie szukamy: wymagałaby sprawdzania ich licencji, a `provenance.md` i tak wyklucza ich użycie.
3. **Wynik płaski**: elementy `<span>` się nie zagnieżdżają. Znak ucieczki w napisie zamyka `hl-str`, dostaje
   własny `hl-esc`, po nim `hl-str` otwiera się od nowa. Sąsiednie tokeny tej samej klasy tworzą jeden element.
   Sprawdzian H3 jest dzięki temu prosty, a `gopdfgo` nie musi znać głębokości.
4. **Koniec wiersza zamyka element**: `\n` stoi zawsze poza `<span>`; komentarz albo napis przez kilka wierszy
   dostaje osobny element w każdym wierszu. Wiersze wyniku można przetwarzać osobno (numery wierszy, podział
   na strony w PDF).
5. **Ucieczka** dotyczy wyłącznie `&`, `<`, `>` (na `&amp;`, `&lt;`, `&gt;`): wynik nie trafia do atrybutów,
   a cudzysłów w treści jest bezpieczny. Odwrócenie ucieczki w H2 to trzy zamiany.
6. **Silnik**: jeden przebieg bajt po bajcie. Język opisuje tabela (`table` w `highlight/table.go`): słowa
   z klasą, komentarze wierszowe i blokowe (z zagnieżdżaniem albo bez), rodzaje napisów (ograniczniki,
   znak ucieczki, czy mogą przejść przez koniec wiersza), funkcja długości sekwencji ucieczki, znaki operatorów,
   rozpoznawanie wywołań. Język spoza tabeli (etap 3) dostaje własną funkcję `func(*writer, []byte)`.
   Nawiasy, przecinek, średnik i kropka nie są operatorami i zostają bez klasy.
7. **Czas liniowy**: niezamknięty komentarz blokowy sięga końca danych, niezamknięty napis jednowierszowy końca
   wiersza; patrzenie w przód przy wywołaniu (`nazwa (`) przechodzi tylko przez spacje i tabulatory, które potem
   i tak są zużywane jako odstęp; sekwencja ucieczki ma ograniczoną długość.
8. **Rodowód**: zasada z `provenance.md` obowiązuje w obecnej postaci (potwierdzona).
9. **Próg fuzzingu** przed wydaniem: 24 h na kandydacie do wydania, jak w `gomdgo` (jego ADR 0012).

## Skutki

+ wzorce korpusu są czytelne: jeden wiersz źródła to jeden wiersz wyniku;
+ sprawdzian bezpieczeństwa nie potrzebuje stosu;
− wynik jest nieco dłuższy przy wielowierszowych komentarzach i napisach (powtórzone `<span class="…">`);
− kolorowanie jest leksykalne: lokalna zmienna o nazwie `len` w Go też dostaje `hl-builtin`.
