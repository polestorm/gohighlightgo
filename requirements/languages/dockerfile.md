# Wymagania: Dockerfile

Źródło: Dockerfile reference (dokumentacja Dockera): instrukcje, „Environment replacement”, komentarze.

| Id | Wymaganie |
|---|---|
| LANG-441 | Nazwy języka: `dockerfile`, `docker`, `containerfile`. |
| LANG-442 | Instrukcje (`FROM`, `RUN`, `COPY`, `ENV` … i `AS`) bez rozróżniania wielkości liter dostają `hl-kw`. |
| LANG-443 | Komentarz `#` do końca wiersza dostaje `hl-com`. |
| LANG-444 | Zmienne `$NAZWA` i `${NAZWA}` dostają `hl-esc`; napisy `"…"` (ze wstawkami) i `'…'` `hl-str`; operatory powłoki `hl-op`. |
