package highlight_test

import (
	"fmt"

	"github.com/polestorm/gohighlightgo/highlight"
)

func ExampleHTML() {
	out, ok := highlight.HTML("go", []byte(`fmt.Println("a < b")`))
	fmt.Println(ok)
	fmt.Println(string(out))
	// Output:
	// true
	// fmt.<span class="hl-fn">Println</span>(<span class="hl-str">"a &lt; b"</span>)
}

func ExampleHTML_unknown() {
	out, ok := highlight.HTML("cobol", []byte("DISPLAY 'HELLO'."))
	fmt.Println(out == nil, ok)
	// Output:
	// true false
}
