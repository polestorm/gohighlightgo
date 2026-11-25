# Wymagania: Metal Shading Language

Źródło: Metal Shading Language Specification (Apple, 3.x). Definicja: `highlight/lang/metal.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2361 | Nazwy języka: `metal`, `msl`. |
| LANG-2362 | Wszystko jak w C++ (`requirements/languages/cpp.md`), a dodatkowo słowa z listy keyword (`kernel`, `vertex`, `fragment`, `device`, `threadgroup`, `constant` …) dostają `hl-kw`; słowa z listy type (`float2`, `float3`, `float4`, `float4x4`, `half4`, `uint2` …) dostają `hl-type`; słowa z listy builtin (`thread_position_in_grid`, `position`, `stage_in`, `buffer`) dostają `hl-builtin`. |
