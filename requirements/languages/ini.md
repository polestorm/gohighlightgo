# Wymagania: INI

Źródło: zwyczajowy format INI (Windows `GetPrivateProfileString`, Python `configparser`, git-config).

| Id | Wymaganie |
|---|---|
| LANG-461 | Nazwy języka: `ini`, `cfg`, `conf`, `desktop`, `gitconfig`, `editorconfig`, `systemd`, `service`, `reg`. |
| LANG-462 | Komentarz `;` albo `#` na początku wiersza dostaje `hl-com`; sekcja `[nazwa]` dostaje `hl-tag`. |
| LANG-463 | Klucz przed `=` albo `:` dostaje `hl-attr`, separator `hl-op`. |
| LANG-464 | Wartość w cudzysłowie dostaje `hl-str`, liczba `hl-num`, `true false yes no on off` `hl-builtin`, odwołanie `${…}` `hl-esc`. |
