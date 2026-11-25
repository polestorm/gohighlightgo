package highlight

import (
	"bytes"
	"strings"
)

// INI i pliki properties (Java Properties.load) oraz AsciiDoc (AsciiDoc Language Documentation); wymagania:
// requirements/languages/ini.md, properties.md, asciidoc.md.

var langINI = language{
	names: []string{"ini", "cfg", "conf", "desktop", "gitconfig", "editorconfig", "systemd", "service", "reg"},
	scan:  configScanner{}.scan,
}

var langProperties = language{
	names: []string{"properties", "props"},
	scan:  configScanner{properties: true}.scan,
}

var langAsciiDoc = language{
	names: []string{"asciidoc", "adoc", "asc"},
	scan:  scanAsciiDoc,
}

// configScanner koloruje plik klucz–wartość wiersz po wierszu: komentarz na początku wiersza (; i #, w properties
// # i !) dostaje hl-com, sekcja [nazwa] hl-tag, klucz przed = albo : hl-attr, separator hl-op, wartość: napis
// w cudzysłowie hl-str, liczba hl-num, true false yes no on off hl-builtin, ${…} hl-esc. W properties \ z znakiem
// to hl-esc, a \ na końcu wiersza przedłuża wartość na następny wiersz.
type configScanner struct{ properties bool }

func (c configScanner) scan(w *writer, src []byte) {
	continued := false
	for i := 0; i < len(src); {
		e := lineEnd(src, i)
		line := src[i:e]
		k := 0
		for k < len(line) && (line[k] == ' ' || line[k] == '\t') {
			k++
		}
		w.emit(none, line[:k])
		rest := line[k:]
		switch {
		case continued:
			c.value(w, rest)
		case len(rest) == 0:
		case rest[0] == ';' && !c.properties || rest[0] == '#' || rest[0] == '!' && c.properties:
			w.emit(comment, rest)
		case rest[0] == '[' && !c.properties:
			end := bytes.IndexByte(rest, ']')
			if end < 0 {
				end = len(rest) - 1
			}
			w.emit(tag, rest[:end+1])
			w.emit(none, rest[end+1:])
		default:
			sep := bytes.IndexAny(rest, "=:")
			if sep < 0 && c.properties {
				sep = bytes.IndexAny(rest, " \t")
			}
			if sep < 0 {
				w.emit(attribute, rest)
				break
			}
			key := bytes.TrimRight(rest[:sep], " \t")
			w.emit(attribute, key)
			w.emit(none, rest[len(key):sep])
			w.emit(operator, rest[sep:sep+1])
			c.value(w, rest[sep+1:])
		}
		continued = c.properties && bytes.HasSuffix(bytes.TrimRight(line, "\r"), []byte{'\\'}) && !continuedEscape(line)
		i = e
		if i < len(src) {
			w.emit(none, src[i:i+1])
			i++
		}
	}
}

// continuedEscape: \ na końcu wiersza jest sam znakiem ucieczki (\\), a nie przedłużeniem, przy parzystej
// liczbie odwrotnych ukośników.
func continuedEscape(line []byte) bool {
	line = bytes.TrimRight(line, "\r")
	n := 0
	for n < len(line) && line[len(line)-1-n] == '\\' {
		n++
	}
	return n%2 == 0
}

// value zapisuje wartość klucza.
func (c configScanner) value(w *writer, v []byte) {
	k := 0
	for k < len(v) && (v[k] == ' ' || v[k] == '\t') {
		k++
	}
	w.emit(none, v[:k])
	v = v[k:]
	trimmed := strings.TrimSpace(string(v))
	switch {
	case len(trimmed) >= 2 && (trimmed[0] == '"' || trimmed[0] == '\'') && trimmed[len(trimmed)-1] == trimmed[0]:
		w.emit(str, v)
		return
	case isConfigNumber(trimmed):
		w.emit(number, v)
		return
	}
	switch strings.ToLower(trimmed) {
	case "true", "false", "yes", "no", "on", "off":
		w.emit(builtin, v)
		return
	}
	start := 0
	for j := 0; j < len(v); j++ {
		n := 0
		switch {
		case v[j] == '$' && j+1 < len(v) && v[j+1] == '{':
			n = bracketLen(v, j, '{', '}', j+1)
		case c.properties && v[j] == '\\' && j+1 < len(v) && v[j+1] == 'u':
			n = 2 + run(v, j+2, 4, isHex)
		case c.properties && v[j] == '\\' && j+1 < len(v):
			n = 1 + utf8Len(v, j+1)
		case c.properties && v[j] == '\\':
			n = 1 // \ na końcu wiersza: przedłużenie wartości
		}
		if n > 0 {
			w.emit(none, v[start:j])
			w.emit(escape, v[j:j+n])
			j += n - 1
			start = j + 1
		}
	}
	w.emit(none, v[start:])
}

func isConfigNumber(s string) bool {
	if s == "" {
		return false
	}
	s = strings.TrimLeft(s, "+-")
	dot := false
	for i := 0; i < len(s); i++ {
		switch {
		case isDigit(s[i]):
		case s[i] == '.' && !dot:
			dot = true
		default:
			return false
		}
	}
	return s != "" && s != "."
}

// scanAsciiDoc koloruje dokument AsciiDoc wiersz po wierszu: tytuły (= na początku wiersza) hl-kw, wpisy
// atrybutów :nazwa: hl-attr, atrybuty bloku [ … ] hl-meta, ograniczniki bloków (----, ...., ====, ****, ____, ++++,
// |===) hl-meta, komentarze // i bloki ////, znaczniki list hl-op, kod `…` hl-str i odwołania {nazwa} hl-esc. Blok
// ---- po [source,język] jest kolorowany tym językiem (bez AsciiDoc w AsciiDoc), bez niego dostaje hl-str, tak
// samo blok ....
func scanAsciiDoc(w *writer, src []byte) {
	pending := "" // język z ostatniego [source,…]
	for i := 0; i < len(src); {
		e := lineEnd(src, i)
		line := bytes.TrimRight(src[i:e], "\r")
		next := e
		switch {
		case isDelimiter(line, '/') && len(line) >= 4:
			next = adocBlock(w, src, i, e, line, comment, "")
		case isDelimiter(line, '-') && len(line) >= 4, isDelimiter(line, '.') && len(line) >= 4:
			next = adocBlock(w, src, i, e, line, str, pending)
			pending = ""
		case bytes.HasPrefix(line, []byte("//")):
			w.emit(comment, src[i:e])
		case isDelimiter(line, '=') && len(line) >= 4, isDelimiter(line, '*') && len(line) >= 4,
			isDelimiter(line, '_') && len(line) >= 4, isDelimiter(line, '+') && len(line) >= 4,
			bytes.HasPrefix(line, []byte("|===")):
			w.emit(meta, src[i:e])
		case len(line) > 1 && line[0] == '=' && (adocTitle(line)):
			w.emit(keyword, src[i:e])
		case len(line) > 2 && line[0] == ':' && bytes.IndexByte(line[1:], ':') > 0:
			n := bytes.IndexByte(line[1:], ':') + 2
			w.emit(attribute, src[i:i+n])
			adocInline(w, src[i+n:e])
		case len(line) > 1 && line[0] == '[' && line[len(line)-1] == ']':
			w.emit(meta, src[i:e])
			inside := string(line[1 : len(line)-1])
			if name, rest, ok := strings.Cut(inside, ","); ok && (name == "source" || name == "") {
				lang, _, _ := strings.Cut(rest, ",")
				pending = strings.TrimSpace(lang)
			}
		default:
			k := 0
			for k < len(line) && (line[k] == ' ' || line[k] == '\t') {
				k++
			}
			m := k
			for m < len(line) && (line[m] == '*' || line[m] == '-' || line[m] == '.') {
				m++
			}
			if m > k && m < len(line) && line[m] == ' ' {
				w.emit(none, src[i:i+k])
				w.emit(operator, src[i+k:i+m])
				adocInline(w, src[i+m:e])
			} else {
				adocInline(w, src[i:e])
			}
		}
		i = next
		if i < len(src) {
			w.emit(none, src[i:i+1])
			i++
		}
	}
}

// adocTitle: od 1 do 6 znaków = i odstęp.
func adocTitle(line []byte) bool {
	n := run(line, 0, 7, func(b byte) bool { return b == '=' })
	return n >= 1 && n <= 6 && n < len(line) && line[n] == ' '
}

// isDelimiter: wiersz złożony z samych znaków c.
func isDelimiter(line []byte, c byte) bool {
	return len(line) > 0 && len(bytes.Trim(line, string(c))) == 0
}

// adocBlock zapisuje blok od wiersza otwierającego src[i:e] do wiersza takiego samego jak otwierający (albo do
// końca danych) i zwraca koniec wiersza zamykającego. Treść: język lang, jeśli jest znany i nie jest AsciiDoc,
// inaczej klasa c. Każdy wiersz jest czytany raz.
func adocBlock(w *writer, src []byte, i, e int, open []byte, c class, lang string) int {
	w.emit(meta, src[i:e])
	if e >= len(src) {
		return e
	}
	body := e + 1
	closeStart, closeEnd := len(src), len(src)
	for p := body; p < len(src); {
		q := lineEnd(src, p)
		if bytes.Equal(bytes.TrimRight(src[p:q], "\r"), open) {
			closeStart, closeEnd = p, q
			break
		}
		if q >= len(src) {
			break
		}
		p = q + 1
	}
	w.emit(none, src[e:e+1])
	content := src[body:closeStart]
	var inner *language
	if lang != "" {
		inner = lookup(lang)
	}
	if inner != nil && inner.names[0] != "asciidoc" && inner.names[0] != "markdown" {
		inner.scan(w, content)
		w.close()
	} else {
		w.emit(c, content)
	}
	w.emit(meta, src[closeStart:closeEnd])
	return closeEnd
}

// adocInline: kod `…` w tym samym wierszu (hl-str) i odwołanie do atrybutu {nazwa} (hl-esc). Po nieudanym
// szukaniu zamykającego ` reszta wiersza nie jest już przeszukiwana, więc koszt jest liniowy.
func adocInline(w *writer, s []byte) {
	start := 0
	noBacktick := false
	for j := 0; j < len(s); j++ {
		n, c := 0, none
		switch {
		case s[j] == '`' && !noBacktick:
			if k := bytes.IndexByte(s[j+1:], '`'); k >= 0 {
				n, c = k+2, str
			} else {
				noBacktick = true
			}
		case s[j] == '{':
			for k := j + 1; k < len(s) && k < j+64; k++ {
				if s[k] == '}' {
					if k > j+1 {
						n, c = k+1-j, escape
					}
					break
				}
				if !isWordByte(s[k]) && s[k] != '-' {
					break
				}
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
