package highlight

// Haki dla Makefile, CMake, LaTeX-a, Vim script i HCL (definicje w lang/make.hl, cmake.hl, latex.hl, vim.hl,
// hcl.hl).

// variablesSpecial: rozwinięcia $(…) i ${…} (do nawiasu zamykającego w tym samym wierszu, najwyżej 256 bajtów)
// oraz zmienne automatyczne Make $@ $< $^ $? $* $% $+ $| dostają hl-esc; nazwa przed : na początku wiersza (cel
// reguły Make) dostaje hl-fn (GNU make manual, 6.1 i 10.5.3; CMake, „cmake-language”, Variable References).
func variablesSpecial(t *table, w *writer, src []byte, i int) int {
	b := src[i]
	if b == '$' && i+1 < len(src) {
		switch next := src[i+1]; next {
		case '(':
			if n := bracketLen(src, i, '(', ')', i+1); n > 0 {
				w.emit(escape, src[i:i+n])
				return i + n
			}
		case '{':
			if n := bracketLen(src, i, '{', '}', i+1); n > 0 {
				w.emit(escape, src[i:i+n])
				return i + n
			}
		case '@', '<', '^', '?', '*', '%', '+', '|', '$':
			w.emit(escape, src[i:i+2])
			return i + 2
		}
		return i
	}
	if (i == 0 || src[i-1] == '\n') && isWordByte(b) {
		j := i
		for j < len(src) && (isWordByte(src[j]) || src[j] == '.' || src[j] == '-' || src[j] == '/' || src[j] == '%') {
			j++
		}
		if j < len(src) && src[j] == ':' && (j+1 >= len(src) || src[j+1] != '=') {
			w.emit(call, src[i:j])
			return j
		}
	}
	return i
}

// latexSpecial: \ z jednym znakiem spoza liter (\%, \\, \{, \$) to znak ucieczki (hl-esc), żeby % po \ nie
// zaczynał komentarza; polecenie \nazwa obsługuje sigil.
func latexSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '\\' || i+1 >= len(src) || isWordByte(src[i+1]) && !isDigit(src[i+1]) || src[i+1] >= 0x80 || isSpace(src[i+1]) {
		return i
	}
	w.emit(escape, src[i:i+2])
	return i + 2
}

// vimSpecial: " na początku polecenia (po odstępach od początku wiersza albo po |) to komentarz do końca wiersza
// (Vim, „:help :comment”); w innym miejscu zaczyna napis.
func vimSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '"' {
		return i
	}
	k := i
	for k > 0 && (src[k-1] == ' ' || src[k-1] == '\t') {
		k--
	}
	if k == 0 || src[k-1] == '\n' || src[k-1] == '|' {
		e := lineEnd(src, i)
		w.emit(comment, src[i:e])
		return e
	}
	return i
}

// hclSpecial: dokument <<EOT i <<-EOT (HCL Native Syntax Specification, „Template Expressions”, heredoc).
func hclSpecial(t *table, w *writer, src []byte, i int) int {
	if !hasPrefix(src[i:], "<<") {
		return i
	}
	return scanHeredoc(t, w, src, i, i+2, noHook, true)
}

// noHook to hak, który niczego nie rozpoznaje (reszta wiersza po znaczniku dokumentu bez haków).
func noHook(t *table, w *writer, src []byte, i int) int { return i }
