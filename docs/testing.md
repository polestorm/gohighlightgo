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
| definicje | `TestDefinitionErrors`, `TestRegister`, `TestRegisterConcurrent`, `FuzzDefinition` | błędy z numerem wiersza, rejestracja i `extends`, równoległość z `HTML`; fuzzing dowolnej definicji z dowolnym kodem: wynik bezpieczny i odwracalny (DEF-004) |
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

## Odniesienie: QLMarkdown 1.5.7 (w środku `highlight` 4.21, wersja odczytana z `libwrapper_highlight.dylib`)

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

## H6: porównanie z odniesieniem (2026-11-25)

`tools/compare` (osobny moduł, Go 1.25 przez chroma): chroma 2.27.0 przez publiczne API (lekser i formater HTML
z klasami, bez `<pre>`), jako czarna skrzynka. highlight tylko wewnątrz QLMarkdown 1.5.7, więc mierzony jako
cały proces: koszt kolorowania to różnica czasu `--syntax-highlight on` i `off` (mediana z 15). Przewaga
`gohighlightgo` według wzoru z `.claude/CLAUDE.md`.

| Wielkość | gohighlightgo | chroma | highlight | przewaga nad chroma | przewaga nad highlight |
|---|--:|--:|--:|--:|--:|
| języki | 19 | 297 | 263 | −93,6% | −92,8% |
| dodane do programu (`-trimpath -ldflags="-s -w"`) | 201 504 B | 4 773 952 B | 3 416 096 B (biblioteka arm64) | 95,8% | 94,1% |
| wiersze Go biblioteki | 3 124 | 31 199 (+ 279 plików definicji, 2,4 MB) | — | 90,0% | — |
| czas, korpus jako 21 bloków (14,5 KB), `BenchmarkRivals`, mediana z 5 | 0,156 ms | 13,26 ms | — | 98,8% | — |
| czas, 1 MB Go | 10,05 ms | 690,5 ms | — | 98,5% | — |
| pamięć przydzielona, korpus | 0,139 MB | 17,50 MB | — | 99,2% | — |
| pamięć przydzielona, 1 MB Go | 4,20 MB | 819,4 MB | — | 99,5% | — |
| pamięć szczytowa procesu, 1 MB Go (2026-11-25, `/usr/bin/time -l`, jedno wywołanie; ten sam proces bez kolorowania 18 MB) | 22 MB | 93 MB | — | 76% | — |
| pamięć szczytowa dodana przez kolorowanie, 1 MB Go (2026-11-26; Go: proces testowy 22 i 93 MB wobec 18 MB; highlight: `qlmarkdown_cli` z `--syntax-highlight on` i `off`, 36 i 25 MB, 2 przebiegi) | 4 MB | 75 MB | 11 MB | 94,7% | 63,6% |
| czas kolorowania, dokument 15 KB kodu (różnica procesów) | 0,4 ms | — | 34,3 ms | — | 98,8% |
| czas kolorowania, dokument 1,3 MB kodu | 16,6 ms | — | 2078,6 ms | — | 99,2% |
| wynik dla korpusu | 65 916 B, 1 817 elementów | 117 394 B, 4 317 elementów | — | 43,9% | — |

Uwagi: chroma zamyka element `<span class="w">` (odstęp) po znaku `\n`, więc jej elementy przechodzą przez koniec
wiersza; nasze nie (ADR 0002). Obie biblioteki i highlight escapują HTML z kodu (`TestRivalsSafety`; highlight
sprawdzony ręcznie w pięciu językach). Rozmiar highlight to część arm64 `libwrapper_highlight.dylib` z
interpreterem Lua, bez 1,5 MB definicji języków i 1,2 MB motywów.

## Odporność na złośliwe dane: porównanie z konkurencją (2026-11-25)

Narzędzie: `tools/compare/cmd/rivalprobe` (moduł `tools/compare`), uruchomienie
`go run ./cmd/rivalprobe -page ../../out/probe.html > probe.json`; strona postępu odświeża się co 2 s, wynik to
JSON. Każde wywołanie chroma i highlight biegnie w osobnym procesie (własna grupa procesów), zabijanym po 15 s,
więc zawieszenie jest wynikiem, a pomiar idzie dalej. Nasza biblioteka biegnie w procesie sondy.

Sprawdzenia:

- **tekst**: korpus i siedem trudnych wejść (CRLF, samotny CR, tabulatory, brak końcowego `\n`, błędny UTF-8, NUL,
  tekst spoza ASCII); po zdjęciu znaczników ma zostać wejście; dla błędnego UTF-8 i NUL U+FFFD liczy się jako
  zachowanie, jednakowo dla obu bibliotek; chroma bez domyślnej zamiany końców wiersza (`EnsureLF: false`);
  drugi przebieg ma dać ten sam wynik; highlight pominięty, bo jego wynik przechodzi przez Markdown;
- **czas**: jedenaście wzorców (`"`, `'`, `` ` ``, `/*`, `(*`, `<!--`, `<`, `${`, `#`, `\`, `[[`, każdy z `a `),
  ok. 4 kB i ×16 (ok. 64 kB); „gorzej niż liniowo”, gdy duży przebieg trwa ponad 200 ms i ponad ×96 małego
  (próg jak w `TestPathological`); „wolno”, gdy liniowo, ale ponad 1 s; highlight jako cały program
  `qlmarkdown_cli` na dokumencie z jednym blokiem, więc czas zawiera start programu.

Pomiar: MacBook arm64, cztery procesy naraz, chroma 2.27.0, QLMarkdown 1.5.7, gohighlightgo przed v0.1.0;
13 min 35 s, nic nie pominięto.

| | gohighlightgo | chroma | highlight |
|---|---:|---:|---:|
| języki: tekst / czas | 298 / 298 | 200 / 200 | — / 134 |
| zawieszenie albo ponad 15 s (64 kB) | 0 z 3 278 | 101 z 2 200 (59 języków) | 8 z 1 474 (8 języków) |
| czas gorszy niż liniowy | 0 | 11 (7 języków: m.in. Makefile, just) | 15 (6 języków, m.in. nginx) |
| liniowo, ale ponad 1 s | 0 | 119 (38 języków) | 1 (Ruby) |
| zawieszenie na zwykłym tekście spoza ASCII, błędnym UTF-8, NUL | 0 | JSONata | — |
| awaria procesu | 0 | YANG (`/*` powtórzone: `panic: unknown state comment`) | 0 |
| tekst zmieniony | 0 | 15 języków (samotny CR; tabulatory i ostatni wiersz w diff, make, just) | — |
| różny wynik drugiego przebiegu | 0 | 0 | — |

Wcześniejsze przebiegi tego samego dnia: limit 3 s dawał 138 przekroczeń w chroma; powtórzone z limitem 15 s
i dwoma procesami dały 93 przekroczenia, 16 przypadków gorszych niż liniowe i 29 wolnych, żadnego fałszywego
alarmu. Liczby zależą od obciążenia maszyny; przed kolejnym wydaniem powtórzyć na spokojnym sprzęcie.

## Powtarzalne porównanie: zamrożone wejścia i `rivalbench` (2026-11-26)

Każdy wiersz tabeli porównania w README ma się dać powtórzyć u siebie, także przez autorów konkurencji. Dlatego
wejścia pomiarów leżą jako pliki w `tools/compare/inputs` (opis po angielsku w jego `README.md`):

- `corpus/`: 21 plików w 19 językach, skopiowane z `testdata/corpus` na `c5ac60f` (tych samych używał pomiar H6);
  korpus biblioteki rośnie z każdym językiem, a ten już nie;
- `text/`: siedem trudnych wejść odwracalności; `hostile/`: 11 wzorców złośliwych × 2 rozmiary; `xss/`: wejście
  `TestRivalsSafety`. Test `TestInputsMatch` w `cmd/rivalprobe` pilnuje, że pliki są bajt w bajt tym, czego
  używa sonda, a `TestRivalsSafety` i `BenchmarkRivals` czytają je wprost.

`go run ./cmd/rivalbench` w `tools/compare` mierzy wszystko poza wrogimi danymi jednym przebiegiem: rozmiar
(programy z `probes/` minus `probes/none`), czas i przydziały (`testing.Benchmark`, mediana z 5), dokument
1,3 MB (u bibliotek Go czas pokolorowania 1 848 bloków, u highlight polecenie QLMarkdown `on` minus `off`, mediana
z 15), pamięć szczytową dodaną (`getrusage` procesu z `probes/` minus `probes/none`, u highlight `on` minus `off`,
mediana z 15) i rozmiar wyniku. Wynik zapisany w `tools/compare/results/rivalbench-2026-11-26.md` i `.json`.

Przebieg 2026-11-26, Apple M4 (10 rdzeni), zasilanie sieciowe, go1.26.6, chroma 2.27.0, QLMarkdown 1.5.7
(highlight 4.21); około 4 minut:

| Wielkość | gohighlightgo | chroma | highlight | przewaga nad chroma | nad highlight |
|---|--:|--:|--:|--:|--:|
| języki | 298 | 297 | 263 | — | — |
| dodane do programu | 536 kB | 4 774 kB | 4 964 kB (biblioteka 3 416 kB, definicje i motywy 1 548 kB) | 88,8% | 89,2% |
| czas, 21 plików (14,5 kB) | 0,17 ms | 13 ms | — | 98,7% | — |
| czas, 1 MB Go | 11 ms | 687 ms | — | 98,4% | — |
| czas, 1,3 MB kodu | 15 ms | 1 138 ms | 2 256 ms | 98,7% | 99,3% |
| pamięć przydzielona, 21 plików | 0,14 MB | 16,93 MB | — | 99,2% | — |
| pamięć przydzielona, 1 MB Go | 4,20 MB | 819,52 MB | — | 99,5% | — |
| pamięć szczytowa dodana, 1 MB Go | 7 MB | 54 MB | 10 MB | 87,5% | 33,4% |
| wynik dla 21 plików | 65,9 kB | 117,4 kB | — | 43,9% | — |

Różnice wobec wcześniejszych liczb z README i tego dokumentu:

- definicje i motywy highlight: 1 548 kB to suma rozmiarów plików; wcześniejsze 2 740 kB podawało `du`, które liczy
  zajęte bloki dysku (dużo małych plików), więc całość highlight spada z 6 156 do 4 964 kB;
- pamięć szczytowa: wcześniej proces testowy Go (4 i 75 MB), teraz programy z `probes/`, które trzymają też całe
  wejście i wynik (u nas +3 MB), a u chromy wynik zależy od chwili sprzątania pamięci; mediana z 15 zamiast
  jednego przebiegu;
- dokument 1,3 MB: wcześniej u nas różnica procesów `gomdgo-mermaid` (17 ms), czego nie da się powtórzyć bez
  łatki z `docs/adapters/`; teraz czas samego kolorowania bloków przez API, co pozwala zmierzyć też chromę;
- czasy w granicach rozrzutu (0,16 → 0,17 ms, 10,0 → 11 ms, 690 → 687 ms).

Ręczne sprawdzenie poleceń z `inputs/README.md` (2026-11-26): chroma YANG na `block-comment-small.txt` kończy
się `panic: unknown state comment`; chroma JSONata na 37-bajtowym `non-ascii.txt` po ponad minucie nadal działa i zajmuje
3,6 GB (zatrzymany ręcznie); chroma make: 35 ms dla 4 kB i 2 725 ms dla 64 kB (z uruchomieniem procesu), nasz
program 12 ms dla 64 kB; highlight z `${` w Dart (`dollar-brace-large.txt`) nadal działa po 20 s (zatrzymany).

## Tabela pomiarów

| Data | Commit | Języki | H2–H4 | Binarka `gomdgo-mermaid` bez / z | Czas bez / z | Uwagi |
|---|---|--:|---|---|---|---|
| 2026-11-25 | szkielet | 0 | brak | 8 021 346 B / — | 4,2 ms / — (24 KB) | punkt wyjścia |
| 2026-11-25 | etap 1 (`fa272fa`) | 1 (Go) | testy przechodzą; fuzzing 2 min, 11,0 mln wywołań, bez błędu | — | — | samo kolorowanie: 124 MB/s, 2 alokacje na wywołanie (`BenchmarkCorpus/go`, korpus Go powtórzony do 1 MB, mediana nieliczona, 1 przebieg `-benchtime 2s`); wynik 4–5 razy dłuższy od wejścia; 1 120 wierszy Go w `highlight/`, z czego testy około połowy; pomiar u odbiorców w etapie 4 |
| 2026-11-25 | etap 2 (`e4ac7a7`) | 11 | testy przechodzą; 77 przypadków patologicznych; fuzzing 3 min, 3,2 mln wejść × 11 języków, bez błędu | — | — | samo kolorowanie 87–122 MB/s (Go 119, C 94, C++ 95, Java 106, JS 116, TS 87, Python 90, Rust 92, SQL 87, JSON 122, Pascal 104), 2–3 alokacje na wywołanie (`BenchmarkCorpus`, korpus języka powtórzony do 1 MB, 1 przebieg `-benchtime 1s`); pierwsza wersja C++ miała 2941 alokacji na wywołanie (ogranicznik napisu surowego budowany dla każdego napisu), poprawione |
| 2026-11-25 | etap 3 (`762c9f2`) | 19 | testy przechodzą; 142 przypadki patologiczne; fuzzing 4 min, 3,5 mln wejść × 19 języków, bez błędu (wcześniejszy przebieg znalazł zgubienie `''` po `<<` w powłoce, poprawione) | — | — | samo kolorowanie (`BenchmarkCorpus`, 1 przebieg `-benchtime 1s`, MB/s): Go 111, C 92, C++ 90, Java 101, JS 102, TS 83, Python 86, Rust 89, SQL 80, JSON 117, Pascal 96, XML 307, HTML 143, CSS 122, YAML 164, TOML 150, sh 132, diff 633, Markdown 482; 2–4 alokacje na wywołanie, Markdown 22; spadek Go względem etapu 2 (119 → 111) po dodaniu śledzenia poprzedniego tokenu dla wyrażeń regularnych JS mieści się w rozrzucie pojedynczych przebiegów (99–119 tego dnia); pierwsze wersje CSS, YAML i Markdown miały 1562, 1596 i 6287 alokacji (`strings.ToLower`, `strings.Fields`), poprawione; 3 762 wiersze Go w `highlight/` |
| 2026-11-25 | etap 4 (`c5ac60f`) | 19 | testy przechodzą | 8 021 282 B / 8 216 690 B | 3,8 / 3,0 ms (25 KB); 9,5 / 26,1 ms (1,3 MB kodu) | motywy, polecenie, strona przeglądu; H5 zmierzone (sekcja „H5”); rozszerzenie `gopeekgo` 6 676 224 / 6 828 496 B; adaptery jeszcze nie w repozytoriach odbiorców |
| 2026-11-25 | etap 5 (na `c5ac60f`) | 19 | testy przechodzą | — | — | H6 zmierzone (sekcja „H6”): wobec chroma 98,5–98,8% mniej czasu, 95,8% mniej dodanego rozmiaru; wobec highlight 99,2% mniej czasu kolorowania na 1,3 MB kodu |
| 2026-11-25 | etap 7 (`20e6e3a`) | 26 | testy przechodzą; 159 przypadków patologicznych (jeden wykrył czas kwadratowy w haku Swifta, poprawione); fuzzing `FuzzHTML` 3 min i `FuzzDefinition` 3,5 min bez błędu | — | — | definicje w plikach (ADR 0008): przeniesienie 11 języków bez zmiany żadnego wzorca, prędkość bez zmian (Go 109–117, C 98, Python 84, Rust 91, SQL 85 MB/s); odczyt definicji przy starcie 83 µs, 166 KB; dodane do programu 201 → 268 KB (26 języków; usunięcie `fmt` z biblioteki zaoszczędziło 186 KB); porównanie czasu z chroma powtórzone pod obciążeniem maszyny (inny proces, średnie obciążenie 16–22), więc tylko stosunek: korpus 75 razy, 1 MB Go 75 razy mniej czasu; pomiar bezwzględny do powtórzenia |
| 2026-11-25 | etap 8 (`c576646`) | 30 | testy przechodzą; 179 przypadków patologicznych (jeden wykrył czas kwadratowy przy serii `@@@@` w nowych znakach przed nazwą, poprawione); fuzzing lokalny 80 s i 35 s z `-parallel 2` (reguła z globalnego CLAUDE.md), jeden znaleziony błąd: literał `%` w Rubym z ogranicznikiem spoza ASCII przecinał znak UTF-8, poprawione | — | — | PHP, Ruby, Haskell, Zig; nowe dyrektywy `ident-part`, `upper-type`, `sigil`, `line-string`, `interp=hash`, kilka haków na język; haki `haskell`, `ruby`, `php` (dokumenty `<<~`/`<<<` przez wspólną funkcję, reszta wiersza kopią tabeli bez dokumentów) |
| 2026-11-25 | etap 8, partia 1 (`c7c98f7`) | 43 | testy przechodzą; 198 przypadków patologicznych; fuzzing lokalny 2 × 90 s z `-parallel 2`, jeden znaleziony błąd: niezamknięte `y!…` w Perlu cofało koniec literału w środek znaku UTF-8, poprawione (ziarno zostaje) | — | — | 13 języków: Perl, Objective-C, Elixir, Erlang, Clojure, F#, OCaml, Julia, MATLAB, Groovy, PowerShell, Visual Basic, Fortran; haki `perl`, `elixir`, `matlab`, `dotted`; dyrektywa `call-after-paren`, opcja `escape : char=`, wstawki `$(…)`; separator opcji to ostatni samotny `:`; dodane do programu 318 KB |
| 2026-11-25 | etap 8, partia 2 (`3f15534`) | 55 | testy przechodzą; 219 przypadków patologicznych; fuzzing lokalny 2 × 90 s z `-parallel 2` bez błędu | — | — | 12 języków: Makefile, Dockerfile, CMake, INI, properties, Protobuf, GraphQL, HCL, Nix, LaTeX, Vim script, AsciiDoc; skanery Go dla INI i properties oraz AsciiDoc (bloki `[source,…]` w swoim języku); haki `variables`, `latex`, `vim`, `hcl`; opcja `interp=dollar-braces`; wczytywanie definicji leniwe (start: 62 µs dla 35 plików); chroma nie zna AsciiDoc; dodane do programu 335 KB |
| 2026-11-25 | etap 8, partia 3 (`e9edaa8`) | 72 | testy przechodzą; 233 przypadki patologiczne; fuzzing lokalny 90 s z `-parallel 2` bez błędu | — | — | 17 języków: Elm, Nim, Crystal, D, Ada, Prolog, Scheme, Racket, Common Lisp, Solidity, Verilog, VHDL, GLSL, Tcl, AWK, asembler x86 i ARM; haki `lisp`, `verilog-number`, `tcl`, `asm-label`; chroma nie zna tylko AsciiDoc (Common Lisp pod nazwą `common-lisp`); dodane do programu 352 KB |
| 2026-11-25 | etap 8, partia 4 (`06a5e66`) | 91 | testy przechodzą; 249 przypadków patologicznych (dwa wykryły czas kwadratowy: IRI w SPARQL i ogranicznik w cudzysłowie przy dokumentach w powłoce liczyły koniec wiersza przy każdej próbie; poprawione); fuzzing lokalny 90 s z `-parallel 2` bez błędu | — | — | 19 języków: COBOL, Smalltalk, Modula-2, Oberon, MoonScript, fish, csh, pliki wsadowe, XQuery, SPARQL, Cypher, PL/SQL, T-SQL, Cython, Mojo, V, Odin, Hare, Gleam; haki `cobol`, `batch`, `iri`, `rem` (REM jako słowo, nie przedrostek: poprawiony też Visual Basic); chroma nie zna AsciiDoc, Cyphera, Oberonu, PL/SQL i XQuery; dodane do programu 385 KB |
| 2026-11-25 | etap 8, partia 5 (`222a2fc`) | 107 | testy przechodzą; 261 przypadków patologicznych; fuzzing lokalny 90 s z `-parallel 2` bez błędu | — | — | 16 języków: Jinja, Handlebars, Liquid, ERB, szablony Go, CSV, TSV, JSON5, Hjson, KDL, RON, Jsonnet, Starlark, Bicep, Rego, Puppet; wspólny skaner szablonów (HTML poza wstawkami, wewnątrz tabela wyrażeń albo Ruby), skaner CSV i TSV; dyrektywa `key-names`; `:` na końcu wiersza jest wartością; identyfikatory wymagań do 4 cyfr; chroma nie zna AsciiDoc, Cyphera, Hjson, Liquid, Oberonu, PL/SQL, RON, TSV i XQuery; dodane do programu 419 KB |
| 2026-11-25 | etap 8, partia 6 (na `222a2fc`) | 124 | testy przechodzą; 273 przypadki patologiczne (jeden wykrył czas kwadratowy przy `[[` bez zamknięcia w Org, poprawione); fuzzing lokalny 90 s z `-parallel 2` bez błędu | — | — | 17 języków: Org, reStructuredText, Textile (skanery Go, bloki kodu w swoim języku), Dhall, Haxe, ActionScript, DFM, Luau, Teal, F*, Lean, Agda, Idris, PureScript, ReScript, Reason, Coq; aliasy `ansible` (YAML) i `packer` (HCL); chroma nie zna 16 naszych języków; dodane do programu 435 KB |
| 2026-11-25 | etap 8, partia 7 (na `74d8e06`) | 154 | testy przechodzą; 273 przypadków patologicznych; fuzzing lokalny ≤ 90 s z `-parallel 2` bez błędu | — | — | Apex, Arduino (alias C++), AutoHotkey, AutoIt, Ballerina, Chapel, Clean, CoffeeScript, Pony, Roc, Factor, Forth, Hack, J, Janet, LiveScript, Logtalk, Mercury, Nemerle, OpenSCAD, Pike, PostScript, QML, Raku, Red, REXX, SAS, Stata, Vala, Wolfram Language, Zeek; komentarze blokowe sprawdzane przed wierszowymi (poprawka dla #= w Julii, ### w CoffeeScripcie, #cs w AutoIt); chroma nie zna 29 naszych języków; dodane do programu 452 KB |
| 2026-11-25 | etap 8, partia 8 (na `74d8e06`) | 185 | testy przechodzą; 273 przypadków patologicznych; fuzzing lokalny ≤ 90 s z `-parallel 2` bez błędu | — | — | ABAP, ANTLR, AppleScript, BibTeX, Brainfuck, Cap'n Proto, Ceylon, Eiffel, Fennel, GDScript, Gherkin, gnuplot, Hy, Io, LLVM IR, Modelica, nginx, NSIS, PL/pgSQL, PromQL, RPM spec, Standard ML, Thrift, Turtle, WGSL, YANG, SourcePawn, Smali, POV-Ray, pliki strefy DNS, HTTP; aliasy jsx, tsx, vue, svelte, scss, less, systemd, reg; definicje, wymagania i korpus z generatora (wymagania wyprowadzone z definicji); chroma nie zna 31 naszych języków; dodane do programu 486 KB |
| 2026-11-25 | etap 8, partia 9 (na `74d8e06`) | 233 | testy przechodzą; 273 przypadków patologicznych; fuzzing lokalny ≤ 90 s z `-parallel 2`: `FuzzDefinition` znalazł ogranicznik niebędący poprawnym UTF-8 (wynik przestawał być poprawnym UTF-8); poprawione, ziarno zachowane | — | — | 47 języków z tabeli (Apache, Caddyfile, Berry, BNF, EBNF, Boo, ChaiScript, Dylan, FreeBASIC, QBasic, Isabelle, JSONata, lighttpd, MiniZinc, MLIR, OpenEdge ABL, Pig Latin, pkg-config, Power Query M, PRQL, Scilab, Sieve, SNOBOL4, TableGen, TypoScript, Whiley, Z80, Kusto, DAX, just, Meson, Ninja, Kconfig, Devicetree, GNU ld, GN, Carbon, Koka, Wren, Squirrel, AngelScript, Pine Script, q, APL, BQN, Logo, ink) i EJS (skaner szablonów); ograniczniki komentarzy i napisów mogą być spoza ASCII (⍝ w APL); chroma nie zna 50 naszych języków; dodane do programu 503 KB |
| 2026-11-25 | etap 8, partia 10 (na `74d8e06`) | 266 | testy przechodzą; 273 przypadków patologicznych; fuzzing lokalny ≤ 90 s z `-parallel 2` bez błędu | — | — | HLSL (wydzielony z aliasu GLSL), CUDA, OpenCL C, Metal, Graphviz DOT, PlantUML, Mermaid, ASN.1, CUE, Pkl, Nickel, Smithy, Typst, roff, jq, Varnish VCL, HAProxy, Nushell, Elvish, Maxima, newLISP, MoonBit, Cairo, Move, Vyper, Clarity, G-code, Structured Text, OpenQASM, Q#, PureBasic, Xojo, AMPL; chroma nie zna 76 naszych języków; dodane do programu 519 KB |
| 2026-11-25 | etap 8, partia 11 (na `74d8e06`) | 298 | testy przechodzą; 283 przypadków patologicznych; fuzzing lokalny ≤ 90 s z `-parallel 2` bez błędu | — | — | asemblery 6502, 68000 i AVR, Inno Setup, LilyPond, Maple, MEL, Icon, PL/I, TTCN-3, Oz, Frink, Gambas, BlitzBasic, CFML, CLIPS, Clipper/xBase, Web IDL, Cassandra CQL, ArangoDB AQL, mcfunction, sed, Xorg, Turing, ALGOL, BCPL, CFEngine 3, ISC dhcpd, Igor Pro, UnrealScript, Papyrus, LSL; 10 nowych przypadków patologicznych; chroma nie zna 98 naszych języków; dodane do programu 536 KB |
