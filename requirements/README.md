# Wymagania

Każde wymaganie ma identyfikator `PREFIKS-nnn` (od `LANG-1001` cztery cyfry). Testy i pliki korpusu odwołują się do nich przez
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
Protobuf 481–484, GraphQL 491–494, HCL 501–504, Nix 511–514, LaTeX 521–524, Vim 531–534, AsciiDoc 541–544,
Elm 551–554, Nim 561–564, Crystal 571–574, D 581–584, Ada 591–594, Prolog 601–604, Scheme 611–614,
Racket 621–622, Common Lisp 631–634, Solidity 641–644, Verilog 651–654, VHDL 661–664, GLSL 671–674, Tcl 681–684,
AWK 691–694, asembler x86 701–704, asembler ARM 711–714, COBOL 721–724, Smalltalk 731–734, Modula-2 741–744,
Oberon 751–754, MoonScript 761–764, fish 771–774, csh 781–784, pliki wsadowe 791–794, XQuery 801–804,
SPARQL 811–814, Cypher 821–824, PL/SQL 831–832, T-SQL 841–842, Cython 851–852, Mojo 861–862, V 871–874,
Odin 881–884, Hare 891–894, Gleam 901–904, Jinja 911–914, Handlebars 921–924, Liquid 931–933, ERB 941–944,
szablony Go 951–954, CSV 961–963, TSV 971–972, JSON5 981–984, Hjson 991–994, KDL 1001–1004, RON 1011–1014,
Jsonnet 1021–1024, Starlark 1031–1032, Bicep 1041–1044, Rego 1051–1054, Puppet 1061–1064; następny język od
`LANG-1071`). Każdy język z `Languages()` musi mieć plik
w `languages/` nazwany jego nazwą główną (`TestLanguageRequirements`).

Pliki powstają razem z etapami w `../docs/roadmap.md`.
