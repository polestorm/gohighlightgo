# Wymagania: Devicetree source

Źródło: Devicetree Specification (devicetree.org, v0.4). Definicja: `highlight/lang/dts.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2181 | Nazwy języka: `dts`, `dtsi`, `devicetree`. |
| LANG-2182 | Słowa z listy keyword (`/dts-v1/`, `/plugin/`, `/include/`, `/delete-node/`, `/delete-property/`) dostają `hl-kw`; słowa z listy builtin (`okay`, `disabled`, `compatible`, `reg`, `status`, `interrupts`) dostają `hl-builtin`; nazwa po `&` dostaje `hl-attr`. |
| LANG-2183 | Komentarze `//` do końca wiersza, `/* … */` dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2184 | Liczby dostają `hl-num`; znaki `=<>;{}@:` dostają `hl-op`; konstrukcje haka `c-preprocessor` jak w opisie haka w `docs/definitions.md`. |
