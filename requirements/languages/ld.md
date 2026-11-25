# Wymagania: GNU linker script

Źródło: GNU ld manual, 3 Linker Scripts. Definicja: `highlight/lang/ld.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2191 | Nazwy języka: `ld`, `linkerscript`, `lds`. |
| LANG-2192 | Słowa z listy keyword (`ENTRY`, `SECTIONS`, `MEMORY`, `OUTPUT_FORMAT`, `OUTPUT_ARCH`, `INPUT` …) dostają `hl-kw`; słowa z listy builtin (`ADDR`, `SIZEOF`, `LOADADDR`) dostają `hl-builtin`. |
| LANG-2193 | Komentarze `/* … */` dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-2194 | Liczby dostają `hl-num`; znaki `=+-*/>:;{}()` dostają `hl-op`. |
