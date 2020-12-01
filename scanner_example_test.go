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
		s.WhileCond(unicode.IsSpace)

		if s.WhileCond(unicode.IsLetter) {
			fmt.Println(s.Text())
			continue
		}

		if s.WhileCond(unicode.IsNumber) {
			fmt.Println(s.Text())
			continue
		}

		if s.Match("'") && s.Until("'") && s.Match("'") {
			fmt.Println(s.Join(3).Text)
			continue
		}

		if s.MatchCond(unicode.IsPunct) || s.MatchCond(unicode.IsSymbol) {
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

	for ; s.WhileCond(unicode.IsLetter); s.Match(" ") {
		fmt.Println(s.Text())
	}

	// Output:
	// The
	// quick
	// fox
}
