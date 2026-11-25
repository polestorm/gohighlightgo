# Wymagania: PowerShell

Źródło: PowerShell 7.4: about_Language_Keywords, about_Quoting_Rules, about_Comments, about_Operators. Definicja: `highlight/lang/powershell.hl`.

| Id | Wymaganie |
|---|---|
| LANG-401 | Nazwy języka: `powershell`, `ps1`, `pwsh`, `posh`. |
| LANG-402 | Słowa kluczowe (bez rozróżniania wielkości liter) dostają `hl-kw`; `$true $false $null` i każda zmienna `$nazwa` dostają `hl-attr`; operator albo parametr `-nazwa` (`-eq`, `-Path`) dostaje `hl-op`; nazwa polecenia `Verb-Noun` jest jedną nazwą. |
| LANG-403 | Komentarz `#` do końca wiersza i `<# … #>` dostają `hl-com`. |
| LANG-404 | Napisy `"…"` (ucieczka odwrotnym apostrofem `` ` ``, wstawki `$x`, `$(…)`) i `'…'` (`''` w środku) oraz here-strings `@"…"@` i `@'…'@` dostają `hl-str`. |
| LANG-405 | Liczby `hl-num`; operatory `hl-op`. |
