package scanner_test

import (
	"fmt"
	"strings"

	"github.com/ofabricio/scanner"
)

func ExampleParser() {

	src := `
root =
    entry+

entry =
    word '=' '\n' stmt '\n'

stmt =
	term+ or*

term =
	ident qualifier?

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
	s.Match("^\\s+")

	var tokens []Token

	for s.More() {

		if s.Match("^\n+") {
			tokens = append(tokens, Token{"NL", "NL"})
			continue
		}

		if s.Match("^\\w+") {
			tokens = append(tokens, Token{s.Text(), "WORD"})
			continue
		}

		if s.String("'") {
			tokens = append(tokens, Token{s.Text(), "STRING"})
			continue
		}

		if s.Match("^[=(|)+*?]") {
			tokens = append(tokens, Token{s.Text(), "SYMBOL"})
			continue
		}

		if s.Match(".") {
			tokens = append(tokens, Token{s.Text(), "INVALID"})
			continue
		}
	}

	_ = Parse(tokens)

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
	// qualifier
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

func Parse(tokens []Token) *AST {
	parser := &Parser{tokens: tokens}
	return parser.Parse()
}

type Token struct {
	Text string
	Type string
}

type AST struct {
	Token    Token
	children []*AST
}

type Parser struct {
	tokens []Token
	cur    Token
	idx    int
	ast    *AST
}

func (t *Parser) Next() {
	if t.More() {
		t.cur = t.tokens[t.idx]
		t.idx++
	}
}

func (t *Parser) More() bool {
	return t.idx < len(t.tokens)
}

func (t *Parser) MatchType(typ string) bool {
	return t.match(t.cur.Type, typ)
}

func (t *Parser) Match(txt string) bool {
	return t.match(t.cur.Text, txt)
}

func (t *Parser) match(a, b string) bool {
	if a == b {
		fmt.Println(t.cur.Text)
		t.Next()
		return true
	}
	return false
}

// entry = word '=' '\n' stmt '\n'
func (t *Parser) Entry() bool {
	return t.Word() && t.Match("=") && t.MatchType("NL") && t.Stmt() && t.MatchType("NL")
}

// stmt = term+ or*
func (t *Parser) Stmt() bool {
	return t.OneToMany(t.Term) && t.ZeroToMany(t.Or)
}

// term = ident qualifier?
func (t *Parser) Term() bool {
	return t.Ident() && (t.Quantifier() || true)
}

// ident = word | string | '(' or ')'
func (t *Parser) Ident() bool {
	return t.Word() || t.String() || (t.Match("(") && t.Stmt() && t.Match(")"))
}

// or = '|' stmt
func (t *Parser) Or() bool {
	return t.Match("|") && t.Stmt()
}

func (t *Parser) Quantifier() bool {
	return t.Match("+") || t.Match("*") || t.Match("?")
}

func (t *Parser) Word() bool {
	return t.MatchType("WORD")
}

func (t *Parser) String() bool {
	return t.MatchType("STRING")
}

func (t *Parser) ZeroToMany(fn func() bool) bool {
	return t.OneToMany(fn) || true
}

func (t *Parser) OneToMany(fn func() bool) bool {
	count := 0
	for fn() {
		count++
	}
	return count > 0
}

func (t *Parser) Parse() *AST {
	t.Next()
	t.OneToMany(t.Entry)
	if t.More() {
		fmt.Println("INVALID", t.cur.Text)
	}
	return t.ast
}
