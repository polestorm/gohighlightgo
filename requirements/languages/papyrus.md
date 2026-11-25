# Wymagania: Papyrus

Źródło: Creation Kit wiki: Papyrus language reference (Bethesda). Definicja: `highlight/lang/papyrus.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2961 | Nazwy języka: `papyrus`, `psc`. |
| LANG-2962 | Słowa z listy keyword (`Scriptname`, `extends`, `Function`, `EndFunction`, `Event`, `EndEvent` …) dostają `hl-kw`; słowa z listy type (`Int`, `Float`, `Bool`, `String`, `Actor`, `ObjectReference` …) dostają `hl-type`; słowa z listy builtin (`True`, `False`, `None`, `Debug`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2963 | Komentarze `;` do końca wiersza, `{; … ;}`, `{ … }` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2964 | Liczby dostają `hl-num`; znaki `+-*/%=<>!&|.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
