# Wymagania: Org

Źródło: Org Mode manual (9.6): „Document Structure”, „Markup for Rich Contents”, „Working with Source Code”.

| Id | Wymaganie |
|---|---|
| LANG-1071 | Nazwy języka: `org`, `orgmode`. |
| LANG-1072 | Nagłówek (gwiazdki w kolumnie 1 i odstęp) dostaje `hl-kw`; słowo kluczowe `#+NAZWA:` `hl-meta`; komentarz `# …` `hl-com`; znaczniki list `-`, `+`, `1.` `hl-op`. |
| LANG-1073 | Blok `#+BEGIN_SRC język` … `#+END_SRC` (bez względu na wielkość liter): ograniczniki `hl-meta`, treść kolorowana tym językiem (bez dokumentów w dokumentach), a bez znanego języka `hl-str`; blok `#+BEGIN_EXAMPLE` dostaje `hl-str`. |
| LANG-1074 | `~kod~` i `=dosłownie=` (po odstępie albo na początku) oraz odnośnik `[[…]]` dostają `hl-str`. |
