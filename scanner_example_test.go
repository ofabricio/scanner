package scanner_test

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ofabricio/scanner"
)

func Example() {

	src := `
		fn main() {
			a = 1
			b = 'Hello, World!'
		}
	`

	s := scanner.NewScanner(strings.NewReader(src))

	for s.More() {

		if s.WhileCond(unicode.IsSpace) {
			continue
		}

		if s.WhileCond(unicode.IsLetter) {
			fmt.Println(s.Text())
			continue
		}

		if s.WhileCond(unicode.IsNumber) {
			fmt.Println(s.Text())
			continue
		}

		m := s.Mark()
		if s.Match("'") && s.Until("'") && s.Match("'") {
			fmt.Println(m.Text())
			continue
		}

		if s.MatchCond(unicode.IsPunct) || s.MatchCond(unicode.IsSymbol) {
			fmt.Println(s.Text())
			continue
		}

	}

	// Output:
	// fn
	// main
	// (
	// )
	// {
	// a
	// =
	// 1
	// b
	// =
	// 'Hello, World!'
	// }
}

func Example_validating_strings() {

	src := `
		a = 'Hello World
		b = 'Hi''There'
	`

	s := scanner.NewScanner(strings.NewReader(src))

	for s.More() {

		if s.WhileCond(unicode.IsSpace) {
			continue
		}

		if s.WhileCond(unicode.IsLetter) || s.Match("=") {
			fmt.Println(s.Text())
			continue
		}

		m := s.Mark()

		if s.Match("'") && s.UntilCond(scanner.Any('\'', '\n')) && s.Match("'") && !m.Left("'") {
			fmt.Println(m.Text())
			continue
		}

		fmt.Println("INVALID", m.Text())
	}

	// Output:
	// a
	// =
	// INVALID 'Hello World
	// b
	// =
	// 'Hi'
	// INVALID 'There'
}

func ExampleScanner() {

	s := scanner.NewScanner(strings.NewReader("The quick fox"))

	for ; s.WhileCond(unicode.IsLetter); s.Match(" ") {
		fmt.Println(s.Text())
	}

	// Output:
	// The
	// quick
	// fox
}
