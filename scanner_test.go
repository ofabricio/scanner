package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// #region Equal

func TestScannerEqual(t *testing.T) {
	tt := []struct {
		give string
		when string
		then bool
	}{
		{give: ``, when: ``, then: true},
		{give: `a`, when: ``, then: true},
		{give: `a`, when: `a`, then: true},
		{give: `a`, when: `b`, then: false},
		{give: `a`, when: `ab`, then: false},
		{give: ``, when: `a`, then: false},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.Equal(tc.when), tc)
	}
}

func BenchmarkScannerEqual(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.Equal("a")
	}
}

func TestScannerEqualByte(t *testing.T) {
	tt := []struct {
		give string
		when byte
		then bool
	}{
		{give: `a`, when: 'a', then: true},
		{give: `a`, when: 'b', then: false},
		{give: ``, when: 'b', then: false},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.EqualByte(tc.when), tc)
	}
}

func BenchmarkScannerEqualByte(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.EqualByte('a')
	}
}

func TestScannerEqualRune(t *testing.T) {
	tt := []struct {
		give string
		when rune
		then bool
	}{
		{give: `世`, when: '世', then: true},
		{give: `界`, when: '世', then: false},
		{give: ``, when: 'b', then: false},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.EqualRune(tc.when), tc)
	}
}

func BenchmarkScannerEqualRune(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.EqualRune('a')
	}
}

func TestScannerEqualByteBy(t *testing.T) {
	f := func(v byte) bool {
		return v == 'a'
	}
	tt := []struct {
		give string
		when func(byte) bool
		then bool
	}{
		{give: `a`, when: f, then: true},
		{give: `b`, when: f, then: false},
		{give: ``, when: f, then: false},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.EqualByteBy(tc.when), tc)
	}
}

func BenchmarkScannerEqualByteBy(b *testing.B) {
	f := func(v byte) bool {
		return v == 'a'
	}
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.EqualByteBy(f)
	}
}

func TestScannerEqualRuneBy(t *testing.T) {
	f := func(v rune) bool {
		return v == '世'
	}
	tt := []struct {
		give string
		when func(rune) bool
		then bool
	}{
		{give: `世`, when: f, then: true},
		{give: `b`, when: f, then: false},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.EqualRuneBy(tc.when), tc)
	}
}

func BenchmarkScannerEqualRuneBy(b *testing.B) {
	f := func(v rune) bool {
		return v == 'a'
	}
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.EqualRuneBy(f)
	}
}

// #endregion Equal

// #region Match

func TestScannerMatch(t *testing.T) {
	tt := []struct {
		give string
		when string
		then bool
		exp  string
	}{
		{give: `abc`, when: "abc", then: true, exp: "abc"},
		{give: `aaa`, when: "abc", then: false, exp: ""},
		{give: `a`, when: "abc", then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.Match(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatch(b *testing.B) {
	x := Scanner(`abc`)
	for i := 0; i < b.N; i++ {
		s := x
		s.Match("abc")
	}
}

func TestScannerMatchByte(t *testing.T) {
	tt := []struct {
		give string
		when byte
		then bool
		exp  string
	}{
		{give: `a`, when: 'a', then: true, exp: "a"},
		{give: `a`, when: 'b', then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchByte(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchByte(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchByte('a')
	}
}

func TestScannerMatchRune(t *testing.T) {
	tt := []struct {
		give string
		when rune
		then bool
		exp  string
	}{
		{give: `a`, when: 'a', then: true, exp: "a"},
		{give: `b`, when: 'a', then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchRune(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchRune(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchRune('a')
	}
}

func TestScannerMatchByteBy(t *testing.T) {
	f := func(v byte) bool {
		return v == 'a'
	}
	tt := []struct {
		give string
		when func(byte) bool
		then bool
		exp  string
	}{
		{give: `a`, when: f, then: true, exp: "a"},
		{give: `b`, when: f, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchByteBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchByteBy(b *testing.B) {
	f := func(v byte) bool {
		return v == 'a'
	}
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchByteBy(f)
	}
}

func TestScannerMatchRuneBy(t *testing.T) {
	f := func(v rune) bool {
		return v == '世'
	}
	tt := []struct {
		give string
		when func(rune) bool
		then bool
		exp  string
	}{
		{give: `世`, when: f, then: true, exp: "世"},
		{give: `b`, when: f, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchRuneBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchRuneBy(b *testing.B) {
	f := func(v rune) bool {
		return v == 'a'
	}
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchRuneBy(f)
	}
}

// #endregion Match

// #region Until

func TestScannerMatchUntil(t *testing.T) {
	tt := []struct {
		give string
		when string
		then bool
		exp  string
	}{
		{give: `abc.`, when: ".", then: true, exp: "abc"},
		{give: `a.b..cd...`, when: "...", then: true, exp: "a.b..cd"},
		{give: `abc?`, when: ".", then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntil(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntil(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntil(".")
	}
}

func TestScannerMatchUntilByte(t *testing.T) {
	tt := []struct {
		give string
		when byte
		then bool
		exp  string
	}{
		{give: `abc.`, when: '.', then: true, exp: "abc"},
		{give: `abc?`, when: '.', then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilByte(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilByte(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilByte('.')
	}
}

func TestScannerMatchUntilRune(t *testing.T) {
	tt := []struct {
		give string
		when rune
		then bool
		exp  string
	}{
		{give: `abc.`, when: '.', then: true, exp: "abc"},
		{give: `abc?`, when: '.', then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilRune(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilRune(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilRune('.')
	}
}

func TestScannerMatchUntilByteBy(t *testing.T) {
	f := func(v byte) bool {
		return v == '.'
	}
	tt := []struct {
		give string
		when func(byte) bool
		then bool
		exp  string
	}{
		{give: `abc.`, when: f, then: true, exp: "abc"},
		{give: `abc?`, when: f, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilByteBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilByteBy(b *testing.B) {
	f := func(v byte) bool {
		return v == '.'
	}
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilByteBy(f)
	}
}

func TestScannerMatchUntilRuneBy(t *testing.T) {
	f := func(v rune) bool {
		return v == '.'
	}
	tt := []struct {
		give string
		when func(rune) bool
		then bool
		exp  string
	}{
		{give: `abc.`, when: f, then: true, exp: "abc"},
		{give: `abc?`, when: f, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilRuneBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilRuneBy(b *testing.B) {
	f := func(v rune) bool {
		return v == '.'
	}
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilRuneBy(f)
	}
}

func TestScannerMatchUntilAny(t *testing.T) {
	tt := []struct {
		give string
		when []string
		then bool
		exp  string
	}{
		{give: `abc.`, when: []string{".", ","}, then: true, exp: "abc"},
		{give: `abc,`, when: []string{".", ","}, then: true, exp: "abc"},
		{give: `abc?`, when: []string{".", ","}, then: false, exp: ""},
		{give: `xxxAAxxxAAA`, when: []string{"AAA", "BBB"}, then: true, exp: "xxxAAxxx"},
		{give: `xxxBBxxxBBB`, when: []string{"AAA", "BBB"}, then: true, exp: "xxxBBxxx"},
		{give: `xxxAAxxxBBB`, when: []string{"AAA", "BBB"}, then: true, exp: "xxxAAxxx"},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilAny(tc.when[0], tc.when[1]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilAny(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilAny(".", ",")
	}
}

func TestScannerMatchUntilAnyByte(t *testing.T) {
	tt := []struct {
		give string
		when []byte
		then bool
		exp  string
	}{
		{give: `abc.`, when: []byte{'.', ','}, then: true, exp: "abc"},
		{give: `abc,`, when: []byte{'.', ','}, then: true, exp: "abc"},
		{give: `abc?`, when: []byte{'.', ','}, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilAnyByte(tc.when[0], tc.when[1]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilAnyByte(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilAnyByte('.', ',')
	}
}

func TestScannerMatchUntilAnyByte3(t *testing.T) {
	tt := []struct {
		give string
		when []byte
		then bool
		exp  string
	}{
		{give: `abc.`, when: []byte{'.', ',', ';'}, then: true, exp: "abc"},
		{give: `abc,`, when: []byte{'.', ',', ';'}, then: true, exp: "abc"},
		{give: `abc;`, when: []byte{'.', ',', ';'}, then: true, exp: "abc"},
		{give: `abc?`, when: []byte{'.', ',', ';'}, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilAnyByte3(tc.when[0], tc.when[1], tc.when[2]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilAnyByte3(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilAnyByte3('.', ',', ';')
	}
}

func TestScannerMatchUntilAnyRune(t *testing.T) {
	tt := []struct {
		give string
		when []rune
		then bool
		exp  string
	}{
		{give: `abc.`, when: []rune{'.', ','}, then: true, exp: "abc"},
		{give: `abc,`, when: []rune{'.', ','}, then: true, exp: "abc"},
		{give: `abc?`, when: []rune{'.', ','}, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilAnyRune(tc.when[0], tc.when[1]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilAnyRune(b *testing.B) {
	x := Scanner(`abc.`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilAnyRune('.', ',')
	}
}

func TestScannerMatchUntilEsc(t *testing.T) {
	tt := []struct {
		give string
		when []string
		then bool
		exp  string
	}{
		{give: `abc"`, when: []string{`"`, `\`}, then: true, exp: `abc`},
		{give: `a\"bc"`, when: []string{`"`, `\`}, then: true, exp: `a\"bc`},
		{give: `abc?`, when: []string{`"`, `\`}, then: false, exp: ""},
		{give: `xxxAAABBBxxxBBB`, when: []string{`BBB`, `AAA`}, then: true, exp: "xxxAAABBBxxx"},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilEsc(tc.when[0], tc.when[1]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilEsc(b *testing.B) {
	x := Scanner(`abc"`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilEsc(`"`, `\`)
	}
}

func TestScannerMatchUntilEscByte(t *testing.T) {
	tt := []struct {
		give string
		when []byte
		then bool
		exp  string
	}{
		{give: `abc"`, when: []byte{'"', '\\'}, then: true, exp: `abc`},
		{give: `a\"bc"`, when: []byte{'"', '\\'}, then: true, exp: `a\"bc`},
		{give: `abc?`, when: []byte{'"', '\\'}, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilEscByte(tc.when[0], tc.when[1]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchUntilEscByte(b *testing.B) {
	x := Scanner(`abc"`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilEscByte('"', '\\')
	}
}

func TestScannerMatchUntilEscRune(t *testing.T) {
	tt := []struct {
		give string
		when []rune
		then bool
		exp  string
	}{
		{give: `abc"`, when: []rune{'"', '\\'}, then: true, exp: `abc`},
		{give: `a\"bc"`, when: []rune{'"', '\\'}, then: true, exp: `a\"bc`},
		{give: `abc?`, when: []rune{'"', '\\'}, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchUntilEscRune(tc.when[0], tc.when[1]), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarScannerkMatchUntilEscRune(b *testing.B) {
	x := Scanner(`abc"`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchUntilEscRune('"', '\\')
	}
}

// #endregion Until

// #region While

func TestScannerMatchWhileByteBy(t *testing.T) {
	f := func(v byte) bool {
		return v == 'a'
	}
	tt := []struct {
		give string
		when func(byte) bool
		then bool
		exp  string
	}{
		{give: `aaa`, when: f, then: true, exp: "aaa"},
		{give: `aaab`, when: f, then: true, exp: "aaa"},
		{give: `bbb`, when: f, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchWhileByteBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchWhileByteBy(b *testing.B) {
	f := func(v byte) bool {
		return v == 'a'
	}
	x := Scanner(`aaab`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchWhileByteBy(f)
	}
}

func TestScannerMatchWhileRuneBy(t *testing.T) {
	f := func(v rune) bool {
		return v == 'a'
	}
	tt := []struct {
		give string
		when func(rune) bool
		then bool
		exp  string
	}{
		{give: `aaa`, when: f, then: true, exp: "aaa"},
		{give: `aaab`, when: f, then: true, exp: "aaa"},
		{give: `bbb`, when: f, then: false, exp: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.MatchWhileRuneBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerMatchWhileRuneBy(b *testing.B) {
	f := func(v rune) bool {
		return v == 'a'
	}
	x := Scanner(`aaab`)
	for i := 0; i < b.N; i++ {
		s := x
		s.MatchWhileRuneBy(f)
	}
}

// #endregion While

// #region Token

func TestScannerToken(t *testing.T) {
	tt := []struct {
		give string
		then string
	}{
		{give: `abc`, then: "abc"},
		{give: ``, then: ""},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		s.Advance(len(s))
		assert.Equal(t, tc.then, s.Token(m), tc)
	}
}

func BenchmarkScannerToken(b *testing.B) {
	ini := Scanner(`abc`)
	end := ini
	end.Advance(3)
	for i := 0; i < b.N; i++ {
		_ = end.Token(ini)
	}
}

func TestScannerTokenByteBy(t *testing.T) {
	f := func(v byte) bool {
		return v == 'a'
	}
	tt := []struct {
		give string
		when func(byte) bool
		then string
		exp  string
	}{
		{give: `aaa`, when: f, then: "aaa", exp: ""},
		{give: `bbb`, when: f, then: "", exp: "bbb"},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.TokenByteBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.String(), tc)
	}
}

func BenchmarkScannerTokenByteBy(b *testing.B) {
	f := func(v byte) bool {
		return v == 'a'
	}
	x := Scanner(`aaa`)
	for i := 0; i < b.N; i++ {
		s := x
		_ = s.TokenByteBy(f)
	}
}

func TestScannerTokenRuneBy(t *testing.T) {
	f := func(v rune) bool {
		return v == 'a'
	}
	tt := []struct {
		give string
		when func(rune) bool
		then string
		exp  string
	}{
		{give: `aaa`, when: f, then: "aaa", exp: ""},
		{give: `bbb`, when: f, then: "", exp: "bbb"},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.TokenRuneBy(tc.when), tc)
		assert.Equal(t, tc.exp, s.String(), tc)
	}
}

func BenchmarkScannerTokenRuneBy(b *testing.B) {
	f := func(v rune) bool {
		return v == 'a'
	}
	x := Scanner(`aaa`)
	for i := 0; i < b.N; i++ {
		s := x
		_ = s.TokenRuneBy(f)
	}
}

func TestScannerTokenFor(t *testing.T) {
	tt := []struct {
		give string
		then string
		exp  string
	}{
		{give: `aaa`, then: "a", exp: "aa"},
		{give: `bbb`, then: "", exp: "bbb"},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		f := func() bool {
			return s.MatchByte('a')
		}
		assert.Equal(t, tc.then, s.TokenFor(f), tc)
		assert.Equal(t, tc.exp, s.String(), tc)
	}
}

func BenchmarkScannerTokenFor(b *testing.B) {
	x := Scanner(`aaa`)
	for i := 0; i < b.N; i++ {
		s := x
		_ = s.TokenFor(func() bool {
			return s.MatchByte('a')
		})
	}
}

func TestScannerTokenWith(t *testing.T) {
	f := func(s *Scanner) bool {
		return s.MatchByte('a')
	}
	tt := []struct {
		give string
		then string
		exp  string
	}{
		{give: `aaa`, then: "a", exp: "aa"},
		{give: `bbb`, then: "", exp: "bbb"},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.TokenWith(f), tc)
		assert.Equal(t, tc.exp, s.String(), tc)
	}
}

func BenchmarkScannerTokenWith(b *testing.B) {
	f := func(s *Scanner) bool {
		return s.MatchByte('a')
	}
	x := Scanner(`aaa`)
	for i := 0; i < b.N; i++ {
		s := x
		_ = s.TokenWith(f)
	}
}

// #endregion Token

// #region Movement

func TestScannerNext(t *testing.T) {
	s := Scanner(`abc`)
	s.Next()
	assert.Equal(t, `bc`, s.String())
	s.Next()
	assert.Equal(t, `c`, s.String())
	s.Next()
	assert.Equal(t, ``, s.String())
}

func BenchmarkScannerNext(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.Next()
	}
}

func TestScannerNextRune(t *testing.T) {
	s := Scanner(`世界`)
	s.NextRune()
	assert.Equal(t, `界`, s.String())
	s.NextRune()
	assert.Equal(t, ``, s.String())
}

func BenchmarkScannerNextRune(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.NextRune()
	}
}

func TestScannerAdvance(t *testing.T) {
	s := Scanner(`abc`)
	s.Advance(1)
	assert.Equal(t, `bc`, s.String())
	s.Advance(2)
	assert.Equal(t, ``, s.String())
}

func BenchmarkScannerAdvance(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.Advance(1)
	}
}

func TestScannerMark(t *testing.T) {
	s := Scanner(`abc`)
	m := s.Mark()
	s.Next()
	assert.Equal(t, `bc`, s.String())
	assert.Equal(t, `abc`, m.String())
}

func BenchmarkScannerMark(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.Mark()
	}
}

func TestScannerBack(t *testing.T) {
	s := Scanner(`abc`)
	m := s.Mark()
	s.Next()
	s.Back(m)
	assert.Equal(t, `abc`, s.String())
}

func BenchmarkScannerBack(b *testing.B) {
	x := Scanner(`a`)
	s := x.Mark()
	for i := 0; i < b.N; i++ {
		x.Back(s)
	}
}

func TestScannerMore(t *testing.T) {
	tt := []struct {
		give string
		then bool
	}{
		{give: `a`, then: true},
		{give: ``, then: false},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		assert.Equal(t, tc.then, s.More(), tc)
	}
}

func BenchmarkScannerMore(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.More()
	}
}

// #endregion Movement

// #region Miscellaneous

func TestScannerCurr(t *testing.T) {
	s := Scanner(`abc`)
	assert.Equal(t, byte('a'), s.Curr())
	s.Next()
	assert.Equal(t, byte('b'), s.Curr())
}

func BenchmarkScannerCurr(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.Curr()
	}
}

func TestScannerCurrRune(t *testing.T) {
	s := Scanner(`世界`)
	assert.Equal(t, '世', s.CurrRune())
	s.Advance(len("世"))
	assert.Equal(t, '界', s.CurrRune())
}

func BenchmarkScannerCurrRune(b *testing.B) {
	x := Scanner(`a`)
	for i := 0; i < b.N; i++ {
		s := x
		s.CurrRune()
	}
}

func TestScannerString(t *testing.T) {
	s := Scanner(`abc`)
	assert.Equal(t, `abc`, s.String())
}

func BenchmarkScannerString(b *testing.B) {
	x := Scanner(`abc`)
	for i := 0; i < b.N; i++ {
		s := x
		_ = s.String()
	}
}

func TestScannerBytes(t *testing.T) {
	s := Scanner(`abc`)
	assert.Equal(t, []byte{'a', 'b', 'c'}, s.Bytes())
}

func BenchmarkScannerBytes(b *testing.B) {
	x := Scanner(`abc`)
	for i := 0; i < b.N; i++ {
		s := x
		_ = s.Bytes()
	}
}

// #endregion Miscellaneous

// #region Utils

func TestScannerUtilMatchString(t *testing.T) {
	tt := []struct {
		give string
		when byte
		then bool
		exp  string
	}{
		{give: `""`, when: '"', then: true, exp: `""`},
		{give: `"a"`, when: '"', then: true, exp: `"a"`},
		{give: `"ab"`, when: '"', then: true, exp: `"ab"`},
		{give: `"abc"`, when: '"', then: true, exp: `"abc"`},
		{give: `"ab\"cd"`, when: '"', then: true, exp: `"ab\"cd"`},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.UtilMatchString(tc.when), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerUtilMatchString(b *testing.B) {
	x := Scanner(`"abc"`)
	for i := 0; i < b.N; i++ {
		s := x
		s.UtilMatchString('"')
	}
}

func TestScannerUtilMatchOpenCloseCount(t *testing.T) {
	tt := []struct {
		give string
		when byte
		then bool
		exp  string
	}{
		{give: `{}`, then: true, exp: `{}`},
		{give: `{{}}`, then: true, exp: `{{}}`},
		{give: `{{}{}{{}}}`, then: true, exp: `{{}{}{{}}}`},
		{give: `{}{}`, then: true, exp: `{}`},
		{give: `{`, then: false, exp: ``},
		{give: `}`, then: false, exp: ``},
		{give: `}{`, then: false, exp: ``},
	}
	for _, tc := range tt {
		s := Scanner(tc.give)
		m := s.Mark()
		assert.Equal(t, tc.then, s.UtilMatchOpenCloseCount('{', '}'), tc)
		assert.Equal(t, tc.exp, s.Token(m), tc)
	}
}

func BenchmarkScannerUtilMatchOpenCloseCount(b *testing.B) {
	x := Scanner(`{{}}`)
	for i := 0; i < b.N; i++ {
		s := x
		s.UtilMatchOpenCloseCount('{', '}')
	}
}

// #endregion Utils
