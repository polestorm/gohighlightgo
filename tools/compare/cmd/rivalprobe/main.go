// Polecenie rivalprobe sprawdza niezmienniki gohighlightgo na naszej bibliotece i na konkurentach, na tych
// samych wejściach:
//
//   - text: odwracalność (po zdjęciu znaczników zostaje wejście) i determinizm, na korpusie i trudnych wejściach;
//     gohighlightgo i chroma;
//   - ours-time, chroma-time, highlight-time: czas dla wzorców złośliwych, wejście ×16.
//
// Każde wywołanie konkurenta biegnie w osobnym procesie (z własną grupą procesów) z twardym limitem czasu,
// więc zawieszony lekser zostaje zabity i trafia do wyników. highlight (André Simon) mierzymy jako czarną
// skrzynkę przez qlmarkdown_cli z QLMarkdown: dokument Markdown z jednym blokiem kodu; czas obejmuje start
// procesu i Markdown, co tylko utrudnia wykrycie czasu gorszego niż liniowy. Postęp: strona HTML co 2 s.
//
//	cd tools/compare && go run ./cmd/rivalprobe -page ../../out/probe.html > probe.json
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/polestorm/gohighlightgo/tools/compare"
)

// patterns: wzorce złośliwe; każdy powtórzony z "a " daje wiele niezamkniętych konstrukcji.
var patterns = []string{"\"", "'", "`", "/*", "(*", "<!--", "<", "${", "#", "\\", "[["}

// patternNames: nazwy plików wzorców w inputs/hostile (<nazwa>-small.txt, <nazwa>-large.txt).
var patternNames = []string{"dquote", "squote", "backtick", "block-comment", "paren-star", "html-comment", "lt",
	"dollar-brace", "hash", "backslash", "double-bracket"}

// tricky: trudne wejścia dla odwracalności.
var tricky = map[string][]byte{
	"CRLF line ends":       []byte("a = 1\r\nb = \"x\"\r\n"),
	"lone carriage return": []byte("a\rb\r"),
	"tabs":                 []byte("\tif x {\n\t\treturn \"a\\tb\"\n\t}"),
	"no final newline":     []byte("x = 1"),
	"invalid UTF-8":        []byte("s = \"\xff\xfe\" // \xc3\n"),
	"NUL byte":             []byte("a\x00b \"c\x00\""),
	"non-ASCII text":       []byte("// zażółć 🙂\nx = \"日本\"\n"),
}

const (
	smallReps = 2000  // ok. 4–10 kB
	bigReps   = 32000 // ×16
	ratioMax  = 96    // jak TestPathological: ×16 danych może kosztować najwyżej ×96 czasu
	slowMs    = 1000  // liniowy, ale ponad 1 s dla dużego wejścia
	minMs     = 200   // poniżej tego czasu nie oceniamy proporcji
)

// refNames: nazwa języka w highlight, gdy różni się od naszej (jak w tools/review).
var refNames = map[string]string{"pascal": "delphi", "objectivec": "objc", "fortran": "f90", "make": "makefile",
	"latex": "tex", "protobuf": "pb", "commonlisp": "lisp", "armasm": "arm", "prolog": "pro", "scheme": "scm",
	"vhdl": "vhd", "moonscript": "moon", "batch": "bat", "puppet": "pp"}

const qlm = "/Applications/QLMarkdown.app/Contents/Resources/qlmarkdown_cli"
const langDefs = "/Applications/QLMarkdown.app/Contents/Resources/highlight/langDefs"

var tagRe = regexp.MustCompile(`<[^>]*>`)

func strip(b []byte) []byte {
	return []byte(html.UnescapeString(tagRe.ReplaceAllString(string(b), "")))
}

// sameText: wynik po zdjęciu znaczników zachowuje wejście. Dla błędnego UTF-8 i NUL dopuszczamy zamianę na
// U+FFFD (reguła gohighlightgo z ADR 0002, chroma robi podobnie); oceniamy obie biblioteki tak samo.
func sameText(out, in []byte, lenient bool) bool {
	got := strip(out)
	if bytes.Equal(got, in) {
		return true
	}
	return lenient && bytes.Equal(fffd(got), fffd(in))
}

func fffd(b []byte) []byte {
	s := strings.ToValidUTF8(strings.ReplaceAll(string(b), "\x00", "�"), "�")
	for strings.Contains(s, "��") {
		s = strings.ReplaceAll(s, "��", "�")
	}
	return []byte(s)
}

func lenientCase(k string) bool { return k == "invalid UTF-8" || k == "NUL byte" }

func input(p, reps int) []byte { return []byte(strings.Repeat(patterns[p]+"a ", reps)) }

func main() {
	child := flag.String("child", "", "internal: case|time")
	lang := flag.String("lang", "", "internal")
	pat := flag.Int("pat", 0, "internal")
	kase := flag.String("case", "", "internal")
	root := flag.String("root", "../..", "repository root")
	page := flag.String("page", "../../out/probe.html", "progress page")
	stages := flag.String("stages", "text,ours-time,chroma-time,highlight-time", "stages to run")
	only := flag.String("only", "", "file with lines `lang \"pattern\"`: time only these")
	workers := flag.Int("workers", 4, "parallel processes")
	timeout := flag.Duration("timeout", 15*time.Second, "limit for one rival process")
	budget := flag.Duration("budget", 18*time.Minute, "total time limit")
	flag.Parse()
	switch *child {
	case "case":
		childCase(*root, *lang, *kase)
	case "time":
		childTime(*lang, *pat)
	default:
		run(*root, *page, *stages, *only, *workers, *timeout, *budget)
	}
}

// childCase: chroma, jedno wejście; JSON [tekst zmieniony, różny drugi wynik, błąd].
func childCase(root, lang, k string) {
	in := tricky[k]
	if k == "corpus" {
		in = corpusOf(root)[lang]
	}
	lx := compare.ChromaLexer(lang)
	c, err := compare.ChromaKeepLF(lx, in)
	if err != nil {
		json.NewEncoder(os.Stdout).Encode([3]bool{false, false, true})
		return
	}
	c2, _ := compare.ChromaKeepLF(lx, in)
	json.NewEncoder(os.Stdout).Encode([3]bool{!sameText(c, in, lenientCase(k)), !bytes.Equal(c, c2), false})
}

// childTime: chroma, małe i duże wejście; JSON [ms małe, ms duże].
func childTime(lang string, p int) {
	lx := compare.ChromaLexer(lang)
	var r [2]float64
	for i, in := range [][]byte{input(p, smallReps), input(p, bigReps)} {
		t := time.Now()
		compare.Chroma(lx, in)
		r[i] = ms(time.Since(t))
	}
	json.NewEncoder(os.Stdout).Encode(r)
}

func ms(d time.Duration) float64 { return d.Seconds() * 1000 }

func corpusOf(root string) map[string][]byte {
	samples, err := compare.Corpus(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	m := map[string][]byte{}
	for _, s := range samples {
		m[s.Lang] = append(m[s.Lang], s.Code...)
	}
	return m
}

// limited uruchamia proces we własnej grupie i zabija całą grupę po przekroczeniu limitu.
func limited(timeout time.Duration, name string, args ...string) (out []byte, elapsed time.Duration, timedOut bool, err error) {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var b bytes.Buffer
	cmd.Stdout = &b
	start := time.Now()
	if err := cmd.Start(); err != nil {
		return nil, 0, false, err
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case err = <-done:
		return b.Bytes(), time.Since(start), false, err
	case <-time.After(timeout):
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		<-done
		return nil, time.Since(start), true, nil
	}
}

// state: postęp i wyniki (biblioteka → rodzaj → przypadki).
type state struct {
	mu       sync.Mutex
	start    time.Time
	phase    string
	done     int
	total    int
	skipped  int
	finished bool
	langs    map[string]int
	finding  map[string]map[string][]string
}

func (s *state) add(lib, kind, item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.finding[lib] == nil {
		s.finding[lib] = map[string][]string{}
	}
	s.finding[lib][kind] = append(s.finding[lib][kind], item)
}

func (s *state) step() {
	s.mu.Lock()
	s.done++
	s.mu.Unlock()
}

// judge ocenia parę pomiarów czasu.
func judge(st *state, lib, item string, small, big float64) {
	item = fmt.Sprintf("%s %.0f→%.0f ms", item, small, big)
	switch {
	case big > minMs && big > ratioMax*max(small, 0.05):
		st.add(lib, "superlinear time (input ×16, time ×96 or more)", item)
	case big > slowMs:
		st.add(lib, "linear, but over 1 s for the large input", item)
	}
}

type job struct {
	kind, lang string
	p          int
	kase       string
}

func run(root, page, stages, only string, workers int, timeout, budget time.Duration) {
	want := map[string]bool{}
	for _, s := range strings.Split(stages, ",") {
		want[strings.TrimSpace(s)] = true
	}
	corpus := corpusOf(root)
	var all, withChroma, withHL []string
	for l := range corpus {
		all = append(all, l)
		if compare.ChromaLexer(l) != nil {
			withChroma = append(withChroma, l)
		}
		name := l
		if n, ok := refNames[l]; ok {
			name = n
		}
		if _, err := os.Stat(filepath.Join(langDefs, name+".lang")); err == nil {
			withHL = append(withHL, l)
		}
	}
	sort.Strings(all)
	sort.Strings(withChroma)
	sort.Strings(withHL)
	picks := map[string]bool{}
	if only != "" {
		b, err := os.ReadFile(only)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
			l, q, _ := strings.Cut(line, " ")
			q, _, _ = strings.Cut(q, " ")
			for i, pt := range patterns {
				if fmt.Sprintf("%q", pt) == q {
					picks[fmt.Sprint(l, "\x00", i)] = true
				}
			}
		}
	}
	picked := func(l string, p int) bool { return only == "" || picks[fmt.Sprint(l, "\x00", p)] }

	st := &state{start: time.Now(), langs: map[string]int{}, finding: map[string]map[string][]string{}}
	var jobs []job
	if want["text"] {
		st.langs["gohighlightgo (text)"] = len(all)
		st.langs["chroma (text)"] = len(withChroma)
		for _, l := range withChroma {
			jobs = append(jobs, job{kind: "case", lang: l, kase: "corpus"})
			for k := range tricky {
				jobs = append(jobs, job{kind: "case", lang: l, kase: k})
			}
		}
	}
	for p := range patterns {
		if want["chroma-time"] {
			for _, l := range withChroma {
				if picked(l, p) {
					jobs = append(jobs, job{kind: "chroma-time", lang: l, p: p})
				}
			}
		}
		if want["highlight-time"] {
			for _, l := range withHL {
				if picked(l, p) {
					jobs = append(jobs, job{kind: "highlight-time", lang: l, p: p})
				}
			}
		}
	}
	if want["chroma-time"] {
		st.langs["chroma (time)"] = len(withChroma)
	}
	if want["highlight-time"] {
		st.langs["highlight (time)"] = len(withHL)
	}
	st.total = len(jobs)
	if want["text"] {
		st.total += len(all)
	}
	if want["ours-time"] {
		st.total += len(all)
		st.langs["gohighlightgo (time)"] = len(all)
	}

	stop := make(chan struct{})
	pageDone := make(chan struct{})
	go func() {
		defer close(pageDone)
		for {
			writePage(page, st, workers, timeout, budget)
			select {
			case <-stop:
				writePage(page, st, workers, timeout, budget)
				return
			case <-time.After(time.Second):
			}
		}
	}()

	// Nasza biblioteka biegnie w tym procesie: zawieszenie zatrzymałoby cały pomiar, co też byłoby wynikiem.
	if want["text"] {
		st.phase = "gohighlightgo: round-trip and determinism"
		for _, l := range all {
			inputs := map[string][]byte{"corpus": corpus[l]}
			for k, v := range tricky {
				inputs[k] = v
			}
			for k, in := range inputs {
				o := compare.Ours(l, in)
				if !sameText(o, in, lenientCase(k)) {
					st.add("gohighlightgo", "text changed ("+k+")", l)
				}
				if !bytes.Equal(o, compare.Ours(l, in)) {
					st.add("gohighlightgo", "different output on second run ("+k+")", l)
				}
				if !utf8.Valid(o) {
					st.add("gohighlightgo", "invalid UTF-8 in output ("+k+")", l)
				}
			}
			st.step()
		}
	}
	if want["ours-time"] {
		st.phase = "gohighlightgo: time"
		for _, l := range all {
			for p := range patterns {
				if !picked(l, p) && only != "" {
					continue
				}
				var r [2]float64
				for i, in := range [][]byte{input(p, smallReps), input(p, bigReps)} {
					t := time.Now()
					compare.Ours(l, in)
					r[i] = ms(time.Since(t))
				}
				judge(st, "gohighlightgo", fmt.Sprintf("%s %q", l, patterns[p]), r[0], r[1])
			}
			st.step()
		}
	}

	st.phase = "rivals: one process per call"
	self, _ := os.Executable()
	tmp, _ := os.MkdirTemp("", "rivalprobe")
	defer os.RemoveAll(tmp)
	deadline := st.start.Add(budget)
	ch := make(chan job)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for j := range ch {
				if time.Now().After(deadline) {
					st.mu.Lock()
					st.skipped++
					st.mu.Unlock()
					st.step()
					continue
				}
				runJob(st, j, w, self, root, tmp, timeout)
				st.step()
			}
		}(w)
	}
	for _, j := range jobs {
		ch <- j
	}
	close(ch)
	wg.Wait()
	st.mu.Lock()
	st.finished = true
	st.phase = "finished"
	st.mu.Unlock()
	close(stop)
	<-pageDone
	st.mu.Lock()
	defer st.mu.Unlock()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(map[string]any{"date": time.Now().Format("2006-01-02"), "elapsed": time.Since(st.start).Round(time.Second).String(),
		"skipped": st.skipped, "languages": st.langs, "timeout": timeout.String(), "workers": workers, "findings": st.finding})
}

func runJob(st *state, j job, w int, self, root, tmp string, timeout time.Duration) {
	switch j.kind {
	case "case":
		out, _, to, err := limited(timeout, self, "-child", "case", "-root", root, "-lang", j.lang, "-case", j.kase)
		var r [3]bool
		switch {
		case to:
			st.add("chroma", "hang, killed after "+timeout.String()+" ("+j.kase+")", j.lang)
		case err != nil || json.Unmarshal(out, &r) != nil:
			st.add("chroma", "crash ("+j.kase+")", j.lang)
		default:
			if r[0] {
				st.add("chroma", "text changed ("+j.kase+")", j.lang)
			}
			if r[1] {
				st.add("chroma", "different output on second run ("+j.kase+")", j.lang)
			}
			if r[2] {
				st.add("chroma", "error ("+j.kase+")", j.lang)
			}
		}
	case "chroma-time":
		item := fmt.Sprintf("%s %q", j.lang, patterns[j.p])
		out, _, to, err := limited(timeout, self, "-child", "time", "-lang", j.lang, "-pat", fmt.Sprint(j.p))
		var r [2]float64
		switch {
		case to:
			st.add("chroma", "hang or over "+timeout.String()+" (killed)", item)
		case err != nil || json.Unmarshal(out, &r) != nil:
			st.add("chroma", "crash", item)
		default:
			judge(st, "chroma", item, r[0], r[1])
		}
	case "highlight-time":
		item := fmt.Sprintf("%s %q", j.lang, patterns[j.p])
		name := j.lang
		if n, ok := refNames[j.lang]; ok {
			name = n
		}
		var r [2]float64
		for i, reps := range []int{smallReps, bigReps} {
			doc := filepath.Join(tmp, fmt.Sprintf("w%d-%d.md", w, i))
			os.WriteFile(doc, []byte("~~~"+name+"\n"+string(input(j.p, reps))+"\n~~~\n"), 0o644)
			_, d, to, err := limited(timeout, qlm, "--syntax-highlight", "on", doc)
			if to {
				st.add("highlight", "hang or over "+timeout.String()+" (killed)", item)
				return
			}
			if err != nil {
				st.add("highlight", "crash", item)
				return
			}
			r[i] = ms(d)
		}
		judge(st, "highlight", item, r[0], r[1])
	}
}

func writePage(path string, st *state, workers int, timeout, budget time.Duration) {
	st.mu.Lock()
	defer st.mu.Unlock()
	el := time.Since(st.start)
	pct := float64(st.done) / float64(max(st.total, 1)) * 100
	eta := "—"
	if st.done > 0 && !st.finished {
		eta = (time.Duration(float64(el)/float64(st.done)*float64(st.total-st.done)) / time.Second * time.Second).String()
	}
	var rows strings.Builder
	libs := make([]string, 0, len(st.finding))
	for l := range st.finding {
		libs = append(libs, l)
	}
	sort.Strings(libs)
	for _, lib := range libs {
		var kinds []string
		for k := range st.finding[lib] {
			kinds = append(kinds, k)
		}
		sort.Strings(kinds)
		for _, k := range kinds {
			v := st.finding[lib][k]
			show, more := v, ""
			if len(show) > 30 {
				show, more = show[:30], " …"
			}
			fmt.Fprintf(&rows, "<tr><td>%s</td><td>%s</td><td class=n>%d</td><td>%s%s</td></tr>", lib, html.EscapeString(k), len(v),
				html.EscapeString(strings.Join(show, ", ")), more)
		}
	}
	var counts []string
	for k, v := range st.langs {
		counts = append(counts, fmt.Sprintf("%s: %d", k, v))
	}
	sort.Strings(counts)
	refresh, status := `<meta http-equiv="refresh" content="2">`, "running"
	if st.finished {
		refresh, status = "", "finished"
	}
	doc := fmt.Sprintf(`<!doctype html><meta charset=utf-8>%s<title>Rival probe</title>
<style>body{font:15px system-ui;margin:24px;max-width:1200px;background:#fff;color:#111}
.bar{height:22px;background:#e5e7eb;border-radius:6px;overflow:hidden}.bar div{height:100%%;background:#2563eb;width:%.1f%%}
table{border-collapse:collapse;width:100%%;margin-top:16px}td,th{border-bottom:1px solid #ddd;padding:6px;text-align:left;vertical-align:top}
.n{text-align:right}.big{font-size:28px;font-weight:600}</style>
<h1>Invariants: gohighlightgo, chroma, highlight</h1>
<p class=big>%.1f%% — %s</p><div class=bar><div></div></div>
<p>Phase: <b>%s</b>. Done %d of %d steps. Elapsed %s, estimated remaining <b>%s</b>
(budget %s, %d processes, %s per rival process). Skipped for the budget: %d.<br>Languages: %s.</p>
<table><tr><th>Library</th><th>Finding</th><th class=n>Count</th><th>Cases</th></tr>%s</table>`,
		refresh, pct, pct, status, st.phase, st.done, st.total, el.Round(time.Second), eta, budget, workers, timeout,
		st.skipped, strings.Join(counts, ", "), rows.String())
	tmp := path + ".tmp"
	os.WriteFile(tmp, []byte(doc), 0o644)
	os.Rename(tmp, path)
}
