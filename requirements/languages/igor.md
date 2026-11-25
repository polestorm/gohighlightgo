# Wymagania: Igor Pro

Źródło: Igor Pro manual: Programming (WaveMetrics). Definicja: `highlight/lang/igor.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2941 | Nazwy języka: `igor`, `igorpro`, `ipf`. |
| LANG-2942 | Słowa z listy keyword (`Function`, `End`, `Macro`, `EndMacro`, `Proc`, `Variable` …) dostają `hl-kw`; słowa z listy builtin (`Print`, `Make`, `Display`, `numpnts`, `sin`, `cos` …) dostają `hl-builtin`; nazwa po `#` dostaje `hl-meta` (bez rozróżniania wielkości liter). |
| LANG-2943 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2944 | Liczby dostają `hl-num`; znaki `+-*/^=<>!&|?:,()[]` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
