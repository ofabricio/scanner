package scanner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unicode"
)

func TestIs(t *testing.T) {

	s := NewScanner(strings.NewReader("H"))

	Equal(t, s.Is(Any('H')), true)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestIsB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hi"))

	Equal(t, s.Is(Any('H')), true)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestIsEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Is(Any('H')), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestMatch(t *testing.T) {

	s := NewScanner(strings.NewReader("H"))

	Equal(t, s.Match(Any('H')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestMatchB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hi"))

	Equal(t, s.Match(Any('H')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestMatchEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Match(Any('H')), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestWhile(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.While(unicode.IsLetter), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestWhileB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello World"))

	Equal(t, s.While(unicode.IsLetter), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)

	s.Match(Any(' '))

	Equal(t, s.While(unicode.IsLetter), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 7)
	Equal(t, s.More(), false)
}

func TestWhilEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.While(unicode.IsLetter), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestUntil(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello!"))

	Equal(t, s.Until(unicode.IsPunct), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestUntilStart(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until(Any('H')), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestUntilB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello! World."))

	Equal(t, s.Until(Any('!')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)

	Equal(t, s.Until(Any('.')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "! World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 6)
	Equal(t, s.More(), true)
}

func TestUntilEnd(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until(Any('!')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestUntilEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Until(unicode.IsPunct), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestExact(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Exact("Hello"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestExactB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello World"))

	Equal(t, s.Exact("Hello"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)

	Equal(t, s.Exact(" World"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), " World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 6)
	Equal(t, s.More(), false)
}

func TestExactEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Exact("Hello"), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestFind(t *testing.T) {

	s := NewScanner(strings.NewReader("There are no uninteresting things, only uninterested people"))

	Equal(t, s.Find("things"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "There are no uninteresting ")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestFindBegining(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Find("Hello"), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestFindEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Find("Hello"), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestString(t *testing.T) {

	s := NewScanner(strings.NewReader("'Hello'"))

	if s.Exact("'") && s.Until(Any('\'')) && s.Exact("'") {
	}

	token := s.Join(3)

	Equal(t, token.Text, "'Hello'")
	Equal(t, token.Row, 1)
	Equal(t, token.Col, 1)
	Equal(t, s.More(), false)
}

func TestMatcherAny(t *testing.T) {

	s := NewScanner(strings.NewReader("banana"))

	s.While(Any('b', 'a', 'n'))

	Equal(t, s.Text(), "banana")
}

func Equal(t *testing.T, got, exp interface{}) {
	if !reflect.DeepEqual(got, exp) {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("\n[error] %s:%d\nExp:\n%v\nGot:\n%v\n", fn, line, exp, got)
	}
}
