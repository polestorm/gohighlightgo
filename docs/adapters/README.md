# Adaptery u odbiorców (łatki)

Kod dołączający kolorowanie w `gomdgo/cmd/gomdgo-mermaid` (znacznik `highlight`, ADR 0001, punkt 9) i w
`gopeekgo` (domyślnie włączone, znacznik `nohighlight` wyłącza: ADR 0003 w `gopeekgo`). Łatki leżą tutaj, bo
odbiorcy wskazują w `go.mod` tylko wydane wersje bibliotek, a `gohighlightgo` nie ma jeszcze wydania. Po pierwszym
wydaniu łatkę stosuje się w repozytorium odbiorcy i dopisuje zależność:

```sh
cd ../gomdgo && git apply ../gohighlightgo/docs/adapters/gomdgo-mermaid.patch
cd cmd/gomdgo-mermaid && go get github.com/polestorm/gohighlightgo@v0.1.0

cd ../gopeekgo && git apply ../gohighlightgo/docs/adapters/gopeekgo.patch
go get github.com/polestorm/gohighlightgo@v0.1.0
```

Do tego czasu oba warianty buduje się lokalnie przez `go.work` (bez commita), tak jak przy pomiarze H5
z 2026-11-25 (`docs/testing.md`).

Blok kodu wychodzi jako `<pre class="hl"><code class="language-…">`: postać jak w `gomdgo`, z klasą `hl`, w której
działa arkusz motywu (ADR 0006). `gomdgo-mermaid` używa motywu jasnego, `gopeekgo` motywu `auto`.
