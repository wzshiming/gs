package parser

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseUnary() (expr ast.Expr) {
	tok := s.tok

	switch {
	case tok.IsOperator():
		expr = s.parseUnaryPre()
	case tok.IsKeywork():
		expr = s.parseKeywork()
	default:
		switch tok {
		case token.EOF:
		case token.INVALID:
			s.errors(fmt.Errorf("Undefined value %v", s.val))
		default:
			expr = s.parseLiteral()
		}

	}

	if expr == nil {
		return
	}

	return s.parseUnarySuf(expr)
}
