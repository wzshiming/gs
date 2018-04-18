package gs

type parser struct {
	scanner
}

func NewParser(s string) *parser {
	return &parser{
		scanner: *NewScanner(s),
	}
}

func (s *parser) Parse() []Expr {
	ex := []Expr{}
	for s.tok != 0 {
		pe := s.ParseExpr()
		if pe == nil {
			break
		}
		ex = append(ex, pe)

	}
	return ex
}

func (s *parser) ParseExpr() Expr {
	return s.parseBinaryExpr(1)
}

func (s *parser) parseUnaryExpr() Expr {

	switch {
	case s.tok.IsOperator():
		switch s.tok {
		case ADD, SUB:
			tok := s.tok
			s.scan()
			b := &OperatorUnary{
				Op: tok,
				X:  s.parseUnaryExpr(),
			}
			return b
		case RPAREN, RBRACE:
			return nil
		case LPAREN:
			s.scan()
			b := s.ParseExpr()
			s.scan()
			return b
		case LBRACE:
			s.scan()
			b := s.Parse()
			s.scan()
			return &BraceExpr{
				List: b,
			}
		case COMMA:
		}
	case s.tok.IsKeywork():
		switch s.tok {
		case IF:
			s.scan()
			cond := s.ParseExpr()
			body := s.ParseExpr()
			var els Expr

			if s.tok == ELSE {
				s.scan()
				els = s.ParseExpr()
			}
			return &IfExpr{
				Cond: cond,
				Body: body,
				Else: els,
			}
		}

	default:
		b := &Literal{
			Type:  s.tok,
			Value: s.val,
		}
		s.scan()
		return b
	}

	return nil
}

func (s *parser) parseBinaryExpr(pre int) Expr {
	x := s.parseUnaryExpr()
	if x == nil {
		return x
	}

	for {
		op := s.tok
		if !op.IsOperator() {
			break
		}
		op2 := op.Precedence()
		if op2 < pre {
			break
		}
		s.scan()
		y := s.parseBinaryExpr(op2 + 1)
		x = &OperatorBinary{
			X:  x,
			Op: op,
			Y:  y,
		}
	}
	return x
}
