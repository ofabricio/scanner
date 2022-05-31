# Scanner

A text scanner.

## Example

Here an example of a very simple JSON iterator using this scanner.
It parses a JSON and prints its keys and values.

```go
package main

import (
    "fmt"
    "unicode"
    . "github.com/ofabricio/scanner"
)

func main() {

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
```

## Operators

#### Equal

- [x] Equal(string) bool
- [x] EqualByte(byte) bool
- [x] EqualRune(rune) bool
- [x] EqualByteBy(func(byte) bool) bool
- [x] EqualRuneBy(func(rune) bool) bool

#### Match

- [x] Match(string) bool
- [x] MatchByte(byte) bool
- [x] MatchRune(rune) bool
- [x] MatchByteBy(func(byte) bool) bool
- [x] MatchRuneBy(func(rune) bool) bool

#### Until

- [x] MatchUntil(string) bool
- [x] MatchUntilByte(byte) bool
- [x] MatchUntilRune(rune) bool
- [x] MatchUntilByteBy(func(byte) bool) bool
- [x] MatchUntilRuneBy(func(rune) bool) bool
- [x] MatchUntilAny(a, b string) bool
- [x] MatchUntilAnyByte(a, b byte) bool
- [x] MatchUntilAnyRune(a, b rune) bool
- [x] MatchUntilAnyByte3(a, b, c byte) bool
- [x] MatchUntilEsc(v, esc string) bool
- [x] MatchUntilEscByte(v, esc byte) bool
- [x] MatchUntilEscRune(v, esc rune) bool

#### While

- [x] MatchWhileByteBy(func(byte) bool) bool
- [x] MatchWhileRuneBy(func(rune) bool) bool

#### Token

- [x] Token(int) string
- [x] TokenByteBy(func(byte) bool) string
- [x] TokenRuneBy(func(rune) bool) string
- [x] TokenFor(func() bool) string

#### Movement

- [x] Next()
- [x] NextRune()
- [x] Advance(int)
- [x] Mark() Scanner
- [x] Back(Scanner)
- [x] More() bool

#### Miscellaneous

- [x] Curr() byte
- [x] CurrRune() rune
- [x] String() string
- [x] Bytes() []byte

#### Utils

- [x] UtilMatchString(quote byte) bool
- [x] UtilMatchOpenCloseCount(o, c byte) bool
- [ ] UtilMatchInteger() bool
- [ ] UtilMatchFloat() bool
- [ ] UtilMatchNumber() bool
- [ ] UtilMatchHex() bool
