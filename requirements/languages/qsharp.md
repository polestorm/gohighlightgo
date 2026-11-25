# Wymagania: Q#

Źródło: Q# language reference (Microsoft Quantum). Definicja: `highlight/lang/qsharp.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2621 | Nazwy języka: `qsharp`, `qs`. |
| LANG-2622 | Słowa z listy keyword (`namespace`, `open`, `operation`, `function`, `let`, `mutable` …) dostają `hl-kw`; słowa z listy type (`Qubit`, `Result`, `Int`, `Double`, `Bool`, `String` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `Zero`, `One`, `PauliX`, `PauliZ` …) dostają `hl-builtin`; nazwa po `@` dostaje `hl-meta`. |
| LANG-2623 | Komentarze `//` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2624 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
