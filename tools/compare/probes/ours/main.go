// Command ours highlights standard input with gohighlightgo; the language is the first argument.
package main

import (
	"io"
	"os"

	"github.com/polestorm/gohighlightgo/highlight"
)

func main() {
	src, _ := io.ReadAll(os.Stdin)
	out, _ := highlight.HTML(os.Args[1], src)
	os.Stdout.Write(out)
}
