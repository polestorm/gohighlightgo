package highlight

import (
	"bytes"
	"strings"
)

// Haki dla Haskella, Ruby'ego i PHP (definicje w lang/haskell.hl, lang/ruby.hl, lang/php.hl).

// haskellPragma: pragma {-# … #-} (Haskell 2010 Report, 12 „Compiler Pragmas”) dostaje hl-meta; musi stać przed
// komentarzem blokowym {- -}, dlatego jest hakiem.
func haskellPragma(t *table, w *writer, src []byte, i int) int {
	if !hasPrefix(src[i:], "{-#") {
		return i
	}
	end := blockEnd(src, i+3, "", "#-}", false)
	w.emit(meta, src[i:end])
	return end
}

// rubySpecial: komentarz =begin … =end, dokument <<~KONIEC, literały %w[…] i podobne, symbole :nazwa i klucze
// nazwa: (Ruby, „Ruby Syntax”: comments, literals). Dokumenty obsługuje tylko na najwyższym poziomie.
func rubySpecial(t *table, w *writer, src []byte, i int) int {
	if j := rubyInner(t, w, src, i); j > i {
		return j
	}
	if hasPrefix(src[i:], "<<") {
		return scanHeredoc(t, w, src, i, i+2, rubyInner, true)
	}
	return i
}

// rubyInner to rubySpecial bez dokumentów: dla reszty wiersza po znaczniku dokumentu.
func rubyInner(t *table, w *writer, src []byte, i int) int {
	b := src[i]
	switch {
	case b == '=' && hasPrefix(src[i:], "=begin") && (i == 0 || src[i-1] == '\n'):
		end := len(src)
		if n := bytes.Index(src[i:], []byte("\n=end")); n >= 0 {
			end = lineEnd(src, i+n+1)
		}
		w.emit(comment, src[i:end])
		return end
	case b == '%':
		return percentLiteral(w, src, i)
	case b == ':' && i+1 < len(src) && (i == 0 || !isWordByte(src[i-1]) && src[i-1] != ':'):
		if j := t.identEnd(src, i+1); j > i+1 {
			if j < len(src) && (src[j] == '?' || src[j] == '!' || src[j] == '=') && (j+1 >= len(src) || src[j+1] != '=') {
				j++
			}
			w.emit(attribute, src[i:j])
			return j
		}
	case isWordByte(b) && !isDigit(b):
		if j := t.identEnd(src, i); j > i && j+1 < len(src) && src[j] == ':' && src[j+1] != ':' && isSpace(src[j+1]) {
			w.emit(attribute, src[i:j+1]) // klucz w skrócie {nazwa: wartość}
			return j + 1
		}
	}
	return i
}

// percentLiteral: %w[…], %i(…), %q{…}, %Q<…>, %r{…}, %s(…), %x(…) oraz %(…), %[…], %{…} po odstępie albo
// nawiasie (inaczej % to reszta z dzielenia). Nawiasy liczą zagnieżdżenie, \ chroni następny znak; niezamknięty
// literał sięga końca danych. Całość dostaje hl-str.
func percentLiteral(w *writer, src []byte, i int) int {
	k := i + 1
	if k < len(src) && strings.IndexByte("qQwWiIrsx", src[k]) >= 0 {
		k++
	} else if i > 0 && !isSpace(src[i-1]) && strings.IndexByte("(,=[", src[i-1]) < 0 {
		return i
	}
	if k >= len(src) {
		return i
	}
	open := src[k]
	if open >= 0x80 {
		return i // ogranicznik tylko z ASCII: zamknięcie na bajcie znaku wielobajtowego przecięłoby znak
	}
	closeByte, nests := open, false
	switch open {
	case '(':
		closeByte, nests = ')', true
	case '[':
		closeByte, nests = ']', true
	case '{':
		closeByte, nests = '}', true
	case '<':
		closeByte, nests = '>', true
	default:
		if isWordByte(open) || isSpace(open) || k == i+1 {
			return i // %(…) bez litery tylko z nawiasem
		}
	}
	depth := 1
	j := k + 1
	for ; j < len(src) && depth > 0; j++ {
		switch c := src[j]; {
		case c == '\\':
			j++
		case nests && c == open:
			depth++
		case c == closeByte:
			depth--
		}
	}
	j = min(j, len(src))
	w.emit(str, src[i:j])
	return j
}

// phpSpecial: przejścia między HTML a PHP (<?php, <?=, ?>), atrybuty #[…] i dokumenty <<<KONIEC (PHP Manual,
// „Basic syntax”, „Strings”, „Attributes”). HTML poza znacznikami PHP koloruje skaner HTML.
func phpSpecial(t *table, w *writer, src []byte, i int) int {
	if j := phpInner(t, w, src, i); j > i {
		return j
	}
	if hasPrefix(src[i:], "<<<") {
		return scanHeredoc(t, w, src, i, i+3, phpInner, false)
	}
	return i
}

// phpInner to phpSpecial bez dokumentów.
func phpInner(t *table, w *writer, src []byte, i int) int {
	if i == 0 {
		// Dane zaczynają się od HTML: HTML do pierwszego znacznika PHP. Sprawdzane tylko na samym początku,
		// bez cofania się przez odstępy, żeby koszt był stały.
		if j := htmlBeforePHP(src, i); j > i {
			markup{html: true}.scan(w, src[i:j])
			return j
		}
	}
	switch {
	case hasPrefix(src[i:], "<?php"):
		w.emit(meta, src[i:i+5])
		return i + 5
	case hasPrefix(src[i:], "<?="), hasPrefix(src[i:], "<?"):
		n := 2
		if hasPrefix(src[i:], "<?=") {
			n = 3
		}
		w.emit(meta, src[i:i+n])
		return i + n
	case hasPrefix(src[i:], "?>"):
		w.emit(meta, src[i:i+2])
		end := len(src)
		if n := bytes.Index(src[i+2:], []byte("<?")); n >= 0 {
			end = i + 2 + n
		}
		markup{html: true}.scan(w, src[i+2:end])
		return end
	case hasPrefix(src[i:], "#["):
		end := t.identEnd(src, i+2)
		w.emit(meta, src[i:end])
		return end
	}
	return i
}

// htmlBeforePHP: gdy od src[i] stoi HTML (< z literą albo <!), zwraca początek pierwszego znacznika <? (albo
// koniec danych); inaczej i.
func htmlBeforePHP(src []byte, i int) int {
	k := i
	if k+1 >= len(src) || src[k] != '<' || src[k+1] == '?' || !(src[k+1] == '!' || isNameStart(src, k+1)) {
		return i
	}
	if n := bytes.Index(src[k:], []byte("<?")); n >= 0 {
		return k + n
	}
	return len(src)
}

// scanHeredoc obsługuje dokument od znacznika w src[i] (<<~, <<- i << w Rubym, <<< w PHP) z ogranicznikiem od
// src[k]: nazwa albo nazwa w cudzysłowie lub apostrofach. Znacznik i ogranicznik dostają hl-meta, reszta wiersza
// jest kolorowana kopią tabeli z hakiem inner (bez dokumentów, więc bez zagnieżdżania), wiersze treści hl-str,
// a wiersz z ogranicznikiem (po wcięciu, z dowolnym dalszym ciągiem bez znaków nazwy) hl-meta. Bez nazwy po
// znaczniku nic nie robi. Każdy wiersz jest czytany raz, więc koszt jest liniowy.
func scanHeredoc(t *table, w *writer, src []byte, i, k int, inner hookFunc, ruby bool) int {
	if ruby && k < len(src) && (src[k] == '~' || src[k] == '-') {
		k++
	}
	for !ruby && k < len(src) && (src[k] == ' ' || src[k] == '\t') {
		k++
	}
	start := k
	var delim []byte
	if k < len(src) && (src[k] == '\'' || src[k] == '"') {
		q := src[k]
		n := bytes.IndexByte(src[k+1:lineEnd(src, k)], q)
		if n <= 0 {
			return i
		}
		delim, k = src[k+1:k+1+n], k+2+n
	} else {
		e := t.identEnd(src, k)
		if e == k || ruby && !(src[k] >= 'A' && src[k] <= 'Z' || src[k] == '_') {
			return i // << bez nazwy (albo z małą literą w Rubym) to operator
		}
		delim, k = src[k:e], e
	}
	w.emit(operator, src[i:start])
	w.emit(meta, src[start:k])
	e := lineEnd(src, k)
	rest := *t
	rest.specials = []hookFunc{inner}
	rest.scan(w, src[k:e])
	w.close()
	if e >= len(src) {
		return e
	}
	w.emit(none, src[e:e+1])
	for p := e + 1; p < len(src); {
		q := lineEnd(src, p)
		line := bytes.TrimLeft(src[p:q], " \t")
		if bytes.HasPrefix(line, delim) && (len(line) == len(delim) || !isWordByte(line[len(delim)])) {
			w.emit(none, src[p:q-len(line)])
			w.emit(meta, line[:len(delim)])
			rest2 := *t
			rest2.specials = []hookFunc{inner}
			rest2.scan(w, line[len(delim):])
			w.close()
			return q
		}
		w.emit(str, src[p:q])
		if q >= len(src) {
			return q
		}
		w.emit(none, src[q:q+1])
		p = q + 1
	}
	return len(src)
}

func isWordByte(b byte) bool {
	return b == '_' || isDigit(b) || b|0x20 >= 'a' && b|0x20 <= 'z'
}
