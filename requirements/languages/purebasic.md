# Wymagania: PureBasic

Źródło: PureBasic reference manual (6.x). Definicja: `highlight/lang/purebasic.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2631 | Nazwy języka: `purebasic`, `pb`, `pbi`. |
| LANG-2632 | Słowa z listy keyword (`If`, `ElseIf`, `Else`, `EndIf`, `For`, `To` …) dostają `hl-kw`; słowa z listy builtin (`#True`, `#False`, `#Null`, `Debug`, `OpenWindow`, `MessageRequester` …) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2633 | Komentarze `;` do końca wiersza dostają `hl-com`; napisy `"…"`; `~"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2634 | Liczby dostają `hl-num`; znaki `+-*/=<>&|!.\` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
