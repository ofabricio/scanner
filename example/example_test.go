package example

import (
	"fmt"
	"unicode"

	. "github.com/ofabricio/scanner"
)

func Example() {

	j := Json{`{ "one": "hello", "two": "world" }`}

	j.Iterate(func(k, v string) {
		fmt.Println(k, v)
	})

	// Output:
	// "one" "hello"
	// "two" "world"
}

type Json struct {
	Scanner
}

func (j *Json) Iterate(f func(k, v string)) {
	if j.Match("{") {
		for j.WS() && !j.Match("}") {
			k, _ := j.TokenFor(j.MatchString), j.Match(":")
			j.WS()
			v, _ := j.TokenFor(j.MatchString), j.Match(",")
			f(k, v)
		}
	}
}

func (j *Json) WS() bool {
	return j.MatchWhileRuneBy(unicode.IsSpace) || true
}

func (j *Json) MatchString() bool {
	return j.Scanner.UtilMatchString('"')
}
