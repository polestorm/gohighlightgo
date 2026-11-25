# Wymagania: OpenCL C

Źródło: The OpenCL C Specification (3.0). Definicja: `highlight/lang/opencl.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2351 | Nazwy języka: `opencl`, `ocl`. |
| LANG-2352 | Wszystko jak w C (`requirements/languages/c.md`), a dodatkowo słowa z listy keyword (`__kernel`, `kernel`, `__global`, `global`, `__local`, `local` …) dostają `hl-kw`; słowa z listy type (`half`, `float2`, `float4`, `int2`, `int4`, `uchar` …) dostają `hl-type`; słowa z listy builtin (`get_global_id`, `get_local_id`, `get_group_id`, `get_global_size`, `barrier`, `CLK_LOCAL_MEM_FENCE`) dostają `hl-builtin`. |
