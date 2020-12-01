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
	s := &Scanner{data: data, cursor: Mark{row: 1}}
	s.Next()
	s.Mark()
	return s
}

// Until advances the cursor until the string matches.
// When trying to match a string at the cursor's position
// it does not match, ie, the cursor does not move.
// Returns true if the cursor moved.
func (t *Scanner) Until(s string) bool {
	t.Mark()
	for t.More() && !t.Equal(s) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// UntilCond advances the cursor until the condition matches.
// UntilCond is like Until bur for a custom condition.
// When trying to UntilCond something at the cursor's position
// it does not match, ie, the cursor doesn't move.
// Returns true if the cursor moved.
func (t *Scanner) UntilCond(cond MatcherFunc) bool {
	t.Mark()
	for t.More() && !t.EqualCond(cond) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// While advances the cursor while the string matches.
// Returns true if the cursor moved.
func (t *Scanner) While(s string) bool {
	t.Mark()
	for t.More() && t.Equal(s) {
		for range s {
			t.Next()
		}
	}
	return t.Matched() && t.Save()
}

// WhileCond advances the cursor while the condition matches.
// Returns true if the cursor moved.
func (t *Scanner) WhileCond(cond MatcherFunc) bool {
	t.Mark()
	for t.More() && t.EqualCond(cond) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// Match advances the cursor if the string matches.
// Returns true if the cursor moved.
func (t *Scanner) Match(s string) bool {
	t.Mark()
	if t.More() && t.Equal(s) {
		for range s {
			t.Next()
		}
	}
	return t.Matched() && t.Save()
}

// MatchCond advances the cursor if the condition matches.
// Returns true if the cursor moved.
func (t *Scanner) MatchCond(cond MatcherFunc) bool {
	t.Mark()
	if t.More() && t.EqualCond(cond) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// Equal tests if a string matches.
// Equal does not move the cursor.
func (t *Scanner) Equal(s string) bool {
	return bytes.HasPrefix(t.data[t.cursor.disp:], []byte(s))
}

// EqualCond tests the current character at cursor's position.
// EqualCond does not move the cursor.
func (t *Scanner) EqualCond(cond MatcherFunc) bool {
	return cond(t.char)
}

// Next moves the cursor to the next position.
func (t *Scanner) Next() {
	if t.char == '\n' {
		t.cursor.row++
		t.cursor.col = 0
	}
	t.cursor.col++
	t.cursor.disp += t.size
	r, s := utf8.DecodeRune(t.data[t.cursor.disp:])
	t.char = r
	t.size = s
}

// Mark marks the begining of a token.
func (t *Scanner) Mark() {
	t.mark = t.cursor
}

// More tells if the scanner still has data ahead of the cursor.
func (t *Scanner) More() bool {
	return t.cursor.disp < len(t.data)
}

// Matched reports whether there is a match or not.
func (t *Scanner) Matched() bool {
	return t.cursor.disp-t.mark.disp > 0
}

// Save saves the current mark.
func (t *Scanner) Save() bool {
	t.marks = append(t.marks, t.mark)
	return true
}

// Join joins the last count tokens.
func (t *Scanner) Join(count int) Token {
	m := t.marks[len(t.marks)-count]
	txt := string(t.data[m.disp:t.cursor.disp])
	return Token{Text: txt, Row: m.row, Col: m.col}
}

// Text returns the current matched token.
func (t *Scanner) Text() string {
	return string(t.data[t.mark.disp:t.cursor.disp])
}

// Col returns the current cursor column.
func (t *Scanner) Col() int {
	return t.mark.col
}

// Row returns the current cursor line.
func (t *Scanner) Row() int {
	return t.cursor.row
}

// Any tests if any character matches.
func Any(r ...rune) MatcherFunc {
	return func(ru rune) bool {
		for _, v := range r {
			if ru == v {
				return true
			}
		}
		return false
	}
}

// Scanner is a scanner.
type Scanner struct {
	data []byte

	char rune // Current character.
	size int  // Current character size.

	cursor Mark

	mark  Mark
	marks []Mark
}

// Token is a token.
type Token struct {
	Row  int
	Col  int
	Text string
}

// MatcherFunc is a matcher function.
type MatcherFunc func(rune) bool

// Mark is a mark in the scanner.
// Each time a token matches a mark is set at its begining.
type Mark struct {
	disp int // Displacement/Offset/Cursor/CurrentIndex.
	row  int // Mark line.
	col  int // Mark column.
}
