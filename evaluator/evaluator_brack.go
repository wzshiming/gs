package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalBrack(t *ast.Brack, s *value.Scope) value.Value {
	x := ev.eval(t.X, s)
	y := ev.eval(t.Y, s)
	switch tx := x.(type) {
	case *value.ValueTuple:
		switch ty := y.(type) {
		case value.ValueNumber:
			l := int(ty.Int())
			i := tx.Len()
			if i <= l {
				ev.errorsPos(t.X.GetPos(), fmt.Errorf("Index out of range."))
				return value.ValueNil
			}
			return tx.Index(l)
		default:
			ev.errorsPos(t.X.GetPos(), fmt.Errorf("Indexes must be Numbers."))
			return value.ValueNil
		}
	default:
		switch ty := y.(type) {
		case value.ValueNumber:
			l := int(ty.Int())
			if l != 0 {
				ev.errorsPos(t.X.GetPos(), fmt.Errorf("Index out of range."))
				return value.ValueNil
			}
			return tx
		default:
			ev.errorsPos(t.X.GetPos(), fmt.Errorf("Indexes must be Numbers."))
			return value.ValueNil
		}
	}

	return value.ValueNil
}