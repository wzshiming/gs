package parser

import (
	"github.com/wzshiming/gs/ast"
)

func (s *parser) parseLiteral() *ast.Literal {
	b := &ast.Literal{
		Pos:   s.pos,
		Type:  s.tok,
		Value: s.val,
	}
	s.scan()
	return b
}
