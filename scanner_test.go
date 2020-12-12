package scanner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestNewLines(t *testing.T) {

	s := NewScanner(strings.NewReader("\nHello"))

	Equal(t, s.Match("Hello"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 2)
	Equal(t, s.Col(), 1)
}

func TestUntil(t *testing.T) {

	s := NewScanner(strings.NewReader("There are no uninteresting things, only uninterested people"))

	Equal(t, s.Match("uninterested"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "There are no uninteresting things, only uninterested")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilBegining(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Match("Hello"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilEnd(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello"))

	Equal(t, s.Match("World"), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilCond(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello!"))

	Equal(t, s.Match("!"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "Hello!")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilCondStart(t *testing.T) {

	s := NewScanner(strings.NewReader("123Hello"))

	Equal(t, s.Match("\\d+"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "123")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestUntilCondB(t *testing.T) {

	s := NewScanner(strings.NewReader("Hello,World "))

	Equal(t, s.Match(","), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "Hello,")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)

	Equal(t, s.Match(" "), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "World ")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 7)
}

func TestWhile(t *testing.T) {

	s := NewScanner(strings.NewReader("NanNanNanNan Batman!"))

	Equal(t, s.Match("(Nan)+"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "NanNanNanNan")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestWhileB(t *testing.T) {

	s := NewScanner(strings.NewReader("aaa\nbbb\nccc"))

	Equal(t, s.Match("aaa"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "aaa")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)

	Equal(t, s.Match("bbb"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "bbb")
	Equal(t, s.Row(), 2)
	Equal(t, s.Col(), 1)

	Equal(t, s.Match("ccc"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "ccc")
	Equal(t, s.Row(), 3)
	Equal(t, s.Col(), 1)
}

func TestString(t *testing.T) {

	s := NewScanner(strings.NewReader("'Hello'"))

	Equal(t, s.String("'"), true)

	Equal(t, s.More(), false)
	Equal(t, s.Text(), "'Hello'")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestStringEmpty(t *testing.T) {

	s := NewScanner(strings.NewReader("''"))

	Equal(t, s.String("'"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), "''")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestStringEscape(t *testing.T) {

	s := NewScanner(strings.NewReader(`'\'Hello\''`))

	Equal(t, s.String("'"), true)

	Equal(t, s.Moved(), true)
	Equal(t, s.More(), false)
	Equal(t, s.Text(), `'\'Hello\''`)
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
}

func TestStringInvalid(t *testing.T) {

	s := NewScanner(strings.NewReader(`'Hello`))

	Equal(t, s.String("'"), false)

	Equal(t, s.Moved(), false)
	Equal(t, s.More(), true)
	Equal(t, s.Text(), "")
	Equal(t, s.Row(), 1)
	Equal(t, s.Col(), 1)
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

func Equal(t *testing.T, got, exp interface{}) {
	if !reflect.DeepEqual(got, exp) {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("\n[error] %s:%d\nExp:\n%v\nGot:\n%v\n", fn, line, exp, got)
	}
}
