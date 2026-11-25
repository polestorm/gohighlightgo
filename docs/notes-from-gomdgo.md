# Uwagi z sesji gomdgo (2026-11-25)

Projekt powstał z pytania właściciela, jak dodać kolorowanie składni do podglądu Quick Look (`gopeekgo`),
który zastępuje QLMarkdown.

## Jak robi to konkurencja

**QLMarkdown 1.5.7** (sprawdzone w paczce aplikacji, bez czytania kodu):
- używa programu `highlight` André Simona, wbudowanego jako biblioteka C++ z interpreterem Lua 5.5;
- 263 definicje języków w plikach Lua i 97 motywów; definicja języka to dane (grupy słów kluczowych,
  ograniczniki napisów i komentarzy, kilka wyrażeń regularnych), nie gramatyka;
- wynik: `<pre class='hl'><code class="language-go"><span class="hl kwa">func</span> …`, kilkanaście klas,
  kolory w CSS motywu; kolorowanie dzieje się przy konwersji, nie w oknie podglądu.

Z wiedzy ogólnej, niesprawdzone w tej sesji:
- `goldmark` nie koloruje sam; robi to wtyczka oparta na `chroma` (leksery z wyrażeń regularnych i stanów);
- `highlight.js` i Prism działają podobnie, w JavaScripcie u czytelnika;
- GitHub i edytory używają pełnych gramatyk (TextMate, tree-sitter): dokładniejsze i o rząd wielkości większe.

## Punkt dołączenia w gomdgo

- `markdown.Config.CodeBlocks` to `map[string]CodeBlockFunc`, klucz to pierwsze słowo napisu informacyjnego;
  od 2026-11-25 **bez rozróżniania wielkości liter** (ADR 0018 w `gomdgo`, decyzja właściciela; wcześniej
  z rozróżnianiem). Funkcja dostaje `CodeBlock{Lang, Info, Code, Line}` i zwraca gotowy HTML,
  który trafia do dokumentu bez zmian (gwarancje bezpieczeństwa `gomdgo` go nie obejmują, stąd kryterium H3).
- Nie ma klucza „każdy język”: adapter wpisuje do mapy wszystkie nazwy z `Languages()` (są małymi literami),
  a ```` ```Go ```` trafia do tej samej funkcji co ```` ```go ````. `CodeBlock.Lang` zachowuje zapis z dokumentu,
  więc klasa `language-Go` zostaje taka, jak w zwykłym bloku. Zmiana w `gomdgo` zeruje tam licznik fuzzingu K5.
- Zwykły blok `gomdgo` zapisuje jako `<pre><code class="language-X">…</code></pre>`; adapter powinien zachować
  tę postać, żeby CSS odbiorcy działał tak samo z kolorowaniem i bez.

## Odbiorcy

- `../gopeekgo`: pakiet `preview` buduje `Config` z diagramami; tam dojdzie plik ze znacznikiem `highlight`
  i dopisanie `highlight.CSS` do arkusza strony.
- `../gomdgo/cmd/gomdgo-mermaid`: osobny moduł z `internal/cli.Options{CodeBlocks, Head}`; ten sam wzór.
- `../gopressgo`: do sprawdzenia, czy `gopdfgo` rysuje `<span>` z kolorem w `<pre>`.
