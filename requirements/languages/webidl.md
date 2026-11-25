# Wymagania: Web IDL

Źródło: Web IDL Standard (WHATWG). Definicja: `highlight/lang/webidl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2831 | Nazwy języka: `webidl`, `idl`. |
| LANG-2832 | Słowa z listy keyword (`interface`, `partial`, `dictionary`, `enum`, `typedef`, `callback` …) dostają `hl-kw`; słowa z listy type (`any`, `boolean`, `byte`, `octet`, `short`, `long` …) dostają `hl-type`; słowa z listy builtin (`true`, `false`, `null`, `Infinity`, `NaN`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-2833 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-2834 | Liczby dostają `hl-num`; znaki `=:;,?<>()[]{}.` dostają `hl-op`. |
