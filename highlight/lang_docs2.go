package highlight

import (
	"bytes"
	"strings"
)

// Org, reStructuredText i Textile (wymagania: requirements/languages/org.md, rst.md, textile.md). Kolorowane jest
// źródło dokumentu wiersz po wierszu; bloki kodu z nazwą języka koloruje ten język (bez dokumentów
// w dokumentach, żeby zagnieżdżenie było ograniczone).

var langOrg = language{names: []string{"org", "orgmode"}, scan: scanOrg}

var langRST = language{names: []string{"rst", "restructuredtext", "rest"}, scan: scanRST}

var langTextile = language{names: []string{"textile"}, scan: scanTextile}

// docLanguages: języki dokumentów, których nie kolorujemy wewnątrz bloków kodu innych dokumentów.
var docLanguages = map[string]bool{"markdown": true, "asciidoc": true, "org": true, "rst": true, "textile": true}

// codeBlock koloruje treść bloku kodu językiem lang albo, gdy go nie ma (albo to dokument), klasą c.
func codeBlock(w *writer, content []byte, lang string, c class) {
	if l := lookup(lang); l != nil && !docLanguages[l.names[0]] {
		l.scan(w, content)
		w.close()
		return
	}
	w.emit(c, content)
}

// lines przechodzi przez wiersze src: dla każdego wywołuje f z początkiem i końcem wiersza (bez \n) i zapisuje
// znaki \n; f zwraca początek następnego wiersza do przetworzenia (albo e, gdy przetworzył tylko ten).
func lines(w *writer, src []byte, f func(i, e int) int) {
	for i := 0; i < len(src); {
		e := lineEnd(src, i)
		next := f(i, e)
		if next > e {
			i = next // f zapisał kilka wierszy razem z ich \n
			continue
		}
		i = e
		if i < len(src) {
			w.emit(none, src[i:i+1])
			i++
		}
	}
}

// scanOrg: nagłówki (gwiazdki i odstęp w kolumnie 1) hl-kw, słowa kluczowe #+NAZWA: hl-meta, komentarz # hl-com,
// bloki #+BEGIN_SRC język … #+END_SRC w swoim języku, #+BEGIN_EXAMPLE hl-str, znaczniki list hl-op, ~kod~ i
// =dosłownie= hl-str, odnośniki [[…]] hl-str (Org Mode manual, „Document Structure”, „Markup for Rich Contents”).
func scanOrg(w *writer, src []byte) {
	lines(w, src, func(i, e int) int {
		line := src[i:e]
		trimmed := bytes.TrimLeft(line, " \t")
		upper := strings.ToUpper(string(trimmed))
		switch {
		case len(line) > 1 && line[0] == '*' && bytes.IndexFunc(line, func(r rune) bool { return r != '*' }) > 0 &&
			line[bytes.IndexFunc(line, func(r rune) bool { return r != '*' })] == ' ':
			w.emit(keyword, line)
		case strings.HasPrefix(upper, "#+BEGIN_"):
			kind, arg, _ := strings.Cut(upper[len("#+BEGIN_"):], " ")
			lang := strings.Fields(strings.ToLower(arg))
			c := class(0)
			switch kind {
			case "SRC":
				c = str
			case "EXAMPLE", "VERSE":
				c = str
			default:
				w.emit(meta, line)
				return e
			}
			name := ""
			if kind == "SRC" && len(lang) > 0 {
				name = lang[0]
			}
			return delimitedBlock(w, src, i, e, "#+END_"+kind, name, c)
		case strings.HasPrefix(upper, "#+"):
			w.emit(meta, line)
		case len(line) > 0 && line[0] == '#' && (len(line) == 1 || line[1] == ' '):
			w.emit(comment, line)
		default:
			k := len(line) - len(trimmed)
			w.emit(none, line[:k])
			if n := orgListMarker(trimmed); n > 0 {
				w.emit(operator, trimmed[:n])
				trimmed = trimmed[n:]
			}
			orgInline(w, trimmed)
		}
		return e
	})
}

// delimitedBlock zapisuje blok od wiersza otwierającego src[i:e] do wiersza, który po odstępach zaczyna się od
// end (bez względu na wielkość liter), i zwraca początek wiersza po nim. Treść koloruje codeBlock.
func delimitedBlock(w *writer, src []byte, i, e int, end, lang string, c class) int {
	w.emit(meta, src[i:e])
	if e >= len(src) {
		return e
	}
	w.emit(none, src[e:e+1])
	body := e + 1
	closeStart, closeEnd := len(src), len(src)
	for p := body; p < len(src); {
		q := lineEnd(src, p)
		t := bytes.TrimLeft(src[p:q], " \t")
		if len(t) >= len(end) && strings.EqualFold(string(t[:len(end)]), end) {
			closeStart, closeEnd = p, q
			break
		}
		if q >= len(src) {
			break
		}
		p = q + 1
	}
	codeBlock(w, src[body:closeStart], lang, c)
	w.emit(meta, src[closeStart:closeEnd])
	if closeEnd < len(src) {
		w.emit(none, src[closeEnd:closeEnd+1])
		return closeEnd + 1
	}
	return closeEnd
}

// orgListMarker: -, +, 1. albo 1) z odstępem po nim.
func orgListMarker(s []byte) int {
	if len(s) > 1 && (s[0] == '-' || s[0] == '+') && s[1] == ' ' {
		return 1
	}
	n := run(s, 0, 9, isDigit)
	if n > 0 && n+1 < len(s) && (s[n] == '.' || s[n] == ')') && s[n+1] == ' ' {
		return n + 1
	}
	return 0
}

// orgInline: ~kod~, =dosłownie= (znacznik na początku albo po odstępie) i odnośnik [[…]] dostają hl-str. Po
// nieudanym szukaniu zamknięcia danego znacznika (albo ]]) reszta wiersza nie jest przeszukiwana drugi raz.
func orgInline(w *writer, s []byte) {
	start := 0
	failed := [2]bool{}
	noLink := false
	for j := 0; j < len(s); j++ {
		n := 0
		switch b := s[j]; {
		case (b == '~' || b == '=') && (j == 0 || s[j-1] == ' ' || s[j-1] == '('):
			slot := 0
			if b == '=' {
				slot = 1
			}
			if failed[slot] {
				break
			}
			if k := bytes.IndexByte(s[j+1:], b); k > 0 {
				n = k + 2
			} else {
				failed[slot] = true
			}
		case b == '[' && j+1 < len(s) && s[j+1] == '[' && !noLink:
			if k := bytes.Index(s[j:], []byte("]]")); k > 0 {
				n = k + 2
			} else {
				noLink = true
			}
		}
		if n > 0 {
			w.emit(none, s[start:j])
			w.emit(str, s[j:j+n])
			j += n - 1
			start = j + 1
		}
	}
	w.emit(none, s[start:])
}

// scanRST: tytuły (wiersz nad podkreśleniem) hl-kw, podkreślenia i nadkreślenia hl-meta, dyrektywy .. nazwa::
// hl-meta (blok po .. code-block:: język w swoim języku), cele .. _nazwa: i przypisy hl-attr, komentarze ..
// hl-com, pola :nazwa: hl-attr, blok dosłowny po :: hl-str, “kod“ hl-str, role :rola:`…` hl-attr i hl-str
// (reStructuredText Markup Specification).
func scanRST(w *writer, src []byte) {
	lines(w, src, func(i, e int) int {
		line := src[i:e]
		trimmed := bytes.TrimLeft(line, " \t")
		indent := len(line) - len(trimmed)
		nextLine := []byte(nil)
		if e < len(src) {
			nextLine = bytes.TrimRight(src[e+1:lineEnd(src, e+1)], "\r")
		}
		switch {
		case isAdornment(line):
			w.emit(meta, line)
		case len(trimmed) > 0 && indent == 0 && isAdornment(nextLine) && len(nextLine) >= len(bytes.TrimRight(line, " \r")):
			w.emit(keyword, line)
		case bytes.HasPrefix(trimmed, []byte(".. ")):
			rest := trimmed[3:]
			if k := bytes.Index(rest, []byte("::")); k > 0 && !bytes.ContainsAny(rest[:k], " \t") {
				name := string(rest[:k])
				w.emit(none, line[:indent])
				w.emit(meta, trimmed[:3+k+2])
				arg := bytes.TrimSpace(rest[k+2:])
				w.emit(none, trimmed[3+k+2:])
				switch name {
				case "code-block", "code", "sourcecode":
					lang := ""
					if f := bytes.Fields(arg); len(f) > 0 {
						lang = strings.ToLower(string(f[0]))
					}
					return indentedBlock(w, src, e, indent, lang, str)
				}
				return e
			}
			if len(rest) > 0 && (rest[0] == '_' || rest[0] == '[') {
				w.emit(none, line[:indent])
				w.emit(attribute, trimmed)
				return e
			}
			w.emit(comment, line)
		case len(trimmed) > 1 && trimmed[0] == ':' && bytes.IndexByte(trimmed[1:], ':') > 0:
			k := bytes.IndexByte(trimmed[1:], ':') + 2
			w.emit(none, line[:indent])
			w.emit(attribute, trimmed[:k])
			rstInline(w, trimmed[k:])
		default:
			w.emit(none, line[:indent])
			if n := rstListMarker(trimmed); n > 0 {
				w.emit(operator, trimmed[:n])
				trimmed = trimmed[n:]
			}
			rstInline(w, trimmed)
			if bytes.HasSuffix(bytes.TrimRight(line, " \r"), []byte("::")) {
				return indentedBlock(w, src, e, indent, "", str)
			}
		}
		return e
	})
}

// isAdornment: wiersz z co najmniej dwóch takich samych znaków interpunkcji ze zbioru reST.
func isAdornment(line []byte) bool {
	line = bytes.TrimRight(line, " \r")
	if len(line) < 2 || strings.IndexByte("=-`:.'\"~^_*+#<>", line[0]) < 0 {
		return false
	}
	return len(bytes.Trim(line, string(line[0]))) == 0
}

// indentedBlock: po wierszu kończącym się w src[e] zapisuje blok wierszy pustych albo wciętych głębiej niż base
// (z opcjami dyrektywy :nazwa: na początku jako hl-attr) i zwraca jego koniec (początek następnego wiersza do
// przetworzenia albo e, gdy bloku nie ma). Treść koloruje codeBlock.
func indentedBlock(w *writer, src []byte, e, base int, lang string, c class) int {
	if e >= len(src) {
		return e
	}
	p := e + 1
	end := p
	for q := p; q < len(src); {
		r := lineEnd(src, q)
		line := src[q:r]
		t := bytes.TrimLeft(line, " \t")
		if len(bytes.TrimSpace(line)) > 0 && len(line)-len(t) <= base {
			break
		}
		end = r
		if r >= len(src) {
			break
		}
		q = r + 1
	}
	if end == p {
		return e
	}
	w.emit(none, src[e:e+1])
	body := src[p:end]
	// opcje dyrektywy (:linenos:) na początku bloku
	for len(body) > 0 {
		r := bytes.IndexByte(body, '\n')
		if r < 0 {
			r = len(body)
		}
		t := bytes.TrimSpace(body[:r])
		if len(t) < 2 || t[0] != ':' || bytes.IndexByte(t[1:], ':') < 0 {
			break
		}
		w.emit(attribute, body[:r])
		if r < len(body) {
			w.emit(none, body[r:r+1])
			r++
		}
		body = body[r:]
	}
	codeBlock(w, body, lang, c)
	if end < len(src) {
		w.emit(none, src[end:end+1])
		return end + 1
	}
	return end
}

// rstListMarker: -, *, + albo 1. z odstępem.
func rstListMarker(s []byte) int {
	if len(s) > 1 && (s[0] == '-' || s[0] == '*' || s[0] == '+') && s[1] == ' ' {
		return 1
	}
	n := run(s, 0, 9, isDigit)
	if n > 0 && n+1 < len(s) && (s[n] == '.' || s[n] == ')') && s[n+1] == ' ' {
		return n + 1
	}
	return 0
}

// rstInline: “kod“ hl-str, rola :nazwa:`tekst` (hl-attr i hl-str), `tekst`_ i `tekst` hl-str. Po nieudanym
// szukaniu zamknięcia reszta wiersza nie jest przeszukiwana drugi raz.
func rstInline(w *writer, s []byte) {
	start := 0
	noLiteral, noBacktick := false, false
	for j := 0; j < len(s); j++ {
		n, c := 0, str
		switch {
		case hasPrefix(s[j:], "``") && !noLiteral:
			if k := bytes.Index(s[j+2:], []byte("``")); k >= 0 {
				n = k + 4
			} else {
				noLiteral = true
			}
		case s[j] == ':' && j+1 < len(s) && isWordByte(s[j+1]):
			k := j + 1
			for k < len(s) && (isWordByte(s[k]) || s[k] == '-') {
				k++
			}
			if k+1 < len(s) && s[k] == ':' && s[k+1] == '`' {
				w.emit(none, s[start:j])
				w.emit(attribute, s[j:k+1])
				j = k
				start = k + 1
				continue
			}
		case s[j] == '`' && !noBacktick:
			if k := bytes.IndexByte(s[j+1:], '`'); k >= 0 {
				n = k + 2
				if j+n < len(s) && s[j+n] == '_' {
					n++
				}
			} else {
				noBacktick = true
			}
		}
		if n > 0 {
			w.emit(none, s[start:j])
			w.emit(c, s[j:j+n])
			j += n - 1
			start = j + 1
		}
	}
	w.emit(none, s[start:])
}

// scanTextile: znaczniki bloków h1.–h6. (cały wiersz hl-kw), bc. i pre. (treść hl-str), bq., p., fn1. (hl-meta),
// znaczniki list * i # hl-op, @kod@ hl-str, odnośnik "tekst":adres (adres hl-str) (Textile Language
// Specification, textile-lang.com).
func scanTextile(w *writer, src []byte) {
	lines(w, src, func(i, e int) int {
		line := src[i:e]
		dot := bytes.IndexByte(line, '.')
		tag := ""
		if dot > 0 && dot+1 <= len(line) && (dot+1 == len(line) || line[dot+1] == ' ') {
			tag = string(line[:dot])
		}
		switch {
		case len(tag) == 2 && tag[0] == 'h' && tag[1] >= '1' && tag[1] <= '6':
			w.emit(keyword, line)
		case tag == "bc" || tag == "pre":
			w.emit(meta, line[:dot+1])
			w.emit(str, line[dot+1:])
		case tag == "bq" || tag == "p" || tag == "notextile" || strings.HasPrefix(tag, "fn") && len(tag) > 2:
			w.emit(meta, line[:dot+1])
			textileInline(w, line[dot+1:])
		default:
			n := 0
			for n < len(line) && (line[n] == '*' || line[n] == '#') {
				n++
			}
			if n > 0 && n < len(line) && line[n] == ' ' {
				w.emit(operator, line[:n])
				line = line[n:]
			}
			textileInline(w, line)
		}
		return e
	})
}

// textileInline: @kod@ i adres po ": w odnośniku "tekst":adres dostają hl-str.
func textileInline(w *writer, s []byte) {
	start := 0
	noCode := false
	for j := 0; j < len(s); j++ {
		n := 0
		switch {
		case s[j] == '@' && !noCode:
			if k := bytes.IndexByte(s[j+1:], '@'); k > 0 {
				n = k + 2
			} else {
				noCode = true
			}
		case s[j] == '"' && j+1 < len(s) && s[j+1] == ':':
			k := j + 2
			for k < len(s) && !isSpace(s[k]) {
				k++
			}
			if k > j+2 {
				w.emit(none, s[start:j+2])
				w.emit(str, s[j+2:k])
				j = k - 1
				start = k
				continue
			}
		}
		if n > 0 {
			w.emit(none, s[start:j])
			w.emit(str, s[j:j+n])
			j += n - 1
			start = j + 1
		}
	}
	w.emit(none, s[start:])
}
