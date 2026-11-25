# Wymagania: reStructuredText

Źródło: reStructuredText Markup Specification (Docutils).

| Id | Wymaganie |
|---|---|
| LANG-1081 | Nazwy języka: `rst`, `restructuredtext`, `rest`. |
| LANG-1082 | Tytuł (wiersz nad podkreśleniem co najmniej tej długości) dostaje `hl-kw`; podkreślenie i nadkreślenie `hl-meta`; dyrektywa `.. nazwa::` `hl-meta`; cel `.. _nazwa:` i przypis `.. [1]` `hl-attr`; inny wiersz `.. …` `hl-com`; pole `:nazwa:` `hl-attr`; znaczniki list `hl-op`. |
| LANG-1083 | Blok wcięty po `.. code-block:: język` (także `code`, `sourcecode`) jest kolorowany tym językiem, z opcjami `:nazwa:` na początku jako `hl-attr`; blok wcięty po wierszu kończącym się `::` dostaje `hl-str`. |
| LANG-1084 | ``` ``kod`` ```, `` `tekst` `` i `` `tekst`_ `` dostają `hl-str`; rola `:nazwa:` przed `` ` `` `hl-attr`. |
