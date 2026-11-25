# Wymagania: Kotlin

Źródło: Kotlin language specification (Kotlin 1.9), rozdział „Syntax and grammar”, „Keywords and operators”.
Definicja: `highlight/lang/kotlin.hl`.

| Id | Wymaganie |
|---|---|
| LANG-201 | Nazwy języka: `kotlin`, `kt`, `kts`. |
| LANG-202 | Słowa kluczowe twarde, miękkie i modyfikatory (bez częstych nazw: `get`, `set`, `field`, `value`) dostają `hl-kw`; typy podstawowe `hl-type`; `true false null` i funkcje biblioteki standardowej z listy dostają `hl-builtin`. |
| LANG-203 | Komentarze `//` i `/* … */`, które się zagnieżdżają, dostają `hl-com`. |
| LANG-204 | Napis `"…"` z ucieczkami i wstawkami `$nazwa`, `${…}` (`hl-esc`), napis surowy `"""…"""` (wiele wierszy, bez ucieczek, ze wstawkami) i znak `'…'` dostają `hl-str`. |
| LANG-205 | Adnotacja `@Nazwa` dostaje `hl-meta`; liczby `hl-num`; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
