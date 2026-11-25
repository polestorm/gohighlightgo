# Wymagania: Xojo

Źródło: Xojo documentation: language reference. Definicja: `highlight/lang/xojo.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2641 | Nazwy języka: `xojo`, `realbasic`. |
| LANG-2642 | Słowa z listy keyword (`Dim`, `Var`, `As`, `New`, `If`, `Then` …) dostają `hl-kw`; słowa z listy type (`Integer`, `Double`, `Boolean`, `String`, `Text`, `Variant` …) dostają `hl-type`; słowa z listy builtin (`True`, `False`, `System`, `MessageBox`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2643 | Komentarze `//` do końca wiersza, `'` do końca wiersza dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2644 | Liczby dostają `hl-num`; znaki `+-*/\^=<>&.` dostają `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
