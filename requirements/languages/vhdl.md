# Wymagania: VHDL

Źródło: IEEE 1076-2019, 15 „Lexical elements”, 15.10 „Reserved words”. Definicja: `highlight/lang/vhdl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-661 | Nazwy języka: `vhdl`, `vhd`. |
| LANG-662 | Słowa zastrzeżone (bez rozróżniania wielkości liter) dostają `hl-kw`; typy z pakietów standardowych `hl-type`; `true false` i funkcje z listy `hl-builtin`. |
| LANG-663 | Komentarz `--` i `/* … */` dostają `hl-com`. |
| LANG-664 | Napis `"…"` z `""` jako ucieczką, napis bitowy `x"FF"`, `b"1010"` i znak `'1'` dostają `hl-str`; atrybut po apostrofie (`clk'event`) `hl-meta`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
