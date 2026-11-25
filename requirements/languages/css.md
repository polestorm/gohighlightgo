# Wymagania: CSS

Źródło: CSS Syntax Module Level 3, Selectors Level 4, CSS Values and Units Level 4, CSS Cascading and
Inheritance Level 4 (W3C).

| Id | Wymaganie |
|---|---|
| LANG-131 | Nazwa języka: `css`. |
| LANG-132 | W selektorze nazwa elementu dostaje `hl-tag`, `.klasa` i `#id` dostają `hl-type`, `:pseudoklasa` i `::pseudoelement` dostają `hl-builtin`, kombinatory `> + ~` i `*` dostają `hl-op`. |
| LANG-133 | Wewnątrz `{}` identyfikator na początku instrukcji, po którym stoi `:`, to właściwość (także `--własna`) i dostaje `hl-attr`; do `;` albo `}` trwa wartość, w której liczby z jednostkami i `%` dostają `hl-num`, kolory `#fff` dostają `hl-num`, słowa `inherit initial unset auto none …` dostają `hl-builtin`, a nazwa przed `(` (`rgb(`, `var(`, `calc(`) dostaje `hl-fn`. W CSS zagnieżdżonym selektor z pseudoklasą na początku instrukcji wewnątrz `{}` (`a:hover {`) jest brany za właściwość. |
| LANG-134 | Reguła `@nazwa` i `!important` dostają `hl-kw`; nagłówek reguły `@` do `{` albo `;` jest kolorowany jak wartość. |
| LANG-135 | Komentarz `/* … */` dostaje `hl-com`; napisy w cudzysłowach i apostrofach `hl-str` z ucieczkami (`\` i do sześciu cyfr szesnastkowych albo `\` i jeden znak) `hl-esc`; adres w `url(…)` bez cudzysłowu dostaje `hl-str`. |
