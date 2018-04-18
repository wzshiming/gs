package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/scanner"
	"github.com/wzshiming/gs/token"
)

type parser struct {
	scanner.Scanner
}

func NewParser(s string) *parser {
	return &parser{
		Scanner: *scanner.NewScanner(s),
	}
}

func (s *parser) Parse() []ast.Expr {
	ex := []ast.Expr{}
	for s.Tok != 0 {
		pe := s.ParseExpr()
		if pe == nil {
			break
		}
		ex = append(ex, pe)

	}
	return ex
}

func (s *parser) ParseExpr() ast.Expr {
	return s.parseBinaryExpr(1)
}

func (s *parser) parseUnaryExpr() ast.Expr {

	switch {
	case s.Tok.IsOperator():
		switch s.Tok {
		case token.ADD, token.SUB:
			tok := s.Tok
			s.Scan()
			b := &ast.OperatorUnary{
				Op: tok,
				X:  s.parseUnaryExpr(),
			}
			return b
		case token.RPAREN, token.RBRACE:
			return nil
		case token.LPAREN:
			s.Scan()
			b := s.ParseExpr()
			s.Scan()
			return b
		case token.LBRACE:
			s.Scan()
			b := s.Parse()
			s.Scan()
			return &ast.BraceExpr{
				List: b,
			}
		case token.COMMA:
		}
	case s.Tok.IsKeywork():
		switch s.Tok {
		case token.IF:
			s.Scan()
			cond := s.ParseExpr()
			body := s.ParseExpr()
			var els ast.Expr

			if s.Tok == token.ELSE {
				s.Scan()
				els = s.ParseExpr()
			}
			return &ast.IfExpr{
				Cond: cond,
				Body: body,
				Else: els,
			}
		}

	default:
		b := &ast.Literal{
			Type:  s.Tok,
			Value: s.Val,
		}
		s.Scan()
		return b
	}

	return nil
}

func (s *parser) parseBinaryExpr(pre int) ast.Expr {
	x := s.parseUnaryExpr()
	if x == nil {
		return x
	}

	for {
		op := s.Tok
		if !op.IsOperator() {
			break
		}
		op2 := op.Precedence()
		if op2 < pre {
			break
		}
		s.Scan()
		y := s.parseBinaryExpr(op2 + 1)
		x = &ast.OperatorBinary{
			X:  x,
			Op: op,
			Y:  y,
		}
	}
	return x
}
