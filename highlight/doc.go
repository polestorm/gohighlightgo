// Package highlight turns source code into HTML with syntax highlighting.
//
// [HTML] escapes the code and wraps its tokens in <span class="hl-…"> elements: keywords (hl-kw), type
// names (hl-type), predeclared names (hl-builtin), strings (hl-str) and escapes in them (hl-esc), numbers
// (hl-num), comments (hl-com), operators (hl-op) and names before a parenthesis (hl-fn); hl-tag, hl-attr,
// hl-ins, hl-del and hl-meta are reserved for markup, diffs and preprocessor lines. Elements never nest
// and never contain a line break, so every output line can be handled on its own. The caller adds the
// <pre><code> wrapper; ThemeCSS gives a stylesheet for the classes (themes "light", the default, "dark"
// and "auto") that applies inside an element with the class "hl".
//
// Highlighting is lexical: one pass over the bytes, driven by a table per language, without regular
// expressions, in linear time also for hostile input. The output is deterministic and safe by construction:
// removing the tags and the three character references (&amp;, &lt;, &gt;) gives back the input.
//
// The package has no dependencies outside the standard library and does not know about any Markdown parser;
// a caller connects it to one through that parser's hook for code blocks.
package highlight
