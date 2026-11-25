# Wymagania: Modelica

Źródło: Modelica Language Specification 3.6, chapter 2 Lexical Structure. Definicja: `highlight/lang/modelica.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1691 | Nazwy języka: `modelica`, `mo`. |
| LANG-1692 | Słowa z listy keyword (`model`, `class`, `record`, `block`, `connector`, `type` …) dostają `hl-kw`; słowa z listy type (`Real`, `Integer`, `Boolean`, `String`) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `time`, `pi`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-1693 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`, wielowierszowy dostają `hl-str`. |
| LANG-1694 | Liczby dostają `hl-num`; znaki `+-*/^=<>:.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
