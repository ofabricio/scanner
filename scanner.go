package scanner

import (
	"bytes"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

// NewScanner creates a new scanner.
func NewScanner(r io.Reader) *Scanner {
	data, _ := ioutil.ReadAll(r)
	s := &Scanner{data: data, cursor: cursor{Row: 1}}
	s.Next()
	s.Mark()
	return s
}

// Until advances the cursor until the string matches.
// Until always return true. To know if Until moved
// the cursor check for Moved() after Until.
func (t *Scanner) Until(s ...string) bool {
	any := func() bool {
		for _, ss := range s {
			if t.Equal(ss) {
				return true
			}
		}
		return false
	}
	t.Mark()
	for t.More() && !any() {
		t.Next()
	}
	return t.Moved() || true
}

// UntilCond is like Until.
func (t *Scanner) UntilCond(cond MatcherFunc) bool {
	t.Mark()
	for t.More() && !t.EqualCond(cond) {
		t.Next()
	}
	return t.Moved() || true
}

// While advances the cursor while the string matches.
// Returns true if the cursor moved.
func (t *Scanner) While(s string) bool {
	t.Mark()
	for t.More() && t.Equal(s) {
		t.next(len(s))
	}
	return t.Moved()
}

// WhileCond is like While.
func (t *Scanner) WhileCond(cond MatcherFunc) bool {
	t.Mark()
	for t.More() && t.EqualCond(cond) {
		t.Next()
	}
	return t.Moved()
}

// Match advances the cursor if the string matches.
// Returns true if the cursor moved.
func (t *Scanner) Match(s string) bool {
	t.Mark()
	if t.More() && t.Equal(s) {
		t.next(len(s))
	}
	return t.Moved()
}

// MatchCond is like Match.
func (t *Scanner) MatchCond(cond MatcherFunc) bool {
	t.Mark()
	if t.More() && t.EqualCond(cond) {
		t.Next()
	}
	return t.Moved()
}

// Equal tests if a string matches.
// Equal does not move the cursor.
func (t *Scanner) Equal(s string) bool {
	return bytes.HasPrefix(t.data[t.cursor.disp:], []byte(s))
}

// EqualCond is like Equal.
func (t *Scanner) EqualCond(cond MatcherFunc) bool {
	return cond(t.char)
}

// Scan runs a scan function.
// You can use this to provide your own scanner implementation.
func (t *Scanner) Scan(fn ScanFunc) bool {
	t.Mark()
	fn(t)
	return t.Moved()
}

// next moves the cursor by n runes.
func (t *Scanner) next(n int) {
	for n > 0 {
		n--
		t.Next()
	}
}

// Next moves the cursor to the next position.
func (t *Scanner) Next() {
	if t.char == '\n' {
		t.cursor.Row++
		t.cursor.Col = 0
	}
	t.cursor.Col++
	t.cursor.disp += t.size
	r, s := utf8.DecodeRune(t.data[t.cursor.disp:])
	t.char = r
	t.size = s
}

// Mark marks the begining of a token.
func (t *Scanner) Mark() Mark {
	t.mark.cursor = t.cursor
	t.mark.scan = t
	return t.mark
}

// More tells if the scanner still has data ahead of the cursor.
func (t *Scanner) More() bool {
	return t.cursor.disp < len(t.data)
}

// Moved reports whether there is a match or not.
func (t *Scanner) Moved() bool {
	return t.cursor.disp > t.mark.disp
}

// Left tests if the left side of a marker matches a string.
func (t *Scanner) Left(s string) bool {
	return t.mark.Left(s)
}

// Text returns the current matched token.
func (t *Scanner) Text() string {
	return t.mark.Text()
}

// Col returns the column of the current match.
func (t *Scanner) Col() int {
	return t.mark.Col
}

// Row returns the line number of the current match.
func (t *Scanner) Row() int {
	return t.mark.Row
}

// Scanner is a scanner.
type Scanner struct {
	data []byte

	char rune // Current character.
	size int  // Current character size.

	cursor cursor

	mark Mark
}

// Mark is a mark in the scanner.
// Each time a token matches a mark is set at its begining.
type Mark struct {
	cursor
	scan *Scanner
}

type cursor struct {
	disp int // Displacement/Offset/Cursor/CurrentIndex.
	Row  int // Cursor line.
	Col  int // Cursor column.
}

// Text returns a token starting from a mark.
func (t Mark) Text() string {
	return string(t.scan.data[t.disp:t.scan.cursor.disp])
}

// Left tests if the left side of a marker matches a string.
func (t Mark) Left(s string) bool {
	return bytes.HasSuffix(t.scan.data[:t.disp], []byte(s))
}

// MatcherFunc is a matcher function.
type MatcherFunc func(rune) bool

// ScanFunc is a scan function useful for customizing the scanner.
type ScanFunc func(*Scanner)

func (t *Scanner) String(r string) bool {
	m := t.Mark()
	if t.Match(r) && t.UntilEsc(r) && t.Match(r) {
		t.mark = m
		return true
	}
	return false
}

func (t *Scanner) UntilEsc(r string) bool {
	m := t.Mark()
	for t.Until(r, "\n") && t.Mark().Left("\\") && t.Match(r) {
	}
	t.mark = m
	return t.Moved() || true
}
