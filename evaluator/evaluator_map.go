package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
	ffmt "gopkg.in/ffmt.v1"
)

func (ev *Evaluator) evalMap(m *ast.Map, s *value.Scope) value.Value {
	switch t := m.Body.(type) {
	case *ast.Brace:
		mr := value.Map{}

		list := t.List
		if len(list) != 1 {
			ev.errorsPos(m.Pos, fmt.Errorf("Map syntax error"))
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
				mr[x] = y
			default:
				ffmt.P(t1)
			}
		}
		return mr
	default:
	}
	return value.Nil
}
