# Wymagania: Solidity

Źródło: dokumentacja Solidity (0.8.x): „Layout of a Source File”, „Types”, „Units and Globally Available Variables”. Definicja: `highlight/lang/solidity.hl`.

| Id | Wymaganie |
|---|---|
| LANG-641 | Nazwy języka: `solidity`, `sol`. |
| LANG-642 | Słowa kluczowe i modyfikatory dostają `hl-kw`; typy (`address`, `uint256`, `bytes32` …) `hl-type`; `true false`, zmienne globalne (`msg`, `block`, `tx`), funkcje (`require`, `keccak256`) i jednostki (`ether`, `days`) `hl-builtin`. |
| LANG-643 | Komentarze `//` i `/* … */` dostają `hl-com`. |
| LANG-644 | Napisy `"…"` i `'…'` z ucieczkami, także z przedrostkiem `unicode` i `hex`, dostają `hl-str`; liczby `hl-num`; nazwa przed `(` `hl-fn`. |
