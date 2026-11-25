# Wymagania: TypoScript

Źródło: TYPO3 TypoScript reference (13.x): syntax. Definicja: `highlight/lang/typoscript.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2091 | Nazwy języka: `typoscript`, `typo3`. |
| LANG-2092 | Słowa z listy keyword (`page`, `config`, `lib`, `plugin`, `tt_content`, `TEXT` …) dostają `hl-kw`; słowa z listy builtin (`TSFE`, `GP`, `DB`, `LLL`) dostają `hl-builtin`. |
| LANG-2093 | Komentarze `#` do końca wiersza, `//` do końca wiersza, `/* … */` dostają `hl-com`. |
| LANG-2094 | Liczby dostają `hl-num`; znaki `=<>{}()[]:` dostają `hl-op`. |
