package scanner

import (
	"io"
	"io/ioutil"
	"unicode/utf8"
)

// NewScanner creates a new scanner.
func NewScanner(r io.Reader) *Scanner {
	data, _ := ioutil.ReadAll(r)
	s := &Scanner{data: data, row: 1, col: 0}
	s.Next()
	s.Mark()
	return s
}

// Next returns the current character and
// moves the cursor to the next position.
func (t *Scanner) Next() {
	if t.char == '\n' {
		t.row++
		t.col = 0
	}
	t.disp += t.size
	// if t.disp < len(t.data) {
	r, s := utf8.DecodeRune(t.data[t.disp:])
	t.char = r
	t.size = s
	t.col++
	// }
}

// While moves the cursor forward while the condition matches.
// Returns true if the cursor moved.
func (t *Scanner) While(cond MatchFunc) bool {
	t.Mark()
	for t.Match(cond) {
	}
	return t.Matched()
}

// Until moves the cursor forward until the condition matches.
// Returns true if the cursor moved.
func (t *Scanner) Until(cond MatchFunc) bool {
	return t.While(Not(cond))
}

// Is matches the param with the current character.
func (t *Scanner) Is(cond MatchFunc) bool {
	return cond(t.Curr())
}

// Match matches the current character with the condition
// and moves the cursor ahead when it matches.
func (t *Scanner) Match(cond MatchFunc) bool {
	if t.More() && t.Is(cond) {
		t.Next()
		return true
	}
	return false
}

// Exact matches an exact match.
func (t *Scanner) Exact(v string) bool {
	t.Mark()
	for _, r := range v {
		if !t.Match(Any(r)) {
			break
		}
	}
	return t.Len() == len(v)
}

// Mark marks the begining of a token.
func (t *Scanner) Mark() {
	t.mark = t.disp
	t.markCol = t.col
}

// More tells if the scanner still has data ahead of the cursor.
func (t *Scanner) More() bool {
	return t.disp < len(t.data)
}

// Matched reports whether there is a match or not.
func (t *Scanner) Matched() bool {
	return t.Len() > 0
}

// Len reports the token length.
func (t *Scanner) Len() int {
	return t.disp - t.mark
}

// Text returns the current token.
func (t *Scanner) Text() string {
	m := t.mark
	return string(t.data[m:t.disp])
}

// Curr returns the current character.
func (t *Scanner) Curr() rune {
	return t.char
}

func (t *Scanner) Col() int {
	return t.markCol
}

func (t *Scanner) Row() int {
	return t.row
}

// Any is a match function that tests if any character matches.
func Any(r ...rune) MatchFunc {
	return func(ru rune) bool {
		for _, v := range r {
			if ru == v {
				return true
			}
		}
		return false
	}
}

// Not negates a condition.
func Not(cond MatchFunc) MatchFunc {
	return func(r rune) bool {
		return !cond(r)
	}
}

// Scanner is a scanner.
type Scanner struct {
	data []byte
	disp int // Displacement/Offset/Cursor/CurrentIndex.

	char rune // Current character.
	size int  // Current character size.

	mark    int // Marks the start of a token.
	markCol int // Save the cursor column at the moment of a mark.

	row int // Cursor line.
	col int // Cursor column.
}

// Token is a token.
type Token struct {
	Text string
	Type string
	Row  int
	Col  int
}

// MatchFunc is a match function.
type MatchFunc func(rune) bool
