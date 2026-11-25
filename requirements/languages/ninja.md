# Wymagania: Ninja

Źródło: The Ninja build system manual: syntax reference. Definicja: `highlight/lang/ninja.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2161 | Nazwy języka: `ninja`. |
| LANG-2162 | Słowa z listy keyword (`rule`, `build`, `default`, `pool`, `include`, `subninja`) dostają `hl-kw`; słowa z listy builtin (`command`, `description`, `depfile`, `deps`, `generator`, `restat` …) dostają `hl-builtin`; nazwa przed `=` dostaje `hl-attr`. |
| LANG-2163 | Komentarze `#` do końca wiersza dostają `hl-com`. |
| LANG-2164 | Liczby dostają `hl-num`; znaki `:|=` dostają `hl-op`; konstrukcje haka `variables` jak w opisie haka w `docs/definitions.md`. |
