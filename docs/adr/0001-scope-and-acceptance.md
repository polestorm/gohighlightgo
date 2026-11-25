# 0001 Zakres, API i kryteria przyjęcia

Status: przyjęta · 2026-11-25 (polecenie właściciela: osobny projekt, agnostyczny wobec `gomdgo`, łatwy do
dołączenia, np. przełącznikiem przy kompilacji; licencja Apache 2.0). Lista klas, zapis niepoprawnego UTF-8
i próg fuzzingu rozstrzygnięte w ADR 0002; lista języków (punkt 2) czeka w `open-questions.md`.

## Kontekst

Podgląd Markdown na macOS (`gopeekgo`) zastępuje rozszerzenie QLMarkdown. Tamto koloruje bloki kodu, nasze nie.
Pomiar z 2026-11-25 (`docs/testing.md`) pokazał, że kolorowanie u konkurencji kosztuje ponad 9 MB na dysku
(biblioteka C++ z interpreterem Lua, 263 definicje języków, 97 motywów). Chcemy tej funkcji w rodzinie, ale
w naszym stylu: mały kod, zero zależności, bezpieczny wynik, czas liniowy.

## Decyzja

### Zakres

1. Biblioteka zamienia (język, kod) na fragment HTML: tekst kodu po ucieczce HTML, pocięty elementami
   `<span class="…">`. Nie dodaje `<pre>` ani `<code>`; to robi wywołujący.
2. Kilkanaście do kilkudziesięciu języków, nie setki (zmienione w ADR 0009: cel to liczba języków na poziomie
   konkurencji). Pierwsza lista (kolejność w `docs/roadmap.md`):
   Go, C, C++, Java, JavaScript, TypeScript, Python, Rust, SQL, JSON, potem HTML/XML, CSS, YAML, TOML,
   powłoka, diff, Markdown, Pascal/Delphi. Język nieznany: `ok == false`, wywołujący zapisuje blok zwyczajnie.
3. Dokładność leksykalna: słowa kluczowe, typy i nazwy wbudowane, napisy, komentarze, liczby, operatory,
   nazwa przed nawiasem jako wywołanie. Bez analizy składni (nie odróżniamy typu od zmiennej).
4. Poza zakresem: wykrywanie języka z treści, numery wierszy, wyróżnianie wierszy, kolory wpisane w `style`,
   wyjście inne niż HTML (do rozważenia później, np. dla `gopdfgo`).

### Budowa

5. Jeden skaner bajt po bajcie sterowany tabelą opisu języka: grupy słów kluczowych, ograniczniki napisów
   i ucieczki, komentarze wierszowe i blokowe (w tym zagnieżdżane), postać liczb, znaki operatorów.
6. Języki, które nie mieszczą się w tabeli (HTML/XML, YAML, powłoka, diff, Markdown), dostają własne funkcje
   na tym samym interfejsie.
7. Bez `regexp`. Każde wyszukiwanie w przód, które może zawieść, ma ograniczony koszt albo pamięta niepowodzenie;
   każdy nowy przypadek złośliwy trafia do `TestPathological`.

### API (robocze)

```go
package highlight

// HTML returns code escaped for HTML with <span class="…"> around tokens. ok is false when lang is unknown.
func HTML(lang string, code []byte) (out []byte, ok bool)

// Languages returns every accepted name and alias, sorted.
func Languages() []string

// CSS holds the colours for the classes, for light and dark appearance.
const CSS = "…" // zastąpione przez ThemeCSS, Themes i DefaultTheme (ADR 0006)
```

Nazwa języka jest porównywana bez rozróżniania wielkości liter. Klasy mają stały przedrostek (propozycja `hl-`)
i stałą, krótką listę: `hl-kw` (słowo kluczowe), `hl-type`, `hl-builtin`, `hl-str`, `hl-esc`, `hl-num`,
`hl-com`, `hl-op`, `hl-fn`, `hl-tag`, `hl-attr`, `hl-ins`, `hl-del`, `hl-meta`. Potwierdzona w ADR 0002.

### Dołączanie (agnostyczność)

8. Biblioteka nie importuje `gomdgo`. Połączenie to kilka wierszy u odbiorcy: dla każdej nazwy z `Languages()`
   wpis w `markdown.Config.CodeBlocks`, który woła `highlight.HTML` i owija wynik w
   `<pre><code class="language-…">`.
9. Opcjonalność przy kompilacji należy do odbiorcy: adapter leży w pliku `//go:build highlight`, a plik
   `//go:build !highlight` daje pustą mapę. `go build -tags highlight` dołącza kolorowanie, zwykłe `go build` nie.
   Wpis `require` w `go.mod` odbiorcy zostaje w obu wariantach; dlatego robimy to w modułach, które i tak mają
   zależności (`gopeekgo`, `gomdgo/cmd/gomdgo-mermaid`), a nie w module głównym `gomdgo` (jego ADR 0015).

### Kryteria przyjęcia

| # | Kryterium |
|---|---|
| H1 | każdy obsługiwany język ma plik korpusu z oczekiwanym wynikiem i wymaganie w `requirements/` |
| H2 | zdjęcie znaczników i odwrócenie ucieczek z wyniku daje wejście bajt w bajt (dla każdego wejścia; niepoprawny UTF-8 zamieniony na U+FFFD, ADR 0002) |
| H3 | w wyniku występują wyłącznie `<span class="K">` i `</span>` dla K ze stałej listy, w parach, bez zagnieżdżeń (ADR 0002) |
| H4 | czas liniowy dla danych złośliwych; fuzzing bez paniki (24 h na kandydacie do wydania, ADR 0002) |
| H5 | koszt u odbiorcy zmierzony przed i po: czas konwersji, pamięć, rozmiar binarki `gomdgo-mermaid` i rozszerzenia `gopeekgo` budowanych ze znacznikiem `highlight` i bez |
| H6 | porównanie z odniesieniem (`highlight`, `chroma`) jako czarną skrzynką: rozmiar, czas, liczba języków |

H5 jest wymaganiem właściciela z 2026-11-25: po wykonaniu sprawdzić, jak zmieniła się prędkość i wielkość
programu. Punkt wyjścia jest w `docs/testing.md`.

## Skutki

- Odbiorca sam dostarcza CSS do strony (`highlight.CSS` albo własny); biblioteka nie zna strony.
- `gopdfgo` zobaczy kolory tylko wtedy, gdy obsługuje `<span>` z klasą i kolorem tekstu wewnątrz `<pre>`;
  do sprawdzenia w tamtym projekcie.
