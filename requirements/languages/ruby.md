# Wymagania: Ruby

Źródło: dokumentacja Ruby 3.3, „Ruby Syntax” (słowa kluczowe, literały, komentarze). Definicja:
`highlight/lang/ruby.hl`.

| Id | Wymaganie |
|---|---|
| LANG-271 | Nazwy języka: `ruby`, `rb`. |
| LANG-272 | Słowa kluczowe dostają `hl-kw`; `true false nil self` i metody z listy (`puts`, `require`, `attr_accessor` …) `hl-builtin`; nazwa z wielkiej litery (stała, klasa, moduł) `hl-type`; `@zmienna`, `@@zmienna` i `$zmienna` dostają `hl-attr`. |
| LANG-273 | Komentarz `#` do końca wiersza i komentarz `=begin` … `=end` (na początku wiersza) dostają `hl-com`. |
| LANG-274 | Napisy `"…"` i `` `…` `` z ucieczkami i wstawkami `#{…}` oraz `'…'` dostają `hl-str`; literały `%w[…]`, `%i(…)`, `%q{…}`, `%Q<…>`, `%r{…}` i `%(…)` (po odstępie albo nawiasie) w całości `hl-str`; dokument `<<~KONIEC`, `<<-KONIEC`, `<<KONIEC` (nazwa z wielkiej litery albo w cudzysłowie) jak w PHP (`LANG-264`); wyrażenie regularne `/…/` jak w JavaScripcie (`LANG-048`). |
| LANG-275 | Symbol `:nazwa` (także z `?`, `!`, `=`) i klucz w skrócie `nazwa:` dostają `hl-attr`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
