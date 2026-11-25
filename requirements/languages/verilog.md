# Wymagania: Verilog i SystemVerilog

Źródło: IEEE 1800-2017, 5 „Lexical conventions” i aneks B „Keywords”. Definicja: `highlight/lang/verilog.hl`.

| Id | Wymaganie |
|---|---|
| LANG-651 | Nazwy języka: `verilog`, `systemverilog`, `sv`, `svh`, `vh`. |
| LANG-652 | Słowa kluczowe dostają `hl-kw`; typy danych `hl-type`; zadanie systemowe `$display` `hl-builtin`; dyrektywa kompilatora `` `define `` `hl-meta`. |
| LANG-653 | Komentarze `//` i `/* … */` dostają `hl-com`. |
| LANG-654 | Napis `"…"` dostaje `hl-str`; liczba z rozmiarem i podstawą (`8'hFF`, `'b1010`, `4'sd5`, `1'bx`) i zwykła `hl-num`; nazwa przed `(` `hl-fn`. |
