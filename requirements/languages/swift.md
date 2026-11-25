# Wymagania: Swift

Źródło: The Swift Programming Language (Swift 5.9), „Lexical Structure”. Definicja: `highlight/lang/swift.hl`.

| Id | Wymaganie |
|---|---|
| LANG-211 | Nazwa języka: `swift`. |
| LANG-212 | Słowa kluczowe deklaracji, instrukcji, wyrażeń i kontekstowe z listy dostają `hl-kw`; typy biblioteki standardowej `hl-type`; `true false nil` i funkcje z listy `hl-builtin`; `$0` jest nazwą. |
| LANG-213 | Komentarze `//` i `/* … */`, które się zagnieżdżają, dostają `hl-com`. |
| LANG-214 | Napisy `"…"` i `"""…"""` (wiele wierszy) dostają `hl-str`, a ucieczki (`\n`, `\u{…}`) i wstawka `\( … )` `hl-esc`; napis surowy `#"…"#`, `##"…"##` i `#"""…"""#` dostaje w całości `hl-str`. |
| LANG-215 | `#if`, `#available` i inne `#nazwa` oraz atrybuty `@nazwa` dostają `hl-meta`; liczba nie zaczyna się od kropki; operatory `hl-op`; nazwa przed `(` dostaje `hl-fn`. |
