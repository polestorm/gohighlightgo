# Wymagania

Każde wymaganie ma identyfikator `PREFIKS-nnn`. Testy i pliki korpusu odwołują się do nich przez
`REQ: PREFIKS-nnn`; test `TestRequirementsCovered` sprawdza pokrycie w obie strony.

| Prefiks | Plik | Zakres |
|---|---|---|
| `API` | `api.md` | publiczne API, nazwy języków i aliasy |
| `OUT` | `output.md` | postać wyniku: klasy, zagnieżdżanie, ucieczki |
| `SAFE` | `safety.md` | bezpieczeństwo wyniku (H2, H3) |
| `DET` | `determinism.md` | powtarzalność, czas liniowy, odporność (H4) |
| `DEF` | `definitions.md` | format definicji języków, `Register` |
| `LANG` | `languages/<język>.md` | co kolorujemy w danym języku, ze wskazaniem rozdziału jego specyfikacji |

Numeracja `LANG` jest wspólna dla wszystkich języków: nowy język zaczyna od następnej wolnej dziesiątki
(Go 001–009, C 011–017, C++ 021–024, Java 031–036, JavaScript 041–048, TypeScript 051–052, Python 061–066,
Rust 071–077, SQL 081–085, JSON 091–094, Pascal 101–107, XML 111–115, HTML 121–123, CSS 131–135,
YAML 141–146, TOML 151–154, powłoka 161–166, diff 171–173, Markdown 181–186, C# 191–195, Kotlin 201–205,
Swift 211–215, Lua 221–225, Scala 231–235, Dart 241–245, R 251–255, PHP 261–265, Ruby 271–275, Haskell 281–285,
Zig 291–295, Perl 301–305, Objective-C 311–313, Elixir 321–325, Erlang 331–335, Clojure 341–345, F# 351–355,
OCaml 361–365, Julia 371–375, MATLAB 381–385, Groovy 391–395, PowerShell 401–405, Visual Basic 411–415,
Fortran 421–425, Makefile 431–434, Dockerfile 441–444, CMake 451–454, INI 461–464, properties 471–474,
Protobuf 481–484, GraphQL 491–494, HCL 501–504, Nix 511–514, LaTeX 521–524, Vim 531–534, AsciiDoc 541–544;
następny język od `LANG-551`). Każdy język z `Languages()` musi mieć plik
w `languages/` nazwany jego nazwą główną (`TestLanguageRequirements`).

Pliki powstają razem z etapami w `../docs/roadmap.md`.
