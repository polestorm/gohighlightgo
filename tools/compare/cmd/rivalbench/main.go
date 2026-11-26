// Polecenie rivalbench powtarza pomiary z tabeli porównania w README na zamrożonych wejściach z inputs/
// i wypisuje wyniki jako tabelę Markdown (oraz JSON z -json). Każdy może je uruchomić u siebie:
//
//	cd tools/compare && go run ./cmd/rivalbench
//
// Mierzy gohighlightgo i chroma przez ich API w Go oraz highlight (André Simon) jako czarną skrzynkę przez
// polecenie QLMarkdown, gdy QLMarkdown jest zainstalowany. Wejścia i sposób pomiaru opisuje inputs/README.md.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/polestorm/gohighlightgo/tools/compare"
)

// refNames: nazwa języka w highlight, gdy różni się od naszej (jak w rivalprobe).
var refNames = map[string]string{"pascal": "delphi"}

type result struct {
	Date, Machine, Go, Chroma, QLMarkdown string
	Rows                                  []row
}

type row struct {
	Name                    string
	Ours, Chroma, Highlight string
	OursV, ChromaV, HighV   float64 `json:",omitempty"`
}

func main() {
	testing.Init()
	inputs := flag.String("inputs", "inputs", "directory with the frozen inputs")
	root := flag.String("root", "../..", "repository root (for the language count)")
	runs := flag.Int("runs", 5, "runs of each benchmark; the median is reported")
	docRuns := flag.Int("doc-runs", 15, "runs of each process for the document and peak-memory rows")
	app := flag.String("qlmarkdown", "/Applications/QLMarkdown.app", "QLMarkdown application (highlight); empty to skip")
	jsonOut := flag.String("json", "", "also write the results to this JSON file")
	flag.Parse()

	corpus, err := compare.CorpusDir(filepath.Join(*inputs, "corpus"))
	check(err)
	qlm := filepath.Join(*app, "Contents/Resources/qlmarkdown_cli")
	hasHL := *app != "" && exists(qlm)

	res := result{Date: time.Now().Format("2006-01-02"), Machine: machine(), Go: runtime.Version(), Chroma: chromaVersion()}
	if hasHL {
		res.QLMarkdown = plist(filepath.Join(*app, "Contents/Info.plist"))
	}
	add := func(r row) {
		res.Rows = append(res.Rows, r)
		fmt.Fprintf(os.Stderr, "%-34s %14s %14s %14s\n", r.Name, r.Ours, r.Chroma, r.Highlight)
	}
	na := "not measured"

	// Liczba języków.
	all, err := compare.Corpus(*root)
	check(err)
	hlLangs := na
	if hasHL {
		defs, _ := filepath.Glob(filepath.Join(*app, "Contents/Resources/highlight/langDefs/*.lang"))
		hlLangs = fmt.Sprint(len(defs))
	}
	add(row{Name: "Languages", Ours: fmt.Sprint(compare.PrimaryLanguages(all)), Chroma: fmt.Sprint(compare.ChromaLanguages()), Highlight: hlLangs})

	// Rozmiar dodany do programu: probes/ours i probes/chroma minus probes/none.
	tmp, err := os.MkdirTemp("", "rivalbench")
	check(err)
	defer os.RemoveAll(tmp)
	bin := map[string]string{}
	for _, p := range []string{"none", "ours", "chroma"} {
		bin[p] = filepath.Join(tmp, p)
		cmd := exec.Command("go", "build", "-trimpath", "-ldflags=-s -w", "-o", bin[p], "./probes/"+p)
		cmd.Stderr = os.Stderr
		check(cmd.Run())
	}
	base := size(bin["none"])
	hlSize, hlSizeV := na, 0.0
	if hasHL {
		lib := filepath.Join(*app, "Contents/Frameworks/libwrapper_highlight.dylib")
		thin := filepath.Join(tmp, "hl.dylib")
		libSize := size(lib)
		if exec.Command("lipo", "-thin", runtime.GOARCH, lib, "-output", thin).Run() == nil {
			libSize = size(thin)
		}
		files := treeSize(filepath.Join(*app, "Contents/Resources/highlight/langDefs")) +
			treeSize(filepath.Join(*app, "Contents/Resources/highlight/themes"))
		hlSizeV = float64(libSize+files) / 1000
		hlSize = fmt.Sprintf("%s kB (library %s kB, definition and theme files %s kB)", kb(libSize+files), kb(libSize), kb(files))
	}
	ov, cv := float64(size(bin["ours"])-base)/1000, float64(size(bin["chroma"])-base)/1000
	add(row{Name: "Size added to a program", Ours: kb(size(bin["ours"])-base) + " kB", Chroma: kb(size(bin["chroma"])-base) + " kB",
		Highlight: hlSize, OursV: ov, ChromaV: cv, HighV: hlSizeV})

	// Czas i przydziały: 21 plików korpusu osobno oraz pliki Go sklejone do 1 MB.
	goMB := compare.Repeat(corpus, "go", 1<<20)
	ours := func(parts []compare.Sample) func() {
		return func() {
			for _, p := range parts {
				compare.Ours(p.Lang, p.Code)
			}
		}
	}
	chr := func(parts []compare.Sample) func() {
		return func() {
			for _, p := range parts {
				if _, err := compare.Chroma(compare.ChromaLexer(p.Lang), p.Code); err != nil {
					check(err)
				}
			}
		}
	}
	oneMB := []compare.Sample{{Lang: "go", Code: goMB}}
	co, cc := bench(*runs, ours(corpus)), bench(*runs, chr(corpus))
	mo, mc := bench(*runs, ours(oneMB)), bench(*runs, chr(oneMB))
	n := 0
	for _, s := range corpus {
		n += len(s.Code)
	}
	add(row{Name: fmt.Sprintf("Time, %d files (%.1f kB)", len(corpus), float64(n)/1000), Ours: msf(co.ns), Chroma: msf(cc.ns), Highlight: na, OursV: co.ns / 1e6, ChromaV: cc.ns / 1e6})
	add(row{Name: "Time, 1 MB of Go", Ours: msf(mo.ns), Chroma: msf(mc.ns), Highlight: na, OursV: mo.ns / 1e6, ChromaV: mc.ns / 1e6})

	// Dokument: korpus jako bloki kodu, powtórzony do 1,3 MB. Dla bibliotek Go: czas pokolorowania wszystkich
	// bloków; dla highlight: czas polecenia QLMarkdown z kolorowaniem minus bez niego.
	doc, blocks := document(corpus, 1_300_000)
	docFile := filepath.Join(tmp, "doc.md")
	check(os.WriteFile(docFile, doc, 0o644))
	do, dc := bench(*runs, ours(blocks)), bench(*runs, chr(blocks))
	hlDoc, hlDocV := na, 0.0
	if hasHL {
		on := median(*docRuns, func() float64 { t, _ := proc(nil, qlm, "--syntax-highlight", "on", docFile); return t })
		off := median(*docRuns, func() float64 { t, _ := proc(nil, qlm, "--syntax-highlight", "off", docFile); return t })
		hlDocV = on - off
		hlDoc = fmt.Sprintf("%.0f ms", hlDocV)
	}
	add(row{Name: fmt.Sprintf("Time, %.1f MB of code", float64(len(doc))/1e6), Ours: msf(do.ns), Chroma: msf(dc.ns), Highlight: hlDoc, OursV: do.ns / 1e6, ChromaV: dc.ns / 1e6, HighV: hlDocV})

	add(row{Name: fmt.Sprintf("Memory allocated, %d files", len(corpus)), Ours: mbf(co.bytes), Chroma: mbf(cc.bytes), Highlight: na, OursV: co.bytes / 1e6, ChromaV: cc.bytes / 1e6})
	add(row{Name: "Memory allocated, 1 MB of Go", Ours: mbf(mo.bytes), Chroma: mbf(mc.bytes), Highlight: na, OursV: mo.bytes / 1e6, ChromaV: mc.bytes / 1e6})

	// Pamięć szczytowa dodana: proces koloruje 1 MB Go raz; minus ten sam proces bez kolorowania.
	peak := func(name string, args ...string) float64 {
		return median(*docRuns, func() float64 { _, rss := proc(goMB, name, args...); return rss })
	}
	pBase := peak(bin["none"], "go")
	po, pc := peak(bin["ours"], "go")-pBase, peak(bin["chroma"], "go")-pBase
	hlPeak, hlPeakV := na, 0.0
	if hasHL {
		md := filepath.Join(tmp, "go.md")
		check(os.WriteFile(md, append(append([]byte("`````go\n"), goMB...), "\n`````\n"...), 0o644))
		rss := func(mode string) float64 {
			return median(*docRuns, func() float64 { _, r := proc(nil, qlm, "--syntax-highlight", mode, md); return r })
		}
		hlPeakV = rss("on") - rss("off")
		hlPeak = fmt.Sprintf("%.0f MB", hlPeakV)
	}
	add(row{Name: "Peak memory added, 1 MB of Go", Ours: fmt.Sprintf("%.0f MB", po), Chroma: fmt.Sprintf("%.0f MB", pc), Highlight: hlPeak, OursV: po, ChromaV: pc, HighV: hlPeakV})

	// Rozmiar wyniku dla korpusu.
	var so, sc int
	for _, s := range corpus {
		so += len(compare.Ours(s.Lang, s.Code))
		c, err := compare.Chroma(compare.ChromaLexer(s.Lang), s.Code)
		check(err)
		sc += len(c)
	}
	add(row{Name: fmt.Sprintf("Output for the %d files", len(corpus)), Ours: fmt.Sprintf("%.1f kB", float64(so)/1000), Chroma: fmt.Sprintf("%.1f kB", float64(sc)/1000), Highlight: na, OursV: float64(so) / 1000, ChromaV: float64(sc) / 1000})

	printMarkdown(res)
	if *jsonOut != "" {
		b, _ := json.MarshalIndent(res, "", "  ")
		check(os.WriteFile(*jsonOut, b, 0o644))
	}
}

type benchResult struct{ ns, bytes float64 }

// bench uruchamia testing.Benchmark runs razy i zwraca medianę czasu i bajtów przydzielonych na wywołanie.
func bench(runs int, f func()) benchResult {
	var ns, by []float64
	for i := 0; i < runs; i++ {
		r := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				f()
			}
		})
		ns = append(ns, float64(r.NsPerOp()))
		by = append(by, float64(r.AllocedBytesPerOp()))
	}
	return benchResult{med(ns), med(by)}
}

// proc uruchamia proces z danym wejściem; zwraca czas w ms i największą pamięć rezydentną w MB.
func proc(stdin []byte, name string, args ...string) (float64, float64) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = bytes.NewReader(stdin)
	t := time.Now()
	check(cmd.Run())
	elapsed := float64(time.Since(t).Microseconds()) / 1000
	rss := float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss)
	if runtime.GOOS == "darwin" {
		rss /= 1 << 20 // bajty
	} else {
		rss /= 1 << 10 // kilobajty
	}
	return elapsed, rss
}

// document składa dokument Markdown z bloków kodu korpusu, powtórzonych do co najmniej n bajtów, i zwraca też
// te bloki jako próbki.
func document(corpus []compare.Sample, n int) ([]byte, []compare.Sample) {
	var b bytes.Buffer
	var blocks []compare.Sample
	for b.Len() < n {
		for _, s := range corpus {
			name := s.Lang
			if r, ok := refNames[name]; ok {
				name = r
			}
			fmt.Fprintf(&b, "`````%s\n%s\n`````\n\n", name, bytes.TrimRight(s.Code, "\n"))
			blocks = append(blocks, s)
		}
	}
	return b.Bytes(), blocks
}

func median(n int, f func() float64) float64 {
	v := make([]float64, n)
	for i := range v {
		v[i] = f()
	}
	return med(v)
}

func med(v []float64) float64 {
	s := append([]float64(nil), v...)
	sort.Float64s(s)
	return s[len(s)/2]
}

func printMarkdown(r result) {
	fmt.Printf("Measured %s on %s, %s; chroma %s; QLMarkdown %s.\n\n", r.Date, r.Machine, r.Go, r.Chroma, orDash(r.QLMarkdown))
	fmt.Println("| | gohighlightgo | chroma | highlight | advantage over chroma | over highlight |")
	fmt.Println("|---|--:|--:|--:|--:|--:|")
	for _, x := range r.Rows {
		fmt.Printf("| %s | %s | %s | %s | %s | %s |\n", x.Name, x.Ours, x.Chroma, x.Highlight, adv(x.OursV, x.ChromaV), adv(x.OursV, x.HighV))
	}
}

// adv: przewaga gohighlightgo, (odniesienie − gohighlightgo) / odniesienie × 100%.
func adv(ours, ref float64) string {
	if ref == 0 {
		return ""
	}
	return fmt.Sprintf("%.1f%%", (ref-ours)/ref*100)
}

func orDash(s string) string {
	if s == "" {
		return "not installed"
	}
	return s
}

func msf(ns float64) string {
	ms := ns / 1e6
	if ms < 10 {
		return fmt.Sprintf("%.2f ms", ms)
	}
	return fmt.Sprintf("%.0f ms", ms)
}

func mbf(b float64) string { return fmt.Sprintf("%.2f MB", b/1e6) }

func kb(b int64) string {
	s := fmt.Sprint((b + 500) / 1000)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + " " + s[i:]
	}
	return s
}

func size(p string) int64 {
	st, err := os.Stat(p)
	check(err)
	return st.Size()
}

func treeSize(dir string) int64 {
	var n int64
	filepath.WalkDir(dir, func(_ string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			if i, err := d.Info(); err == nil {
				n += i.Size()
			}
		}
		return nil
	})
	return n
}

func exists(p string) bool { _, err := os.Stat(p); return err == nil }

func machine() string {
	cpu := runtime.GOARCH
	if out, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output(); err == nil {
		cpu = strings.TrimSpace(string(out))
	} else if b, err := os.ReadFile("/proc/cpuinfo"); err == nil {
		for _, l := range strings.Split(string(b), "\n") {
			if strings.HasPrefix(l, "model name") {
				cpu = strings.TrimSpace(l[strings.Index(l, ":")+1:])
				break
			}
		}
	}
	return fmt.Sprintf("%s, %d cores, %s/%s", cpu, runtime.NumCPU(), runtime.GOOS, runtime.GOARCH)
}

func chromaVersion() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, d := range bi.Deps {
			if d.Path == "github.com/alecthomas/chroma/v2" {
				return d.Version
			}
		}
	}
	return "unknown"
}

func plist(p string) string {
	out, err := exec.Command("plutil", "-extract", "CFBundleShortVersionString", "raw", p).Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "rivalbench:", err)
		os.Exit(1)
	}
}
