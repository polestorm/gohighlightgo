# Wymagania: asembler ARM

Źródło: Arm Architecture Reference Manual (A32, T32, A64) i dialekt ARM w GNU as. Definicja: `highlight/lang/armasm.hl`.

| Id | Wymaganie |
|---|---|
| LANG-711 | Nazwy języka: `armasm`, `arm`, `aarch64`, `arm64`. |
| LANG-712 | Instrukcje z listy (bez rozróżniania wielkości liter) dostają `hl-kw`; rejestry (`r0`–`r15`, `x0`–`x30`, `w0`–`w15`, `sp`, `lr`, `pc`) `hl-builtin`; dyrektywa `.nazwa` `hl-meta`. |
| LANG-713 | Komentarze `@`, `//`, `;` i `/* … */` dostają `hl-com`. |
| LANG-714 | Etykieta `nazwa:` na początku wiersza dostaje `hl-fn`; napisy `hl-str`; liczby `hl-num`; `#` i inne operatory `hl-op`. |
