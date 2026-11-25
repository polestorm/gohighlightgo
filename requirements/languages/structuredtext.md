# Wymagania: IEC 61131-3 Structured Text

Źródło: IEC 61131-3 (3rd edition), 7.3 Structured Text. Definicja: `highlight/lang/structuredtext.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2601 | Nazwy języka: `structuredtext`, `iecst`, `scl`. |
| LANG-2602 | Słowa z listy keyword (`PROGRAM`, `END_PROGRAM`, `FUNCTION`, `END_FUNCTION`, `FUNCTION_BLOCK`, `END_FUNCTION_BLOCK` …) dostają `hl-kw`; słowa z listy type (`BOOL`, `BYTE`, `WORD`, `DWORD`, `INT`, `DINT` …) dostają `hl-type`; słowa z listy builtin (`TRUE`, `FALSE`) dostają `hl-builtin` (bez rozróżniania wielkości liter). |
| LANG-2603 | Komentarze `//` do końca wiersza, `(* … *)` dostają `hl-com`; napisy `'…'` z podwojonym ogranicznikiem jako ucieczką; `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2604 | Liczby dostają `hl-num`; znaki `:=+-*/<>.;,[]#` dostają `hl-op`. |
