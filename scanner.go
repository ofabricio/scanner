package scanner

import (
	"unicode/utf8"
	"unsafe"
)

type Scanner string

// #region Equal

// Equal tests the current token given a string.
func (s Scanner) Equal(v string) bool {
	if len(s) < len(v) {
		return false
	}
	i := 0
	for ; i < len(v); i++ {
		if s[i] != v[i] {
			break
		}
	}
	return i == len(v)
}

// EqualByte tests the current token given a byte.
func (s Scanner) EqualByte(v byte) bool {
	return s.Curr() == v
}

// EqualRune tests the current token given a rune.
func (s Scanner) EqualRune(v rune) bool {
	return s.CurrRune() == v
}

// EqualByteBy tests the current token given a byte function.
func (s Scanner) EqualByteBy(f func(byte) bool) bool {
	return f(s.Curr())
}

// EqualRuneBy tests the current token given a rune function.
func (s Scanner) EqualRuneBy(f func(rune) bool) bool {
	return f(s.CurrRune())
}

// EqualByteRange tests the current byte given a byte range.
func (s Scanner) EqualByteRange(a, b byte) bool {
	c := s.Curr()
	return c >= a && c <= b
}

// #endregion Equal

// #region Match

// Match matches a token given a string.
func (s *Scanner) Match(v string) bool {
	if s.Equal(v) {
		*s = (*s)[len(v):]
		return true
	}
	return false
}

// MatchByte matches a token given a byte.
func (s *Scanner) MatchByte(v byte) bool {
	if s.Curr() == v {
		*s = (*s)[1:]
		return true
	}
	return false
}

// MatchRune matches a token given a rune.
func (s *Scanner) MatchRune(v rune) bool {
	if r, size := utf8.DecodeRuneInString(s.String()); r == v {
		*s = (*s)[size:]
		return true
	}
	return false
}

// MatchByteBy matches a token given a byte function.
func (s *Scanner) MatchByteBy(f func(byte) bool) bool {
	if f(s.Curr()) {
		*s = (*s)[1:]
		return true
	}
	return false
}

// MatchRuneBy matches a token given a rune function.
func (s *Scanner) MatchRuneBy(f func(rune) bool) bool {
	if r, size := utf8.DecodeRuneInString(s.String()); f(r) {
		*s = (*s)[size:]
		return true
	}
	return false
}

// #endregion Match

// #region Until

// MatchUntil matches until v matches.
func (s *Scanner) MatchUntil(v string) bool {
	ss := *s
	for a, b := 0, 0; a < len(ss); a++ {
		if ss[a] == v[b] {
			b++
			if b == len(v) {
				*s = ss[a-b+1:]
				return true
			}
			continue
		}
		b = 0
	}
	return false
}

// MatchUntilByte matches until v matches.
func (s *Scanner) MatchUntilByte(v byte) bool {
	ss := *s
	for i := 0; i < len(ss); i++ {
		if ss[i] == v {
			*s = ss[i:]
			return true
		}
	}
	return false
}

// MatchUntilRune matches until v matches.
func (s *Scanner) MatchUntilRune(v rune) bool {
	ss := *s
	for i := 0; i < len(ss); {
		r, size := utf8.DecodeRuneInString(ss[i:].String())
		if r == v {
			*s = ss[i:]
			return true
		}
		i += size
	}
	return false
}

// MatchUntilByteBy matches until f matches.
func (s *Scanner) MatchUntilByteBy(f func(byte) bool) bool {
	ss := *s
	for i := 0; i < len(ss); i++ {
		if f(ss[i]) {
			*s = ss[i:]
			return true
		}
	}
	return false
}

// MatchUntilRuneBy matches until f matches.
func (s *Scanner) MatchUntilRuneBy(f func(rune) bool) bool {
	ss := *s
	for i := 0; i < len(ss); {
		r, size := utf8.DecodeRuneInString(ss[i:].String())
		if f(r) {
			*s = ss[i:]
			return true
		}
		i += size
	}
	return false
}

// MatchUntilAny matches until either a or b matches.
func (s *Scanner) MatchUntilAny(a, b string) bool {
	ss := *s
	for si, ai, bi := 0, 0, 0; si < len(ss); si++ {
		if ss[si] == a[ai] {
			ai++
			if ai == len(a) {
				*s = ss[si-ai+1:]
				return true
			}
			continue
		}
		if ss[si] == b[bi] {
			bi++
			if bi == len(b) {
				*s = ss[si-bi+1:]
				return true
			}
			continue
		}
		ai, bi = 0, 0
	}
	return false
}

// MatchUntilAnyByte matches until either a or b matches.
func (s *Scanner) MatchUntilAnyByte(a, b byte) bool {
	ss := *s
	for i := 0; i < len(ss); i++ {
		if ss[i] == a || ss[i] == b {
			*s = ss[i:]
			return true
		}
	}
	return false
}

// MatchUntilAnyByte3 matches until either a, b or c matches.
func (s *Scanner) MatchUntilAnyByte3(a, b, c byte) bool {
	ss := *s
	for i := 0; i < len(ss); i++ {
		if ss[i] == a || ss[i] == b || ss[i] == c {
			*s = ss[i:]
			return true
		}
	}
	// If last param is 0 the caller wants
	// whatever matched until EOF.
	if c == 0 {
		*s = ss[len(ss):]
		return true
	}
	return false
}

// MatchUntilAnyByte4 matches until any argument matches.
func (s *Scanner) MatchUntilAnyByte4(a, b, c, d byte) bool {
	ss := *s
	for i := 0; i < len(ss); i++ {
		if ss[i] == a || ss[i] == b || ss[i] == c || ss[i] == d {
			*s = ss[i:]
			return true
		}
	}
	return false
}

// MatchUntilAnyByte5 matches until any argument matches.
func (s *Scanner) MatchUntilAnyByte5(a, b, c, d, e byte) bool {
	ss := *s
	for i := 0; i < len(ss); i++ {
		if ss[i] == a || ss[i] == b || ss[i] == c || ss[i] == d || ss[i] == e {
			*s = ss[i:]
			return true
		}
	}
	// If last param is 0 the caller wants
	// whatever matched until EOF.
	if e == 0 {
		*s = ss[len(ss):]
		return true
	}
	return false
}

// MatchUntilAnyRune matches until either a or b matches.
func (s *Scanner) MatchUntilAnyRune(a, b rune) bool {
	ss := *s
	for i := 0; i < len(ss); {
		r, size := utf8.DecodeRuneInString(ss[i:].String())
		if r == a || r == b {
			*s = ss[i:]
			return true
		}
		i += size
	}
	return false
}

// MatchUntilEsc matches until v matches and
// escapes v if esc matches.
func (s *Scanner) MatchUntilEsc(v, esc string) bool {
	ss := *s
	for ss.More() {
		if ss.Match(esc) && ss.Match(v) {
			continue
		}
		if ss.Equal(v) {
			*s = ss
			return true
		}
		ss.Advance(1)
	}
	return false
}

// MatchUntilEscByte matches until v matches and
// escapes v if esc matches.
func (s *Scanner) MatchUntilEscByte(v, esc byte) bool {
	ss := *s
	var prev byte
	for i := 0; i < len(ss); i++ {
		if ss[i] == v && prev != esc {
			*s = ss[i:]
			return true
		}
		prev = ss[i]
	}
	return false
}

// MatchUntilEscRune matches until v matches and
// escapes v if esc matches.
func (s *Scanner) MatchUntilEscRune(v, esc rune) bool {
	ss := *s
	var prev rune
	for i := 0; i < len(ss); {
		r, size := utf8.DecodeRuneInString(ss[i:].String())
		if r == v && prev != esc {
			*s = ss[i:]
			return true
		}
		i += size
		prev = r
	}
	return false
}

// #endregion Until

// #region While

// MatchWhileAnyByte4 matches while any argument matches.
func (s *Scanner) MatchWhileAnyByte4(a, b, c, d byte) bool {
	i := 0
	ss := *s
	for ; i < len(ss); i++ {
		if ss[i] != a && ss[i] != b && ss[i] != c && ss[i] != d {
			break
		}
	}
	if i > 0 {
		*s = ss[i:]
		return true
	}
	return false
}

// MatchWhileByteLTE matches while the current byte is less than or equal to a.
func (s *Scanner) MatchWhileByteLTE(a byte) bool {
	i := 0
	ss := *s
	for i < len(ss) && ss[i] <= a {
		i++
	}
	// Had a match?
	if i > 0 {
		*s = ss[i:]
		return true
	}
	return false
}

// MatchWhileByteBy matches while f matches.
func (s *Scanner) MatchWhileByteBy(f func(byte) bool) bool {
	i := 0
	ss := *s
	for ; i < len(ss); i++ {
		if !f(ss[i]) {
			break
		}
	}
	if i > 0 {
		*s = ss[i:]
		return true
	}
	return false
}

// MatchWhileRuneBy matches while f matches.
func (s *Scanner) MatchWhileRuneBy(f func(rune) bool) bool {
	i := 0
	ss := *s
	for i < len(ss) {
		r, size := utf8.DecodeRuneInString(ss[i:].String())
		if !f(r) {
			break
		}
		i += size
	}
	if i > 0 {
		*s = ss[i:]
		return true
	}
	return false
}

// #endregion While

// #region Token

// Token returns a token given a start position.
func (end Scanner) Token(ini Scanner) string {
	return ini[:len(ini)-len(end)].String()
}

// TokenByteBy returns a token given a byte function.
func (s *Scanner) TokenByteBy(f func(byte) bool) string {
	m := *s
	s.MatchWhileByteBy(f)
	m = m[:len(m)-len(*s)]
	return *(*string)(unsafe.Pointer(&m))
}

// TokenRuneBy returns a token given a rune function.
func (s *Scanner) TokenRuneBy(f func(rune) bool) string {
	m := *s
	s.MatchWhileRuneBy(f)
	m = m[:len(m)-len(*s)]
	return *(*string)(unsafe.Pointer(&m))
}

// TokenFor returns a token given a match function.
func (s *Scanner) TokenFor(f func() bool) string {
	m := *s
	f()
	return m[:len(m)-len(*s)].String()
}

func (s *Scanner) TokenWith(f func(*Scanner) bool) string {
	m := *s
	f(s)
	return m[:len(m)-len(*s)].String()
}

// #endregion Token

// #region Movement

// Next moves to the next byte.
func (s *Scanner) Next() {
	*s = (*s)[1:]
}

// Next moves to the next rune.
func (s *Scanner) NextRune() {
	ss := *s
	_, size := utf8.DecodeRuneInString(*(*string)(unsafe.Pointer(&ss)))
	*s = ss[size:]
}

// Advance advances n bytes.
func (s *Scanner) Advance(n int) {
	*s = (*s)[n:]
}

// Mark returns a mark.
func (s Scanner) Mark() Scanner {
	return s
}

// Back sets the scanner back to a mark.
func (s *Scanner) Back(m Scanner) {
	*s = m
}

// More tells if there are more bytes to scan.
func (s Scanner) More() bool {
	return len(s) > 0
}

// #endregion Movement

// #region Miscellaneous

// Curr returns the current byte.
func (s Scanner) Curr() byte {
	if len(s) == 0 {
		return 0
	}
	return s[0]
}

// CurrRune returns the current rune.
func (s Scanner) CurrRune() rune {
	r, _ := utf8.DecodeRuneInString(s.String())
	return r
}

// String returns the scanner text.
func (s Scanner) String() string {
	return *(*string)(unsafe.Pointer(&s))
}

// Bytes returns the scanner bytes.
func (s Scanner) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// #endregion Miscellaneous

// #region Util

// UtilMatchString matches a string given a quote.
func (s *Scanner) UtilMatchString(quote byte) bool {
	if ss := *s; len(ss) > 1 && ss.EqualByte(quote) {
		for i := 1; i < len(ss); i++ {
			if ss[i] == quote && ss[i-1] != '\\' {
				*s = ss[i+1:]
				return true
			}
		}
	}
	return false
}

// UtilMatchOpenCloseCount matches open and close by counting them.
// Also skips strings by quote.
func (s *Scanner) UtilMatchOpenCloseCount(open, clos, quote byte) bool {
	if ss := *s; len(ss) > 0 && ss[0] == open {
		c := 0
		for i := 0; i < len(ss); i++ {
			if ss[i] == open {
				c++
				continue
			}
			if ss[i] == clos {
				if c--; c == 0 {
					*s = ss[i+1:]
					break
				}
			}
			// Skip string.
			if ss[i] == quote {
				for i = i + 1; i < len(ss); i++ {
					if ss[i] == quote && ss[i-1] != '\\' {
						break
					}
				}
				continue
			}

		}
		return c == 0
	}
	return false
}

// #endregion Utils
