// Program bez kolorowania: kopiuje standardowe wejście na wyjście. Punkt odniesienia dla rozmiaru.
package main

import (
	"io"
	"os"
)

func main() {
	src, _ := io.ReadAll(os.Stdin)
	os.Stdout.Write(src)
}
