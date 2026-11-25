# Wymagania: Pig Latin

Źródło: Apache Pig 0.17 documentation: Pig Latin Basics. Definicja: `highlight/lang/pig.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2011 | Nazwy języka: `pig`, `piglatin`. |
| LANG-2012 | Słowa z listy keyword (`LOAD`, `STORE`, `DUMP`, `FILTER`, `BY`, `FOREACH` …) dostają `hl-kw`; słowa z listy type (`int`, `long`, `float`, `double`, `chararray`, `bytearray` …) dostają `hl-type`; słowa z listy builtin (`COUNT`, `SUM`, `AVG`, `MIN`, `MAX`, `CONCAT` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr` (bez rozróżniania wielkości liter). |
| LANG-2013 | Komentarze `--` do końca wiersza, `/* … */` dostają `hl-com`; napisy `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2014 | Liczby dostają `hl-num`; znaki `=<>!+-*/%:.#` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
