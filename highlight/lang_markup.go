package highlight

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// XML według Extensible Markup Language (XML) 1.0 (piąte wydanie), HTML według HTML Living Standard
// (rozdział 13, „The HTML syntax”); wymagania: requirements/languages/xml.md i html.md.

var langXML = language{
	names: []string{"xml", "xsd", "xsl", "xslt", "svg", "plist"},
	scan:  markup{}.scan,
}

var langHTML = language{
	names: []string{"html", "htm", "xhtml", "vue", "svelte"},
	scan:  markup{html: true}.scan,
}

// markup koloruje XML i HTML; w HTML treść <script> i <style> koloruje JavaScript i CSS.
type markup struct{ html bool }

func (m markup) scan(w *writer, src []byte) {
	for i := 0; i < len(src); {
		b := src[i]
		switch {
		case b == '<' && hasPrefix(src[i:], "<!--"):
			end := blockEnd(src, i+4, "", "-->", false)
			w.emit(comment, src[i:end])
			i = end
		case b == '<' && hasPrefix(src[i:], "<![CDATA["):
			w.emit(meta, src[i:i+9])
			j := i + 9
			n := bytes.Index(src[j:], []byte("]]>"))
			if n < 0 {
				w.emit(none, src[j:])
				return
			}
			w.emit(none, src[j:j+n])
			w.emit(meta, src[j+n:j+n+3])
			i = j + n + 3
		case b == '<' && i+1 < len(src) && (src[i+1] == '?' || src[i+1] == '!'):
			closing := ">"
			if src[i+1] == '?' {
				closing = "?>"
			}
			end := blockEnd(src, i+2, "", closing, false)
			w.emit(meta, src[i:end])
			i = end
		case b == '<' && i+1 < len(src) && (src[i+1] == '/' || isNameStart(src, i+1)):
			i = m.tag(w, src, i)
		case b == '&':
			if n := entityLen(src, i); n > 0 {
				w.emit(escape, src[i:i+n])
				i += n
			} else {
				w.emit(none, src[i:i+1])
				i++
			}
		default:
			j := i + 1
			for j < len(src) && src[j] != '<' && src[j] != '&' {
				j++
			}
			w.emit(none, src[i:j])
			i = j
		}
	}
}

// tag koloruje znacznik od < w src[i] i zwraca jego koniec. Nazwa znacznika z < (albo </) i zamykające > lub />
// dostają hl-tag, nazwy atrybutów hl-attr, wartości hl-str, odwołania znakowe w wartościach hl-esc.
func (m markup) tag(w *writer, src []byte, i int) int {
	j := i + 1
	closingTag := src[j] == '/'
	if closingTag {
		j++
	}
	k := nameEnd(src, j)
	name := src[j:k]
	w.emit(tag, src[i:k])
	afterEq := false
	for k < len(src) {
		c := src[k]
		switch {
		case isSpace(c):
			e := k + 1
			for e < len(src) && isSpace(src[e]) {
				e++
			}
			w.emit(none, src[k:e])
			k = e
		case c == '>':
			w.emit(tag, src[k:k+1])
			k++
			if m.html && !closingTag {
				return m.rawText(w, src, k, name)
			}
			return k
		case c == '/' && k+1 < len(src) && src[k+1] == '>':
			w.emit(tag, src[k:k+2])
			return k + 2
		case c == '<':
			return k // urwany znacznik: dalej zwykła treść
		case c == '=':
			w.emit(none, src[k:k+1])
			k++
			afterEq = true
		case c == '"' || c == '\'':
			k = attributeValue(w, src, k)
			afterEq = false
		default:
			e := k + 1
			for e < len(src) && !isSpace(src[e]) && src[e] != '>' && src[e] != '<' &&
				(afterEq || src[e] != '=' && src[e] != '/' && src[e] != '"' && src[e] != '\'') {
				e++
			}
			if afterEq {
				w.emit(str, src[k:e])
			} else {
				w.emit(attribute, src[k:e])
			}
			k = e
			afterEq = false
		}
	}
	return k
}

// rawText: treść <script> i <style> w HTML, do znacznika zamykającego (bez względu na wielkość liter),
// koloruje JavaScript albo CSS. Dla innych znaczników nic nie robi.
func (m markup) rawText(w *writer, src []byte, i int, name []byte) int {
	var inner func(*writer, []byte)
	switch {
	case bytes.EqualFold(name, []byte("script")):
		inner = lookup("javascript").scan
	case bytes.EqualFold(name, []byte("style")):
		inner = langCSS.scan
	default:
		return i
	}
	end := i
	for end < len(src) {
		n := bytes.IndexByte(src[end:], '<')
		if n < 0 {
			end = len(src)
			break
		}
		end += n
		if len(src)-end >= 2+len(name) && src[end+1] == '/' && bytes.EqualFold(src[end+2:end+2+len(name)], name) {
			break
		}
		end++
	}
	inner(w, src[i:end])
	return end
}

// attributeValue zapisuje wartość atrybutu w cudzysłowie od src[i] (może obejmować wiele wierszy).
func attributeValue(w *writer, src []byte, i int) int {
	q := src[i]
	start := i
	j := i + 1
	for j < len(src) && src[j] != q {
		if src[j] == '&' {
			if n := entityLen(src, j); n > 0 {
				w.emit(str, src[start:j])
				w.emit(escape, src[j:j+n])
				j += n
				start = j
				continue
			}
		}
		j++
	}
	if j < len(src) {
		j++
	}
	w.emit(str, src[start:j])
	return j
}

// entityLen zwraca długość odwołania znakowego od & w src[i] (&nazwa;, &#123;, &#x1F;) albo 0. Nazwa ma
// najwyżej 32 znaki, więc sprawdzenie jest ograniczone.
func entityLen(src []byte, i int) int {
	j := i + 1
	ok := func(b byte) bool { return isDigit(b) || b|0x20 >= 'a' && b|0x20 <= 'z' }
	if j < len(src) && src[j] == '#' {
		j++
		ok = isDigit
		if j < len(src) && src[j]|0x20 == 'x' {
			j++
			ok = isHex
		}
	}
	n := run(src, j, 32, ok)
	if n == 0 || j+n >= len(src) || src[j+n] != ';' {
		return 0
	}
	return j + n + 1 - i
}

// isNameStart mówi, czy w src[i] zaczyna się nazwa XML: litera, _ albo :.
func isNameStart(src []byte, i int) bool {
	r, _ := utf8.DecodeRune(src[i:])
	return r == '_' || r == ':' || isIdentStart(r)
}

// nameEnd zwraca koniec nazwy XML od src[i]: litery, cyfry, - . _ :.
func nameEnd(src []byte, i int) int {
	for i < len(src) {
		r, size := utf8.DecodeRune(src[i:])
		if !(r == '-' || r == '.' || r == '_' || r == ':' || isIdentPart(r) || r >= utf8.RuneSelf && unicode.IsMark(r)) {
			break
		}
		i += size
	}
	return i
}
