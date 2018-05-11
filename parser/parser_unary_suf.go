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
				Args: s.parseUnaryPre(),
			}
			return expr
		case token.LBRACK:
			expr = &ast.Brack{
				Pos: pos,
				X:   expr,
				Y:   s.parseUnaryPre(),
			}
			return expr
		default:
			return expr
		}

	case tok.IsKeywork():
		return expr

	default:
		switch s.tok {
		case token.EOF:
			return expr
		case token.SEMICOLON:
			return expr
		default:
			expr = &ast.Call{
				Pos:  pos,
				Name: expr,
				Args: s.parseExpr(),
			}
			return s.parseUnarySuf(expr)
		}

	}
}
