package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalFor(t *ast.For, s value.Assigner) value.Value {
	ss := s.Child()
	ev.eval(t.Init, ss)
	i := 0

	ex := value.Nil
	for {
		loop := ev.eval(t.Cond, ss)
		vb, ok := loop.(value.Bool)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'for'"))
			return value.Nil
		}

		if !vb {
			if i != 0 && t.Else != nil {
				return ev.eval(t.Else, ss)
			}
			return ex
		}

		ex = ev.eval(t.Body, ss)
		if t.Next != nil {
			ev.eval(t.Next, ss)
		}
	}
}
