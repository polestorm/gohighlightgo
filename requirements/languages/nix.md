# Wymagania: Nix

Źródło: Nix Reference Manual (2.18), „Nix Language”.

| Id | Wymaganie |
|---|---|
| LANG-511 | Nazwa języka: `nix`. |
| LANG-512 | Słowa kluczowe dostają `hl-kw`; `true false null` i funkcje z listy `hl-builtin`; nazwa może zawierać `-` i `'`. |
| LANG-513 | Komentarze `#` i `/* … */` dostają `hl-com`. |
| LANG-514 | Napis `"…"` i napis z wcięciem `''…''` (oba wielowierszowe) ze wstawkami `${…}` dostają `hl-str`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` `hl-fn`. |
