# Wymagania: pkg-config files

Źródło: pkg-config(1) manual, metadata file syntax. Definicja: `highlight/lang/pkgconfig.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2021 | Nazwy języka: `pkgconfig`, `pc`. |
| LANG-2022 | Słowa z listy keyword (`Name`, `Description`, `Version`, `URL`, `Requires`, `Requires.private` …) dostają `hl-kw`; nazwa przed `:` albo `=` dostaje `hl-attr`. |
| LANG-2023 | Komentarze `#` do końca wiersza dostają `hl-com`. |
| LANG-2024 | Liczby dostają `hl-num`; znaki `=:<>` dostają `hl-op`; konstrukcje haka `variables` jak w opisie haka w `docs/definitions.md`. |
