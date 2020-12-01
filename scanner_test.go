package scanner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unicode"
)

func TestScannerIs(t *testing.T) {

	s := NewScanner(strings.NewReader("H"))

	Equal(t, s.Is(Any('H')), true)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestScannerIsB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hi"))

	Equal(t, s.Is(Any('H')), true)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestScannerIsEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Is(Any('H')), false)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerMatch(t *testing.T) {

	s := NewScanner(strings.NewReader("H"))

	Equal(t, s.Match(Any('H')), true)
	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerMatchB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hi"))

	Equal(t, s.Match(Any('H')), true)
	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "H")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestScannerMatchEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Match(Any('H')), false)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerWhile(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.While(unicode.IsLetter), true)
	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerWhileB(t *testing.T) {

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

func TestScannerWhilEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.While(unicode.IsLetter), false)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerUntil(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello!"))

	Equal(t, s.Until(unicode.IsPunct), true)
	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), true)
}

func TestScannerUntilB(t *testing.T) {

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

func TestScannerUntilEnd(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Until(Any('!')), true)
	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerUntilEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Until(unicode.IsPunct), false)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerExact(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Exact("Hello"), true)
	Equal(t, s.Matched(), true)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerExactB(t *testing.T) {

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

func TestScannerExactEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader(""))

	Equal(t, s.Exact("Hello"), false)
	Equal(t, s.Matched(), false)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
	Equal(t, s.More(), false)
}

func TestScannerString(t *testing.T) {

	s := NewScanner(strings.NewReader("'Hello'"))

	if s.Exact("'") && s.Until(Any('\'')) && s.Exact("'") {
	}

	token := s.Join(3)

	Equal(t, token.Text, "'Hello'")
	Equal(t, token.Row, 1)
	Equal(t, token.Col, 1)
	Equal(t, s.More(), false)
}

func Equal(t *testing.T, got, exp interface{}) {
	if !reflect.DeepEqual(got, exp) {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("\n[error] %s:%d\nExp:\n%v\nGot:\n%v\n", fn, line, exp, got)
	}
}
