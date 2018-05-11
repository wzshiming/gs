package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalFunc(t *ast.Func, s value.Assigner) value.Value {
	fun := t.Func
	//ss := s.Child()
	vf := &value.Func{
		Scope: s,
		Body:  t.Body,
	}
	switch t0 := fun.(type) {
	case *ast.Call: // func name a,b
		switch t1 := t0.Name.(type) {
		case *ast.Literal: // func name a,b
			s.SetLocal(value.String(t1.Value), vf)
			//t1.Value
		default: // func typ.name a,b
		}
		vf.Args = t0.Args
	default: // func a,b
	}

	return vf
}
