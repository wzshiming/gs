package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/scanner"
	"github.com/wzshiming/gs/token"
)

type parser struct {
	scanner *scanner.Scanner

	tok token.Token
	val string
	pos position.Pos
	err error
}

func NewParser(fset *position.FileSet, filename string, src []rune) *parser {
	file := fset.AddFile(filename, 1, len(src)-1)
	p := &parser{
		scanner: scanner.NewScanner(file, src),
	}
	p.scan()
	return p
}

func (s *parser) scan() {
	s.pos, s.tok, s.val, s.err = s.scanner.Scan()
}

func (s *parser) Parse() []ast.Expr {
	ex := []ast.Expr{}
	for s.tok != 0 {
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

func (s *parser) parseUnaryExpr() (expr ast.Expr) {
	tok := s.tok
	pos := s.pos
	switch {
	case tok.IsOperator():
		switch s.tok {
		case token.ADD, token.SUB, token.ELLIPSIS:
			s.scan()
			expr = &ast.OperatorPreUnary{
				Pos: pos,
				Op:  tok,
				X:   s.parseUnaryExpr(),
			}

		case token.RPAREN, token.RBRACE:
			// return nil
		case token.LPAREN:
			s.scan()
			b := s.ParseExpr()
			s.scan()
			expr = b
		case token.LBRACE:

			s.scan()
			b := s.Parse()
			s.scan()
			expr = &ast.BraceExpr{
				Pos:  pos,
				List: b,
			}
		case token.COMMA:
		}
	case tok.IsKeywork():
		switch tok {
		case token.IF:
			s.scan()
			cond := s.ParseExpr()
			body := s.ParseExpr()
			var els ast.Expr

			if s.tok == token.ELSE {
				s.scan()
				els = s.ParseExpr()
			}
			expr = &ast.IfExpr{
				Pos:  pos,
				Cond: cond,
				Body: body,
				Else: els,
			}
		}

	default:
		b := &ast.Literal{
			Pos:   pos,
			Type:  s.tok,
			Value: s.val,
		}
		s.scan()
		expr = b
	}

loop:
	for {
		tok := s.tok
		pos := s.pos
		switch {
		case tok.IsOperator():
			switch s.tok {
			case token.INC, token.DEC, token.ELLIPSIS:

				expr = &ast.OperatorSufUnary{
					Pos: pos,
					Op:  tok,
					X:   expr,
				}
				s.scan()
			default:
				break loop
			}

		default:
			break loop
		}
	}

	return expr
}

func (s *parser) parseBinaryExpr(pre int) ast.Expr {
	x := s.parseUnaryExpr()
	if x == nil {
		return x
	}

	for {
		op := s.tok
		pos := s.pos
		if !op.IsOperator() {
			break
		}
		op2 := op.Precedence()
		if op2 < pre {
			break
		}
		s.scan()
		y := s.parseBinaryExpr(op2 + 1)
		x = &ast.OperatorBinary{
			Pos: pos,
			X:   x,
			Op:  op,
			Y:   y,
		}
	}
	return x
}
