// Package compare compares gohighlightgo with chroma (a Go library) on this project's corpus. It is a separate
// module, so that the library itself has no dependencies. chroma is used only through its public API, as a
// black box.
package compare

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/polestorm/gohighlightgo/highlight"
)

// chromaNames: nazwa leksera chroma dla naszej nazwy głównej, gdy się różni.
var chromaNames = map[string]string{"sh": "bash", "pascal": "objectpascal", "commonlisp": "common-lisp", "gotemplate": "go-text-template"}

// Sample is a corpus file without its first line (the requirement references).
type Sample struct {
	Lang, Path string
	Code       []byte
}

// Corpus reads testdata/corpus under the repository root.
func Corpus(root string) ([]Sample, error) { return CorpusDir(filepath.Join(root, "testdata/corpus")) }

// CorpusDir reads the files dir/<language>/*.txt, each without its first line, in the order of their names.
func CorpusDir(dir string) ([]Sample, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*/*.txt"))
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	var out []Sample
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		_, rest, _ := bytes.Cut(b, []byte("\n"))
		out = append(out, Sample{Lang: filepath.Base(filepath.Dir(f)), Path: f, Code: rest})
	}
	return out, nil
}

// Ours highlights code with gohighlightgo.
func Ours(lang string, code []byte) []byte {
	out, _ := highlight.HTML(lang, code)
	return out
}

// ChromaLexer returns chroma's lexer for a gohighlightgo language name, or nil.
func ChromaLexer(lang string) chroma.Lexer {
	name := lang
	if n, ok := chromaNames[lang]; ok {
		name = n
	}
	return lexers.Get(name)
}

var chromaFormatter = chromahtml.New(chromahtml.WithClasses(true), chromahtml.PreventSurroundingPre(true))

// Chroma highlights code with chroma: the language's lexer and the HTML formatter with classes (no inline
// styles) and without <pre>, which gives output of the same form as gohighlightgo's.
func Chroma(lexer chroma.Lexer, code []byte) ([]byte, error) {
	it, err := lexer.Tokenise(nil, string(code))
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	if err := chromaFormatter.Format(&b, styles.Fallback, it); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// ChromaKeepLF works like Chroma but without converting line ends to LF (the EnsureLF option), so that a check
// of text preservation is about the lexer, not about that default.
func ChromaKeepLF(lexer chroma.Lexer, code []byte) ([]byte, error) {
	it, err := lexer.Tokenise(&chroma.TokeniseOptions{State: "root", EnsureLF: false}, string(code))
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	if err := chromaFormatter.Format(&b, styles.Fallback, it); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// ChromaLanguages returns the number of chroma's lexers.
func ChromaLanguages() int { return len(lexers.GlobalLexerRegistry.Lexers) }

// PrimaryLanguages returns the number of languages in samples (without aliases).
func PrimaryLanguages(samples []Sample) int {
	seen := map[string]bool{}
	for _, s := range samples {
		seen[s.Lang] = true
	}
	return len(seen)
}

// Repeat joins the code of the samples of a language (or of all, when lang is empty) to at least n bytes.
func Repeat(samples []Sample, lang string, n int) []byte {
	var one []byte
	for _, s := range samples {
		if lang == "" || s.Lang == lang {
			one = append(one, s.Code...)
		}
	}
	if len(one) == 0 {
		return nil
	}
	return bytes.Repeat(one, n/len(one)+1)
}

// Spans counts the <span elements in out.
func Spans(out []byte) int { return strings.Count(string(out), "<span") }
