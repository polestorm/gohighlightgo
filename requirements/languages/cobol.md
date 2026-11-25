# Wymagania: COBOL

Źródło: ISO/IEC 1989:2023 (słowa zastrzeżone, format źródła) i IBM Enterprise COBOL Language Reference. Definicja: `highlight/lang/cobol.hl`.

| Id | Wymaganie |
|---|---|
| LANG-721 | Nazwy języka: `cobol`, `cbl`, `cob`, `cpy`. |
| LANG-722 | Słowa zastrzeżone z listy (bez rozróżniania wielkości liter) dostają `hl-kw`; stałe figuratywne (`ZERO`, `SPACES`, `HIGH-VALUES` …) `hl-builtin`; nazwa może zawierać `-`. |
| LANG-723 | Komentarz `*>` do końca wiersza i wiersz z `*` albo `/` w kolumnie 7 formatu stałego (kolumny 1–6 to cyfry albo odstępy) dostają `hl-com`. |
| LANG-724 | Napisy `"…"` i `'…'` z podwojonym ogranicznikiem dostają `hl-str`; liczby `hl-num`; operatory `hl-op`. |
