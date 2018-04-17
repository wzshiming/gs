package gs

type parser struct {
	scanner
}

func NewParser(s string) *parser {
	return &parser{
		scanner: *NewScanner(s),
	}
}

func (s *parser) Parse() Expr {
	return s.parseBinary(1)
}
func (s *parser) parseUnary() Expr {
	switch {
	case s.ch == '+', s.ch == '-':
		op := s.operator()
		s.next()
		b := &OperatorUnary{
			Op: op,
			X:  s.parseUnary(),
		}
		return b
	case s.ch == '(':
		s.next()
		b := s.parseBinary(1)
		s.next()
		return b
	case s.ch >= '0' && s.ch <= '9':
		return &Literal{
			Type:  NUMBER,
			Value: s.scanNumber(),
		}
	default:
		return &Literal{
			Type:  IDENT,
			Value: s.scanIdent(),
		}
	}
	return nil
}

func (s *parser) parseBinary(pre int) Expr {
	x := s.parseUnary()
	for {
		op := s.operator()
		op2 := op.Precedence()
		if op2 < pre {
			return x
		}
		s.next()
		y := s.parseBinary(op2 + 1)
		x = &OperatorBinary{
			X:  x,
			Op: op,
			Y:  y,
		}
	}
}
