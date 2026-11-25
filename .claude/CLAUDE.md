# gohighlightgo — instrukcje dla Claude Code

Projekt: kolorowanie składni kodu źródłowego w czystym Go. Wejście: nazwa języka i kod; wynik: ten sam kod
zapisany bezpiecznie w HTML, z elementami `<span class="…">`. Cele: zero zależności poza biblioteką standardową,
wynik bezpieczny z konstrukcji, powtarzalny co do bajta, czas liniowy także dla danych złośliwych, mały kod.
Projekt jest agnostyczny wobec odbiorców: **nie importuje `gomdgo`** ani żadnego innego parsera Markdown.

Stan (2026-11-25): **etap 1 gotowy**. Silnik tabelowy, `HTML`, `Languages` i pierwszy język (Go) z korpusem,
wymaganiami i testami H1–H4; postać wyniku ustala ADR 0002. Następny jest etap 2. Zacznij od `README.md`,
potem `docs/adr/0001-scope-and-acceptance.md`, `docs/adr/0002-engine-and-output.md`, `docs/roadmap.md`,
`docs/provenance.md`, `docs/notes-from-gomdgo.md` i `docs/adr/open-questions.md`.

## Po co i w jakim układzie

- Rodzina: `../gomdgo` (Markdown → HTML), `../gomermaidgo` (Mermaid → SVG), `../gopdfgo` (HTML → PDF),
  `../gopressgo` (potok do PDF), `../gopeekgo` (podgląd Quick Look na macOS). Wszystkie powstają tak samo:
  od zera, czarna skrzynka jako odniesienie, korpus, wymagania, datowane pomiary. Wzoruj układ testów
  i dokumentów na `../gomdgo/docs/testing.md` i `../gomdgo/requirements/`.
- Pierwsi odbiorcy: `../gopeekgo` (podgląd) i polecenie `cmd/gomdgo-mermaid` w `../gomdgo`. Kolorowanie ma być
  u nich **opcjonalne przy kompilacji**: adapter leży w module odbiorcy, w pliku ze znacznikiem
  `//go:build highlight` (ADR 0001, punkt „Dołączanie”). Tutaj nie ma żadnego kodu zależnego od odbiorcy.
- Odniesienie do porównań: program `highlight` (André Simon), którego używa QLMarkdown, oraz `chroma`.
  Wolno je uruchamiać jako czarną skrzynkę; kodu i definicji języków nie czytamy (`docs/provenance.md`).

## Kryteria (skrót ADR 0001)

| # | Kryterium | Gdzie mierzone |
|---|---|---|
| H1 | każdy obsługiwany język ma plik korpusu i wymaganie; testy przechodzą | `TestCorpus`, `TestRequirementsCovered` |
| H2 | wynik po zdjęciu znaczników i odwróceniu ucieczek jest równy wejściu, bajt w bajt | `TestRoundTrip`, `FuzzHTML` |
| H3 | w wyniku są tylko `<span class="…">` z klas ze stałej listy | `TestSafeOutput`, `FuzzHTML` |
| H4 | czas liniowy dla danych złośliwych; fuzzing bez paniki | `TestPathological`, `FuzzHTML` |
| H5 | koszt u odbiorcy zmierzony i zapisany: czas, pamięć, rozmiar binarki z kolorowaniem i bez | `docs/testing.md` |
| H6 | porównanie z odniesieniem: rozmiar, czas, liczba języków | `tools/compare` (osobny moduł), `docs/testing.md` |

Po każdym etapie dopisz wiersz do tabeli pomiarów w `docs/testing.md` (data, commit, liczby) i zaktualizuj
kolumnę „Stan” w `docs/roadmap.md`.

## Jak pracujemy

- etapy w kolejności z `docs/roadmap.md`;
- najpierw test (plik korpusu z oczekiwanym wynikiem), potem kod;
- wymagania piszemy razem z etapem, w plikach według prefiksów z `requirements/README.md`;
- układ repozytorium jak w `gomdgo` (jego ADR 0016): w katalogu głównym tylko `README.md`, `LICENSE`, `NOTICE`,
  `go.mod`; biblioteka to pakiet `highlight` w `highlight/` (import `github.com/polestorm/gohighlightgo/highlight`);
  ten plik leży w `.claude/`;
- `go.mod` wymaga Go 1.22 i nie ma sekcji `require`; narzędzia z zależnościami (porównanie z odniesieniem)
  żyją w osobnych modułach w `tools/`;
- repozytorium `github.com/polestorm/gohighlightgo` (prywatne); moduły `github.com/polestorm/*` są w `GOPRIVATE`;
  commit i push tylko na prośbę właściciela.

Zasady:
- **czysty rodowód** (`docs/provenance.md`): opisy języków piszemy ze specyfikacji tych języków. Nie czytaj,
  nie kopiuj i nie przenoś kodu ani definicji języków z innych narzędzi do kolorowania (highlight, chroma,
  pygments, highlight.js, Prism, rouge, gramatyki TextMate i tree-sitter);
- **zero zależności** w module głównym;
- **agnostyczność**: żadnego importu `gomdgo`; przykład połączenia jest w `README.md` i w testach przykładowych
  wolno go pokazać tylko jako tekst;
- nazwy plików, pakietów i identyfikatorów po angielsku; komentarze wewnętrzne i dokumentacja po polsku;
  `README.md` i komentarze dokumentacyjne publicznego API po angielsku (pokazuje je pkg.go.dev);
  komunikaty błędów po angielsku;
- każda decyzja = plik w `docs/adr/`; pytania bez decyzji w `docs/adr/open-questions.md`;
- liczby są datowane i podpisane sposobem pomiaru; porównania z odniesieniem podajemy jako przewagę
  `gohighlightgo` w procentach, `(odniesienie − gohighlightgo) / odniesienie × 100%`; wartość ujemną też zapisujemy;
- licencja: Apache 2.0;
- opisy commitów po angielsku; przed commitem muszą przechodzić `go vet ./...` i `go test ./...`;
- opis commita kończy się jedną linią `Made with the help of Claude Code` (po pustym wierszu); nie dopisuj
  `Co-Authored-By` ani `Claude-Session` (decyzja właściciela z 2026-11-24, przyjęta w `gomdgo`).
