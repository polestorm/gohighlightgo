# Wymagania: csh i tcsh

Źródło: tcsh(1) (6.24): „Lexical structure” i „Builtin commands”. Definicja: `highlight/lang/csh.hl`.

| Id | Wymaganie |
|---|---|
| LANG-781 | Nazwy języka: `csh`, `tcsh`. |
| LANG-782 | Słowa kluczowe i polecenia sterujące dostają `hl-kw`; polecenia wbudowane z listy `hl-builtin`; zmienna `$nazwa` `hl-attr`. |
| LANG-783 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-784 | Napisy `"…"` (ze wstawkami), `'…'` i `` `…` `` dostają `hl-str`; operatory `hl-op`. |
