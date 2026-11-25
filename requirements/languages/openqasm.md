# Wymagania: OpenQASM

Źródło: OpenQASM 3.0 Live Specification (openqasm.com). Definicja: `highlight/lang/openqasm.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2611 | Nazwy języka: `openqasm`, `qasm`. |
| LANG-2612 | Słowa z listy keyword (`OPENQASM`, `include`, `qubit`, `bit`, `qreg`, `creg` …) dostają `hl-kw`; słowa z listy type (`int`, `uint`, `float`, `angle`, `bool`, `complex` …) dostają `hl-type`; słowa z listy builtin (`h`, `x`, `y`, `z`, `cx`, `cz` …) dostają `hl-builtin`. |
| LANG-2613 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2614 | Liczby dostają `hl-num`; znaki `+-*/%&|^!<>=:?~.@[]->` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
