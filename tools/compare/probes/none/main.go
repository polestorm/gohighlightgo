// Command none copies standard input to standard output without highlighting: the baseline for size and memory.
package main

import (
	"io"
	"os"
)

func main() {
	src, _ := io.ReadAll(os.Stdin)
	os.Stdout.Write(src)
}
