# Wymagania: powłoka

Źródło: POSIX.1-2024, Shell and Utilities, rozdział 2 „Shell Command Language”, oraz Bash Reference Manual
(wersja 5.2) dla `[[ ]]`, `function`, `local`, `declare`, `source`, `<<<`.

| Id | Wymaganie |
|---|---|
| LANG-161 | Nazwy języka: `sh`, `bash`, `shell`, `zsh`, `ksh`. |
| LANG-162 | Słowo w miejscu nazwy polecenia (na początku wiersza, po `\| & ; (`, po `$(` i po `if then else elif do while until ! {`) dostaje `hl-kw`, gdy to słowo zastrzeżone (2.4), `hl-builtin`, gdy to polecenie wbudowane (2.14 i Bash), a inaczej `hl-fn`. Przypisanie `NAZWA=` w tym miejscu dostaje `hl-attr` i nie kończy miejsca polecenia. Te same słowa w argumentach nie dostają klasy; argument złożony z samych cyfr dostaje `hl-num`. |
| LANG-163 | Napis `'…'` dostaje `hl-str` bez ucieczek; napis `"…"` dostaje `hl-str`, a w nim `\` przed `$ \` " \\` i końcem wiersza oraz rozwinięcia `hl-esc`; `` `…` `` dostaje `hl-str`; `\` poza napisem z następnym znakiem dostaje `hl-esc` (2.2). |
| LANG-164 | Rozwinięcia `$nazwa`, `$1`, `$?` i inne parametry specjalne oraz `${…}` (w tym samym wierszu) dostają `hl-esc`; w `$( … )` ograniczniki dostają `hl-esc`, a treść jest kolorowana jak polecenia, także wewnątrz `"…"` i z zagnieżdżeniem (2.5, 2.6). |
| LANG-165 | Dokument `<<KONIEC`, `<<-KONIEC`, `<<'KONIEC'`: operator dostaje `hl-op`, ogranicznik `hl-meta`, wiersze od następnego do wiersza równego ogranicznikowi (przy `<<-` po pominięciu tabulatorów) `hl-str`, a ten wiersz `hl-meta` (2.7.4). |
| LANG-166 | Komentarz `#` na początku słowa dostaje `hl-com`; operatory `\| & ; < >` dostają `hl-op`. |
