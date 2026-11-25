# Wymagania: G-code

Źródło: ISO 6983 and the RepRap G-code reference. Definicja: `highlight/lang/gcode.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2591 | Nazwy języka: `gcode`, `nc`, `ngc`. |
| LANG-2592 | Słowa z listy keyword (`G0`, `G00`, `G1`, `G01`, `G2`, `G02` …) dostają `hl-kw`; słowa z listy builtin (`M0`, `M00`, `M1`, `M01`, `M2`, `M02` …) dostają `hl-builtin`. |
| LANG-2593 | Komentarze `;` do końca wiersza, `( … )` dostają `hl-com`. |
| LANG-2594 | Liczby dostają `hl-num`. |
