package highlight

import "bytes"

// Różnice według GNU Diffutils Manual (rozdziały „Unified Format” i „Normal Format”) oraz formatu git diff
// (dokumentacja git-diff, „Generating patch text”); wymagania: requirements/languages/diff.md.
var langDiff = language{
	names: []string{"diff", "patch"},
	scan:  scanDiff,
}

// diffHeaders: wiersze nagłówka poza fragmentem (hunk), które dostają hl-meta.
var diffHeaders = []string{"+++ ", "--- ", "diff ", "index ", "new file", "deleted file", "similarity ",
	"dissimilarity ", "rename ", "copy ", "old mode", "new mode", "Binary files", "Only in ", "*** "}

// scanDiff koloruje różnice wiersz po wierszu. Wewnątrz fragmentu (po @@ -a,b +c,d @@, dopóki liczby wierszy
// się nie wyczerpią) + to hl-ins, - to hl-del, \ to hl-com, a wiersz zaczynający się od --- albo +++ jest
// treścią, nie nagłówkiem. Poza fragmentem nagłówki dostają hl-meta, a + > i - < (format zwykły) hl-ins i hl-del.
func scanDiff(w *writer, src []byte) {
	oldLeft, newLeft := 0, 0
	for i := 0; i < len(src); {
		e := lineEnd(src, i)
		line := src[i:e]
		c := none
		inHunk := oldLeft > 0 || newLeft > 0
		first := byte(0)
		if len(line) > 0 {
			first = line[0]
		}
		switch {
		case inHunk && first == '+':
			c = inserted
			newLeft--
		case inHunk && first == '-':
			c = deleted
			oldLeft--
		case inHunk && (first == ' ' || first == 0 || first == '\r'):
			oldLeft--
			newLeft--
		case first == '\\':
			c = comment
		case hasPrefix(line, "@@"):
			oldLeft, newLeft = hunkCounts(line)
			end := len(line)
			for k := 2; k+1 < len(line); k++ {
				if line[k] == '@' && line[k+1] == '@' {
					end = k + 2
					break
				}
			}
			w.emit(meta, line[:end])
			w.emit(none, line[end:])
			line = nil
		default:
			oldLeft, newLeft = 0, 0
			for _, h := range diffHeaders {
				if hasPrefix(line, h) {
					c = meta
					break
				}
			}
			if c == none && string(bytes.TrimRight(line, "\r")) == "---" {
				c = meta // separator zmiany w formacie zwykłym
			}
			if c == none {
				switch first {
				case '+', '>':
					c = inserted
				case '-', '<':
					c = deleted
				}
			}
		}
		oldLeft, newLeft = max(oldLeft, 0), max(newLeft, 0)
		w.emit(c, line)
		i = e
		if i < len(src) {
			w.emit(none, src[i:i+1])
			i++
		}
	}
}

// hunkCounts odczytuje liczby wierszy z nagłówka @@ -a,b +c,d @@ (bez ,b liczba wynosi 1).
func hunkCounts(line []byte) (old, new int) {
	count := func(k int) (int, int) {
		for k < len(line) && isDigit(line[k]) {
			k++
		}
		if k >= len(line) || line[k] != ',' {
			return 1, k
		}
		n, k := 0, k+1
		for ; k < len(line) && isDigit(line[k]); k++ {
			if n < 1<<30 {
				n = n*10 + int(line[k]-'0')
			}
		}
		return n, k
	}
	k := 2
	for k < len(line) && line[k] == ' ' {
		k++
	}
	if k < len(line) && line[k] == '-' {
		old, k = count(k + 1)
	}
	for k < len(line) && line[k] == ' ' {
		k++
	}
	if k < len(line) && line[k] == '+' {
		new, _ = count(k + 1)
	}
	return old, new
}
