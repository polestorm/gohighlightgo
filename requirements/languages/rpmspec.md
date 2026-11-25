# Wymagania: RPM spec

Źródło: RPM Packaging Guide and rpm.org spec file documentation. Definicja: `highlight/lang/rpmspec.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1741 | Nazwy języka: `rpmspec`, `spec`. |
| LANG-1742 | Słowa z listy keyword (`Name`, `Version`, `Release`, `Summary`, `License`, `URL` …) dostają `hl-kw`; słowa z listy builtin (`if`, `else`, `endif`, `define`, `global`, `ifarch` …) dostają `hl-builtin`; nazwa po `%` dostaje `hl-builtin`; nazwa przed `:` dostaje `hl-attr`. |
| LANG-1743 | Komentarze `#` do końca wiersza dostają `hl-com`. |
| LANG-1744 | Liczby dostają `hl-num`; znaki `={}()<>` dostają `hl-op`. |
