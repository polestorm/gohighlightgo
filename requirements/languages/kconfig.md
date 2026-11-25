# Wymagania: Kconfig

Źródło: Linux kernel Kconfig Language documentation. Definicja: `highlight/lang/kconfig.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2171 | Nazwy języka: `kconfig`. |
| LANG-2172 | Słowa z listy keyword (`config`, `menuconfig`, `choice`, `endchoice`, `menu`, `endmenu` …) dostają `hl-kw`; słowa z listy type (`bool`, `tristate`, `string`, `hex`, `int`) dostają `hl-type`; słowa z listy builtin (`y`, `n`, `m`) dostają `hl-builtin`. |
| LANG-2173 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc`; `'…'` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2174 | Liczby dostają `hl-num`; znaki `=!<>&|()` dostają `hl-op`. |
