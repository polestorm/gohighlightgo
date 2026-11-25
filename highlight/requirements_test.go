package highlight

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var reqRef = regexp.MustCompile(`\b[A-Z]{3,5}-\d{3}\b`)

// corpusFiles zwraca pliki korpusu: testdata/corpus/<język>/<nazwa>.txt, każdy z wzorcem <nazwa>.html.
func corpusFiles(t testing.TB) []string {
	files, _ := filepath.Glob("../testdata/corpus/*/*.txt")
	if len(files) == 0 {
		t.Fatal("no corpus files")
	}
	return files
}

// TestRequirementsCovered mierzy kryterium H1: każde wymaganie z requirements/ ma test albo plik korpusu,
// a każde odwołanie „REQ:” wskazuje istniejące wymaganie.
func TestRequirementsCovered(t *testing.T) {
	read := func(path string) string {
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return string(b)
	}
	defined := map[string]bool{}
	docs, _ := filepath.Glob("../requirements/*.md")
	langs, _ := filepath.Glob("../requirements/languages/*.md")
	for _, doc := range append(docs, langs...) {
		for _, line := range strings.Split(read(doc), "\n") {
			if id := reqRef.FindString(line); id != "" && strings.HasPrefix(line, "| "+id) {
				if defined[id] {
					t.Errorf("requirement %s is defined twice", id)
				}
				defined[id] = true
			}
		}
	}
	if len(defined) == 0 {
		t.Fatal("no requirements found")
	}
	corpus := corpusFiles(t)
	tests, _ := filepath.Glob("*_test.go")
	used := map[string]bool{}
	for _, f := range append(append([]string{}, corpus...), tests...) {
		for _, line := range strings.Split(read(f), "\n") {
			if !strings.Contains(line, "REQ:") {
				continue
			}
			for _, id := range reqRef.FindAllString(line, -1) {
				used[id] = true
				if !defined[id] {
					t.Errorf("%s refers to unknown requirement %s", f, id)
				}
			}
		}
	}
	for id := range defined {
		if !used[id] {
			t.Errorf("requirement %s has no corpus file or test", id)
		}
	}
	for _, f := range corpus {
		if first, _, _ := strings.Cut(read(f), "\n"); !strings.Contains(first, "REQ:") {
			t.Errorf("%s: the first line has no REQ reference", f)
		}
	}
}

// TestLanguageRequirements: każdy język z Languages() ma plik wymagań i katalog korpusu (kryterium H1).
func TestLanguageRequirements(t *testing.T) {
	for _, l := range languages {
		name := l.names[0]
		if _, err := os.Stat("../requirements/languages/" + name + ".md"); err != nil {
			t.Errorf("language %s has no requirements file: %v", name, err)
		}
		if m, _ := filepath.Glob("../testdata/corpus/" + name + "/*.txt"); len(m) == 0 {
			t.Errorf("language %s has no corpus files", name)
		}
	}
}
