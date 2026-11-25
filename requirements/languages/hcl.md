# Wymagania: HCL i Terraform

Źródło: HCL Native Syntax Specification i dokumentacja języka Terraform.

| Id | Wymaganie |
|---|---|
| LANG-501 | Nazwy języka: `hcl`, `terraform`, `tf`, `tfvars`. |
| LANG-502 | Typy bloków Terraform i słowa szablonów (`resource`, `variable`, `for`, `if` …) dostają `hl-kw`; `true false null` `hl-builtin`. |
| LANG-503 | Komentarze `#`, `//` i `/* … */` dostają `hl-com`. |
| LANG-504 | Napis `"…"` z ucieczkami i wstawkami `${…}` dostaje `hl-str`; dokument `<<EOT` i `<<-EOT` jak w Rubym (`LANG-274`); nazwa funkcji przed `(` `hl-fn`; liczby `hl-num`; operatory `hl-op`. |
