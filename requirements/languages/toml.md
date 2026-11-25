# Wymagania: TOML

Źródło: TOML v1.0.0.

| Id | Wymaganie |
|---|---|
| LANG-151 | Nazwa języka: `toml`. |
| LANG-152 | Nagłówek tabeli `[a.b]` i tablicy tabel `[[a]]` na początku wiersza dostaje `hl-tag`; klucz przed `=` (bez cudzysłowu, w cudzysłowie albo z kropkami) dostaje `hl-attr`; `=` dostaje `hl-op`. |
| LANG-153 | Napisy podstawowe `"…"` i wielowierszowe `"""…"""` z ucieczkami (`hl-esc`) oraz dosłowne `'…'` i `'''…'''` bez ucieczek dostają `hl-str`. |
| LANG-154 | Liczby całkowite i zmiennoprzecinkowe (także `0x`, `0o`, `0b`, `_`, znak, `inf`, `nan`) oraz daty i czasy (`1979-05-27T07:32:00Z`, `07:32:00`) dostają `hl-num`; `true` i `false` dostają `hl-builtin`; komentarz `#` dostaje `hl-com`. |
