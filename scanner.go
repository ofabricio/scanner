package scanner

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"unicode/utf8"
)

// NewScanner creates a new scanner.
func NewScanner(r io.Reader) *Scanner {
	data, _ := ioutil.ReadAll(r)
	s := &Scanner{data: data, cursor: cursor{Row: 1}}
	s.Next()
	s.Mark()
	s.Space("^\\s+")
	return s
}

// String matches a string like 'Hi' if r is "'".
// This default implementation does not scan multiline strings.
func (t *Scanner) String(r string) bool {
	r = fmt.Sprintf(`^%s(?:\\.|[^%s\n])*%s`, r, r, r)
	return t.Match(r)
}

// Match matches the given regex and
// advances the cursor on match.
func (t *Scanner) Match(regex string) bool {
	return t.Regex(regexp.MustCompile(regex))
}

// Regex matches the given regex and
// advances the cursor on match.
func (t *Scanner) Regex(re *regexp.Regexp) bool {
	t.Mark()
	return t.regex(re)
}

// Mark marks the begining of a token.
func (t *Scanner) Mark() Mark {
	if t.space != nil {
		t.Skip(t.space)
	}
	t.mark.cursor = t.cursor
	t.mark.scan = t
	return t.mark
}

// Space sets what should be skipped.
func (t *Scanner) Space(s string) {
	t.space = regexp.MustCompile(s)
}

// Skip skips a regex. It moves the cursor without Mark().
func (t *Scanner) Skip(re *regexp.Regexp) bool {
	return t.regex(re)
}

func (t *Scanner) regex(re *regexp.Regexp) bool {
	if loc := re.FindIndex(t.data[t.cursor.disp:]); loc != nil {
		count := utf8.RuneCount(t.data[:loc[1]])
		t.next(count)
	}
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

	space *regexp.Regexp
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
