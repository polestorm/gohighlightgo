# Plan etapów

| Etap | Zakres | Stan |
|---|---|---|
| 0 | szkielet: moduł, licencja, ADR 0001, rodowód, punkt wyjścia pomiarów | gotowy (2026-11-25) |
| 1 | silnik tabelowy, klasy, `HTML`, `Languages`, testy H2–H4; pierwszy język: Go | gotowy (2026-11-25; ADR 0002) |
| 2 | języki z tabeli: C, C++, Java, JavaScript, TypeScript, Python, Rust, SQL, JSON, Pascal/Delphi (decyzja właściciela 2026-11-25) | gotowy (2026-11-25; ADR 0003) |
| 3 | języki z własnym kodem: HTML/XML, CSS, YAML, TOML, powłoka, diff, Markdown | gotowy (2026-11-25; ADR 0004) |
| 4 | motywy (ADR 0006), polecenie `cmd/gohighlightgo` (ADR 0005), strona przeglądu `tools/review`, adaptery u odbiorców (`gomdgo/cmd/gomdgo-mermaid` pod znacznikiem `highlight`, `gopeekgo` domyślnie włączone), pomiar H5 | prawie gotowy (2026-11-25): wszystko poza commitem adapterów u odbiorców, które czekają na pierwsze wydanie (łatki w `docs/adapters/`) |
| 5 | `tools/compare`: porównanie z odniesieniem (H6), także polecenia `gohighlightgo` z programem `highlight`, tabela w `README.md` | gotowy (2026-11-25): chroma przez API, highlight w QLMarkdown jako proces |
| 7 | definicje języków w plikach tekstowych (ADR 0008): format, `Register`, `-def`, przeniesienie 11 języków z tabeli; nowe języki z definicji: C#, Kotlin, Swift, Lua, Scala, Dart, R | gotowy (2026-11-25); kolejne języki z definicji w miarę potrzeb |
| 8 | liczba języków na poziomie konkurencji (ADR 0009): kolejne języki z definicji i haków | w toku: 43 języki (2026-11-25: PHP, Ruby, Haskell, Zig; partia 1: Perl, Objective-C, Elixir, Erlang, Clojure, F#, OCaml, Julia, MATLAB, Groovy, PowerShell, Visual Basic, Fortran; partia 2: Makefile, Dockerfile, CMake, INI, properties, Protobuf, GraphQL, HCL, Nix, LaTeX, Vim script, AsciiDoc) — 55 języków; wczytywanie definicji leniwe |
| 6 | fuzzing, pierwsze wydanie | wydanie v0.1.0 przed 24 h fuzzingu (ADR 0007); 24 h do zrobienia przed v0.1.1 |

Szacunek z sesji `gomdgo` (2026-11-25, nie pomiar): etapy 1–2 to jedna sesja pracy, etap 3 jedna do dwóch,
etap 4 pół sesji; rozmiar rzędu 2–3 tys. wierszy Go.
