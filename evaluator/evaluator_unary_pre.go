package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalUnaryPre(t *ast.UnaryPre, s value.Assigner) value.Value {
	lx := ev.eval(t.X, s)

	z, err := lx.UnaryPre(t.Op)
	if err != nil {
		ev.errorsPos(t.Pos, err)
		return value.Nil
	}
	return z
}
