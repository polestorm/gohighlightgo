package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// TestInputsMatch: pliki w inputs/ to dokładnie te wejścia, których używa rivalprobe, więc każdy może je
// sprawdzić bez czytania kodu.
func TestInputsMatch(t *testing.T) {
	files := map[string]string{"CRLF line ends": "crlf.txt", "lone carriage return": "lone-cr.txt", "tabs": "tabs.txt",
		"no final newline": "no-final-newline.txt", "invalid UTF-8": "invalid-utf8.txt", "NUL byte": "nul.txt",
		"non-ASCII text": "non-ascii.txt"}
	if len(files) != len(tricky) {
		t.Fatalf("inputs/text has %d files, rivalprobe %d inputs", len(files), len(tricky))
	}
	for k, f := range files {
		b, err := os.ReadFile(filepath.Join("../../inputs/text", f))
		if err != nil || !bytes.Equal(b, tricky[k]) {
			t.Errorf("inputs/text/%s differs from %q (%v)", f, k, err)
		}
	}
	for p := range patterns {
		for _, c := range []struct {
			size string
			reps int
		}{{"small", smallReps}, {"large", bigReps}} {
			f := filepath.Join("../../inputs/hostile", patternNames[p]+"-"+c.size+".txt")
			b, err := os.ReadFile(f)
			if err != nil || !bytes.Equal(b, input(p, c.reps)) {
				t.Errorf("%s differs from pattern %q ×%d (%v)", f, patterns[p], c.reps, err)
			}
		}
	}
}
