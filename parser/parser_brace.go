package parser

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseBrace() ast.Expr {
	tok := s.tok
	pos := s.pos
	s.scan()
	exprs := s.parse()
	if s.tok != token.RBRACE {
		s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
	}
	s.scan()
	expr := &ast.Brace{
		Pos:  pos,
		List: exprs,
	}
	return expr
}
