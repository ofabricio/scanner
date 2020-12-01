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

// Match matches the current character with the condition
// and moves the cursor ahead when it matches.
// Returns true if the cursor moved.
func (t *Scanner) Match(cond MatcherFunc) bool {
	t.Mark()
	if t.More() && t.Is(cond) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// While moves the cursor forward while the condition matches.
// Returns true if the cursor moved.
func (t *Scanner) While(cond MatcherFunc) bool {
	t.Mark()
	for t.More() && t.Is(cond) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// Until moves the cursor forward until the condition matches.
//
// When trying to Until something at the cursor's position
// it does not match, ie, the cursor doesn't move.
//
// Returns true if the cursor moved.
func (t *Scanner) Until(cond MatcherFunc) bool {
	t.Mark()
	for t.More() && !t.Is(cond) {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// Find advances the cursor until the string matches.
//
// It is like Until() bur for string.
//
// When trying to find a string at the cursor's position
// it does not match, ie, the cursor does not move.
//
// Returns true if the cursor moved.
func (t *Scanner) Find(s string) bool {
	t.Mark()
	for i := bytes.Index(t.data[t.cursor.disp:], []byte(s)); i > 0; i-- {
		t.Next()
	}
	return t.Matched() && t.Save()
}

// Exact matches a string.
// Returns true if the cursor moved.
func (t *Scanner) Exact(s string) bool {
	t.Mark()
	if t.More() && bytes.HasPrefix(t.data[t.cursor.disp:], []byte(s)) {
		for i := len(s); i > 0; i-- {
			t.Next()
		}
	}
	return t.Matched() && t.Save()
}

// Is tests the current character at cursor's position.
// Is does not move the cursor.
func (t *Scanner) Is(cond MatcherFunc) bool {
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
