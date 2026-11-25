# Wymagania: Pascal i Delphi

Źródło: Delphi Language Guide (dokumentacja RAD Studio 12): „Fundamental Syntactic Elements”, „Directives”,
„Data Types”, procedury i funkcje modułu `System`. Wielkość liter nie ma znaczenia.

| Id | Wymaganie |
|---|---|
| LANG-101 | Nazwy języka: `pascal`, `delphi`, `pas`, `objectpascal`. |
| LANG-102 | Słowa zastrzeżone i dyrektywy używane w deklaracjach (`abstract cdecl deprecated dynamic external final forward helper overload override platform private protected public published reference reintroduce safecall sealed static stdcall strict virtual operator`) dostają `hl-kw`. Dyrektywy, które bywają zwykłymi nazwami (`read`, `write`, `name`, `index`, `default`, `message`), nie są kolorowane. |
| LANG-103 | Typy wbudowane modułu `System` (`Integer`, `Cardinal`, `Int64`, `Char`, `UnicodeString`, `Double`, `Currency`, `Variant` …) dostają `hl-type`; `True False Self Result` i procedury `System` (`Assigned`, `Length`, `SetLength`, `Inc`, `WriteLn` …) dostają `hl-builtin`; `Read` i `Write` nie, bo to częste nazwy metod i dyrektywy właściwości. |
| LANG-104 | Komentarze `//`, `{ … }` i `(* … *)` (bez zagnieżdżania) dostają `hl-com`; dyrektywy kompilatora `{$…}` i `(*$…*)` dostają `hl-meta`. |
| LANG-105 | Napis `'…'` dostaje `hl-str` i nie przechodzi przez koniec wiersza; podwojony apostrof `''` dostaje `hl-esc`; kody znaków `#13` i `#$0D` dostają `hl-esc`. |
| LANG-106 | Liczby dziesiętne (z częścią ułamkową i wykładnikiem, bez kropki na początku; `1..10` to dwie liczby), szesnastkowe `$FF`, dwójkowe `%1010` i ósemkowe `&777`, z `_` między cyframi, dostają `hl-num`. |
| LANG-107 | Operatory `+ - * / = < > : @ ^` dostają `hl-op`; identyfikator przed `(`, który nie jest słowem z listy, dostaje `hl-fn`. |
