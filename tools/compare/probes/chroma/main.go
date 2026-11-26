// Command chroma highlights standard input with chroma; the language is the first argument.
package main

import (
	"io"
	"os"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func main() {
	src, _ := io.ReadAll(os.Stdin)
	it, err := lexers.Get(os.Args[1]).Tokenise(nil, string(src))
	if err != nil {
		os.Exit(1)
	}
	html.New(html.WithClasses(true), html.PreventSurroundingPre(true)).Format(os.Stdout, styles.Fallback, it)
}
