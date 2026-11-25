# Wymagania: CUDA C++

Źródło: CUDA C++ Programming Guide (NVIDIA, 12.x): language extensions. Definicja: `highlight/lang/cuda.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2341 | Nazwy języka: `cuda`, `cu`, `cuh`. |
| LANG-2342 | Wszystko jak w C++ (`requirements/languages/cpp.md`), a dodatkowo słowa z listy keyword (`__global__`, `__device__`, `__host__`, `__shared__`, `__constant__`, `__managed__` …) dostają `hl-kw`; słowa z listy type (`dim3`, `cudaError_t`, `cudaStream_t`, `float2`, `float3`, `float4` …) dostają `hl-type`; słowa z listy builtin (`threadIdx`, `blockIdx`, `blockDim`, `gridDim`, `warpSize`, `cudaMalloc` …) dostają `hl-builtin`. |
