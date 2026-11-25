# 0003 Języki z tabeli (etap 2)

Status: przyjęta · 2026-11-25 (Pascal/Delphi w etapie 2: decyzja właściciela; pozostałe punkty: wybory
wykonawcy w ramach ADR 0001 i 0002)

## Kontekst

Etap 2 dodał C, C++, Javę, JavaScript, TypeScript, Pythona, Rusta, SQL, JSON i Pascal/Delphi. Tabela z etapu 1
wystarczała dla Go; pozostałe języki mają konstrukcje, których sama tabela nie opisuje, a kilka miejsc wymagało
wyboru między dokładnością a prostotą.

## Decyzja

1. **Tabela dostaje pola zamiast osobnych skanerów**: słowa bez rozróżniania wielkości liter (`foldCase`),
   dodatkowe znaki identyfikatora (`identExtra`, `$` w JS i Javie), przedrostki napisów (`prefixes`, także surowe
   jak `r` w Pythonie), podwojony ogranicznik jako ucieczka (`quote.doubled`, `''` w SQL i Pascalu), nazwy
   w cudzysłowie bez klasy (`quote.ident`, SQL), reguły liczb (`numberRule`), makra (`macros`), dekoratory
   (`decorators`), szablony JS z wyrażeniami (`templates`) i klucze JSON (`keys`). Sekwencje ucieczki opisuje
   jedna struktura `escapeRule`.
2. **Konstrukcje spoza tabeli obsługuje hak `special`** wołany na początku tokenu: preprocesor C i C++, napisy
   surowe C++ (`R"ogr(…)ogr"`) i Rusta (`r#"…"#`), literały znakowe i czasy życia w Ruście, atrybuty Rusta,
   minus przed liczbą w JSON, dyrektywy i liczby z przedrostkiem w Pascalu. Hak musi spełniać ADR 0002,
   punkt 7; każdy ma przypadki w `TestPathological`.
3. **Szablony JS bez rekurencji**: zagnieżdżone `${ … }` obsługuje stos liczników nawiasów w pętli skanera,
   więc głębokość zagnieżdżenia nie zużywa stosu wywołań.
4. **Wyrażenia regularne JS rozpoznaje poprzedni token** (zmiana z 2026-11-25 na prośbę właściciela; pierwsza
   wersja ich nie rozpoznawała i apostrof w `/'…/` otwierał napis do końca wiersza). Specyfikacja rozstrzyga
   to składnią; tu `/` zaczyna wyrażenie po operatorze, nawiasie otwierającym, przecinku, średniku i kilku
   słowach kluczowych (`LANG-048`), a po wartości jest dzieleniem. Nieudana próba (brak zamknięcia w wierszu)
   wyłącza kolejne próby do końca wiersza, co utrzymuje czas liniowy. Pomyłki zostają w rzadkich miejscach:
   `}` kończący blok, po którym stoi wyrażenie, daje dzielenie.
5. **SQL bez `hl-fn`**: identyfikator przed nawiasem to w SQL najczęściej nazwa tabeli; funkcje wbudowane mają
   `hl-builtin`. Dialekty nie są rozróżniane; lista słów to norma i słowa powszechne w PostgreSQL, SQLite
   i MySQL.
6. **JSON toleruje komentarze** (`//`, `/* */`): bloki opisane jako `json` bywają JSONC, a w poprawnym JSON
   komentarzy nie ma, więc nic się nie psuje.
7. **Pascal bez dyrektyw, które bywają nazwami**: `read`, `write`, `name`, `index`, `default`, `message` nie są
   kolorowane, tak samo procedury `Read` i `Write`.
8. **Słowa kluczowe miękkie i kontekstowe** są kolorowane tylko wtedy, gdy rzadko bywają nazwami (`var`,
   `record`, `yield` w Javie, `async` w JS, `type` w TS). Nie są kolorowane `match` i `case` w Pythonie ani
   `from` i `of` w JS.

## Skutki

+ jedenaście języków w jednym skanerze; kod języka to głównie listy słów (`highlight/lang_*.go`);
+ wydajność 87–122 MB/s we wszystkich językach (`docs/testing.md`);
− kolorowanie pozostaje leksykalne: konwersja `T(x)` dostaje `hl-fn`, `t.0` w Ruście daje liczbę `0`,
  `}` przed wyrażeniem regularnym w JS daje dzielenie;
− plik z przyrostkiem `_js.go` byłby budowany tylko dla `GOOS=js`, dlatego JavaScript leży w `lang_javascript.go`.
