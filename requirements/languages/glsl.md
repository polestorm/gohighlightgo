# Wymagania: GLSL

Źródło: The OpenGL Shading Language 4.60: 3 „Basics”, 4 „Variables and Types”, 8 „Built-in Functions”. Definicja: `highlight/lang/glsl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-671 | Nazwy języka: `glsl`, `vert`, `frag`, `hlsl`. |
| LANG-672 | Kwalifikatory i słowa kluczowe dostają `hl-kw`; typy (`vec3`, `mat4`, `sampler2D` …) `hl-type`; zmienne i funkcje wbudowane z listy `hl-builtin`. |
| LANG-673 | Komentarze `//` i `/* … */` dostają `hl-com`; dyrektywa `#version` i inne jak w C (`LANG-016`). |
| LANG-674 | Liczby (także z przyrostkiem `f` i `u`) dostają `hl-num`; operatory `hl-op`; nazwa przed `(` `hl-fn`. |
