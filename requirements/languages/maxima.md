# Wymagania: Maxima

Źródło: Maxima Manual (5.47): language reference. Definicja: `highlight/lang/maxima.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2521 | Nazwy języka: `maxima`, `mac`. |
| LANG-2522 | Słowa z listy keyword (`if`, `then`, `else`, `elseif`, `for`, `from` …) dostają `hl-kw`; słowa z listy builtin (`true`, `false`, `pi`, `%pi`, `%e`, `%i` …) dostają `hl-builtin`. |
| LANG-2523 | Komentarze `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2524 | Liczby dostają `hl-num`; znaki `+-*/^=<>:!'.$;,` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
