package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseUnarySuf(expr ast.Expr) ast.Expr {

	tok := s.tok
	pos := s.pos
	switch {
	case tok.IsOperator():
		switch s.tok {
		case token.INC, token.DEC, token.ELLIPSIS:

			expr = &ast.UnarySuf{
				Pos: pos,
				Op:  tok,
				X:   expr,
			}
			s.scan()
			return s.parseUnarySuf(expr)
		case token.LPAREN:
			expr = &ast.Call{
				Pos:  pos,
				Name: expr,
				Args: s.parseExpr(),
			}
			return expr
		default:
			return expr
		}

	case tok.IsKeywork():
		return expr

	//case tok.IsLiteral():
	default:
		switch s.tok {
		case token.EOF:
		case token.SEMICOLON:
		default:
			expr = &ast.Call{
				Pos:  pos,
				Name: expr,
				Args: s.parseExpr(),
			}
		}

		return s.parseUnarySuf(expr)
	}

	return expr
}