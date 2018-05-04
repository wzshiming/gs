package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalCall(t *ast.Call, s *value.Scope) value.Value {

	switch t1 := t.Name.(type) {
	case *ast.Literal: // name a,b
		val, ok := s.Get(t1.Value)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("Undefined function %s", t1.Value))
			break
		}
		switch t2 := val.(type) {
		case *value.ValueFunc:

			ss := t2.Scope.NewChildScope()
			ev.evalBinaryBy(ev.eval(t2.Args, ss), ev.toValues(t.Args, s), token.DEFINE)
			ev.stackRet++
			return ev.eval(t2.Body, ss)
		case *value.ValueFuncBuiltin:

			r, err := t2.Call(ev.toValues(t.Args, s))
			if err != nil {
				ev.errorsPos(t.Pos, err)
				return value.ValueNil
			}
			return r
		default:
			ev.errorsPos(t.Pos, fmt.Errorf("Not a function"))
			break
		}

	default: // typ.name a,b
	}
	return value.ValueNil
}
