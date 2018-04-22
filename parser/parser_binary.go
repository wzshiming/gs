package parser

import (
	"github.com/wzshiming/gs/ast"
)

func (s *parser) parseBinary(pre int) ast.Expr {
	x := s.parseUnary()
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
		y := s.parseBinary(op2 + 1)
		x = &ast.Binary{
			Pos: pos,
			X:   x,
			Op:  op,
			Y:   y,
		}
	}
	return x
}
