package highlight

// YAML według YAML Ain't Markup Language, wersja 1.2.2 (schemat podstawowy, „Core Schema”);
// wymagania: requirements/languages/yaml.md.
var langYAML = language{
	names: []string{"yaml", "yml"},
	scan:  scanYAML,
}

var (
	yamlStrings = &table{escape: escapeRule{continuation: true, hexX: 2, u4: true, u8: true}.length}
	yamlDouble  = quote{open: `"`, close: `"`, escapes: true, multiline: true}
	yamlSingle  = quote{open: "'", close: "'", doubled: true, multiline: true}
	yamlWords   = words(wordGroup{builtin, `null Null NULL ~ true True TRUE false False FALSE`})
)

// scanYAML koloruje YAML wiersz po wierszu. Klucz (skalar albo napis przed „: ”) dostaje hl-attr, wskaźniki
// - ? : | > hl-op, kotwice, aliasy i znaczniki (&a, *a, !tag) hl-meta, a treść skalara blokowego (| albo >)
// w wierszach wciętych głębiej niż wiersz ze wskaźnikiem hl-str.
func scanYAML(w *writer, src []byte) {
	flow := 0           // głębokość [] i {}
	blockIndent := -1   // ≥ 0: trwa skalar blokowy, jego treść to wiersze wcięte głębiej
	pendingIndent := -1 // ≥ 0: w tym wierszu stał wskaźnik | albo >, skalar zaczyna się od następnego
	lineIndent := 0
	lineStart := true
	for i := 0; i < len(src); {
		if lineStart {
			j := i
			for j < len(src) && src[j] == ' ' {
				j++
			}
			indent := j - i
			if blockIndent >= 0 {
				if j >= len(src) || src[j] == '\n' || src[j] == '\r' || indent > blockIndent {
					e := lineEnd(src, i)
					w.emit(none, src[i:j])
					w.emit(str, src[j:e])
					i = e
					if i < len(src) {
						w.emit(none, src[i:i+1])
						i++
					}
					continue
				}
				blockIndent = -1
			}
			w.emit(none, src[i:j])
			i, lineIndent, lineStart = j, indent, false
			if indent == 0 && (hasPrefix(src[i:], "---") || hasPrefix(src[i:], "...")) &&
				(i+3 >= len(src) || isSpace(src[i+3])) {
				w.emit(meta, src[i:i+3])
				i += 3
				continue
			}
			if indent == 0 && i < len(src) && src[i] == '%' {
				e := lineEnd(src, i)
				w.emit(meta, src[i:e])
				i = e
				continue
			}
			continue
		}
		b := src[i]
		next := byte('\n')
		if i+1 < len(src) {
			next = src[i+1]
		}
		switch {
		case b == '\n':
			w.emit(none, src[i:i+1])
			i++
			lineStart = true
			if pendingIndent >= 0 {
				blockIndent, pendingIndent = pendingIndent, -1
			}
		case b == ' ' || b == '\t' || b == '\r':
			w.emit(none, src[i:i+1])
			i++
		case b == '#' && (i == 0 || isSpace(src[i-1])):
			e := lineEnd(src, i)
			w.emit(comment, src[i:e])
			i = e
		case b == '"' || b == '\'':
			q := &yamlDouble
			if b == '\'' {
				q = &yamlSingle
			}
			c := str
			if end := yamlStrings.scanQuote(nil, src, i, i, q, false, str); yamlKeyColon(src, end, flow) {
				c = attribute
			}
			i = yamlStrings.scanQuote(w, src, i, i, q, false, c)
		case (b == '-' || b == '?' || b == ':') && (isSpace(next) || flow > 0 && b == ':' && isFlowIndicator(next)):
			w.emit(operator, src[i:i+1])
			i++
		case b == '[' || b == '{':
			flow++
			w.emit(none, src[i:i+1])
			i++
		case b == ']' || b == '}':
			if flow > 0 {
				flow--
			}
			w.emit(none, src[i:i+1])
			i++
		case b == ',' && flow > 0:
			w.emit(none, src[i:i+1])
			i++
		case b == '&' || b == '*' || b == '!':
			j := i + 1
			for j < len(src) && !isSpace(src[j]) && !isFlowIndicator(src[j]) {
				j++
			}
			w.emit(meta, src[i:j])
			i = j
		case (b == '|' || b == '>') && flow == 0:
			j := i + 1
			for j < len(src) && (src[j] == '-' || src[j] == '+' || isDigit(src[j])) {
				j++
			}
			w.emit(operator, src[i:j])
			i = j
			pendingIndent = lineIndent
		default:
			i = yamlPlain(w, src, i, flow)
		}
	}
}

// yamlPlain zapisuje skalar bez cudzysłowu od src[i]: do „: ”, „ #”, końca wiersza, a w [] i {} także do , [ ] { }.
// Przed dwukropkiem to klucz; liczba według schematu podstawowego dostaje hl-num, null i wartości logiczne
// hl-builtin.
func yamlPlain(w *writer, src []byte, i, flow int) int {
	j := i
	for j < len(src) {
		c := src[j]
		if c == '\n' || c == ':' && yamlKeyColon(src, j, flow) || c == '#' && j > i && isSpace(src[j-1]) ||
			flow > 0 && isFlowIndicator(c) && j > i {
			break
		}
		j++
	}
	e := j
	for e > i && (src[e-1] == ' ' || src[e-1] == '\t' || src[e-1] == '\r') {
		e--
	}
	if e == i {
		e = i + 1
	}
	word := src[i:e]
	c := yamlWords[string(word)]
	switch {
	case j < len(src) && src[j] == ':' && e == j:
		c = attribute
	case c == none && (isDigit(word[0]) || word[0] == '+' || word[0] == '-' || word[0] == '.') &&
		isYAMLNumber(word):
		c = number
	}
	w.emit(c, word)
	return e
}

// yamlKeyColon mówi, czy w src[j] stoi dwukropek kończący klucz: po nim odstęp, koniec danych albo,
// w [] i {}, wskaźnik przepływu.
func yamlKeyColon(src []byte, j, flow int) bool {
	if j >= len(src) || src[j] != ':' {
		return false
	}
	return j+1 >= len(src) || isSpace(src[j+1]) || flow > 0 && isFlowIndicator(src[j+1])
}

func isFlowIndicator(b byte) bool { return b == ',' || b == '[' || b == ']' || b == '{' || b == '}' }

// isYAMLNumber: liczby całkowite i zmiennoprzecinkowe schematu podstawowego YAML 1.2 (10.3.2).
func isYAMLNumber(s []byte) bool {
	if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'o') {
		ok := isHex
		if s[1] == 'o' {
			ok = isOctal
		}
		return run(s, 2, len(s), ok) == len(s)-2
	}
	k := 0
	if k < len(s) && (s[k] == '+' || s[k] == '-') {
		k++
	}
	switch string(s[k:]) {
	case ".inf", ".Inf", ".INF":
		return true
	case ".nan", ".NaN", ".NAN":
		return k == 0
	}
	whole := run(s, k, len(s), isDigit)
	k += whole
	frac := 0
	if k < len(s) && s[k] == '.' {
		frac = run(s, k+1, len(s), isDigit)
		k += 1 + frac
	}
	if whole+frac == 0 {
		return false
	}
	if k < len(s) && s[k]|0x20 == 'e' {
		k++
		if k < len(s) && (s[k] == '+' || s[k] == '-') {
			k++
		}
		n := run(s, k, len(s), isDigit)
		if n == 0 {
			return false
		}
		k += n
	}
	return k == len(s)
}
