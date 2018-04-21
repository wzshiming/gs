package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalFor(t *ast.For, s *value.Scope) value.Value {
	ss := s.NewChildScope()
	ev.eval(t.Init, ss)
	i := 0

	var ex value.Value
	for {
		loop := ev.eval(t.Cond, ss)
		vb, ok := loop.(value.ValueBool)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'for'."))
			break
		}

		if !vb {
			break
		}

		ex = ev.eval(t.Body, ss)
		if t.Next != nil {
			ev.eval(t.Next, ss)
		}
	}

	if i != 0 {
		return ex
	}

	if t.Else != nil {
		return ev.eval(t.Else, ss)
	}

	return value.ValueNil
}
