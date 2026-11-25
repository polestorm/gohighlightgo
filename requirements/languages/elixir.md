# Wymagania: Elixir

Źródło: dokumentacja Elixira 1.16, „Syntax reference” i „Operators”. Definicja: `highlight/lang/elixir.hl`.

| Id | Wymaganie |
|---|---|
| LANG-321 | Nazwy języka: `elixir`, `ex`, `exs`. |
| LANG-322 | Słowa zastrzeżone i makra definiujące (`def`, `defmodule`, `case`, `with` …) dostają `hl-kw`; `true false nil` i `__MODULE__` `hl-builtin`; nazwa modułu z wielkiej litery `hl-type`; atrybut `@nazwa` `hl-meta`. |
| LANG-323 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-324 | Napisy `"…"`, `'…'` i wielowierszowe `"""…"""`, `'''…'''` z ucieczkami i wstawkami `#{…}`; sigile `~r/…/`, `~w(…)`, `~s"…"` (litera i dowolny ogranicznik, z modyfikatorami) w całości `hl-str`. |
| LANG-325 | Atom `:nazwa` i klucz `nazwa:` dostają `hl-attr`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
