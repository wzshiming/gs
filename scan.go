package gs

type scan struct {
	buf []rune
	ch  rune
	off int
}

func NewScan(buf string) *scan {
	s := &scan{
		buf: []rune(buf),
	}
	s.next()
	return s
}

func (s *scan) next() {
	if len(s.buf) == s.off {
		s.ch = -1
		s.off = len(s.buf) + 1
		return
	}

	s.ch = s.buf[s.off]
	s.off++

	return
}

func (s *scan) Scan() Expr {
	return s.scanBinary(1)
}

func (s *scan) scanUnary() Expr {
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

func (s *scan) scanBinary(pre int) Expr {
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

func (s *scan) scanOperator() Token {
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

func (s *scan) scanNumber() string {
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
