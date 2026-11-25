# 0004 Języki z własnym skanerem (etap 3)

Status: przyjęta · 2026-11-25 (punkt 2: decyzja właściciela; pozostałe: wybory wykonawcy w ramach ADR 0001–0003)

## Kontekst

Etap 3 dodał XML, HTML, CSS, YAML, TOML, powłokę, diff i Markdown. Ich budowa zależy od miejsca w wierszu,
wcięcia albo stanu (wartość właściwości CSS, fragment różnicy, skalar blokowy YAML), więc tabela z ADR 0003
nie wystarcza. Każdy język ma własną funkcję na tym samym interfejsie `func(*writer, []byte)`.

## Decyzja

1. **Osobna funkcja na język** (`highlight/lang_markup.go`, `lang_css.go`, `lang_yaml.go`, `lang_toml.go`,
   `lang_shell.go`, `lang_diff.go`, `lang_markdown.go`), z pomocniczymi kawałkami silnika tabelowego, gdzie
   pasują (`scanQuote`, `escapeRule`, `scanNumber`, `entityLen`). Zasady czasu liniowego z ADR 0002, punkt 7,
   obowiązują bez zmian; każdy skaner ma przypadki w `TestPathological`.
2. **XML ogólny, bez wiedzy o JPK** (decyzja właściciela z 2026-11-25, zamyka pytanie otwarte 1): przedrostek
   przestrzeni nazw należy do nazwy znacznika, treść elementów zostaje bez klasy, także liczby i daty.
3. **Języki osadzone**: treść `<script>` i `<style>` w HTML kolorują JavaScript i CSS, a treść bloku kodu
   w Markdown język z napisu informacyjnego. Markdown w Markdown nie jest kolorowany (dostaje `hl-str`), bo
   zagnieżdżanie bez ograniczenia dałoby czas większy niż liniowy. Głębokość osadzenia jest więc stała:
   Markdown → HTML → JavaScript.
4. **Stan zamiast parsowania**: CSS rozróżnia selektor i wartość po `{`, `;`, `}` i dwukropku; diff liczy
   wiersze fragmentu z nagłówka `@@`, więc `---` w treści fragmentu nie jest nagłówkiem; YAML pamięta wcięcie
   wiersza ze wskaźnikiem `|` albo `>`; powłoka pamięta miejsce nazwy polecenia, otwarte `$( … )` i `$(( … ))`
   (stos, bez rekurencji) oraz oczekujące dokumenty `<<KONIEC`.
5. **Markdown kolorowany wiersz po wierszu**, bez wyróżnień `*…*`: ich rozpoznanie wymaga algorytmu
   ograniczników z CommonMark, a nie poprawia czytelności źródła na tyle, żeby uzasadnić ten kod. Kod w wierszu
   szuka zamknięcia na liście serii odwrotnych apostrofów liczonej raz na wiersz.
6. **Klasy**: znacznik i nagłówek tabeli TOML dostają `hl-tag`, nazwy atrybutów, właściwości CSS i klucze
   `hl-attr`, wiersze dodane i usunięte `hl-ins` i `hl-del`, rozwinięcia powłoki (`$x`, `$( … )`) `hl-esc`, jak
   `${ … }` w szablonach JS.

## Skutki

+ 19 języków; wszystkie przechodzą H1–H4, a fuzzing objął je od razu (znalazł zgubienie `''` po `<<` w powłoce,
  poprawione; wejście zostaje w `highlight/testdata/fuzz/FuzzHTML/` jako test regresji);
− CSS zagnieżdżony: selektor z pseudoklasą na początku instrukcji wewnątrz `{}` (`a:hover {`) jest brany za
  właściwość;
− powłoka: nazwa polecenia po słowie kluczowym, które go nie zapowiada (np. po `done`), nie dostaje `hl-fn`;
  rozwinięcia w dokumencie `<<KONIEC` nie są kolorowane osobno;
− Markdown: blok HTML jest kolorowany tylko w wierszu, który zaczyna się od znacznika.
