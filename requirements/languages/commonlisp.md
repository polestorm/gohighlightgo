# Wymagania: Common Lisp

Źródło: ANSI Common Lisp (Common Lisp HyperSpec), 2 „Syntax”, 3 „Evaluation and Compilation”. Definicja: `highlight/lang/commonlisp.hl`.

| Id | Wymaganie |
|---|---|
| LANG-631 | Nazwy języka: `commonlisp`, `lisp`, `cl`, `el`, `elisp`, `emacs-lisp`. |
| LANG-632 | Makra i formy specjalne z listy (`defun`, `let`, `loop`, `setf` …) dostają `hl-kw`; funkcje z listy oraz `t` i `nil` `hl-builtin`; słowo kluczowe `:nazwa` `hl-attr`; nazwa tuż po `(` dostaje `hl-fn`. |
| LANG-633 | Komentarz `;` i zagnieżdżany `#| … |#` dostają `hl-com`. |
| LANG-634 | Napis `"…"` i znak `#\a` dostają `hl-str`; cytat `'nazwa` `hl-meta`; liczby `hl-num`. |
