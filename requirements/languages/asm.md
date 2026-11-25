# Wymagania: asembler x86

Źródło: NASM Manual (2.16) i Intel 64 and IA-32 Architectures Software Developer's Manual (lista instrukcji). Definicja: `highlight/lang/asm.hl`.

| Id | Wymaganie |
|---|---|
| LANG-701 | Nazwy języka: `asm`, `nasm`, `x86asm`, `gas`, `s`. |
| LANG-702 | Instrukcje i dyrektywy z listy (bez rozróżniania wielkości liter) dostają `hl-kw`; rejestry `hl-builtin`, także `%rax` w składni GNU; rozmiary (`byte`, `qword`, `ptr`) `hl-type`; dyrektywa `.nazwa` `hl-meta`. |
| LANG-703 | Komentarze `;`, `#` i `/* … */` dostają `hl-com`. |
| LANG-704 | Etykieta `nazwa:` na początku wiersza dostaje `hl-fn`; napisy `hl-str`; liczby `hl-num`. |
