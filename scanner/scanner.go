package scanner

import (
	"github.com/wzshiming/gs/token"
)

type Scanner struct {
	buf []rune
	ch  rune
	off int
	Tok token.Token
	Val string
}

func NewScanner(buf string) *Scanner {
	s := &Scanner{
		buf: []rune(buf),
	}
	s.next()
	s.Scan()
	return s
}

func (s *Scanner) skipSpace() {
	for {
		switch s.ch {
		case ' ', '\n', '\r', '\t':
			s.next()
		default:
			return
		}
	}
}

func (s *Scanner) next() {
	if len(s.buf) <= s.off {
		s.ch = -1
		s.off = len(s.buf) + 1
		return
	}

	s.ch = s.buf[s.off]
	s.off++

	return
}

func (s *Scanner) scanIdent() string {
	off := s.off - 1
	for s.ch >= '0' && s.ch <= '9' ||
		s.ch >= 'a' && s.ch <= 'z' ||
		s.ch >= 'A' && s.ch <= 'Z' ||
		s.ch == '_' {
		s.next()
	}
	return string(s.buf[off : s.off-1])
}

func (s *Scanner) scanNumber() string {
	off := s.off - 1
	for s.ch >= '0' && s.ch <= '9' {
		s.next()
	}
	if s.ch == '.' {
		s.next()
		for s.ch >= '0' && s.ch <= '9' {
			s.next()
		}
	}
	return string(s.buf[off : s.off-1])
}

func (s *Scanner) Scan() {
	s.skipSpace()
	switch {
	case s.ch >= '0' && s.ch <= '9':
		s.Tok = token.NUMBER
		s.Val = s.scanNumber()
		return
	case s.ch >= 'a' && s.ch <= 'z',
		s.ch >= 'A' && s.ch <= 'Z',
		s.ch == '_':

		iden := s.scanIdent()

		if tok := token.LookupKeywork(iden); tok != token.INVALID {
			s.Tok = tok
			s.Val = iden
			return
		}
		s.Tok = token.IDENT
		s.Val = iden
		return

	default:
		op := token.LookupOperator(string([]rune{s.ch}))
		s.Tok = op
		s.Val = ""
		s.next()
		return
	}
}
