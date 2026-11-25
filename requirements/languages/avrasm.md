# Wymagania: AVR assembly

Źródło: AVR Instruction Set Manual (Microchip). Definicja: `highlight/lang/avrasm.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2681 | Nazwy języka: `avrasm`, `avr`. |
| LANG-2682 | Słowa z listy keyword (`ldi`, `ld`, `st`, `lds`, `sts`, `mov` …) dostają `hl-kw`; słowa z listy builtin (`r0`, `r1`, `r16`, `r17`, `r18`, `r24` …) dostają `hl-builtin`; nazwa po `.` dostaje `hl-meta` (bez rozróżniania wielkości liter). |
| LANG-2683 | Komentarze `;` do końca wiersza, `//` do końca wiersza dostają `hl-com`; napisy `"…"` dostają `hl-str`. |
| LANG-2684 | Liczby dostają `hl-num`; znaki `+-*/(),=<>` dostają `hl-op`; konstrukcje haka `asm-label` jak w opisie haka w `docs/definitions.md`. |
