# Wymagania: Clojure

Źródło: Clojure Reference (1.11): „The Reader”, „Special Forms”. Definicja: `highlight/lang/clojure.hl`.

| Id | Wymaganie |
|---|---|
| LANG-341 | Nazwy języka: `clojure`, `clj`, `cljs`, `cljc`, `edn`. |
| LANG-342 | Formy specjalne i makra z listy (`def`, `defn`, `let`, `if`, `fn` …) dostają `hl-kw`; funkcje rdzenia z listy i `true false nil` `hl-builtin`; nazwa może zawierać `* ! ? + < > = - ' / .`. |
| LANG-343 | Komentarz `;` do końca wiersza dostaje `hl-com`. |
| LANG-344 | Napis `"…"` i wyrażenie regularne `#"…"` dostają `hl-str`; słowo kluczowe `:nazwa` (także `::nazwa`) dostaje `hl-attr`. |
| LANG-345 | Nazwa tuż po `(` dostaje `hl-fn` (wywołanie w składni Lispu); liczby `hl-num`. |
