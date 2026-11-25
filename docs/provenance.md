# Rodowód kodu

`gohighlightgo` powstaje od zera. Zasada potwierdzona przez właściciela 2026-11-25 (ADR 0002, punkt 8).

## Źródła, z których piszemy

- specyfikacje i oficjalna dokumentacja języków: listy słów kluczowych, postać napisów, komentarzy i liczb
  (np. The Go Programming Language Specification, ISO C, ECMAScript, dokumentacja Pythona, norma SQL,
  RFC 8259 dla JSON, specyfikacje YAML i TOML, norma HTML).

## Czego nie robimy

Nie czytamy, nie kopiujemy i nie przenosimy kodu ani definicji języków z innych narzędzi do kolorowania:
highlight (André Simon), chroma, pygments, highlight.js, Prism, rouge, gramatyki TextMate i tree-sitter.
Dotyczy to także ich motywów kolorów.

## Co wolno

- uruchamiać cudze narzędzie jako czarną skrzynkę i porównywać jego wynik, czas i rozmiar z naszym;
- czytać publiczną dokumentację użytkową cudzych narzędzi (lista języków, opis opcji).

## Odnotowane zetknięcie

W sesji `gomdgo` z 2026-11-25, przy sprawdzaniu, jak koloruje QLMarkdown, asystent wyświetlił pierwsze
45 wierszy pliku `go.lang` z programu `highlight` (licencja GPL) i jeden wynik HTML tego programu. Posłużyło to
do opisu budowy konkurencji (`notes-from-gomdgo.md`). Opis języka Go w tym projekcie ma powstać ze specyfikacji
Go, bez zaglądania do tamtego pliku; lista klas i postać tabeli są zaprojektowane w ADR 0001 niezależnie.
