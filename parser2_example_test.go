package scanner_test

import (
	"fmt"
	"strings"

	"github.com/ofabricio/scanner"
)

func ExampleParser2() {

	src := `
root =
    entry+

entry =
    word '=' '\n' stmt '\n'

stmt =
	term+ or*

term =
	ident quantifier?

ident =
    word | string | '(' stmt ')'

or =
    '|' stmt

word =
    '[\\w_-]+'

quantifier =
    '*' | '+' | '?'

string =
	'\'(?:\\.|[^\'\n])*\''
	`

	s := scanner.NewScanner(strings.NewReader(src))
	s.Space("^[ \t\r]+")

	_ = Parse2(s)

	// Output:
	// root
	// =
	// NL
	// entry
	// +
	// NL
	// entry
	// =
	// NL
	// word
	// '='
	// '\n'
	// stmt
	// '\n'
	// NL
	// stmt
	// =
	// NL
	// term
	// +
	// or
	// *
	// NL
	// term
	// =
	// NL
	// ident
	// quantifier
	// ?
	// NL
	// ident
	// =
	// NL
	// word
	// |
	// string
	// |
	// '('
	// stmt
	// ')'
	// NL
	// or
	// =
	// NL
	// '|'
	// stmt
	// NL
	// word
	// =
	// NL
	// '[\\w_-]+'
	// NL
	// quantifier
	// =
	// NL
	// '*'
	// |
	// '+'
	// |
	// '?'
	// NL
	// string
	// =
	// NL
	// '\'(?:\\.|[^\'\n])*\''
	// NL
}

func Parse2(s *scanner.Scanner) *AST {
	parser := &Parser2{s: s}
	return parser.Parse()
}

type Parser2 struct {
	s   *scanner.Scanner
	ast *AST
}

func (t *Parser2) More() bool {
	return t.s.More()
}

func (t *Parser2) Match(s string) bool {
	if t.s.Match("^" + s) {
		txt := t.s.Text()
		if txt == "\n" {
			txt = "NL"
		}
		fmt.Println(txt)
		return true
	}
	return false
}

// entry = word '=' '\n' stmt '\n'
func (t *Parser2) Entry() bool {
	t.s.Match("^\\s+")
	return t.Word() && t.Match("=") && t.Match("\n") && t.Stmt() && t.Match("\n")
}

// stmt = term+ or*
func (t *Parser2) Stmt() bool {
	return t.OneToMany(t.Term) && t.ZeroToMany(t.Or)
}

// term = ident qualifier?
func (t *Parser2) Term() bool {
	return t.Ident() && (t.Quantifier() || true)
}

// ident = word | string | '(' or ')'
func (t *Parser2) Ident() bool {
	return t.Word() || t.String() || (t.Match("\\(") && t.Stmt() && t.Match("\\)"))
}

// or = '|' stmt
func (t *Parser2) Or() bool {
	return t.Match("\\|") && t.Stmt()
}

func (t *Parser2) Quantifier() bool {
	return t.Match("[+*?]")
}

func (t *Parser2) Word() bool {
	return t.Match("\\w+")
}

func (t *Parser2) String() bool {
	return t.Match(`'(?:\\.|[^'\n])*'`)
}

func (t *Parser2) ZeroToMany(fn func() bool) bool {
	return t.OneToMany(fn) || true
}

func (t *Parser2) OneToMany(fn func() bool) bool {
	count := 0
	for fn() {
		count++
	}
	return count > 0
}

func (t *Parser2) Parse() *AST {
	t.OneToMany(t.Entry)
	if t.More() {
		t.s.Match(".")
		fmt.Println("INVALID", t.s.Text())
	}
	return t.ast
}
