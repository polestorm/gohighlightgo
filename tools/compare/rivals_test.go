package compare

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const root = "../.."

func samples(t testing.TB) []Sample {
	s, err := Corpus(root)
	if err != nil || len(s) == 0 {
		t.Fatalf("corpus: %v", err)
	}
	return s
}

// TestRivalsLanguages wypisuje liczby języków i sprawdza, że chroma ma lekser dla każdego naszego języka.
func TestRivalsLanguages(t *testing.T) {
	s := samples(t)
	for _, sample := range s {
		if ChromaLexer(sample.Lang) == nil {
			t.Errorf("chroma has no lexer for %s", sample.Lang)
		}
	}
	defs, _ := filepath.Glob("/Applications/QLMarkdown.app/Contents/Resources/highlight/langDefs/*.lang")
	t.Logf("languages: gohighlightgo %d, chroma %d, highlight (QLMarkdown) %d", PrimaryLanguages(s), ChromaLanguages(), len(defs))
}

// TestRivalsOutput porównuje postać wyniku na korpusie: rozmiar fragmentu i liczbę elementów <span>.
func TestRivalsOutput(t *testing.T) {
	var oursBytes, chromaBytes, oursSpans, chromaSpans, in int
	for _, s := range samples(t) {
		o := Ours(s.Lang, s.Code)
		c, err := Chroma(ChromaLexer(s.Lang), s.Code)
		if err != nil {
			t.Fatalf("%s: %v", s.Path, err)
		}
		in += len(s.Code)
		oursBytes += len(o)
		chromaBytes += len(c)
		oursSpans += Spans(o)
		chromaSpans += Spans(c)
	}
	t.Logf("corpus %d B: gohighlightgo %d B, %d spans; chroma %d B, %d spans", in, oursBytes, oursSpans, chromaBytes, chromaSpans)
}

// TestRivalsSafety: obie biblioteki muszą zapisać znaki HTML z kodu jako odwołania.
func TestRivalsSafety(t *testing.T) {
	attack := []byte("</code></pre><script>alert(1)</script> \"<img src=x onerror=alert(1)>\" & <!-- -->")
	for _, s := range samples(t) {
		o := Ours(s.Lang, attack)
		c, err := Chroma(ChromaLexer(s.Lang), attack)
		if err != nil {
			t.Fatal(err)
		}
		for name, out := range map[string][]byte{"gohighlightgo": o, "chroma": c} {
			if bytes.Contains(out, []byte("<script")) || bytes.Contains(out, []byte("<img")) {
				t.Errorf("%s/%s: unescaped HTML: %s", name, s.Lang, out)
			}
		}
	}
}

// BenchmarkRivals: korpus każdego języka powtórzony do około 1 MB oraz całość; chroma z leksera i formatera HTML.
func BenchmarkRivals(b *testing.B) {
	s := samples(b)
	langs := []string{""}
	seen := map[string]bool{}
	for _, sample := range s {
		if !seen[sample.Lang] {
			seen[sample.Lang] = true
			langs = append(langs, sample.Lang)
		}
	}
	for _, lang := range langs {
		name := lang
		if name == "" {
			name = "all"
		}
		// "all": każdy plik korpusu osobno, jak wiele małych bloków w jednym dokumencie; język: jego korpus
		// sklejony do około 1 MB.
		parts := s
		if lang != "" {
			parts = []Sample{{Lang: lang, Code: Repeat(s, lang, 1<<20)}}
		}
		size := 0
		for _, p := range parts {
			size += len(p.Code)
		}
		b.Run(name+"/gohighlightgo", func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, p := range parts {
					Ours(p.Lang, p.Code)
				}
			}
		})
		b.Run(name+"/chroma", func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, p := range parts {
					if _, err := Chroma(ChromaLexer(p.Lang), p.Code); err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}

// TestMain pozwala wypisać wynik obu bibliotek dla jednej próbki: COMPARE_SHOW=go.
func TestMain(m *testing.M) {
	if lang := os.Getenv("COMPARE_SHOW"); lang != "" {
		s, _ := Corpus(root)
		for _, sample := range s {
			if sample.Lang == lang {
				c, _ := Chroma(ChromaLexer(lang), sample.Code)
				fmt.Println(strings.Repeat("=", 20), "gohighlightgo")
				fmt.Println(string(Ours(lang, sample.Code)))
				fmt.Println(strings.Repeat("=", 20), "chroma")
				fmt.Println(string(c))
				break
			}
		}
	}
	os.Exit(m.Run())
}
