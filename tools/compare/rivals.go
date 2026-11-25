// Pakiet compare porównuje gohighlightgo z chroma (biblioteka Go) na korpusie projektu. Osobny moduł, żeby
// biblioteka nie miała zależności (ADR 0001, H6). chroma jest używana wyłącznie przez publiczne API, jako
// czarna skrzynka (docs/provenance.md).
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
var chromaNames = map[string]string{"sh": "bash", "pascal": "objectpascal"}

// Sample to plik korpusu bez pierwszego wiersza (odwołania REQ).
type Sample struct {
	Lang, Path string
	Code       []byte
}

// Corpus czyta testdata/corpus z katalogu głównego repozytorium (root).
func Corpus(root string) ([]Sample, error) {
	files, err := filepath.Glob(filepath.Join(root, "testdata/corpus/*/*.txt"))
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

// Ours koloruje kod biblioteką gohighlightgo.
func Ours(lang string, code []byte) []byte {
	out, _ := highlight.HTML(lang, code)
	return out
}

// ChromaLexer zwraca lekser chroma dla naszej nazwy języka albo nil.
func ChromaLexer(lang string) chroma.Lexer {
	name := lang
	if n, ok := chromaNames[lang]; ok {
		name = n
	}
	return lexers.Get(name)
}

var chromaFormatter = chromahtml.New(chromahtml.WithClasses(true), chromahtml.PreventSurroundingPre(true))

// Chroma koloruje kod biblioteką chroma: lekser języka i formater HTML z klasami (bez stylów w atrybutach),
// bez <pre>, czyli wynik tej samej postaci co nasz.
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

// ChromaLanguages zwraca liczbę lekserów chroma.
func ChromaLanguages() int { return len(lexers.GlobalLexerRegistry.Lexers) }

// PrimaryLanguages zwraca liczbę naszych języków (bez aliasów).
func PrimaryLanguages(samples []Sample) int {
	seen := map[string]bool{}
	for _, s := range samples {
		seen[s.Lang] = true
	}
	return len(seen)
}

// Repeat skleja kod próbek języka (albo wszystkich, gdy lang jest pusty) do co najmniej n bajtów.
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

// Spans liczy elementy <span w wyniku.
func Spans(out []byte) int { return strings.Count(string(out), "<span") }
