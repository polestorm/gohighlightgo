package highlight

// Haki dla Lispów, Veriloga, Tcl i asemblerów (definicje w lang/scheme.hl, racket.hl, commonlisp.hl,
// verilog.hl, tcl.hl, asm.hl, armasm.hl).

// lispSpecial: znak #\a i #\nazwa (hl-str), #t #f #true #false (hl-builtin) oraz wiersz #lang na początku
// (hl-meta) (R7RS, 7.1.1; The Racket Reference, „The Reader”; Common Lisp HyperSpec, 2.4.8.1).
func lispSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '#' || i+1 >= len(src) {
		return i
	}
	switch next := src[i+1]; {
	case next == '\\' && i+2 < len(src):
		j := i + 2 + utf8Len(src, i+2)
		for j < len(src) && isWordByte(src[j]) && isWordByte(src[i+2]) {
			j++ // #\space, #\newline
		}
		w.emit(str, src[i:j])
		return j
	case next == 't' || next == 'f':
		j := i + 2
		for j < len(src) && src[j] >= 'a' && src[j] <= 'z' {
			j++
		}
		switch string(src[i+1 : j]) {
		case "t", "f", "true", "false":
			w.emit(builtin, src[i:j])
			return j
		}
	case hasPrefix(src[i:], "#lang") && atLineStart(src, i):
		e := lineEnd(src, i)
		w.emit(meta, src[i:e])
		return e
	}
	return i
}

// verilogNumber: liczba z rozmiarem i podstawą (8'hFF, 'b1010, 4'sd5, 1'bx) dostaje hl-num (IEEE 1800-2017,
// 5.7.1); cyfry rozmiaru przed apostrofem należą do liczby.
func verilogNumber(t *table, w *writer, src []byte, i int) int {
	j := i
	for j < len(src) && (isDigit(src[j]) || src[j] == '_') {
		j++
	}
	if j >= len(src) || src[j] != '\'' {
		return i
	}
	k := j + 1
	if k < len(src) && (src[k] == 's' || src[k] == 'S') {
		k++
	}
	if k >= len(src) || !(src[k]|0x20 == 'b' || src[k]|0x20 == 'o' || src[k]|0x20 == 'd' || src[k]|0x20 == 'h') {
		return i
	}
	k++
	e := k
	for e < len(src) && (isHex(src[e]) || src[e] == '_' || src[e]|0x20 == 'x' || src[e]|0x20 == 'z' || src[e] == '?') {
		e++
	}
	if e == k {
		return i
	}
	w.emit(number, src[i:e])
	return e
}

// tclSpecial: # na początku polecenia (po odstępach od początku wiersza albo po ;) to komentarz do końca wiersza
// (Tcl, „Tcl(n)”, reguła [10]); w innym miejscu to zwykły znak.
func tclSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '#' {
		return i
	}
	k := i
	for k > 0 && (src[k-1] == ' ' || src[k-1] == '\t') {
		k--
	}
	if k == 0 || src[k-1] == '\n' || src[k-1] == ';' {
		e := lineEnd(src, i)
		w.emit(comment, src[i:e])
		return e
	}
	return i
}

// asmLabel: nazwa z dwukropkiem na początku wiersza (po odstępach) to etykieta (hl-fn), także lokalna .L1:.
func asmLabel(t *table, w *writer, src []byte, i int) int {
	if !atLineStart(src, i) || !(isWordByte(src[i]) || src[i] == '.') || isDigit(src[i]) {
		return i
	}
	j := i + 1
	for j < len(src) && (isWordByte(src[j]) || src[j] == '.' || src[j] == '$') {
		j++
	}
	if j < len(src) && src[j] == ':' {
		w.emit(call, src[i:j])
		return j
	}
	return i
}
