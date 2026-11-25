// Package highlight turns source code into HTML with syntax highlighting.
//
// The package is a skeleton: the scope, the planned API and the acceptance criteria are described in
// docs/adr/0001-scope-and-acceptance.md. It has no dependencies outside the standard library and does
// not know about any Markdown parser; a caller connects it to one through that parser's hook for code
// blocks.
package highlight
