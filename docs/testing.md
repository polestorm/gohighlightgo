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
| motywy | `TestThemes`, `TestThemeContrast` | każda klasa ma regułę w każdym motywie, arkusz powtarzalny, kontrast co najmniej 4,5 : 1 (WCAG 2.2) |
| polecenie | `cmd/gohighlightgo`: `TestRun`, `TestStandaloneAndOutput`, `TestList`, `TestExtensions` | flagi, kody wyjścia, strona `-standalone`, tabela rozszerzeń zgodna z `Languages()` |
| przegląd przez człowieka | `go run ./tools/review -o out/review.html` | jedna strona: każdy plik korpusu w motywie jasnym i ciemnym obok wyniku QLMarkdown (program `highlight` w środku, czarna skrzynka, w ramce), z liczbami (rozmiar, elementy, czas), oceną i uwagami do skopiowania; bez QLMarkdown `-ref ''` |

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

Uwaga: pliki pomiarowe mają mało bloków kodu. Do H5 służy dodatkowo `testdata/h5/code.md` (15 KB, po jednym
bloku na plik korpusu, wszystkie 19 języków).

## H5: koszt u odbiorców (2026-11-25)

Ten sam komputer i sposób co w punkcie wyjścia; `gohighlightgo` na `6c800aa` z motywami, `gomdgo` na `4f39431`,
adaptery z `docs/adapters/` w osobnych kopiach roboczych, zależności lokalnie przez `go.work`. Czas: mediana
z 15 uruchomień procesu `gomdgo-mermaid -standalone -alerts -footnotes -heading-ids plik`; pamięć szczytowa
z `/usr/bin/time -l`.

| Wielkość | bez kolorowania | z kolorowaniem | zmiana |
|---|--:|--:|--:|
| binarka `gomdgo-mermaid` (`go build -trimpath`) | 8 021 282 B | 8 216 690 B | +195 408 B (+2,4%) |
| plik wykonywalny rozszerzenia `gopeekgo` (`macos/build.sh`, `-tags nohighlight` / domyślnie) | 6 676 224 B | 6 828 496 B | +152 272 B (+2,3%) |
| czas, `gomdgo/docs/testing.md` (25 KB, mało kodu) | 3,8 ms | 3,0 ms | w granicach rozrzutu |
| czas, `testdata/h5/code.md` (15 KB kodu) | 2,5 ms | 2,9 ms | +0,4 ms |
| czas, 1,0 MB (40 × `gomdgo/docs/testing.md` + `README.md`) | 11,0 ms | 11,9 ms | +0,9 ms |
| czas, 1,3 MB (85 × `code.md`, sam kod) | 9,5 ms | 26,1 ms | +16,6 ms |
| pamięć szczytowa, 15 KB kodu | 5,9 MB | 7,2 MB | +1,3 MB |
| pamięć szczytowa, 1,3 MB kodu | 17,9 MB | 40,0 MB | +22,1 MB |
| rozmiar wyniku, 1,3 MB kodu | 1 561 973 B | 5 796 314 B | × 3,7 |

Dla porządku QLMarkdown 1.5.7 na tych samych dokumentach z kodem (`--syntax-highlight on --appearance light`,
cały jego potok): 15 KB 47,1 ms i 21,0 MB, 1,3 MB 2211,0 ms i 55,9 MB. Przewaga `gomdgo-mermaid` z kolorowaniem:
czas 93,8% (15 KB) i 98,8% (1,3 MB), pamięć 65,7% i 28,4%. Porównanie samego kolorowania (H6) w etapie 5.

Wnioski: kolorowanie kosztuje około 0,2 MB pliku wykonywalnego; na zwykłych dokumentach czas się nie zmienia,
na samym kodzie rośnie o około 13 ms na MB, a wynik i pamięć rosną z liczbą elementów `<span>`.

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
| 2026-11-25 | etap 2 (`e4ac7a7`) | 11 | testy przechodzą; 77 przypadków patologicznych; fuzzing 3 min, 3,2 mln wejść × 11 języków, bez błędu | — | — | samo kolorowanie 87–122 MB/s (Go 119, C 94, C++ 95, Java 106, JS 116, TS 87, Python 90, Rust 92, SQL 87, JSON 122, Pascal 104), 2–3 alokacje na wywołanie (`BenchmarkCorpus`, korpus języka powtórzony do 1 MB, 1 przebieg `-benchtime 1s`); pierwsza wersja C++ miała 2941 alokacji na wywołanie (ogranicznik napisu surowego budowany dla każdego napisu), poprawione |
| 2026-11-25 | etap 3 (`762c9f2`) | 19 | testy przechodzą; 142 przypadki patologiczne; fuzzing 4 min, 3,5 mln wejść × 19 języków, bez błędu (wcześniejszy przebieg znalazł zgubienie `''` po `<<` w powłoce, poprawione) | — | — | samo kolorowanie (`BenchmarkCorpus`, 1 przebieg `-benchtime 1s`, MB/s): Go 111, C 92, C++ 90, Java 101, JS 102, TS 83, Python 86, Rust 89, SQL 80, JSON 117, Pascal 96, XML 307, HTML 143, CSS 122, YAML 164, TOML 150, sh 132, diff 633, Markdown 482; 2–4 alokacje na wywołanie, Markdown 22; spadek Go względem etapu 2 (119 → 111) po dodaniu śledzenia poprzedniego tokenu dla wyrażeń regularnych JS mieści się w rozrzucie pojedynczych przebiegów (99–119 tego dnia); pierwsze wersje CSS, YAML i Markdown miały 1562, 1596 i 6287 alokacji (`strings.ToLower`, `strings.Fields`), poprawione; 3 762 wiersze Go w `highlight/` |
| 2026-11-25 | etap 4 (na `6c800aa`) | 19 | testy przechodzą | 8 021 282 B / 8 216 690 B | 3,8 / 3,0 ms (25 KB); 9,5 / 26,1 ms (1,3 MB kodu) | motywy, polecenie, strona przeglądu; H5 zmierzone (sekcja „H5”); rozszerzenie `gopeekgo` 6 676 224 / 6 828 496 B; adaptery jeszcze nie w repozytoriach odbiorców |
