# Wymagania: Markdown

Źródło: CommonMark Spec 0.31.2 i GitHub Flavored Markdown Spec 0.29 (listy zadań). Kolorowane jest źródło
dokumentu wiersz po wierszu, bez parsowania; wyróżnienia `*…*` i `**…**` nie są kolorowane.

| Id | Wymaganie |
|---|---|
| LANG-181 | Nazwy języka: `markdown`, `md`. |
| LANG-182 | Nagłówek ATX (1–6 `#` z odstępem) dostaje w całości `hl-kw`; linia pozioma i podkreślenie nagłówka (`===`, `---`, `***`, `___`) dostają `hl-meta` (4.1–4.3). |
| LANG-183 | Ogrodzenie bloku kodu (co najmniej trzy `` ` `` albo `~`, z wcięciem do trzech spacji) i ogrodzenie zamykające (ten sam znak, co najmniej tyle samo) dostają `hl-meta`; treść jest kolorowana językiem z pierwszego słowa napisu informacyjnego, gdy jest on znany i nie jest Markdownem, a inaczej dostaje `hl-str`. Bez zamknięcia blok sięga końca danych (4.5). |
| LANG-184 | Znaczniki cytatu `>` i znaczniki list (`-`, `*`, `+`, `1.`, `1)` z odstępem) dostają `hl-op`; pole zadania `[ ]`, `[x]` dostaje `hl-meta`; w definicji odsyłacza `[etykieta]:` dostaje `hl-attr`, a reszta wiersza `hl-str` (5.1–5.3, 4.7). |
| LANG-185 | Kod w wierszu (seria `` ` `` do serii tej samej długości w tym samym wierszu), adres odnośnika w `](…)` i autolink `<schemat:…>` albo `<adres@poczty>` dostają `hl-str`; ucieczka `\` przed znakiem interpunkcji ASCII i odwołania znakowe dostają `hl-esc` (6.1–6.7). |
| LANG-186 | Wiersz zaczynający się od znacznika HTML (`<nazwa`, `</`, `<!`) jest kolorowany jak XML (`LANG-112`–`LANG-114`) (4.6). |
