# Wymagania: XML

Źródło: Extensible Markup Language (XML) 1.0, piąte wydanie (W3C); w nawiasach rozdział. Kolorowanie ogólne,
bez wiedzy o schematach (pliki JPK, XSD i SVG koloruje się tak samo; decyzja właściciela z 2026-11-25).

| Id | Wymaganie |
|---|---|
| LANG-111 | Nazwy języka: `xml`, `xsd`, `xsl`, `xslt`, `svg`, `plist`. |
| LANG-112 | Znacznik: `<` albo `</` z nazwą (także z przedrostkiem przestrzeni nazw, np. `tns:Faktura`) oraz zamykające `>` albo `/>` dostają `hl-tag`; nazwy atrybutów `hl-attr`; wartości atrybutów w cudzysłowach lub apostrofach (także wielowierszowe) `hl-str` (3.1, 2.3). |
| LANG-113 | Odwołania znakowe `&nazwa;`, `&#123;`, `&#x1F;` w treści i w wartościach atrybutów dostają `hl-esc` (4.1). |
| LANG-114 | Komentarz `<!-- … -->` dostaje `hl-com`; instrukcja przetwarzania `<? … ?>` (także deklaracja XML) i deklaracja `<!DOCTYPE … >` dostają `hl-meta`; w sekcji CDATA ograniczniki `<![CDATA[` i `]]>` dostają `hl-meta`, a treść zostaje bez klasy (2.5–2.8, 2.7). Niezamknięte sięgają końca danych. |
| LANG-115 | Treść elementów zostaje bez klasy; `<`, po którym nie stoi nazwa, `/`, `!` ani `?`, jest zwykłym tekstem. |
