package highlight

import "bytes"

// Haki dla Lua i Swifta (definicje w lang/lua.hl i lang/swift.hl).

// luaSpecial: długie nawiasy Lua (Lua 5.4 Reference Manual, 3.1): napis [[…]] i [==[…]==] (hl-str) oraz
// komentarz --[[…]] i --[==[…]==] (hl-com); koniec to ] z tą samą liczbą = i ]. Bez zamknięcia sięga końca
// danych. Komentarz -- bez długiego nawiasu obsługuje tabela.
func luaSpecial(t *table, w *writer, src []byte, i int) int {
	c := str
	k := i
	if hasPrefix(src[i:], "--[") {
		c, k = comment, i+2
	}
	if k >= len(src) || src[k] != '[' {
		return i
	}
	level := run(src, k+1, len(src), func(b byte) bool { return b == '=' })
	if k+1+level >= len(src) || src[k+1+level] != '[' {
		return i
	}
	end := closeLong(src, k+2+level, level)
	w.emit(c, src[i:end])
	return end
}

// closeLong zwraca koniec długiego nawiasu poziomu level od src[p]: za ], level znaków = i ]. Każdy ] jest
// sprawdzany raz, kosztem najwyżej level+1 bajtów, a bajty = nie są potem czytane drugi raz jako ], więc
// koszt jest liniowy.
func closeLong(src []byte, p, level int) int {
	for p < len(src) {
		n := bytes.IndexByte(src[p:], ']')
		if n < 0 {
			break
		}
		q := p + n + 1
		eq := run(src, q, level, func(b byte) bool { return b == '=' })
		if eq == level && q+level < len(src) && src[q+level] == ']' {
			return q + level + 1
		}
		p = q + eq
	}
	return len(src)
}

// swiftSpecial: napis surowy #"…"#, ##"…"## (The Swift Programming Language, „String Literals”, „Extended
// String Delimiters”), także wielowierszowy #"""…"""#, w całości hl-str; # z nazwą (#if, #available, #selector)
// dostaje hl-meta.
func swiftSpecial(t *table, w *writer, src []byte, i int) int {
	if src[i] != '#' {
		return i
	}
	hashes := run(src, i, len(src), func(b byte) bool { return b == '#' })
	k := i + hashes
	if k < len(src) && src[k] == '"' {
		quotes := 1
		if hasPrefix(src[k:], `"""`) {
			quotes = 3
		}
		closing := bytes.Repeat([]byte{'"'}, quotes)
		closing = append(closing, bytes.Repeat([]byte{'#'}, hashes)...)
		end := len(src)
		if n := bytes.Index(src[k+quotes:], closing); n >= 0 {
			end = k + quotes + n + len(closing)
		}
		w.emit(str, src[i:end])
		return end
	}
	if hashes == 1 {
		if j := t.identEnd(src, i+1); j > i+1 {
			w.emit(meta, src[i:j])
			return j
		}
		return i
	}
	// Seria # bez napisu po niej: zużyta w całości, żeby następna pozycja nie liczyła jej od nowa (czas liniowy).
	w.emit(none, src[i:k])
	return k
}
