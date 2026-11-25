package highlight

import (
	"slices"
	"strings"
)

// DefaultTheme is the name of the theme used when the caller has no preference.
const DefaultTheme = "light"

// style to wygląd jednej klasy w motywie: kolor tekstu, opcjonalnie tło (wiersze diff) i kursywa.
type style struct {
	color, background string
	italic            bool
}

// theme to motyw: kolor tekstu i tło bloku (klasa hl na elemencie, który owija kod) oraz wygląd każdej klasy.
// Kolory są dobrane w tym projekcie (docs/provenance.md: motywów innych narzędzi nie przenosimy); kontrast
// z tłem sprawdza TestThemeContrast.
type theme struct {
	color, background string
	classes           [len(classNames)]style
}

var themes = map[string]*theme{
	"light": {
		color: "#1f2328", background: "#f7f7f8",
		classes: [len(classNames)]style{
			keyword:   {color: "#8a1fa8"},
			typeName:  {color: "#006b73"},
			builtin:   {color: "#9a4a00"},
			str:       {color: "#1b6e20"},
			escape:    {color: "#b0005e"},
			number:    {color: "#0a57a8"},
			comment:   {color: "#636a73", italic: true},
			operator:  {color: "#8a3a12"},
			call:      {color: "#2443ad"},
			tag:       {color: "#9a1a3a"},
			attribute: {color: "#715400"},
			inserted:  {color: "#0f5c26", background: "#e3f3e6"},
			deleted:   {color: "#9c1419", background: "#fbe7e8"},
			meta:      {color: "#55599a"},
		},
	},
	"dark": {
		color: "#dcdde0", background: "#1d1f23",
		classes: [len(classNames)]style{
			keyword:   {color: "#d792f0"},
			typeName:  {color: "#5ccfd6"},
			builtin:   {color: "#f0a861"},
			str:       {color: "#93cf83"},
			escape:    {color: "#ff86bd"},
			number:    {color: "#82b9f5"},
			comment:   {color: "#8f969f", italic: true},
			operator:  {color: "#e6a585"},
			call:      {color: "#8eacff"},
			tag:       {color: "#ff8f9e"},
			attribute: {color: "#e5c66f"},
			inserted:  {color: "#97e0a0", background: "#1d3624"},
			deleted:   {color: "#ffa3a3", background: "#3f1f20"},
			meta:      {color: "#aab2ff"},
		},
	},
}

// Themes returns the names of the built-in themes, sorted, including "auto".
func Themes() []string {
	names := []string{"auto"}
	for name := range themes {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

// ThemeCSS returns a stylesheet that colours the classes written by HTML. Its rules apply inside an element
// with the class "hl" (for example <pre class="hl">), which also gets the theme's text colour and
// background. "light" and "dark" are fixed; "auto" is light, switching to dark when the reader's system
// prefers a dark appearance. ok is false for an unknown name.
func ThemeCSS(name string) (css string, ok bool) {
	if name == "auto" {
		return themeRules(themes["light"]) + "@media (prefers-color-scheme: dark) {\n" +
			themeRules(themes["dark"]) + "}\n", true
	}
	t := themes[name]
	if t == nil {
		return "", false
	}
	return themeRules(t), true
}

// themeRules zapisuje reguły motywu w stałej kolejności klas, więc wynik jest powtarzalny co do bajta.
func themeRules(t *theme) string {
	var b strings.Builder
	b.WriteString(".hl { color: " + t.color + "; background: " + t.background + "; }\n")
	for c := keyword; int(c) < len(classNames); c++ {
		s := t.classes[c]
		b.WriteString(".hl ." + classNames[c] + " { color: " + s.color + ";")
		if s.background != "" {
			b.WriteString(" background: " + s.background + ";")
		}
		if s.italic {
			b.WriteString(" font-style: italic;")
		}
		b.WriteString(" }\n")
	}
	return b.String()
}
