package highlight

// Object Pascal (Delphi) według Delphi Language Guide (dokumentacja RAD Studio 12): słowa zastrzeżone,
// dyrektywy, typy i procedury modułu System; wymagania: requirements/languages/pascal.md.
var langPascal = language{
	names: []string{"pascal", "delphi", "pas", "objectpascal"},
	scan: (&table{
		words: foldWords(words(
			wordGroup{keyword, `and array as asm begin case class const constructor destructor dispinterface
				div do downto else end except exports file finalization finally for function goto if
				implementation in inherited initialization inline interface is label library mod nil not
				object of or out packed procedure program property raise record repeat resourcestring set
				shl shr string then threadvar to try type unit until uses var while with xor
				abstract cdecl deprecated dynamic external final forward helper overload override platform
				private protected public published reference reintroduce safecall sealed static stdcall
				strict virtual operator`},
			wordGroup{typeName, `Boolean ByteBool WordBool LongBool Byte ShortInt SmallInt Word Integer Cardinal
				LongInt LongWord Int64 UInt64 NativeInt NativeUInt Char AnsiChar WideChar AnsiString
				WideString UnicodeString ShortString RawByteString Real Single Double Extended Currency Comp
				Pointer PChar PAnsiChar PWideChar Variant OleVariant`},
			wordGroup{builtin, `True False Self Result Abs Assigned Break Chr Continue Copy Dec Delete Dispose
				Exclude Exit FillChar Halt High Include Inc Insert Length Low New Odd Ord Pos Pred ReadLn Round
				SetLength SizeOf Sqr Sqrt Succ Trunc WriteLn`},
		)),
		foldCase:     true,
		lineComments: []string{"//"},
		blockComments: []blockComment{
			{open: "{", close: "}"},
			{open: "(*", close: "*)"},
		},
		quotes:    []quote{{open: `'`, close: `'`, doubled: true}},
		numbers:   numberRule{noLeadingDot: true},
		operators: "+-*/=<>:@^",
		calls:     true,
		special:   pascalSpecial,
	}).scan,
}

// pascalSpecial: dyrektywy kompilatora {$…} i (*$…*) (hl-meta), kody znaków #13 i #$0D (hl-esc) oraz liczby
// szesnastkowe $FF, dwójkowe %1010 i ósemkowe &777 (hl-num).
func pascalSpecial(t *table, w *writer, src []byte, i int) int {
	next := byte(0)
	if i+1 < len(src) {
		next = src[i+1]
	}
	switch src[i] {
	case '{':
		if next == '$' {
			end := blockEnd(src, i+2, "{", "}", false)
			w.emit(meta, src[i:end])
			return end
		}
	case '(':
		if hasPrefix(src[i:], "(*$") {
			end := blockEnd(src, i+3, "(*", "*)", false)
			w.emit(meta, src[i:end])
			return end
		}
	case '#':
		n := run(src, i+1, unbounded, isDigit)
		if next == '$' {
			n = run(src, i+2, unbounded, isHex)
			if n > 0 {
				n++
			}
		}
		if n > 0 {
			w.emit(escape, src[i:i+1+n])
			return i + 1 + n
		}
	case '$':
		return pascalNumber(w, src, i, isHex)
	case '%':
		return pascalNumber(w, src, i, isBinary)
	case '&':
		return pascalNumber(w, src, i, isOctal)
	}
	return i
}

// pascalNumber zapisuje liczbę z przedrostkiem w src[i] i cyframi spełniającymi ok (z _ między nimi).
func pascalNumber(w *writer, src []byte, i int, ok func(byte) bool) int {
	n := run(src, i+1, unbounded, func(b byte) bool { return ok(b) || b == '_' })
	if n == 0 || !ok(src[i+1]) {
		return i
	}
	w.emit(number, src[i:i+1+n])
	return i + 1 + n
}
