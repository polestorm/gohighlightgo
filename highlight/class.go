package highlight

// class to rodzaj tokenu; none oznacza tekst bez elementu <span>.
type class uint8

const (
	none class = iota
	keyword
	typeName
	builtin
	str
	escape
	number
	comment
	operator
	call
	tag
	attribute
	inserted
	deleted
	meta
)

// classNames to stała lista klas z ADR 0002 (OUT-002), w kolejności stałych powyżej.
var classNames = [...]string{
	none:      "",
	keyword:   "hl-kw",
	typeName:  "hl-type",
	builtin:   "hl-builtin",
	str:       "hl-str",
	escape:    "hl-esc",
	number:    "hl-num",
	comment:   "hl-com",
	operator:  "hl-op",
	call:      "hl-fn",
	tag:       "hl-tag",
	attribute: "hl-attr",
	inserted:  "hl-ins",
	deleted:   "hl-del",
	meta:      "hl-meta",
}
