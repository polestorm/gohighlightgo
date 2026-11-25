# Wymagania: BlitzBasic

Źródło: Blitz3D/BlitzMax language reference. Definicja: `highlight/lang/blitzbasic.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2791 | Nazwy języka: `blitzbasic`, `blitzmax`, `bb`, `bmx`. |
| LANG-2792 | Słowa z listy keyword (`Global`, `Local`, `Const`, `Function`, `EndFunction`, `Type` …) dostają `hl-kw`; słowa z listy builtin (`True`, `False`, `Null`, `Print`, `Input`, `Len` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2793 | Komentarze `'` do końca wiersza dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-2794 | Liczby dostają `hl-num`; znaki `+-*/^=<>:.%#$` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
