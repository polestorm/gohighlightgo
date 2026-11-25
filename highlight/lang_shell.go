package highlight

import (
	"bytes"
	"strings"
)

// Powłoka według POSIX.1-2024, „Shell Command Language” (rozdział 2 tomu Shell and Utilities), z rozszerzeniami
// opisanymi w Bash Reference Manual (wersja 5.2): [[ ]], function, local, declare, $'…'.
// Wymagania: requirements/languages/sh.md.
var langShell = language{
	names: []string{"sh", "bash", "shell", "zsh", "ksh"},
	scan:  scanShell,
}

var shellWords = words(
	wordGroup{keyword, `if then else elif fi case esac for select while until do done in function time
		! { } [[ ]]`},
	wordGroup{builtin, `. : [ alias bg break cd command continue declare echo eval exec exit export false fg
		getopts hash jobs kill let local printf pwd read readonly return set shift source test times trap
		true type typeset ulimit umask unalias unset wait`},
)

// shellKeepsCommand: po tych słowach kluczowych dalej stoi polecenie.
var shellKeepsCommand = map[string]bool{"if": true, "then": true, "else": true, "elif": true, "do": true,
	"while": true, "until": true, "!": true, "{": true, "time": true}

// shellFrame to otwarte $( … ): w kodzie albo w napisie "…", do którego trzeba wrócić po ).
type shellFrame struct {
	inString bool
	arith    bool // $(( … )): wyrażenie arytmetyczne, bez poleceń
	depth    int  // otwarte nawiasy ( wewnątrz
}

type heredoc struct {
	delim     []byte
	stripTabs bool // <<- pomija tabulatory na początku wierszy
}

// scanShell: pierwsze słowo polecenia dostaje hl-fn (słowo kluczowe hl-kw, polecenie wbudowane hl-builtin),
// rozwinięcia $x, ${x}, $( … ) hl-esc, napisy hl-str, dokumenty <<KONIEC hl-str z wierszem końca hl-meta.
func scanShell(w *writer, src []byte) {
	var stack []shellFrame
	var pending []heredoc
	command := true      // następne słowo stoi w miejscu nazwy polecenia
	keepCommand := false // trwa wartość przypisania NAZWA=…; po odstępie wraca miejsce polecenia
	loopWords := 0       // po for, select i case: 3, potem odlicza słowa; przy 1 słowo in to słowo kluczowe
	for i := 0; i < len(src); {
		if len(stack) > 0 && stack[len(stack)-1].arith {
			command = false
		}
		b := src[i]
		next := byte(0)
		if i+1 < len(src) {
			next = src[i+1]
		}
		switch {
		case b == '\n':
			w.emit(none, src[i:i+1])
			i++
			command = true
			if len(pending) > 0 {
				i = scanHeredocs(w, src, i, pending)
				pending = pending[:0]
			}
		case b == ' ' || b == '\t' || b == '\r':
			w.emit(none, src[i:i+1])
			i++
			if keepCommand {
				command, keepCommand = true, false
			}
		case b == '#' && (i == 0 || isShellBreak(src[i-1])):
			e := lineEnd(src, i)
			w.emit(comment, src[i:e])
			i = e
		case b == '\\':
			n := 2
			if next == 0 {
				n = 1
			} else if next >= 0x80 {
				n = 1 + utf8Len(src, i+1)
			}
			w.emit(escape, src[i:i+n])
			i += n
		case b == '\'':
			end := len(src)
			if n := bytes.IndexByte(src[i+1:], '\''); n >= 0 {
				end = i + 2 + n
			}
			w.emit(str, src[i:end])
			i = end
			command = false
		case b == '"':
			var open bool
			i, open = shellDouble(w, src, i, i+1)
			if open {
				stack = append(stack, shellFrame{inString: true})
			}
			command = open
		case b == '`':
			j := i + 1
			for j < len(src) && src[j] != '`' {
				if src[j] == '\\' {
					j++
				}
				j++
			}
			j = min(j+1, len(src))
			w.emit(str, src[i:j])
			i = j
			command = false
		case b == '$' && next == '(' && i+2 < len(src) && src[i+2] == '(':
			w.emit(escape, src[i:i+3])
			i += 3
			stack = append(stack, shellFrame{arith: true})
		case b == '$' && next == '(':
			w.emit(escape, src[i:i+2])
			i += 2
			stack = append(stack, shellFrame{})
			command = true
		case b == '$':
			n := expansionLen(src, i)
			if n == 0 {
				n = 1
				w.emit(none, src[i:i+1])
			} else {
				w.emit(escape, src[i:i+n])
			}
			i += n
			command = false
		case b == '(':
			if len(stack) > 0 {
				stack[len(stack)-1].depth++
			}
			w.emit(none, src[i:i+1])
			i++
			command = true
		case b == ')':
			if top := len(stack) - 1; top >= 0 {
				if stack[top].arith && stack[top].depth == 0 && next == ')' {
					stack = stack[:top]
					w.emit(escape, src[i:i+2])
					i += 2
					command = false
					continue
				}
				if stack[top].depth > 0 {
					stack[top].depth--
				} else {
					frame := stack[top]
					stack = stack[:top]
					w.emit(escape, src[i:i+1])
					i++
					command = false
					if frame.inString {
						var open bool
						i, open = shellDouble(w, src, i, i)
						if open {
							stack = append(stack, shellFrame{inString: true})
							command = true
						}
					}
					continue
				}
			}
			w.emit(none, src[i:i+1])
			i++
			command = false
		case b == '<' && next == '<':
			var h heredoc
			var ok bool
			i, h, ok = shellHeredoc(w, src, i)
			if ok {
				pending = append(pending, h)
			}
		case strings.IndexByte("|&;<>", b) >= 0:
			w.emit(operator, src[i:i+1])
			// po > i < stoi plik, a & po nich (2>&1) to część przekierowania
			if b != '<' && b != '>' && !(b == '&' && i > 0 && (src[i-1] == '>' || src[i-1] == '<')) {
				command = true
			}
			i++
		default:
			j := i + 1
			for j < len(src) && !isShellBreak(src[j]) && strings.IndexByte("$\\\"'`", src[j]) < 0 {
				j++
			}
			word := src[i:j]
			c := none
			if loopWords > 0 {
				loopWords--
			}
			switch {
			case command && shellAssignment(word) > 0:
				n := shellAssignment(word)
				w.emit(attribute, word[:n])
				word = word[n:]
				command, keepCommand = false, true
			case string(word) == "]]" || loopWords == 1 && string(word) == "in":
				c = keyword
			case command:
				c = shellWords[string(word)]
				switch {
				case c == keyword:
					command = shellKeepsCommand[string(word)]
					if string(word) == "for" || string(word) == "select" || string(word) == "case" {
						loopWords = 3
					}
				case c == builtin:
					command = false
				default:
					c, command = call, false
				}
			case strings.Trim(string(word), "0123456789") == "":
				c = number
			}
			w.emit(c, word)
			i = j
		}
	}
}

// shellDouble zapisuje napis "…" od src[start] z treścią od src[j]. W napisie \ przed $ ` " \ i końcem wiersza
// to sekwencja ucieczki, $x i ${x} to rozwinięcia. Zwraca open == true, gdy zatrzymał się na $(.
func shellDouble(w *writer, src []byte, start, j int) (end int, open bool) {
	for j < len(src) {
		switch src[j] {
		case '"':
			w.emit(str, src[start:j+1])
			return j + 1, false
		case '\\':
			if j+1 < len(src) && strings.IndexByte("$`\"\\\n", src[j+1]) >= 0 {
				w.emit(str, src[start:j])
				w.emit(escape, src[j:j+2])
				j += 2
				start = j
				continue
			}
		case '$':
			if j+1 < len(src) && src[j+1] == '(' {
				w.emit(str, src[start:j])
				w.emit(escape, src[j:j+2])
				return j + 2, true
			}
			if n := expansionLen(src, j); n > 0 {
				w.emit(str, src[start:j])
				w.emit(escape, src[j:j+n])
				j += n
				start = j
				continue
			}
		}
		j++
	}
	w.emit(str, src[start:j])
	return j, false
}

// expansionLen zwraca długość rozwinięcia od $ w src[i] albo 0: ${…} (do } w tym samym wierszu, najwyżej
// 256 bajtów; bez } samo ${), $nazwa, $0–$9 i parametry specjalne $? $$ $! $# $@ $* $-.
func expansionLen(src []byte, i int) int {
	if i+1 >= len(src) {
		return 0
	}
	c := src[i+1]
	switch {
	case c == '{':
		for j := i + 2; j < len(src) && j < i+258 && src[j] != '\n'; j++ {
			if src[j] == '}' {
				return j + 1 - i
			}
		}
		return 2
	case c == '_' || c|0x20 >= 'a' && c|0x20 <= 'z':
		j := i + 2
		for j < len(src) && (src[j] == '_' || isDigit(src[j]) || src[j]|0x20 >= 'a' && src[j]|0x20 <= 'z') {
			j++
		}
		return j - i
	case isDigit(c) || strings.IndexByte("?$!#@*-", c) >= 0:
		return 2
	}
	return 0
}

// shellHeredoc rozpoznaje <<, <<- i <<< od src[i]. Przy << i <<- zapisuje operator i ogranicznik (hl-meta)
// i zwraca dokument do odczytania od następnego wiersza.
func shellHeredoc(w *writer, src []byte, i int) (int, heredoc, bool) {
	if hasPrefix(src[i:], "<<<") {
		w.emit(operator, src[i:i+3])
		return i + 3, heredoc{}, false
	}
	j := i + 2
	h := heredoc{}
	if j < len(src) && src[j] == '-' {
		h.stripTabs = true
		j++
	}
	w.emit(operator, src[i:j])
	k := j
	for k < len(src) && (src[k] == ' ' || src[k] == '\t') {
		k++
	}
	w.emit(none, src[j:k])
	start := k
	if k < len(src) && (src[k] == '\'' || src[k] == '"') {
		q := src[k]
		e := quoteBeforeEOL(src, k+1, q)
		if e < 0 {
			return k, heredoc{}, false
		}
		h.delim = src[k+1 : k+1+e]
		k += e + 2
	} else {
		for k < len(src) && !isShellBreak(src[k]) && strings.IndexByte("\"'`", src[k]) < 0 {
			k++
		}
		h.delim = src[start:k]
	}
	if len(h.delim) == 0 {
		return start, heredoc{}, false // bez ogranicznika: reszta to zwykłe tokeny
	}
	w.emit(meta, src[start:k])
	return k, h, true
}

// scanHeredocs zapisuje treść dokumentów od src[i] (początek wiersza), po kolei: wiersze do wiersza równego
// ogranicznikowi dostają hl-str, a wiersz ogranicznika hl-meta. Każdy wiersz jest porównywany raz.
func scanHeredocs(w *writer, src []byte, i int, docs []heredoc) int {
	for _, h := range docs {
		for i < len(src) {
			e := lineEnd(src, i)
			line := src[i:e]
			if h.stripTabs {
				line = bytes.TrimLeft(line, "\t")
			}
			line = bytes.TrimSuffix(line, []byte("\r"))
			isEnd := bytes.Equal(line, h.delim)
			if isEnd {
				w.emit(meta, src[i:e])
			} else {
				w.emit(str, src[i:e])
			}
			i = e
			if i < len(src) {
				w.emit(none, src[i:i+1])
				i++
			}
			if isEnd {
				break
			}
		}
	}
	return i
}

// shellAssignment zwraca długość nazwy z = w przypisaniu NAZWA=wartość albo 0.
func shellAssignment(word []byte) int {
	eq := bytes.IndexByte(word, '=')
	if eq <= 0 {
		return 0
	}
	for k, c := range word[:eq] {
		if !(c == '_' || c|0x20 >= 'a' && c|0x20 <= 'z' || k > 0 && isDigit(c)) {
			return 0
		}
	}
	return eq + 1
}

// isShellBreak: bajty, które kończą słowo powłoki.
func isShellBreak(b byte) bool {
	return isSpace(b) || strings.IndexByte("|&;<>()", b) >= 0
}

// utf8Len zwraca długość znaku UTF-8 od src[i] (co najmniej 1).
func utf8Len(src []byte, i int) int {
	n := 1
	for i+n < len(src) && n < 4 && src[i+n]&0xC0 == 0x80 {
		n++
	}
	return n
}
