# Wymagania: Python

Źródło: The Python Language Reference i The Python Standard Library, wersja 3.12; w nawiasach rozdział.

| Id | Wymaganie |
|---|---|
| LANG-061 | Nazwy języka: `python`, `py`, `python3`. |
| LANG-062 | Słowa kluczowe (Reference 2.3.1) dostają `hl-kw`, z wyjątkiem `True False None`, które razem z `NotImplemented`, `Ellipsis` i funkcjami wbudowanymi (Library, „Built-in Functions”) dostają `hl-builtin`; klasy wbudowane `bool bytearray bytes complex dict float frozenset int list memoryview object range set slice str tuple type` dostają `hl-type`. Słowa kluczowe miękkie (`match`, `case`, `type` jako polecenie) nie są rozpoznawane. |
| LANG-063 | Komentarz `#` sięga końca wiersza (2.1.3). |
| LANG-064 | Napisy `'…'`, `"…"` (jednowierszowe) i `'''…'''`, `"""…"""` (wielowierszowe), z przedrostkami `r u f b br rb fr rf` w dowolnej wielkości liter, dostają `hl-str` razem z przedrostkiem; sekwencje ucieczki (`\` z jednym znakiem, do trzech cyfr ósemkowych, `\x` z dwiema, `\u` z czterema, `\U` z ośmioma cyframi, `\N{…}`) dostają `hl-esc`; w napisie z przedrostkiem `r` nie ma sekwencji ucieczki, ale `\` chroni następny znak przed zamknięciem napisu (2.4.1). Wyrażenia w `{}` napisów `f` nie są kolorowane osobno. |
| LANG-065 | Liczby jak w Go (`LANG-007`) z przyrostkiem `j` liczb urojonych (2.4.5–2.4.7). |
| LANG-066 | Dekorator `@nazwa` dostaje `hl-meta`; operatory `+ - * / % @ & \| ^ ~ < > = ! :` dostają `hl-op`; identyfikator przed `(` dostaje `hl-fn`. |
