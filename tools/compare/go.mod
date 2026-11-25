module github.com/polestorm/gohighlightgo/tools/compare

go 1.25

require (
	github.com/alecthomas/chroma/v2 v2.27.0
	github.com/polestorm/gohighlightgo v0.0.0
)

require github.com/dlclark/regexp2/v2 v2.2.1 // indirect

replace github.com/polestorm/gohighlightgo => ../..
