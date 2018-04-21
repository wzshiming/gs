package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalCall(t *ast.Call, s *value.Scope) value.Value {

	switch t1 := t.Name.(type) {
	case *ast.Literal: // name a,b
		val, ok := s.Get(t1.Value)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("未定义"))
			break
		}
		switch t2 := val.(type) {
		case *value.ValueFunc:

			ti1 := ev.toValues(t.Args, s)

			ti2 := ev.toIdents(t2.Args)

			ss := t2.Scope.NewChildScope()
			for i := 0; i != len(ti1) && i != len(ti2); i++ {
				ss.SetLocal(ti2[i].Value, ti1[i])
			}
			ev.stackRet++
			return ev.eval(t2.Body, ss)
		}

	default: // typ.name a,b
	}
	return value.ValueNil
}
