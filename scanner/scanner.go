package scanner

import (
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

type Scanner struct {
	file *position.File
	buf  []rune
	ch   rune
	off  int
}

func NewScanner(f *position.File, buf []rune) *Scanner {
	s := &Scanner{
		file: f,
		buf:  buf,
	}
	s.next()
	return s
}

func (s *Scanner) skipSpace() {
	for {
		switch s.ch {
		case '\n':
			s.file.AddLine(s.off)
			return
		case ' ', '\r', '\t':
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

func (s *Scanner) scanString() string {
	off := s.off - 1
	ch := s.ch
	s.next()
	for ch != s.ch {
		s.next()
	}
	s.next()
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

func (s *Scanner) scanOperator() token.Token {
	look := token.LookupOperator
	for {
		look0 := look.GetRune(s.ch)
		if look0 == nil {
			return look.Tok
		}
		look = look0
		s.next()
	}
}

func (s *Scanner) Scan() (pos position.Pos, tok token.Token, val string, err error) {
	s.skipSpace()
	pos = s.file.Pos(s.off - 1)
	switch {
	case s.ch == '\'', s.ch == '"', s.ch == '`':
		tok = token.STRING
		val = s.scanString()
		return
	case s.ch >= '0' && s.ch <= '9':
		tok = token.NUMBER
		val = s.scanNumber()
		return
	case s.ch >= 'a' && s.ch <= 'z':
		val = s.scanIdent()

		switch val {
		case "true", "false":
			tok = token.BOOL
		case "nil":
			tok = token.NIL
		default:
			tok = token.LookupKeywork.Get(val)
			if tok == token.INVALID {
				tok = token.IDENT
			}
		}
		return
	case s.ch >= 'A' && s.ch <= 'Z', s.ch == '_':
		tok = token.IDENT
		val = s.scanIdent()
		return

	case s.ch == '\n', s.ch == ';':
		val = string([]rune{s.ch})
		for s.ch == '\n' || s.ch == ';' {
			s.next()
		}
		tok = token.SEMICOLON
		return
	case s.ch == -1:
		val = ""
		tok = token.EOF
		return
	default:
		tok = s.scanOperator()
		val = tok.String()
		if tok == token.PERIOD && s.ch >= '0' && s.ch <= '9' {
			tok = token.NUMBER
			val = "." + s.scanNumber()
		}
		return
	}
}
