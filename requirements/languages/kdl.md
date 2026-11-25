# Wymagania: KDL

Źródło: The KDL Document Language (2.0.0).

| Id | Wymaganie |
|---|---|
| LANG-1001 | Nazwa języka: `kdl`. |
| LANG-1002 | Nazwa właściwości przed `=` dostaje `hl-attr`; `#true #false #null #inf #nan` (i `true false null` z wersji 1) `hl-builtin`; nazwa może zawierać `-` i `.`. |
| LANG-1003 | Komentarze `//` i zagnieżdżany `/* … */` dostają `hl-com`. |
| LANG-1004 | Napisy `"…"`, `"""…"""` i surowe `r#"…"#` dostają `hl-str`; liczby `hl-num`. |
