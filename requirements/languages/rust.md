# Wymagania: Rust

Źródło: The Rust Reference (wydanie 2021) i dokumentacja preludium biblioteki standardowej (`std::prelude`);
w nawiasach rozdział.

| Id | Wymaganie |
|---|---|
| LANG-071 | Nazwy języka: `rust`, `rs`. |
| LANG-072 | Słowa ścisłe, zastrzeżone i słabe (`union`, `macro_rules`) dostają `hl-kw` („Keywords”); typy proste i typy z preludium `String Vec Option Result Box` dostają `hl-type`; `true false Some None Ok Err` dostają `hl-builtin`. |
| LANG-073 | Komentarze `//` i `/* … */`, które się zagnieżdżają, dostają `hl-com` („Comments”). |
| LANG-074 | Napisy `"…"`, `b"…"`, `c"…"` dostają `hl-str` i mogą obejmować wiele wierszy; sekwencje ucieczki (`\` z jednym znakiem, `\x` z dwiema cyframi, `\u{…}`) dostają `hl-esc`. Napisy surowe `r"…"`, `r#"…"#` z dowolną liczbą `#` (także `br`, `cr`) dostają `hl-str` w całości. Literały znakowe `'a'`, `'\n'`, `b'a'` dostają `hl-str` („Tokens”). |
| LANG-075 | Czas życia i etykieta `'nazwa` (apostrof bez zamknięcia po jednym znaku) dostają `hl-meta` („Lifetimes and loop labels”). |
| LANG-076 | Liczby jak w Go (`LANG-007`) z przyrostkami typów (`u8`, `f64` …); liczba nie zaczyna się od kropki, a kropka przed `.` albo identyfikatorem jej nie przedłuża (`1..10`, `1.max(2)`); w `t.0` liczbą jest samo `0` („Number literals”). |
| LANG-077 | Atrybut: `#[` albo `#![` z nazwą atrybutu dostaje `hl-meta` („Attributes”); identyfikator przed `!` (nie `!=`) i przed `(` dostaje `hl-fn` (makra i wywołania); operatory `+ - * / % & \| ^ ! < > = : ? ~ @` dostają `hl-op`. |
