# Wymagania: Java

Źródło: The Java Language Specification, Java SE 21 Edition; w nawiasach rozdział.

| Id | Wymaganie |
|---|---|
| LANG-031 | Nazwa języka: `java`. |
| LANG-032 | Słowa kluczowe (3.9) i słowa kontekstowe `var record sealed permits yield when` dostają `hl-kw`; typy proste i `void` dostają `hl-type`; `true false null` dostają `hl-builtin`. Identyfikator może zawierać `$` (3.8). |
| LANG-033 | Komentarze `//` i `/* … */` dostają `hl-com` (3.7). |
| LANG-034 | Napisy, literały znakowe i bloki tekstu `"""…"""` (wielowierszowe) dostają `hl-str`; sekwencje ucieczki (`\` z jednym znakiem, do trzech cyfr ósemkowych, `\u` z dowolną liczbą `u` i czterema cyframi szesnastkowymi) dostają `hl-esc` (3.3, 3.10.5–3.10.7). |
| LANG-035 | Liczby jak w Go (`LANG-007`) z przyrostkami `L`, `f`, `d` (3.10.1, 3.10.2). |
| LANG-036 | Adnotacja `@Nazwa` dostaje `hl-meta` (9.7); operatory `+ - * / % & \| ^ ! < > = : ? ~` dostają `hl-op`; identyfikator przed `(` dostaje `hl-fn`. |
