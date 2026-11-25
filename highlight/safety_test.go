package highlight

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
	"unicode/utf8"
)

// allowedClasses to lista z OUT-002, przepisana tu niezależnie od biblioteki: sprawdzian nie może ufać
// tablicy, którą sprawdza.
var allowedClasses = map[string]bool{
	"hl-kw": true, "hl-type": true, "hl-builtin": true, "hl-str": true, "hl-esc": true, "hl-num": true,
	"hl-com": true, "hl-op": true, "hl-fn": true, "hl-tag": true, "hl-attr": true, "hl-ins": true,
	"hl-del": true, "hl-meta": true,
}

// checkSafe sprawdza postać wyniku niezależnie od kodu biblioteki (SAFE-002, OUT-002–OUT-005): tylko tekst bez
// < i >, trzy odwołania znakowe, pary <span class="K">…</span> bez zagnieżdżeń, bez pustych elementów
// i bez \n wewnątrz elementu; całość to poprawny UTF-8.
func checkSafe(out []byte) error {
	if !utf8.Valid(out) {
		return errors.New("output is not valid UTF-8")
	}
	const openTag, closeTag = `<span class="`, `</span>`
	open := false
	empty := false
	for i := 0; i < len(out); {
		rest := out[i:]
		switch out[i] {
		case '<':
			switch {
			case bytes.HasPrefix(rest, []byte(closeTag)):
				if !open {
					return fmt.Errorf("offset %d: </span> without an open element", i)
				}
				if empty {
					return fmt.Errorf("offset %d: empty element", i)
				}
				open = false
				i += len(closeTag)
			case bytes.HasPrefix(rest, []byte(openTag)):
				if open {
					return fmt.Errorf("offset %d: nested element", i)
				}
				end := bytes.Index(rest, []byte(`">`))
				if end < 0 {
					return fmt.Errorf("offset %d: unterminated tag", i)
				}
				if class := string(rest[len(openTag):end]); !allowedClasses[class] {
					return fmt.Errorf("offset %d: class %q is not allowed", i, class)
				}
				open, empty = true, true
				i += end + 2
			default:
				return fmt.Errorf("offset %d: raw < or a tag other than span", i)
			}
		case '>':
			return fmt.Errorf("offset %d: raw >", i)
		case '&':
			n := 0
			for _, ref := range []string{"&amp;", "&lt;", "&gt;"} {
				if bytes.HasPrefix(rest, []byte(ref)) {
					n = len(ref)
				}
			}
			if n == 0 {
				return fmt.Errorf("offset %d: & outside &amp; &lt; &gt;", i)
			}
			empty = false
			i += n
		case '\n':
			if open {
				return fmt.Errorf("offset %d: line break inside an element", i)
			}
			i++
		default:
			empty = false
			i++
		}
	}
	if open {
		return errors.New("element left open at the end")
	}
	return nil
}

// plainText odwraca wynik: zdejmuje znaczniki i odwołania znakowe (SAFE-001). Zakłada wynik, który przeszedł
// checkSafe.
func plainText(out []byte) []byte {
	var b bytes.Buffer
	for i := 0; i < len(out); {
		switch {
		case out[i] == '<':
			i += bytes.IndexByte(out[i:], '>') + 1
		case bytes.HasPrefix(out[i:], []byte("&amp;")):
			b.WriteByte('&')
			i += 5
		case bytes.HasPrefix(out[i:], []byte("&lt;")):
			b.WriteByte('<')
			i += 4
		case bytes.HasPrefix(out[i:], []byte("&gt;")):
			b.WriteByte('>')
			i += 4
		default:
			b.WriteByte(out[i])
			i++
		}
	}
	return b.Bytes()
}

// expectedText to wejście po zamianie z OUT-006, zapisanej tu wprost według wymagania.
func expectedText(src []byte) []byte {
	return bytes.ReplaceAll(bytes.ToValidUTF8(src, []byte("�")), []byte{0}, []byte("�"))
}

// trickyInputs: wejścia na granicach skanera, wspólne dla TestRoundTrip, TestSafeOutput i ziaren fuzzingu.
var trickyInputs = []string{
	"",
	"\n",
	"<>&\"'",
	"&amp; &lt; <span class=\"hl-kw\"></span>",
	"\x00\xff\xfe a \xc3",
	"\"\\",
	"'\\",
	"\"a\\\nb\"",
	"`\n\n`",
	"/*",
	"/*/",
	"//",
	"0x",
	"0x.p",
	"1e+",
	".",
	"..5",
	"a (",
	"a\t\t(",
	"a\n(",
	"\"\\U0010",
	"'\\x",
	"\r\n\"a\r\nb",
	"func main()",
	"ż(ź)",
	"R\"(", "u8R\"x(\n)x\"", "r#\"", "br##\"\"#", "'a", "b'", "'\\'", "`${", "`${`${}`}`", "}",
	"{$", "(*$", "#$", "$", "''", "'it''s'", "\"a\":", "-", "-1", "#include <", "  # define x", "@", "@a.b(",
	"r\"\\\"", "\"\\N{", "\"\\u{", "'''", "\"\"\"", "x!=y", "a!", "#[", "#![",
	"<a", "<a b=", "<!--", "<![CDATA[", "<script>", "<style>a{", "&#", "&#x;", "url(", "a:", "{a:", "@", "--x",
	"a: |\n b", "- ", "? ", "&a", "*", "!!", "%", "---", "...", "[a", "[[a]]", "a.b = 1", "$((", "$(", "${",
	"<<a\nx\na", "<<-'a'", "<<<", "for x in", "@@ -1 +1 @@\n+a", "```go\nfunc", "~~~\n", "`` `", "](", "<a:b>",
	"> - [x]", "[a]: b",
}

// REQ: SAFE-001 OUT-006
func TestRoundTrip(t *testing.T) {
	inputs := append([]string{}, trickyInputs...)
	for _, path := range corpusFiles(t) {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		inputs = append(inputs, string(src))
	}
	for _, name := range primaryNames() {
		for _, in := range inputs {
			out, ok := HTML(name, []byte(in))
			if !ok {
				t.Fatalf("%s: language not found", name)
			}
			if got, want := plainText(out), expectedText([]byte(in)); !bytes.Equal(got, want) {
				t.Errorf("%s: round trip differs\ninput: %q\ngot:   %q", name, in, got)
			}
		}
	}
}

// REQ: SAFE-002 OUT-002 OUT-003 OUT-004 OUT-005
func TestSafeOutput(t *testing.T) {
	inputs := append([]string{}, trickyInputs...)
	for _, path := range corpusFiles(t) {
		src, _ := os.ReadFile(path)
		inputs = append(inputs, string(src))
	}
	for _, name := range primaryNames() {
		for _, in := range inputs {
			out, _ := HTML(name, []byte(in))
			if err := checkSafe(out); err != nil {
				t.Errorf("%s: %v\ninput:  %q\noutput: %q", name, err, in, out)
			}
		}
	}
}

// TestCheckSafeRejects pilnuje samego sprawdzianu: każdy z tych wyników musi zostać odrzucony.
func TestCheckSafeRejects(t *testing.T) {
	for _, bad := range []string{
		"<b>x</b>",
		"a < b",
		"a > b",
		"&quot;",
		"&",
		`<span class="hl-x">a</span>`,
		`<span class="hl-kw">a`,
		`a</span>`,
		`<span class="hl-kw"></span>`,
		`<span class="hl-kw"><span class="hl-str">a</span></span>`,
		"<span class=\"hl-com\">a\nb</span>",
		`<span class="hl-kw" onclick="x">a</span>`,
		"\xff",
	} {
		if checkSafe([]byte(bad)) == nil {
			t.Errorf("checkSafe accepted %q", bad)
		}
	}
	if err := checkSafe([]byte("a &lt; <span class=\"hl-kw\">b</span>\n\"'")); err != nil {
		t.Errorf("checkSafe rejected valid output: %v", err)
	}
}

// primaryNames zwraca nazwę główną każdego języka (bez aliasów).
func primaryNames() []string {
	var names []string
	for _, l := range languages {
		names = append(names, l.names[0])
	}
	return names
}
