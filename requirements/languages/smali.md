# Wymagania: Smali

Źródło: smali/baksmali documentation (Dalvik bytecode assembler syntax). Definicja: `highlight/lang/smali.hl`.

| Id | Wymaganie |
|---|---|
| LANG-1811 | Nazwy języka: `smali`. |
| LANG-1812 | Słowa z listy keyword (`invoke-virtual`, `invoke-direct`, `invoke-static`, `invoke-super`, `invoke-interface`, `const` …) dostają `hl-kw`; słowa z listy builtin (`public`, `private`, `protected`, `static`, `final`, `abstract` …) dostają `hl-builtin`; nazwa po `.` dostaje `hl-meta`; nazwa po `:` dostaje `hl-fn`. |
| LANG-1813 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-1814 | Liczby dostają `hl-num`; znaki `->,{}()` dostają `hl-op`. |
