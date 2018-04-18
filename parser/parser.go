package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/scanner"
	"github.com/wzshiming/gs/token"
)

type parser struct {
	scanner *scanner.Scanner
}

func NewParser(fset *position.FileSet, filename string, src []rune) *parser {
	file := fset.AddFile(filename, 1, len(src)-1)
	p := &parser{
		scanner: scanner.NewScanner(file, src),
	}

	return p
}

func (s *parser) Parse() []ast.Expr {
	ex := []ast.Expr{}
	for s.scanner.Tok != 0 {
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
	tok := s.scanner.Tok
	pos := s.scanner.Pos
	switch {
	case tok.IsOperator():
		switch s.scanner.Tok {
		case token.ADD, token.SUB:
			s.scanner.Scan()
			b := &ast.OperatorUnary{
				Pos: pos,
				Op:  tok,
				X:   s.parseUnaryExpr(),
			}
			return b
		case token.RPAREN, token.RBRACE:
			return nil
		case token.LPAREN:
			s.scanner.Scan()
			b := s.ParseExpr()
			s.scanner.Scan()
			return b
		case token.LBRACE:

			s.scanner.Scan()
			b := s.Parse()
			s.scanner.Scan()
			return &ast.BraceExpr{
				Pos:  pos,
				List: b,
			}
		case token.COMMA:
		}
	case tok.IsKeywork():
		switch tok {
		case token.IF:
			s.scanner.Scan()
			cond := s.ParseExpr()
			body := s.ParseExpr()
			var els ast.Expr

			if s.scanner.Tok == token.ELSE {
				s.scanner.Scan()
				els = s.ParseExpr()
			}
			return &ast.IfExpr{
				Pos:  pos,
				Cond: cond,
				Body: body,
				Else: els,
			}
		}

	default:
		b := &ast.Literal{
			Pos:   pos,
			Type:  s.scanner.Tok,
			Value: s.scanner.Val,
		}
		s.scanner.Scan()
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
		op := s.scanner.Tok
		pos := s.scanner.Pos
		if !op.IsOperator() {
			break
		}
		op2 := op.Precedence()
		if op2 < pre {
			break
		}
		s.scanner.Scan()
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
