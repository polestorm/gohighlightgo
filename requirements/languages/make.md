# Wymagania: Makefile

Źródło: GNU make manual (4.4): „Rule Syntax”, „How to Use Variables”, „Conditional Parts of Makefiles”, „Automatic Variables”.

| Id | Wymaganie |
|---|---|
| LANG-431 | Nazwy języka: `make`, `makefile`, `mk`, `mak`. |
| LANG-432 | Dyrektywy (`ifeq`, `ifdef`, `include`, `define`, `export` …) dostają `hl-kw`; funkcje (`patsubst`, `wildcard`, `shell` …) `hl-builtin`. |
| LANG-433 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-434 | Rozwinięcia `$(…)`, `${…}` i zmienne automatyczne `$@ $< $^ $? $* $% $+ $|` oraz `$$` dostają `hl-esc`; nazwa celu przed `:` na początku wiersza dostaje `hl-fn`; napisy w cudzysłowach `hl-str`; operatory przypisania i `:` `hl-op`. |
