# Wymagania: Ada

Źródło: Ada 2022 Reference Manual (ISO/IEC 8652:2023), 2 „Lexical Elements”, 2.9 „Reserved Words”. Definicja: `highlight/lang/ada.hl`.

| Id | Wymaganie |
|---|---|
| LANG-591 | Nazwy języka: `ada`, `adb`, `ads`. |
| LANG-592 | 73 słowa zastrzeżone (bez rozróżniania wielkości liter) dostają `hl-kw`; typy z pakietu Standard `hl-type`; `True False` i procedury Ada.Text_IO z listy `hl-builtin`. |
| LANG-593 | Komentarz `--` do końca wiersza dostaje `hl-com`. |
| LANG-594 | Napis `"…"` z `""` jako ucieczką i znak `'a'` dostają `hl-str`; atrybut po apostrofie (`X'First`) `hl-meta`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
