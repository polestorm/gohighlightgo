package highlight

import (
	"bytes"
	"strings"
)

// Haki dla Perla, Elixira, MATLAB-a i Fortranu (definicje w lang/perl.hl, elixir.hl, matlab.hl, fortran.hl).

// delimited zapisuje literał od src[i] z ogranicznikiem otwierającym w src[k] i zwraca jego koniec: nawiasy ( [ {
// < liczą zagnieżdżenie, inny znak ASCII zamyka się tym samym znakiem, \ chroni następny bajt. parts to liczba
// części (2 dla s/a/b/ i tr/a/b/: druga część zaczyna się od razu albo, przy nawiasach, od nowego nawiasu).
// Niezamknięty literał sięga końca danych. Ogranicznik spoza ASCII, litera, cyfra albo odstęp: zwraca i.
func delimited(w *writer, src []byte, i, k, parts int, c class) int {
	if k >= len(src) {
		return i
	}
	j := k
	for part := 0; part < parts; part++ {
		for j < len(src) && part > 0 && isSpace(src[j]) && strings.IndexByte("([{<", src[k]) >= 0 {
			j++ // s{a} {b}: odstęp między częściami z nawiasami
		}
		if j >= len(src) {
			break
		}
		open := src[j]
		if open >= 0x80 || isWordByte(open) || isSpace(open) {
			if part == 0 {
				return i
			}
			break
		}
		closing, nests := open, false
		switch open {
		case '(':
			closing, nests = ')', true
		case '[':
			closing, nests = ']', true
		case '{':
			closing, nests = '}', true
		case '<':
			closing, nests = '>', true
		}
		depth := 1
		j++
		for ; j < len(src) && depth > 0; j++ {
			switch b := src[j]; {
			case b == '\\':
				j++
			case nests && b == open:
				depth++
			case b == closing:
				depth--
			}
		}
		j = min(j, len(src))
		if depth > 0 {
			break // niezamknięty: do końca danych
		}
		if !nests && part+1 < parts {
			j-- // s/a/b/: środkowy / zamyka pierwszą część i otwiera drugą
		}
	}
	// modyfikatory po literale (s/a/b/gi, qr/x/i)
	for j < len(src) && src[j] >= 'a' && src[j] <= 'z' {
		j++
	}
	w.emit(c, src[i:j])
	return j
}

// perlSpecial: dokumentacja POD (=nazwa na początku wiersza do =cut), operatory cytujące q qq qw qr m s tr y
// z dowolnym ogranicznikiem oraz dokumenty <<"KONIEC", <<'KONIEC', <<~KONIEC (perlop, perlpod).
func perlSpecial(t *table, w *writer, src []byte, i int) int {
	if j := perlInner(t, w, src, i); j > i {
		return j
	}
	if hasPrefix(src[i:], "<<") && i+2 < len(src) && (src[i+2] == '"' || src[i+2] == '\'' || src[i+2] == '~' || src[i+2] >= 'A' && src[i+2] <= 'Z') {
		return scanHeredoc(t, w, src, i, i+2, perlInner, true)
	}
	return i
}

func perlInner(t *table, w *writer, src []byte, i int) int {
	b := src[i]
	if b == '=' && (i == 0 || src[i-1] == '\n') && i+1 < len(src) && isWordByte(src[i+1]) {
		end := len(src)
		if n := bytes.Index(src[i:], []byte("\n=cut")); n >= 0 {
			end = lineEnd(src, i+n+1)
		}
		w.emit(comment, src[i:end])
		return end
	}
	if !isWordByte(b) || i > 0 && (isWordByte(src[i-1]) || src[i-1] == '$' || src[i-1] == '@' || src[i-1] == '%' || src[i-1] == '>' && i > 1 && src[i-2] == '-') {
		return i
	}
	j := i
	for j < len(src) && j < i+2 && src[j] >= 'a' && src[j] <= 'z' {
		j++
	}
	switch string(src[i:j]) {
	case "q", "qq", "qw", "qr", "m":
		if j < len(src) && (src[j] == '=' && j+1 < len(src) && src[j+1] == '>' || src[j] == ',') {
			return i // klucz hasza q => 1
		}
		return delimited(w, src, i, j, 1, str)
	case "s", "tr", "y":
		if j < len(src) && (src[j] == '=' && j+1 < len(src) && src[j+1] == '>' || src[j] == ',') {
			return i
		}
		return delimited(w, src, i, j, 2, str)
	}
	return i
}

// elixirSpecial: sigile ~r/…/, ~w(…), ~s"…" i inne (~ z literą i ogranicznikiem), atomy :nazwa i klucze nazwa:
// (Elixir, „Syntax reference”).
func elixirSpecial(t *table, w *writer, src []byte, i int) int {
	b := src[i]
	switch {
	case b == '~' && i+2 < len(src) && (src[i+1] >= 'a' && src[i+1] <= 'z' || src[i+1] >= 'A' && src[i+1] <= 'Z'):
		k := i + 2
		for k < len(src) && src[k] >= 'A' && src[k] <= 'Z' {
			k++ // ~HTML"""…"""
		}
		return delimited(w, src, i, k, 1, str)
	case b == ':' || isWordByte(b) && !isDigit(b):
		return atomOrKey(t, w, src, i)
	}
	return i
}

// atomOrKey: symbol :nazwa (nie po znaku nazwy ani po :, z ? ! = na końcu) i klucz nazwa: przed odstępem
// dostają hl-attr; wspólne dla Ruby'ego i Elixira.
func atomOrKey(t *table, w *writer, src []byte, i int) int {
	b := src[i]
	if b == ':' && i+1 < len(src) && (i == 0 || !isWordByte(src[i-1]) && src[i-1] != ':') {
		if j := t.identEnd(src, i+1); j > i+1 {
			if j < len(src) && (src[j] == '?' || src[j] == '!' || src[j] == '=') && (j+1 >= len(src) || src[j+1] != '=') {
				j++
			}
			w.emit(attribute, src[i:j])
			return j
		}
		return i
	}
	if isWordByte(b) && !isDigit(b) {
		if j := t.identEnd(src, i); j > i && j+1 < len(src) && src[j] == ':' && src[j+1] != ':' && isSpace(src[j+1]) {
			w.emit(attribute, src[i:j+1])
			return j + 1
		}
	}
	return i
}

// matlabSpecial: apostrof po nazwie, liczbie, ) ] } . albo innym apostrofie to transpozycja (hl-op); inaczej
// zaczyna napis z ” jako znakiem ucieczki, do końca wiersza (MATLAB, „Characters and Strings”). Komentarz blokowy
// %{ … %} w osobnych wierszach.
func matlabSpecial(t *table, w *writer, src []byte, i int) int {
	switch src[i] {
	case '\'':
		if i > 0 && (isWordByte(src[i-1]) || strings.IndexByte(")]}.'", src[i-1]) >= 0) {
			w.emit(operator, src[i:i+1])
			return i + 1
		}
		j := i + 1
		for j < len(src) && src[j] != '\n' {
			if src[j] == '\'' {
				if j+1 < len(src) && src[j+1] == '\'' {
					j += 2
					continue
				}
				j++
				break
			}
			j++
		}
		w.emit(str, src[i:j])
		return j
	case '%':
		if hasPrefix(src[i:], "%{") && atLineStart(src, i) && len(bytes.TrimSpace(src[i+2:lineEnd(src, i)])) == 0 {
			end := len(src)
			if n := bytes.Index(src[i:], []byte("%}")); n >= 0 {
				end = i + n + 2
			}
			w.emit(comment, src[i:end])
			return end
		}
	}
	return i
}

// dottedSpecial: operatory i stałe Fortranu w kropkach (.and., .eq., .not.; .true. i .false. dostają hl-builtin),
// bez względu na wielkość liter (Fortran 2018, 6.2.4 i 7.4.5).
func dottedSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '.' {
		return i
	}
	j := i + 1
	for j < len(src) && j < i+8 && (src[j]|0x20 >= 'a' && src[j]|0x20 <= 'z') {
		j++
	}
	if j == i+1 || j >= len(src) || src[j] != '.' {
		return i
	}
	c := operator
	switch strings.ToLower(string(src[i+1 : j])) {
	case "true", "false":
		c = builtin
	}
	w.emit(c, src[i:j+1])
	return j + 1
}
