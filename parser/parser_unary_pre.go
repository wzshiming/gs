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
		return s.parseParen()
	case token.LBRACK:
		return s.parseBrack()
	case token.LBRACE:
		return s.parseBrace()
	default:
		s.errors(fmt.Errorf("Undefined unary expr %v", s.val))
		return nil
	}

}
