package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalUnarySuf(t *ast.UnarySuf, s value.Assigner) value.Value {
	lx := ev.eval(t.X, s)

	z, err := lx.UnarySuf(t.Op)
	if err != nil {
		ev.errorsPos(t.Pos, err)
		return value.Nil
	}
	return z
}
