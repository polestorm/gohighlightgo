# Testy i pomiary

Testów jeszcze nie ma (etap 0). Układ wzorujemy na `../gomdgo/docs/testing.md`: korpus w `testdata/`,
wymagania w `requirements/`, tabela datowanych pomiarów poniżej.

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
