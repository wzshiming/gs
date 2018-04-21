package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalBinary(t *ast.Binary, s *value.Scope) value.Value {
	lx := ev.eval(t.X, s)
	ly := ev.eval(t.Y, s)

	z, err := lx.Binary(t.Op, ly)
	if err != nil {
		ev.errorsPos(t.Pos, err)
		return value.ValueNil
	}
	return z
}
