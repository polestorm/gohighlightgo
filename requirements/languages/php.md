# Wymagania: PHP

Źródło: PHP Manual (PHP 8.3), Language Reference: „Basic syntax”, „Types”, „Strings”, „Attributes”. Definicja:
`highlight/lang/php.hl`.

| Id | Wymaganie |
|---|---|
| LANG-261 | Nazwa języka: `php`. |
| LANG-262 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`, typy `hl-type`, `true false null` i funkcje z listy `hl-builtin`; zmienna `$nazwa` dostaje `hl-attr`. |
| LANG-263 | Komentarze `//`, `#` i `/* … */` dostają `hl-com`; atrybut `#[Nazwa` dostaje `hl-meta`. |
| LANG-264 | Napisy `"…"` (ucieczki i wstawki `$x`, `${…}`) i `'…'` dostają `hl-str`; dokument `<<<KONIEC` i `<<<'KONIEC'`: znacznik `hl-op`, ogranicznik `hl-meta`, treść `hl-str`, wiersz końca (także wcięty i z `;` po nim) `hl-meta`. |
| LANG-265 | Znaczniki `<?php`, `<?=`, `<?` i `?>` dostają `hl-meta`; tekst poza nimi (także na początku danych, gdy zaczynają się od znacznika HTML) jest kolorowany jak HTML (`LANG-122`). Liczby, operatory i wywołania jak w innych językach. |
