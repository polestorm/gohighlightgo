# 0006 Motywy kolorów

Status: przyjęta · 2026-11-25 (decyzja właściciela: na razie dwa motywy, domyślnie jasny; zamyka pytanie
otwarte 3; budowa motywów to wybór wykonawcy)

## Kontekst

ADR 0001 zapowiadał `const CSS` z kolorami dla wyglądu jasnego i ciemnego. Właściciel chce na start dwóch
motywów, z jasnym jako domyślnym, i możliwości dodawania kolejnych bez przebudowy.

## Decyzja

1. **Motyw to dane w Go** (`highlight/theme.go`): kolor tekstu i tło bloku oraz wygląd każdej z 14 klas (kolor,
   opcjonalnie tło dla `hl-ins` i `hl-del`, kursywa dla komentarzy). Arkusz CSS jest z nich składany w stałej
   kolejności, więc jest powtarzalny co do bajta. Nowy motyw to nowy wpis w mapie `themes`.
2. **API**: `ThemeCSS(nazwa) (css, ok)`, `Themes()` i `DefaultTheme = "light"` zamiast stałej `CSS`. Motywy:
   `light` (domyślny), `dark` i `auto` (jasny, a przy ustawieniu systemu na ciemny wygląd ciemny, przez
   `@media (prefers-color-scheme: dark)`); `auto` służy podglądom, które idą za wyglądem systemu (`gopeekgo`).
3. **Reguły działają wewnątrz elementu z klasą `hl`** (`.hl .hl-kw { … }`), który dostaje kolor tekstu i tło
   motywu. Adapter zapisuje blok jako `<pre class="hl"><code class="language-…">`: postać jak u `gomdgo`, z jedną
   klasą więcej. Arkusz nie dotyka strony poza takim blokiem.
4. **Zwykły CSS**: kolory wpisane w reguły, bez zmiennych CSS (`var(--…)`), bo nie wiadomo, czy obsługuje je
   `gopdfgo`.
5. **Kolory dobrane w projekcie** (`docs/provenance.md`: motywów innych narzędzi nie przenosimy). Każdy kolor ma
   z tłem kontrast co najmniej 4,5 : 1 (WCAG 2.2, 1.4.3), co sprawdza `TestThemeContrast`; `TestThemes`
   sprawdza, że każda klasa z listy ma regułę.

## Skutki

+ kolejny motyw to około 30 wierszy danych i nie wymaga zmian w skanerach ani w postaci wyniku;
+ odbiorca nadal może podać własny arkusz dla tych samych klas;
− zmiana listy klas wymaga uzupełnienia każdego motywu (test to wykryje) i własnych arkuszy odbiorców.
