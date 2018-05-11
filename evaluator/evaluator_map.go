package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalMap(m *ast.Map, s value.Assigner) value.Value {
	switch t := m.Body.(type) {
	case *ast.Brace:
		mr := value.Map{}

		list := t.List
		if len(list) != 1 {
			return mr
		}

		vt, ok := list[0].(*ast.Tuple)
		if ok {
			list = vt.List
		}
		for _, v := range list {
			switch t1 := v.(type) {
			case *ast.Binary:
				if t1.Op != token.COLON {
					ev.errorsPos(t1.Pos, fmt.Errorf("Not a key value pair"))
					break
				}
				x := ev.eval(t1.X, s)
				y := ev.eval(t1.Y, s)
				x = x.Point()
				y = y.Point()
				mr[x] = y
			default:
				ev.errorsPos(v.GetPos(), fmt.Errorf("Not a colon statement"))
			}
		}
		return mr
	default:
	}
	return value.Nil
}
