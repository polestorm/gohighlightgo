# Plan etapów

| Etap | Zakres | Stan |
|---|---|---|
| 0 | szkielet: moduł, licencja, ADR 0001, rodowód, punkt wyjścia pomiarów | gotowy (2026-11-25) |
| 1 | silnik tabelowy, klasy, `HTML`, `Languages`, testy H2–H4; pierwszy język: Go | gotowy (2026-11-25; ADR 0002) |
| 2 | języki z tabeli: C, C++, Java, JavaScript, TypeScript, Python, Rust, SQL, JSON, Pascal/Delphi (decyzja właściciela 2026-11-25) | gotowy (2026-11-25; ADR 0003) |
| 3 | języki z własnym kodem: HTML/XML, CSS, YAML, TOML, powłoka, diff, Markdown | nie rozpoczęty |
| 4 | `highlight.CSS` (jasny i ciemny), dołączenie u odbiorców pod znacznikiem `highlight` (`gopeekgo`, `gomdgo/cmd/gomdgo-mermaid`), pomiar H5 | nie rozpoczęty |
| 5 | `tools/compare`: porównanie z odniesieniem (H6), tabela w `README.md` | nie rozpoczęty |
| 6 | fuzzing, pierwsze wydanie | nie rozpoczęty |

Szacunek z sesji `gomdgo` (2026-11-25, nie pomiar): etapy 1–2 to jedna sesja pracy, etap 3 jedna do dwóch,
etap 4 pół sesji; rozmiar rzędu 2–3 tys. wierszy Go.
