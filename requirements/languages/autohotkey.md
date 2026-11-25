# Wymagania: AutoHotkey

Źródło: AutoHotkey v2 documentation: Concepts and Conventions, Scripting Language. Definicja: `highlight/lang/autohotkey.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1251 | Nazwy języka: `autohotkey`, `ahk`. |
| LANG-1252 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; funkcje wbudowane z listy `hl-builtin`; dyrektywa `#Nazwa` `hl-meta`; `%zmienna%` (początek) `hl-attr`. |
| LANG-1253 | Komentarze `;` i `/* … */` dostają `hl-com`; napisy `"…"` i `'…'` z ucieczką odwrotnym apostrofem `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
