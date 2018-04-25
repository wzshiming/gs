package parser

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseUnaryPre() (expr ast.Expr) {
	tok := s.tok
	pos := s.pos

	switch s.tok {
	case token.SEMICOLON:
		s.scan()
		return s.parseUnary()
	case token.ADD, token.SUB, token.ELLIPSIS:
		s.scan()
		expr = &ast.UnaryPre{
			Pos: pos,
			Op:  tok,
			X:   s.parseUnary(),
		}
		return expr
	case token.RPAREN, token.RBRACK, token.RBRACE:
		return nil
	case token.LPAREN:
		s.scan()
		b := s.parseExpr()
		if s.tok != token.RPAREN {
			s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
		}
		s.scan()
		return b
	case token.LBRACK:
		s.scan()
		b := s.parseExpr()
		if s.tok != token.RBRACK {
			s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
		}
		s.scan()
		return b
	case token.LBRACE:
		s.scan()
		b := s.parse()
		if s.tok != token.RBRACE {
			s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
		}
		s.scan()

		expr = &ast.Brace{
			Pos:  pos,
			List: b,
		}
		return expr
	default:
		s.errors(fmt.Errorf("Undefined unary expr %v", s.val))
		return nil
	}

}
