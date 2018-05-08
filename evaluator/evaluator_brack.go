package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalBrack(t *ast.Brack, s *value.Scope) value.Value {
	x := ev.eval(t.X, s)
	y := ev.eval(t.Y, s)
	x = x.Point()
	y = y.Point()
	switch tx := x.(type) {
	case *value.Tuple:
		switch ty := y.(type) {
		case value.Number:
			l := int(ty.Int())
			i := tx.Len()
			if i <= l {
				ev.errorsPos(t.X.GetPos(), fmt.Errorf("Index out of range"))
				return value.Nil
			}
			return tx.Index(l)
		default:
			ev.errorsPos(t.X.GetPos(), fmt.Errorf("Indexes must be Numbers"))
			return value.Nil
		}
	case value.Map:
		val, ok := tx[y]
		if !ok {
			return value.Nil
		}
		return val
	default:
		switch ty := y.(type) {
		case value.Number:
			l := int(ty.Int())
			if l != 0 {
				ev.errorsPos(t.X.GetPos(), fmt.Errorf("Index out of range"))
				return value.Nil
			}
			return tx
		default:
			ev.errorsPos(t.X.GetPos(), fmt.Errorf("Indexes must be Numbers"))
			return value.Nil
		}
	}
}
