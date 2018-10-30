package scanner

import (
	"fmt"

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

func (s *Scanner) SkipError() {
	s.next()
	for s.ch != '\n' && s.ch != -1 {
		s.next()
	}
}

func (s *Scanner) skipSpace() {
	for {
		switch s.ch {
		case '\n': // Prevents calling parameters from crossing lines
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
	if s.ch == '\n' {
		s.file.AddLine(s.off)
	}
	s.off++
	return
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
	case s.ch > 127:
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
	case s.ch == '#':
		s.SkipError()
		return s.Scan()
	default:
		tok, val = s.scanOperator()

		switch val {
		case "/":
			if s.ch == '/' {
				s.SkipError()
				return s.Scan()
			}

		}
		switch tok {
		case token.PERIOD:
			if s.ch >= '0' && s.ch <= '9' {
				tok = token.NUMBER
				val = "." + s.scanNumber()
			}
			return
		case token.INVALID:
			err = fmt.Errorf("Invalid symbol '%v'", val)
			s.SkipError()
			return
		}

		return
	}
}
