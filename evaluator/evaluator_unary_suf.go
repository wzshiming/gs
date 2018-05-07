package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalUnarySuf(t *ast.UnarySuf, s *value.Scope) value.Value {
	lx := ev.eval(t.X, s)

	z, err := lx.UnarySuf(t.Op)
	if err != nil {
		ev.errorsPos(t.Pos, err)
		return value.Nil
	}
	return z
}
