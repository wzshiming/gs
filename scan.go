package gs

type scanner struct {
	buf []rune
	ch  rune
	off int
}

func NewScan(buf string) *scanner {
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

func (s *scanner) Scan() Expr {
	return s.scanBinary(1)
}

func (s *scanner) scanUnary() Expr {
	switch {
	case s.ch == '+', s.ch == '-':
		op := s.scanOperator()
		s.next()
		b := &OperatorUnary{
			Op: op,
			X:  s.scanUnary(),
		}
		return b
	case s.ch == '(':
		s.next()
		b := s.scanBinary(1)
		s.next()
		return b
	case s.ch >= '0' && s.ch <= '9':
		return &Literal{
			Value: s.scanNumber(),
		}
	default:

	}
	return nil
}

func (s *scanner) scanBinary(pre int) Expr {
	x := s.scanUnary()
	for {
		op := s.scanOperator()
		op2 := op.Precedence()
		if op2 < pre {
			return x
		}
		s.next()
		y := s.scanBinary(op2 + 1)
		x = &OperatorBinary{
			X:  x,
			Op: op,
			Y:  y,
		}
	}
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
