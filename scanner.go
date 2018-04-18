package gs

type scanner struct {
	buf []rune
	ch  rune
	off int
	tok Token
	val string
}

func NewScanner(buf string) *scanner {
	s := &scanner{
		buf: []rune(buf),
	}
	s.next()
	s.scan()
	return s
}

func (s *scanner) skipSpace() {
	for {
		switch s.ch {
		case ' ', '\n', '\r', '\t':
			s.next()
		default:
			return
		}
	}
}

func (s *scanner) next() {
	if len(s.buf) <= s.off {
		s.ch = -1
		s.off = len(s.buf) + 1
		return
	}

	s.ch = s.buf[s.off]
	s.off++

	return
}

func (s *scanner) scanIdent() string {
	off := s.off - 1
	for s.ch >= '0' && s.ch <= '9' ||
		s.ch >= 'a' && s.ch <= 'z' ||
		s.ch >= 'A' && s.ch <= 'Z' ||
		s.ch == '_' {
		s.next()
	}
	return string(s.buf[off : s.off-1])
}

func (s *scanner) scanNumber() string {
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

func (s *scanner) scan() {
	s.skipSpace()
	switch {
	case s.ch >= '0' && s.ch <= '9':
		s.tok = NUMBER
		s.val = s.scanNumber()
		return
	case s.ch >= 'a' && s.ch <= 'z',
		s.ch >= 'A' && s.ch <= 'Z',
		s.ch == '_':

		iden := s.scanIdent()

		if tok := LookupKeywork(iden); tok != INVALID {
			s.tok = tok
			s.val = iden
			return
		}
		s.tok = IDENT
		s.val = iden
		return

	default:
		op := LookupOperator(string([]rune{s.ch}))
		s.tok = op
		s.val = ""
		s.next()
		return
	}
}
