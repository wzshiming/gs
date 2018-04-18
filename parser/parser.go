package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/scanner"
	"github.com/wzshiming/gs/token"
)

type parser struct {
	scanner *scanner.Scanner
	fset    *position.FileSet

	tok token.Token
	val string
	pos position.Pos
	err error
}

func NewParser(fset *position.FileSet, filename string, src []rune) *parser {
	file := fset.AddFile(filename, 1, len(src)-1)
	p := &parser{
		fset:    fset,
		scanner: scanner.NewScanner(file, src),
	}
	p.scan()
	return p
}

func (s *parser) scan() {
	s.pos, s.tok, s.val, s.err = s.scanner.Scan()

	//	ffmt.Mark(s.fset.Position(s.pos), s.tok, s.val)
}

func (s *parser) Parse() []ast.Expr {
	ex := []ast.Expr{}
	for {
		switch s.tok {
		case token.EOF:
			return ex
		case token.RPAREN, token.RBRACK, token.RBRACE:
			s.scan()
			return ex
		}
		pe := s.ParseExpr()
		if pe == nil {
			//			ffmt.Mark(s.fset.Position(s.pos), s.tok, s.val)
			return ex
		}
		ex = append(ex, pe)
	}
}

func (s *parser) ParseExpr() ast.Expr {
	return s.parseBinaryExpr(1)
}

func (s *parser) parsePreUnaryExpr() (expr ast.Expr) {
	tok := s.tok
	pos := s.pos

	switch {
	case tok.IsOperator():
		switch s.tok {
		case token.SEMICOLON:
			s.scan()
			return s.parsePreUnaryExpr()
		case token.ADD, token.SUB, token.ELLIPSIS:
			s.scan()
			expr = &ast.OperatorPreUnary{
				Pos: pos,
				Op:  tok,
				X:   s.parsePreUnaryExpr(),
			}

		case token.RPAREN, token.RBRACE:
			//s.scan()
			// return nil
		case token.LPAREN:
			s.scan()
			b := s.ParseExpr()
			s.scan()
			expr = b
		case token.LBRACE:
			//			ffmt.Mark(s.tok)
			s.scan()
			//	ffmt.Mark(s.tok)
			b := s.Parse()
			//	ffmt.Mark(s.tok)
			s.scan()
			expr = &ast.BraceExpr{
				Pos:  pos,
				List: b,
			}
		//	ffmt.Mark(s.tok)
		case token.COMMA:
		}
	case tok.IsKeywork():
		switch tok {
		case token.IF:
			s.scan()
			cond := s.ParseExpr()
			body := s.ParseExpr()
			var els ast.Expr

			for s.tok == token.SEMICOLON {
				s.scan()
			}

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
		switch tok {
		case token.EOF:
		default:
			b := &ast.Literal{
				Pos:   pos,
				Type:  s.tok,
				Value: s.val,
			}
			s.scan()
			expr = b
		}

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
		case tok.IsKeywork():
			break loop

		case tok.IsLiteral():
			expr = &ast.CallExpr{
				Pos:      pos,
				Name:     expr,
				Argument: s.ParseExpr(),
			}
		default:

			break loop

		}
	}

	return expr
}

func (s *parser) parseBinaryExpr(pre int) ast.Expr {

	x := s.parsePreUnaryExpr()
	if x == nil {
		return x
	}

	for {
		op := s.tok
		pos := s.pos

		op2 := op.Precedence()
		if op2 < pre {
			break
		}
		s.scan()
		y := s.parseBinaryExpr(op2 + 1)

		switch op {
		case token.COMMA:
			if t, ok := x.(*ast.TupleExpr); ok {
				t.List = append(t.List, y)
			} else {
				x = &ast.TupleExpr{
					Pos:  pos,
					List: []ast.Expr{x, y},
				}
			}
		default:
			x = &ast.OperatorBinary{
				Pos: pos,
				X:   x,
				Op:  op,
				Y:   y,
			}
		}
	}
	return x
}
