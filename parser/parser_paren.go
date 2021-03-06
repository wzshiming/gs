package parser

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseParen() ast.Expr {
	tok := s.tok
	pos := s.pos
	s.scan()
	expr := s.parseTuple()
	if s.tok != token.RPAREN {
		s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
	}
	s.scan()
	return expr
}
