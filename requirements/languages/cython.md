# Wymagania: Cython

Źródło: dokumentacja Cython (3.0), „Language basics”. Definicja: `highlight/lang/cython.hl`.

| Id | Wymaganie |
|---|---|
| LANG-851 | Nazwy języka: `cython`, `pyx`, `pxd`. |
| LANG-852 | Wszystko jak w Pythonie (`LANG-062`–`LANG-066`), a dodatkowo `cdef`, `cpdef`, `ctypedef`, `cimport`, `nogil` … dostają `hl-kw`, a typy C (`int`, `double`, `Py_ssize_t`, `bint` …) `hl-type`. |
