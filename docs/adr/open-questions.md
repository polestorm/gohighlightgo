# Pytania otwarte

Wybory robocze z ADR 0001 czekają na potwierdzenie właściciela.

1. **Lista języków pierwszego wydania** i ich kolejność (ADR 0001, punkt 2). Czy Pascal/Delphi i formaty
   księgowe (XML z JPK) mają pierwszeństwo przed np. Rustem?
2. **Nazwy klas**: przedrostek `hl-` i lista z ADR 0001. Alternatywa: klasy zgodne z którymś istniejącym
   narzędziem, żeby działały cudze motywy (wtedy trzeba sprawdzić licencje motywów).
3. **Motywy**: jeden jasny i jeden ciemny arkusz w `highlight.CSS` czy więcej?
4. **Zasada rodowodu** (`docs/provenance.md`): czy przyjmujemy tę samą co w `gomdgo`, rozszerzoną o definicje
   języków cudzych narzędzi?
5. **Niepoprawny UTF-8 na wejściu**: przepisujemy bajty bez zmian (wtedy H2 obowiązuje zawsze) czy zamieniamy
   na U+FFFD (jak robi to `gomdgo`; sprawdzić w jego ADR 0006)?
6. **Wyjście inne niż HTML**: czy `gopdfgo` potrzebuje listy tokenów zamiast HTML?
7. **Próg fuzzingu** przed pierwszym wydaniem (w `gomdgo` to 24 h, jego ADR 0012).
