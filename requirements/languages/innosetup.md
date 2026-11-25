# Wymagania: Inno Setup

Źródło: Inno Setup Help: script format. Definicja: `highlight/lang/innosetup.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2691 | Nazwy języka: `innosetup`, `iss`, `inno`. |
| LANG-2692 | Słowa z listy keyword (`Setup`, `Files`, `Icons`, `Run`, `UninstallRun`, `Registry` …) dostają `hl-kw`; słowa z listy builtin (`AppName`, `AppVersion`, `DefaultDirName`, `OutputBaseFilename`, `Source`, `DestDir` …) dostają `hl-builtin`; nazwa przed `:` albo `=` dostaje `hl-attr` (bez rozróżniania wielkości liter). |
| LANG-2693 | Komentarze `;` do końca wiersza, `//` do końca wiersza, `{ … }` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką; `'…'` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2694 | Liczby dostają `hl-num`; znaki `[]=;:{}` dostają `hl-op`; konstrukcje haka `c-preprocessor` jak w opisie haka w `docs/definitions.md`. |
