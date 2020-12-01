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
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestIsB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hi"))

	Equal(t, s.Is(Any('H')), true)

	Equal(t, s.Matched(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Is(Any('H')), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestMatch(t *testing.T) {

	s := NewScanner(strings.NewReader("H"))

	Equal(t, s.Match(Any('H')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestMatchB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hi"))

	Equal(t, s.Match(Any('H')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhile(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.While(unicode.IsLetter), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhileB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello World"))

	Equal(t, s.While(unicode.IsLetter), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)

	s.Match(Any(' '))

	Equal(t, s.While(unicode.IsLetter), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 7)
}

func TestUntil(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello!"))

	Equal(t, s.Until(unicode.IsPunct), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilStart(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until(Any('H')), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello! World."))

	Equal(t, s.Until(Any('!')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)

	Equal(t, s.Until(Any('.')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "! World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 6)
}

func TestUntilEnd(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until(Any('!')), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestExact(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Exact("Hello"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestExactB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello World"))

	Equal(t, s.Exact("Hello"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)

	Equal(t, s.Exact(" World"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), " World")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 6)
}

func TestFind(t *testing.T) {

	s := NewScanner(strings.NewReader("There are no uninteresting things, only uninterested people"))

	Equal(t, s.Find("things"), true)

	Equal(t, s.Matched(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "There are no uninteresting ")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestFindBegining(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Find("Hello"), false)

	Equal(t, s.Matched(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestString(t *testing.T) {

	s := NewScanner(strings.NewReader("'Hello'"))

	if s.Exact("'") && s.Until(Any('\'')) && s.Exact("'") {
	}

	token := s.Join(3)

	Equal(t, s.More(), false)
	Equal(t, token.Text, "'Hello'")
	Equal(t, token.Row, 1)
	Equal(t, token.Col, 1)
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
