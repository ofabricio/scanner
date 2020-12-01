package scanner

import (
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

// Exact matches an exact match.
func (t *Scanner) Exact(s string) bool {
	return t.While(Exact(s))
}

// Until moves the cursor forward until the condition matches.
// Returns true if the cursor moved.
func (t *Scanner) Until(cond MatcherFunc) bool {
	return t.While(Not(cond))
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

// Match matches the current character with the condition
// and moves the cursor ahead when it matches.
func (t *Scanner) Match(cond MatcherFunc) bool {
	return t.While(Once(cond))
}

// Next moves the cursor to the next character.
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

// Is matches the param with the current character.
func (t *Scanner) Is(cond MatcherFunc) bool {
	return cond(t.char)
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

// Text returns the current token.
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

// Not negates a condition.
func Not(cond MatcherFunc) MatcherFunc {
	return func(r rune) bool {
		return !cond(r)
	}
}

// Exact tests for an exact match of a string.
func Exact(str string) MatcherFunc {
	i := 0
	return func(r rune) bool {
		if i == len(str) {
			return false
		}
		c, s := utf8.DecodeRuneInString(str[i:])
		i += s
		return r == c
	}
}

// Once matches only one time.
// Useful with While() to force a single match. This is a
// trick to not repeat the same code in Match() and While().
// Don't use with Until() since Not(Once) means always.
func Once(cond MatcherFunc) MatcherFunc {
	once := false
	return func(r rune) bool {
		once = !once
		return once && cond(r)
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
