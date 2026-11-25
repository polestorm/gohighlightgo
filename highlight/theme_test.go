package highlight

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// REQ: API-006
func TestThemes(t *testing.T) {
	names := Themes()
	if strings.Join(names, " ") != "auto dark light" {
		t.Errorf("Themes() = %v", names)
	}
	if _, ok := ThemeCSS(DefaultTheme); !ok {
		t.Errorf("default theme %q is missing", DefaultTheme)
	}
	if _, ok := ThemeCSS("solarized"); ok {
		t.Error("unknown theme accepted")
	}
	for _, name := range names {
		css, _ := ThemeCSS(name)
		again, _ := ThemeCSS(name)
		if css != again {
			t.Errorf("%s: stylesheet differs between calls", name)
		}
		// Każda klasa z listy OUT-002 ma regułę; sprawdzian korzysta z niezależnej listy allowedClasses.
		for class := range allowedClasses {
			if !strings.Contains(css, ".hl ."+class+" {") {
				t.Errorf("%s: no rule for %s", name, class)
			}
		}
		if strings.ContainsAny(css, "<>") {
			t.Errorf("%s: stylesheet contains < or >", name)
		}
	}
	auto, _ := ThemeCSS("auto")
	light, _ := ThemeCSS("light")
	dark, _ := ThemeCSS("dark")
	if !strings.HasPrefix(auto, light) || !strings.Contains(auto, "@media (prefers-color-scheme: dark) {\n"+dark+"}") {
		t.Error("auto is not light followed by dark in a media query")
	}
}

// REQ: API-006
// TestThemeContrast: kolor każdej klasy ma z tłem (motywu albo własnym tłem klasy) kontrast co najmniej 4,5 : 1,
// próg WCAG 2.2 dla zwykłego tekstu (kryterium 1.4.3).
func TestThemeContrast(t *testing.T) {
	for name, th := range themes {
		if r := contrast(th.color, th.background); r < 4.5 {
			t.Errorf("%s: text %.2f : 1", name, r)
		}
		for c := keyword; int(c) < len(classNames); c++ {
			s := th.classes[c]
			bg := th.background
			if s.background != "" {
				bg = s.background
			}
			if s.color == "" {
				t.Errorf("%s: %s has no colour", name, classNames[c])
				continue
			}
			if r := contrast(s.color, bg); r < 4.5 {
				t.Errorf("%s: %s %s on %s: %.2f : 1", name, classNames[c], s.color, bg, r)
			}
		}
	}
}

// contrast liczy stosunek kontrastu dwóch kolorów #rrggbb według definicji WCAG 2.2 (względna luminancja).
func contrast(a, b string) float64 {
	la, lb := luminance(a), luminance(b)
	if la < lb {
		la, lb = lb, la
	}
	return (la + 0.05) / (lb + 0.05)
}

func luminance(hex string) float64 {
	channel := func(i int) float64 {
		v, err := strconv.ParseUint(hex[1+2*i:3+2*i], 16, 8)
		if err != nil {
			panic(fmt.Sprintf("bad colour %q", hex))
		}
		c := float64(v) / 255
		if c <= 0.04045 {
			return c / 12.92
		}
		return math.Pow((c+0.055)/1.055, 2.4)
	}
	return 0.2126*channel(0) + 0.7152*channel(1) + 0.0722*channel(2)
}
