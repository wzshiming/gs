package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
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

		switch op {
		case token.COMMA:
			if t, ok := x.(*ast.Tuple); ok {
				if y != nil {
					t.List = append(t.List, y)
				}
			} else {
				x = &ast.Tuple{
					Pos:  pos,
					List: []ast.Expr{x, y},
				}
			}
		default:
			x = &ast.Binary{
				Pos: pos,
				X:   x,
				Op:  op,
				Y:   y,
			}
		}
	}
	return x
}
