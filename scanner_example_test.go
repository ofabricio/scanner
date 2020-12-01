package scanner

import (
	"fmt"
	"strings"
	"unicode"
)

func Example() {

	src := `
		fn main() {
			a = 1
			b = 'Hello, World!'
		}
	`

	s := NewScanner(strings.NewReader(src))

	for s.More() {
		s.While(unicode.IsSpace)

		if s.While(unicode.IsLetter) {
			fmt.Println(s.Text())
			continue
		}

		if s.While(unicode.IsNumber) {
			fmt.Println(s.Text())
			continue
		}

		if s.Exact("'") && s.Until(Any('\'')) && s.Exact("'") {
			fmt.Println(s.Join(3).Text)
			continue
		}

		if s.Match(unicode.IsPunct) || s.Match(unicode.IsSymbol) {
			fmt.Println(s.Text())
			continue
		}

		s.Next()
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

func ExampleScanner() {

	s := NewScanner(strings.NewReader("The quick fox"))

	for ; s.While(unicode.IsLetter); s.Exact(" ") {
		fmt.Println(s.Text())
	}

	// Output:
	// The
	// quick
	// fox
}
