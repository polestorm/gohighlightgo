# Wymagania: Go

Źródło: The Go Programming Language Specification, wersja dla Go 1.22 (<https://go.dev/ref/spec>);
w nawiasach rozdział specyfikacji. Kolorowanie jest leksykalne: identyfikator nie jest rozwiązywany,
więc lokalna zmienna o nazwie predeklarowanej dostaje tę samą klasę co nazwa predeklarowana.

| Id | Wymaganie |
|---|---|
| LANG-001 | Nazwy języka: `go`, `golang`. |
| LANG-002 | 25 słów kluczowych dostaje `hl-kw` („Keywords”): `break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var`. |
| LANG-003 | Predeklarowane typy (`any bool byte comparable complex64 complex128 error float32 float64 int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr`) dostają `hl-type`; predeklarowane stałe, wartość zerowa i funkcje (`true false iota nil append cap clear close complex copy delete imag len make max min new panic print println real recover`) dostają `hl-builtin` („Predeclared identifiers”). |
| LANG-004 | Komentarz `//` sięga końca wiersza, komentarz `/* … */` kończy się na pierwszym `*/` (bez zagnieżdżania); niezamknięty sięga końca danych. Oba dostają `hl-com` („Comments”). |
| LANG-005 | Napis w cudzysłowach i literał znakowy w apostrofach dostają `hl-str`, a sekwencje ucieczki w nich `hl-esc`: `\a \b \f \n \r \t \v \\ \' \"`, `\x` z dwiema cyframi szesnastkowymi, `\u` z czterema, `\U` z ośmioma, trzy cyfry ósemkowe; `\` z innym znakiem też dostaje `hl-esc`. Napis i literał nie przechodzą przez koniec wiersza: niezamknięty kończy się przed `\n` („Rune literals”, „String literals”). |
| LANG-006 | Napis surowy w odwrotnych apostrofach może obejmować wiele wierszy i nie ma sekwencji ucieczki; dostaje `hl-str` („String literals”). |
| LANG-007 | Liczby dostają `hl-num`: dziesiętne, `0b`, `0o`, `0x` i ósemkowe z zerem na początku, z `_` między cyframi, zmiennoprzecinkowe dziesiętne (z wykładnikiem `e`) i szesnastkowe (z wykładnikiem `p`), liczba zaczynająca się od `.` i przyrostek `i` liczb urojonych („Integer literals”, „Floating-point literals”, „Imaginary literals”). |
| LANG-008 | Znaki operatorów `+ - * / % & \| ^ < > = ! : ~` dostają `hl-op`; nawiasy, przecinek, średnik i kropka zostają bez klasy („Operators and punctuation”). |
| LANG-009 | Identyfikator, który nie jest słowem kluczowym ani nazwą predeklarowaną, a po nim (po ewentualnych spacjach i tabulatorach) stoi `(`, dostaje `hl-fn`. Identyfikator to litera Unicode albo `_`, potem litery, cyfry i `_` („Identifiers”). |
