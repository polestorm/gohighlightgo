# Wymagania: Perl

Źródło: perlsyn, perlop („Quote and Quote-like Operators”), perlpod i perlfunc (Perl 5.38). Definicja: `highlight/lang/perl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-301 | Nazwy języka: `perl`, `pl`, `pm`. |
| LANG-302 | Słowa kluczowe i operatory słowne (`eq`, `ne`, `cmp` …) dostają `hl-kw`; funkcje wbudowane z listy `hl-builtin`; zmienne `$x`, `@x`, `%x` (także `$$x`) dostają `hl-attr`. |
| LANG-303 | Komentarz `#` do końca wiersza i dokumentacja POD od `=nazwa` na początku wiersza do `=cut` dostają `hl-com`. |
| LANG-304 | Napisy `"…"` (ucieczki i wstawki `$x`), `'…'` i `` `…` ``; operatory cytujące `q`, `qq`, `qw`, `qr`, `m` (jedna część) oraz `s`, `tr`, `y` (dwie części) z dowolnym ogranicznikiem, z nawiasami zagnieżdżanymi i modyfikatorami (`s/a/b/gi`); dokumenty `<<"KONIEC"`, `<<'KONIEC'`, `<<~KONIEC` jak w Rubym (`LANG-274`); wyrażenie regularne `/…/` jak w JavaScripcie. Wszystko dostaje `hl-str`. |
| LANG-305 | Liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
