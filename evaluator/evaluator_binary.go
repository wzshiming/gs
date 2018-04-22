package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalBinary(t *ast.Binary, s *value.Scope) value.Value {
	lx := ev.eval(t.X, s)
	ly := ev.eval(t.Y, s)
	z, err := ev.evalBinaryBy(lx, ly, t.Op)
	if err != nil {
		ev.errorsPos(t.Pos, err)
		return value.ValueNil
	}
	return z
}

func (ev *Evaluator) evalBinaryBy(x, y value.Value, op token.Token) (value.Value, error) {
	z, err := x.Binary(op, y)
	if err != nil {
		return value.ValueNil, err
	}
	return z, nil
}
