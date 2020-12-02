package scanner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unicode"
)

func TestEqual(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Equal("Hello"), true)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestEqualB(t *testing.T) {

	s := NewScanner(strings.NewReader("World"))

	Equal(t, s.Equal("Hello"), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Equal("Hello"), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestEqualCond(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.EqualCond(unicode.IsLetter), true)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestEqualCondB(t *testing.T) {

	s := NewScanner(strings.NewReader("12345"))

	Equal(t, s.EqualCond(unicode.IsLetter), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestMatch(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello World"))

	Equal(t, s.Match("Hello"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestMatchB(t *testing.T) {

	s := NewScanner(strings.NewReader("World Hello"))

	Equal(t, s.Match("Hello"), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestMatchCond(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.MatchCond(unicode.IsLetter), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestMatchCondB(t *testing.T) {

	s := NewScanner(strings.NewReader("1234"))

	Equal(t, s.MatchCond(unicode.IsLetter), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntil(t *testing.T) {

	s := NewScanner(strings.NewReader("There are no uninteresting things, only uninterested people"))

	Equal(t, s.Until("uninterested"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "There are no uninteresting things, only ")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilBegining(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until("Hello"), true)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilEnd(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until("World"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilCond(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello!"))

	Equal(t, s.UntilCond(unicode.IsPunct), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilCondStart(t *testing.T) {

	s := NewScanner(strings.NewReader("123Hello"))

	Equal(t, s.UntilCond(unicode.IsDigit), true)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilCondB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello,World "))

	Equal(t, s.UntilCond(unicode.IsPunct), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)

	Equal(t, s.UntilCond(unicode.IsSpace), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), ",World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 6)
}

func TestUntilCondEnd(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.UntilCond(unicode.IsPunct), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhile(t *testing.T) {

	s := NewScanner(strings.NewReader("NanNanNanNan Batman!"))

	Equal(t, s.While("Nan"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "NanNanNanNan")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhileB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.While("World"), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhileCond(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.WhileCond(unicode.IsLetter), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhileCondB(t *testing.T) {

	s := NewScanner(strings.NewReader("12345 Hello"))

	Equal(t, s.WhileCond(unicode.IsLetter), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestString(t *testing.T) {

	s := NewScanner(strings.NewReader("'Hello'"))

	m := s.Mark()
	if s.Match("'") && s.Until("'") && s.Match("'") {
	}

	Equal(t, s.More(), false)
	Equal(t, m.Text(), "'Hello'")
	Equal(t, m.Row, 1)
	Equal(t, m.Col, 1)
}

func TestStringEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader("''"))

	m := s.Mark()
	if s.Match("'") && s.Until("'") && s.Match("'") {
	}

	Equal(t, s.More(), false)
	Equal(t, m.Text(), "''")
	Equal(t, m.Row, 1)
	Equal(t, m.Col, 1)
}

func TestMarkLeft(t *testing.T) {

	s := NewScanner(strings.NewReader("He-Man"))

	s.Match("He-")
	m := s.Mark()
	s.Match("Man")
	Equal(t, m.Left("-"), true)
}

func TestMarkLeftB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	m := s.Mark()
	Equal(t, m.Left("a"), false)
}

func TestMatcherAny(t *testing.T) {

	s := NewScanner(strings.NewReader("banana"))

	s.WhileCond(Any('b', 'a', 'n'))

	Equal(t, s.Text(), "banana")
}

func Equal(t *testing.T, got, exp interface{}) {
	if !reflect.DeepEqual(got, exp) {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("\n[error] %s:%d\nExp:\n%v\nGot:\n%v\n", fn, line, exp, got)
	}
}
