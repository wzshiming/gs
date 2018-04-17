package gs

type scanner struct {
	buf []rune
	ch  rune
	off int
}

func NewScanner(buf string) *scanner {
	s := &scanner{
		buf: []rune(buf),
	}
	s.next()
	return s
}

func (s *scanner) next() {
	if len(s.buf) == s.off {
		s.ch = -1
		s.off = len(s.buf) + 1
		return
	}

	s.ch = s.buf[s.off]
	s.off++

	return
}

func (s *scanner) scanOperator() Token {
	switch {
	case s.ch == '+':
		return ADD
	case s.ch == '-':
		return SUB
	case s.ch == '*':
		return MUL
	case s.ch == '/':
		return QUO
	default:
		return 0
	}
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
