# Testy i pomiary

Układ jak w `../gomdgo/docs/testing.md`: korpus w `testdata/`, wymagania w `requirements/`, tabela datowanych
pomiarów na końcu.

## Warstwy

| Warstwa | Gdzie | Co sprawdza |
|---|---|---|
| wymagania | `requirements/*.md`, `requirements/languages/*.md`, `TestRequirementsCovered`, `TestLanguageRequirements` | każde wymaganie ma test albo plik korpusu; każde odwołanie `REQ:` wskazuje istniejące wymaganie; każdy język ma plik wymagań i katalog korpusu (H1) |
| korpus | `testdata/corpus/<język>/*.txt` + `*.html`, `TestCorpus` | wynik co do bajta; język to nazwa katalogu; pierwszy wiersz pliku to komentarz z `REQ:`; wzorce przepisuje `go test -update ./highlight`, różnicę ogląda się w `git diff`. Rozszerzenie `.txt`, żeby `gofmt`, `go vet` i inne narzędzia nie traktowały korpusu jak kodu |
| odwracalność | `TestRoundTrip` | zdjęcie znaczników i trzech odwołań znakowych daje wejście po zamianie niepoprawnego UTF-8 (H2) |
| bezpieczeństwo | `TestSafeOutput`, `TestCheckSafeRejects` | sprawdzian `checkSafe` niezależny od biblioteki (własna kopia listy klas): tylko `<span class="hl-…">` i `</span>`, bez zagnieżdżeń, pustych elementów i `\n` w elemencie, poprawny UTF-8 (H3); drugi test pilnuje, że sprawdzian odrzuca złe wyniki |
| API | `TestHTML`, `TestLanguages`, `TestEscapeOnly`, `TestInvalidUTF8`, `ExampleHTML` | nazwy bez rozróżniania wielkości liter, język nieznany, lista nazw, ucieczka, zamiana na U+FFFD |
| powtarzalność | `TestDeterministic` | 8 gorutyn po 20 wywołań daje ten sam wynik (warto uruchamiać z `-race`) |
| patologiczne | `TestPathological`, lista `pathological` | czas procesora dla dwóch rozmiarów różniących się 16 razy, próg 96 (jak w `gomdgo`); `go test -short` pomija ten test (H4) |
| odporność | `FuzzHTML` | dla każdego języka: bez paniki, `checkSafe`, odwracalność, powtarzalność; ziarna: korpus, `trickyInputs`, małe wersje przypadków patologicznych |
| wydajność | `BenchmarkCorpus` | korpus każdego języka powtórzony do około 1 MB |

`TestPathological` od razu się przydał: pierwsza wersja szukała końca komentarza przez
`strings.Index(string(src[j:]), …)`, a ta konwersja kopiuje resztę danych przy każdym komentarzu. Seria `/*/*/*…`
dała 2,2 ms dla 20 KB i 833 ms dla 320 KB (czynnik 384 przy progu 96). Po zamianie na `bytes.Index` test przechodzi.

## Fuzzing

```sh
go test -run xxx -fuzz FuzzHTML -fuzztime 10m ./highlight
```

Próg przed wydaniem: 24 h na kandydacie do wydania (ADR 0002, punkt 9). Wejścia, które wywołały błąd, Go zapisuje
w `highlight/testdata/fuzz/FuzzHTML/`; trafiają do repozytorium jako testy regresji.

## Punkt wyjścia (przed kolorowaniem)

Pomiar z 2026-11-25, MacBook arm64, Go 1.26.6, `gomdgo` w commicie `587ee59`, `gomermaidgo` v0.1.0.
Po dołączeniu kolorowania te same pomiary trzeba powtórzyć ze znacznikiem `highlight` i bez (kryterium H5).

Sposób: `go build -trimpath` polecenia `gomdgo/cmd/gomdgo-mermaid`; czas całego procesu
(`-standalone -alerts -footnotes -heading-ids plik`), mediana z 15 uruchomień; pamięć z `/usr/bin/time -l`.

| Wielkość | Wartość |
|---|--:|
| binarka `gomdgo-mermaid` | 8 021 346 B |
| plik wykonywalny rozszerzenia Quick Look (Go jako archiwum C + Swift) | 6 676 384 B |
| czas, `gomdgo/docs/testing.md` (24 KB) | 4,2 ms |
| czas, próbka z diagramem `mermaid` (0,2 KB) | 2,9 ms |
| czas, 1,3 MB (40 × `docs/testing.md` + `README.md` z `gomdgo`) | 14,9 ms |
| pamięć szczytowa, próbka | 7,1 MB |
| pamięć szczytowa, 1,3 MB | 31,8 MB |

Uwaga: pliki pomiarowe mają mało bloków kodu. Do H5 potrzebny jest dodatkowo dokument z dużą ilością kodu
w obsługiwanych językach, mierzony w obu wariantach.

## Odniesienie: QLMarkdown 1.5.7 (program `highlight` w środku)

Ten sam dzień i komputer; program `qlmarkdown_cli` z paczki aplikacji, uruchamiany jako czarna skrzynka.

| Wielkość | Wartość |
|---|--:|
| biblioteka kolorowania `libwrapper_highlight.dylib` (x86_64 + arm64) | 6 922 272 B |
| definicje języków: 263 pliki | 1 516 KB |
| motywy: 97 plików | 1 224 KB |
| czas, 24 KB | 21,5 ms |
| czas, próbka z diagramem | 16,5 ms |
| czas, 1,3 MB | 199,8 ms |
| pamięć szczytowa, próbka | 18,2 MB |
| pamięć szczytowa, 1,3 MB | 60,3 MB |

Czasy QLMarkdown obejmują cały jego potok (parser Markdown, kolorowanie), nie samo kolorowanie; diagramów
nie rysuje przy konwersji. Porównanie samego kolorowania wymaga osobnego pomiaru programu `highlight` (H6).

## Tabela pomiarów

| Data | Commit | Języki | H2–H4 | Binarka `gomdgo-mermaid` bez / z | Czas bez / z | Uwagi |
|---|---|--:|---|---|---|---|
| 2026-11-25 | szkielet | 0 | brak | 8 021 346 B / — | 4,2 ms / — (24 KB) | punkt wyjścia |
| 2026-11-25 | etap 1 (`fa272fa`) | 1 (Go) | testy przechodzą; fuzzing 2 min, 11,0 mln wywołań, bez błędu | — | — | samo kolorowanie: 124 MB/s, 2 alokacje na wywołanie (`BenchmarkCorpus/go`, korpus Go powtórzony do 1 MB, mediana nieliczona, 1 przebieg `-benchtime 2s`); wynik 4–5 razy dłuższy od wejścia; 1 120 wierszy Go w `highlight/`, z czego testy około połowy; pomiar u odbiorców w etapie 4 |
| 2026-11-25 | etap 2 (na `fa272fa`) | 11 | testy przechodzą; 77 przypadków patologicznych; fuzzing 3 min, 3,2 mln wejść × 11 języków, bez błędu | — | — | samo kolorowanie 87–122 MB/s (Go 119, C 94, C++ 95, Java 106, JS 116, TS 87, Python 90, Rust 92, SQL 87, JSON 122, Pascal 104), 2–3 alokacje na wywołanie (`BenchmarkCorpus`, korpus języka powtórzony do 1 MB, 1 przebieg `-benchtime 1s`); pierwsza wersja C++ miała 2941 alokacji na wywołanie (ogranicznik napisu surowego budowany dla każdego napisu), poprawione |
