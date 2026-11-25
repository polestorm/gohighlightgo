# Wymagania: q (kdb+)

Źródło: Q for Mortals and the kdb+ reference (code.kx.com). Definicja: `highlight/lang/q.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2271 | Nazwy języka: `q`, `kdb`. |
| LANG-2272 | Słowa z listy keyword (`select`, `exec`, `update`, `delete`, `from`, `where` …) dostają `hl-kw`; słowa z listy builtin (`count`, `sum`, `avg`, `min`, `max`, `first` …) dostają `hl-builtin`; nazwa po `` ` `` dostaje `hl-attr`; nazwa po `.` dostaje `hl-meta`. |
| LANG-2273 | Komentarze `/` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2274 | Liczby dostają `hl-num`; znaki `+-*%&|^!<>=:#$@,_~?'` dostają `hl-op`. |
