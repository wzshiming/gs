package scanner

import (
	"github.com/wzshiming/gs/token"
)

func (s *Scanner) scanOperator() (token.Token, string) {
	off := s.off - 1
	look := token.LookupOperator
	for {
		look0 := look.GetRune(s.ch)
		if look0 == nil {
			return look.Tok, string(s.buf[off : s.off-1])
		}
		look = look0
		s.next()
	}
}
