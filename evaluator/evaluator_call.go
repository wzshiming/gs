package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalCall(t *ast.Call, s value.Assigner) value.Value {

	switch t1 := t.Name.(type) {
	case *ast.Literal: // name a,b
		val, ok := s.Get(value.String(t1.Value))
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("Undefined function %s", t1.Value))
			break
		}
		switch t2 := val.(type) {
		case *value.Func:
			ss := t2.Scope.Child()
			x := value.Nil
			y := value.Nil
			if t2.Args != nil {
				x = ev.evalVar(t2.Args, ss)
			}
			if t.Args != nil {
				y = ev.eval(t.Args, s)
			}
			ev.evalBinaryBy(x, y, token.DEFINE)
			ev.stackRet++
			return ev.eval(t2.Body, ss)
		case *value.FuncBuiltin:

			r, err := t2.Call(ev.toValues(t.Args, s))
			if err != nil {
				ev.errorsPos(t.Pos, err)
				return value.Nil
			}
			return r
		default:
			ev.errorsPos(t.Pos, fmt.Errorf("Not a function"))
			break
		}

	default: // typ.name a,b
	}
	return value.Nil
}
