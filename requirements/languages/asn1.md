# Wymagania: ASN.1

Źródło: ITU-T X.680 (Abstract Syntax Notation One), clause 12 lexical items. Definicja: `highlight/lang/asn1.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2401 | Nazwy języka: `asn1`, `asn`. |
| LANG-2402 | Słowa z listy keyword (`DEFINITIONS`, `BEGIN`, `END`, `IMPORTS`, `EXPORTS`, `FROM` …) dostają `hl-kw`; słowa z listy type (`INTEGER`, `BOOLEAN`, `NULL`, `REAL`, `BIT`, `STRING` …) dostają `hl-type`; słowa z listy builtin (`TRUE`, `FALSE`, `MIN`, `MAX`) dostają `hl-builtin`; inne nazwy z wielkiej litery dostają `hl-type`. |
| LANG-2403 | Komentarze `--` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z podwojonym ogranicznikiem jako ucieczką dostają `hl-str`. |
| LANG-2404 | Liczby dostają `hl-num`; znaki `::=,{}()[]|.` dostają `hl-op`. |
