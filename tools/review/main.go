// Polecenie review składa jedną stronę HTML do przeglądu korpusu przez człowieka: dla każdego pliku z
// testdata/corpus język, wymagania, wynik gohighlightgo w motywie jasnym i ciemnym oraz (opcjonalnie) wynik
// programu odniesienia, z liczbami: rozmiar, liczba elementów, czas. Strona pozwala zaznaczyć ocenę i skopiować
// uwagi.
//
//	go run ./tools/review -o out/review.html
//	go run ./tools/review -ref '' -o out/review.html   # bez odniesienia
//
// Odniesieniem jest domyślnie qlmarkdown_cli z QLMarkdown (w środku program highlight André Simona),
// uruchamiany jako czarna skrzynka na dokumencie Markdown z jednym blokiem kodu; jego cała strona trafia do
// ramki <iframe>, żeby jego arkusz nie mieszał się z naszym (docs/provenance.md: wolno uruchamiać, nie czytamy
// kodu ani definicji). Uruchamiać z katalogu głównego repozytorium.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/polestorm/gohighlightgo/highlight"
)

const defaultRef = "/Applications/QLMarkdown.app/Contents/Resources/qlmarkdown_cli"

// refNames: nazwa języka dla odniesienia, gdy różni się od naszej nazwy głównej.
var refNames = map[string]string{"pascal": "delphi", "objectivec": "objc", "fortran": "f90", "make": "makefile",
	"latex": "tex", "protobuf": "pb", "commonlisp": "lisp", "armasm": "arm", "prolog": "pro", "scheme": "scm",
	"vhdl": "vhd", "moonscript": "moon", "batch": "bat", "puppet": "pp"}

type side struct {
	Light, Dark string // fragment HTML albo cała strona odniesienia
	Bytes       int    // rozmiar fragmentu z kodem
	Spans       int
	Classes     int
	Time        time.Duration
	Err         string
}

type entry struct {
	Lang, Path string
	Reqs       []string
	Source     string
	Ours, Ref  side
}

var (
	reqRef   = regexp.MustCompile(`\b[A-Z]{3,5}-\d{3,4}\b`)
	spanOpen = regexp.MustCompile(`<span class=["']([^"']+)["']`)
	refPre   = regexp.MustCompile(`(?s)<pre class='hl'>.*?</pre>`)
)

func main() {
	out := flag.String("o", "out/review.html", "output `file`")
	ref := flag.String("ref", defaultRef, "reference `program` (qlmarkdown_cli); empty: no reference")
	runs := flag.Int("runs", 20, "timing `runs` for gohighlightgo (median)")
	flag.Parse()

	reqs := loadRequirements()
	files, _ := filepath.Glob("testdata/corpus/*/*.txt")
	if len(files) == 0 {
		fail("no corpus files; run from the repository root")
	}
	sort.Strings(files)
	if *ref != "" {
		if _, err := os.Stat(*ref); err != nil {
			fmt.Fprintf(os.Stderr, "review: reference %s not found, continuing without it\n", *ref)
			*ref = ""
		}
	}
	tmp, err := os.MkdirTemp("", "gohighlightgo-review")
	if err != nil {
		fail(err.Error())
	}
	defer os.RemoveAll(tmp)

	var entries []entry
	var oursTotal, refTotal time.Duration
	for _, path := range files {
		raw, err := os.ReadFile(path)
		if err != nil {
			fail(err.Error())
		}
		lang := filepath.Base(filepath.Dir(path))
		first, rest, _ := strings.Cut(string(raw), "\n")
		e := entry{Lang: lang, Path: path, Reqs: reqRef.FindAllString(first, -1), Source: rest}
		e.Ours = ours(lang, []byte(rest), *runs)
		oursTotal += e.Ours.Time
		if *ref != "" {
			e.Ref = reference(*ref, tmp, lang, rest)
			refTotal += e.Ref.Time
		}
		entries = append(entries, e)
	}
	page := render(entries, reqs, *ref != "", oursTotal, refTotal)
	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		fail(err.Error())
	}
	if err := os.WriteFile(*out, page, 0o644); err != nil {
		fail(err.Error())
	}
	fmt.Printf("%s: %d samples, %.1f MB\n", *out, len(entries), float64(len(page))/1e6)
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "review:", msg)
	os.Exit(1)
}

// ours koloruje próbkę i mierzy medianę czasu z runs przebiegów.
func ours(lang string, src []byte, runs int) side {
	frag, ok := highlight.HTML(lang, src)
	if !ok {
		return side{Err: "nieznany język " + lang}
	}
	times := make([]time.Duration, runs)
	for i := range times {
		start := time.Now()
		highlight.HTML(lang, src)
		times[i] = time.Since(start)
	}
	sort.Slice(times, func(a, b int) bool { return times[a] < times[b] })
	s := side{Bytes: len(frag), Time: times[len(times)/2]}
	s.Spans, s.Classes = countSpans(string(frag))
	pre := `<pre class="hl"><code class="language-` + lang + `">` + string(frag) + "</code></pre>"
	s.Light, s.Dark = pre, pre
	return s
}

// reference uruchamia odniesienie na dokumencie z jednym blokiem kodu, w wyglądzie jasnym i ciemnym.
func reference(prog, tmp, lang, code string) side {
	name := lang
	if n, ok := refNames[lang]; ok {
		name = n
	}
	fence := strings.Repeat("`", max(3, longestRun(code, '`')+1))
	doc := filepath.Join(tmp, lang+".md")
	if err := os.WriteFile(doc, []byte(fence+name+"\n"+strings.TrimSuffix(code, "\n")+"\n"+fence+"\n"), 0o644); err != nil {
		return side{Err: err.Error()}
	}
	var s side
	for _, appearance := range []string{"light", "dark"} {
		start := time.Now()
		page, err := exec.Command(prog, "--syntax-highlight", "on", "--appearance", appearance, doc).Output()
		elapsed := time.Since(start)
		if err != nil {
			return side{Err: fmt.Sprintf("%s: %v", filepath.Base(prog), err)}
		}
		if appearance == "light" {
			s.Light, s.Time = string(page), elapsed
			frag := refPre.FindString(string(page))
			s.Bytes = len(frag)
			s.Spans, s.Classes = countSpans(frag)
		} else {
			s.Dark = string(page)
		}
	}
	return s
}

// countSpans liczy elementy <span class> i różne klasy w fragmencie.
func countSpans(frag string) (spans, classes int) {
	seen := map[string]bool{}
	for _, m := range spanOpen.FindAllStringSubmatch(frag, -1) {
		spans++
		seen[m[1]] = true
	}
	return spans, len(seen)
}

func longestRun(s string, c byte) int {
	best, n := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			n++
			best = max(best, n)
		} else {
			n = 0
		}
	}
	return best
}

// loadRequirements czyta treść wymagań z requirements/ (identyfikator → opis bez znaczników).
func loadRequirements() map[string]string {
	m := map[string]string{}
	docs, _ := filepath.Glob("requirements/*.md")
	langs, _ := filepath.Glob("requirements/languages/*.md")
	for _, doc := range append(docs, langs...) {
		b, _ := os.ReadFile(doc)
		for _, line := range strings.Split(string(b), "\n") {
			id := reqRef.FindString(line)
			if id == "" || !strings.HasPrefix(line, "| "+id) {
				continue
			}
			text := strings.TrimSuffix(strings.TrimPrefix(line, "| "+id+" | "), " |")
			m[id] = strings.ReplaceAll(text, "`", "")
		}
	}
	return m
}

func render(entries []entry, reqs map[string]string, withRef bool, oursTotal, refTotal time.Duration) []byte {
	var b bytes.Buffer
	light, _ := highlight.ThemeCSS("light")
	dark, _ := highlight.ThemeCSS("dark")
	refInfo := "bez odniesienia"
	if withRef {
		refInfo = fmt.Sprintf("odniesienie: QLMarkdown (highlight), łącznie %s na %d procesów; gohighlightgo łącznie %s (mediana na próbkę, w procesie)",
			ms(refTotal), len(entries), us(oursTotal))
	}
	fmt.Fprintf(&b, pageHead, scope(light, "theme-light"), scope(dark, "theme-dark"),
		time.Now().Format("2006-01-02 15:04"), len(entries), html.EscapeString(refInfo))
	b.WriteString(`<section class="legend"><h2>Klasy</h2><pre class="hl theme-light"><code>`)
	for _, c := range []string{"hl-kw", "hl-type", "hl-builtin", "hl-str", "hl-esc", "hl-num", "hl-com", "hl-op",
		"hl-fn", "hl-tag", "hl-attr", "hl-ins", "hl-del", "hl-meta"} {
		fmt.Fprintf(&b, `<span class="%s">%s</span>  `, c, c)
	}
	b.WriteString("</code></pre></section>\n<nav class=\"toc\">")
	for _, e := range entries {
		fmt.Fprintf(&b, `<a href="#%s">%s</a>`, e.Lang+"-"+stem(e.Path), e.Lang)
	}
	b.WriteString("</nav>\n")
	for i, e := range entries {
		id := e.Lang + "-" + stem(e.Path)
		fmt.Fprintf(&b, "<article class=\"card\" id=%q data-id=%q>\n<header><span class=\"num\">%d</span><h3>%s</h3><code class=\"path\">%s</code></header>\n",
			id, id, i+1, html.EscapeString(e.Lang), html.EscapeString(e.Path))
		b.WriteString("<details class=\"reqs\"><summary>Wymagania (" + fmt.Sprint(len(e.Reqs)) + ")</summary><ul>")
		for _, r := range e.Reqs {
			fmt.Fprintf(&b, "<li><b>%s</b> %s</li>", r, html.EscapeString(reqs[r]))
		}
		b.WriteString("</ul></details>\n<table class=\"stats\"><tr><th></th><th>fragment</th><th>elementy &lt;span&gt;</th><th>różne klasy</th><th>czas</th></tr>")
		fmt.Fprintf(&b, "<tr><td>gohighlightgo</td><td>%d B</td><td>%d</td><td>%d</td><td>%s (w procesie)</td></tr>", e.Ours.Bytes, e.Ours.Spans, e.Ours.Classes, us(e.Ours.Time))
		if withRef {
			fmt.Fprintf(&b, "<tr><td>QLMarkdown</td><td>%d B</td><td>%d</td><td>%d</td><td>%s (cały proces)</td></tr>", e.Ref.Bytes, e.Ref.Spans, e.Ref.Classes, ms(e.Ref.Time))
		}
		b.WriteString("</table>\n<div class=\"cols\">\n<section><h4>gohighlightgo</h4>")
		if e.Ours.Err != "" {
			fmt.Fprintf(&b, "<p class=\"error\">%s</p>", html.EscapeString(e.Ours.Err))
		} else {
			b.WriteString(strings.Replace(e.Ours.Light, `class="hl"`, `class="hl theme-light ours"`, 1))
			b.WriteString(strings.Replace(e.Ours.Dark, `class="hl"`, `class="hl theme-dark ours"`, 1))
		}
		b.WriteString("</section>\n")
		if withRef {
			b.WriteString("<section><h4>QLMarkdown (highlight)</h4>")
			if e.Ref.Err != "" {
				fmt.Fprintf(&b, "<p class=\"error\">%s</p>", html.EscapeString(e.Ref.Err))
			} else {
				fmt.Fprintf(&b, "<iframe class=\"ref ref-light\" srcdoc=\"%s\"></iframe>", html.EscapeString(e.Ref.Light))
				fmt.Fprintf(&b, "<iframe class=\"ref ref-dark\" srcdoc=\"%s\"></iframe>", html.EscapeString(e.Ref.Dark))
			}
			b.WriteString("</section>\n")
		}
		b.WriteString("</div>\n<footer>" +
			radio(id, "v", "ok", "OK") + radio(id, "v", "fix", "Do poprawy") + "<span class=\"sep\"></span>" +
			radio(id, "c", "better", "lepsze niż odniesienie") + radio(id, "c", "same", "porównywalne") + radio(id, "c", "worse", "gorsze") +
			"<textarea rows=\"1\" placeholder=\"Uwagi do tej próbki\"></textarea></footer>\n</article>\n")
	}
	b.WriteString(pageTail)
	return b.Bytes()
}

func radio(id, group, value, label string) string {
	return fmt.Sprintf(`<label><input type="radio" name="%s-%s" data-group="%s" value="%s"> %s</label>`, group, id, group, value, label)
}

// scope ogranicza arkusz motywu do bloków z klasą motywu („.hl ” → „.hl.theme-… ”), żeby oba motywy mieściły
// się na jednej stronie. Spacja po .hl odróżnia selektor bloku od nazw klas .hl-kw.
func scope(css, class string) string {
	return strings.ReplaceAll(css, ".hl ", ".hl."+class+" ")
}

func stem(path string) string { return strings.TrimSuffix(filepath.Base(path), ".txt") }

func us(d time.Duration) string { return fmt.Sprintf("%.0f µs", float64(d)/float64(time.Microsecond)) }

func ms(d time.Duration) string { return fmt.Sprintf("%.1f ms", float64(d)/float64(time.Millisecond)) }

const pageHead = `<!doctype html>
<html lang="pl">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>gohighlightgo: przegląd korpusu</title>
<style>
%s
%s
:root { --ink:#1b1f24; --muted:#5d6875; --line:#d5dae1; --paper:#ffffff; --desk:#f3f5f8; --ok:#1f7a4d; --fix:#b3421c; --accent:#2f5d8a; }
* { box-sizing: border-box; }
body { margin:0; background:var(--desk); color:var(--ink); font:15px/1.5 system-ui, -apple-system, "Segoe UI", sans-serif; }
.top { position:sticky; top:0; z-index:5; background:var(--paper); border-bottom:1px solid var(--line); padding:10px 20px; display:flex; flex-wrap:wrap; gap:10px 18px; align-items:center; }
.top h1 { font-size:17px; margin:0 8px 0 0; }
.top .meta { color:var(--muted); font-size:13px; }
.top .tally { font-variant-numeric:tabular-nums; }
.top button { font:inherit; padding:5px 12px; border:1px solid var(--line); border-radius:6px; background:var(--paper); color:var(--ink); cursor:pointer; }
.top button.primary { background:var(--accent); border-color:var(--accent); color:#fff; }
main { max-width:1700px; margin:0 auto; padding:8px 20px 80px; }
h2 { font-size:15px; margin:18px 0 6px; }
.toc { display:flex; flex-wrap:wrap; gap:4px 12px; margin:10px 0 18px; }
.toc a { color:var(--accent); }
.card { background:var(--paper); border:1px solid var(--line); border-left:5px solid var(--line); border-radius:8px; margin:0 0 18px; padding:14px 18px; }
.card.ok { border-left-color:var(--ok); }
.card.fix { border-left-color:var(--fix); }
.card header { display:flex; flex-wrap:wrap; align-items:baseline; gap:10px; }
.card h3 { font-size:17px; margin:0; }
.num, .path { color:var(--muted); }
.path { font-size:12.5px; }
.reqs { margin:6px 0; color:var(--muted); font-size:13px; }
.reqs ul { margin:4px 0; padding-left:18px; max-width:130ch; }
.stats { border-collapse:collapse; font-size:13px; margin:6px 0 10px; font-variant-numeric:tabular-nums; }
.stats th, .stats td { border-bottom:1px solid var(--line); padding:2px 12px 2px 0; text-align:left; }
.stats th { color:var(--muted); font-weight:600; }
.cols { display:grid; grid-template-columns:repeat(2, minmax(0, 1fr)); gap:14px; align-items:start; }
.cols.single { grid-template-columns:1fr; }
@media (max-width: 1100px) { .cols { grid-template-columns:1fr; } }
h4 { font-size:12px; text-transform:uppercase; letter-spacing:.05em; color:var(--muted); margin:0 0 4px; font-weight:600; }
pre.hl { margin:0; padding:10px 12px; border:1px solid var(--line); border-radius:6px; font:12.5px/1.45 ui-monospace, SFMono-Regular, Menlo, monospace; overflow:auto; white-space:pre; }
.legend pre.hl { white-space:pre-wrap; }
iframe.ref { width:100%%; border:1px solid var(--line); border-radius:6px; height:200px; background:#fff; }
body.dark .theme-light.ours, body.dark .ref-light, body:not(.dark) .theme-dark.ours, body:not(.dark) .ref-dark { display:none; }
.error { color:var(--fix); margin:0; }
.card footer { display:flex; flex-wrap:wrap; gap:8px 16px; align-items:center; margin-top:12px; padding-top:10px; border-top:1px solid var(--line); }
.card footer label { cursor:pointer; }
.card footer .sep { width:1px; height:20px; background:var(--line); }
.card footer textarea { flex:1 1 320px; font:inherit; padding:5px 8px; border:1px solid var(--line); border-radius:6px; resize:vertical; min-height:32px; }
body.only-open .card.ok { display:none; }
#copied { color:var(--ok); font-size:13px; }
</style>
</head>
<body>
<div class="top">
  <h1>gohighlightgo: przegląd korpusu</h1>
  <span class="meta">wygenerowano %s · %d próbek · %s</span>
  <span class="tally" id="tally"></span>
  <label><input type="checkbox" id="dark"> motyw ciemny</label>
  <label><input type="checkbox" id="onlyOpen"> ukryj zaakceptowane</label>
  <button class="primary" id="copy">Kopiuj oceny i uwagi</button>
  <button id="reset">Wyczyść oceny</button>
  <span id="copied"></span>
</div>
<main>
`

const pageTail = `</main>
<script>
(function () {
  var KEY = 'gohighlightgo-review';
  var state = {};
  try { state = JSON.parse(localStorage.getItem(KEY) || '{}'); } catch (e) {}
  var cards = Array.prototype.slice.call(document.querySelectorAll('.card'));
  function save() { try { localStorage.setItem(KEY, JSON.stringify(state)); } catch (e) {} }
  function tally() {
    var ok = 0, fix = 0, better = 0, worse = 0;
    cards.forEach(function (c) {
      var s = state[c.dataset.id] || {};
      c.classList.toggle('ok', s.v === 'ok');
      c.classList.toggle('fix', s.v === 'fix');
      if (s.v === 'ok') ok++;
      if (s.v === 'fix') fix++;
      if (s.c === 'better') better++;
      if (s.c === 'worse') worse++;
    });
    document.getElementById('tally').textContent = 'OK ' + ok + ' · do poprawy ' + fix + ' · bez oceny ' + (cards.length - ok - fix) +
      ' · lepsze ' + better + ' · gorsze ' + worse;
  }
  cards.forEach(function (c) {
    var id = c.dataset.id, s = state[id] || {};
    var note = c.querySelector('textarea');
    note.value = s.note || '';
    c.querySelectorAll('input[type=radio]').forEach(function (r) {
      r.checked = s[r.dataset.group] === r.value;
      r.addEventListener('change', function () { (state[id] = state[id] || {})[r.dataset.group] = r.value; save(); tally(); });
    });
    note.addEventListener('input', function () { (state[id] = state[id] || {}).note = note.value; save(); });
  });
  function fit(f) {
    try { var d = f.contentDocument; if (d && d.body) { f.style.height = (d.documentElement.scrollHeight + 4) + 'px'; } } catch (e) {}
  }
  document.querySelectorAll('iframe.ref').forEach(function (f) { f.addEventListener('load', function () { fit(f); }); });
  var dark = document.getElementById('dark');
  var prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
  try { var saved = localStorage.getItem(KEY + '-dark'); if (saved !== null) { prefersDark = saved === '1'; } } catch (e) {}
  dark.checked = prefersDark;
  document.body.classList.toggle('dark', prefersDark);
  dark.addEventListener('change', function () {
    document.body.classList.toggle('dark', dark.checked);
    try { localStorage.setItem(KEY + '-dark', dark.checked ? '1' : '0'); } catch (e) {}
    document.querySelectorAll('iframe.ref').forEach(fit);
  });
  document.getElementById('onlyOpen').addEventListener('change', function (e) { document.body.classList.toggle('only-open', e.target.checked); });
  document.getElementById('reset').addEventListener('click', function () {
    state = {}; save();
    cards.forEach(function (c) { c.querySelector('textarea').value = ''; c.querySelectorAll('input[type=radio]').forEach(function (r) { r.checked = false; }); });
    tally();
  });
  document.getElementById('copy').addEventListener('click', function () {
    var lines = [];
    cards.forEach(function (c) {
      var s = state[c.dataset.id] || {};
      var v = s.v === 'ok' ? 'OK' : s.v === 'fix' ? 'DO POPRAWY' : 'bez oceny';
      var cmp = s.c === 'better' ? ', lepsze' : s.c === 'same' ? ', porównywalne' : s.c === 'worse' ? ', gorsze' : '';
      lines.push(c.dataset.id + ': ' + v + cmp + (s.note ? ' — ' + s.note.trim() : ''));
    });
    var text = lines.join('\n');
    var done = function () { document.getElementById('copied').textContent = 'skopiowano ' + lines.length + ' pozycji'; };
    if (navigator.clipboard && navigator.clipboard.writeText) {
      navigator.clipboard.writeText(text).then(done, function () { window.prompt('Skopiuj:', text); });
    } else { window.prompt('Skopiuj:', text); }
  });
  tally();
})();
</script>
</body>
</html>
`
