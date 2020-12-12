package scanner_test

import (
	"fmt"
	"strings"

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

		if s.Match(`^\w+`) {
			fmt.Println(s.Text())
			continue
		}

		if s.String("'") {
			fmt.Println(s.Text())
			continue
		}

		if s.Match(".") {
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

		if s.Match(`^\w+`) {
			fmt.Println(s.Text())
			continue
		}

		if s.Match(`^=`) {
			fmt.Println(s.Text())
			continue
		}

		m := s.Mark()

		if s.String("'") && !m.Left("'") {
			fmt.Println(m.Text())
			continue
		}

		s.Match(".+")
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

	for s.Match("\\w+") {
		fmt.Println(s.Text())
	}

	// Output:
	// The
	// quick
	// fox
}

func Example_string_escape() {

	src := `
		'Apple
		'Apple'
		'\'Apple\''
		'\'Apple'
		'Apple\''
		'Apple \'Grape\' Mango'
		'Apple \' Grape'
		''
		'\''
		'\'\''

		"Apple
		"Apple"
		"\"Grape\""
		"\"Grape"
		"Grape\""
		"Apple \"Grape\" Mango"
		"Apple \" Grape"
		""
		"\""
		"\"\""
	`

	s := scanner.NewScanner(strings.NewReader(src))

	for s.More() {
		m := s.Mark()
		if s.String("'") || s.String("\"") {
			fmt.Println(s.Text())
		} else if s.Match(".*") {
			fmt.Println("INVALID", m.Text())
		}
	}

	// Output:
	// INVALID 'Apple
	// 'Apple'
	// '\'Apple\''
	// '\'Apple'
	// 'Apple\''
	// 'Apple \'Grape\' Mango'
	// 'Apple \' Grape'
	// ''
	// '\''
	// '\'\''
	// INVALID "Apple
	// "Apple"
	// "\"Grape\""
	// "\"Grape"
	// "Grape\""
	// "Apple \"Grape\" Mango"
	// "Apple \" Grape"
	// ""
	// "\""
	// "\"\""
}
