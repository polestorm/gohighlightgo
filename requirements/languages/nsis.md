# Wymagania: NSIS

Źródło: NSIS Users Manual (3.x): script reference. Definicja: `highlight/lang/nsis.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1711 | Nazwy języka: `nsis`, `nsi`, `nsh`. |
| LANG-1712 | Słowa z listy keyword (`Name`, `OutFile`, `InstallDir`, `Section`, `SectionEnd`, `Function` …) dostają `hl-kw`; słowa z listy builtin (`INSTDIR`, `TEMP`, `SMPROGRAMS`, `PROGRAMFILES`, `HKLM`, `HKCU` …) dostają `hl-builtin`; nazwa po `$` dostaje `hl-attr`; nazwa po `!` dostaje `hl-meta` (bez rozróżniania wielkości liter). |
| LANG-1713 | Komentarze `;` do końca wiersza, `#` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, ze wstawkami `hl-esc`; `'…'`; `` `…` `` dostają `hl-str`. |
| LANG-1714 | Liczby dostają `hl-num`; znaki `+-*/=<>|&` dostają `hl-op`. |
