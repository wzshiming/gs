package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalIf(t *ast.If, s *value.Scope) value.Value {
	ss := s.NewChildScope()
	ev.eval(t.Init, ss)
	loop := ev.eval(t.Cond, ss)
	vb, ok := loop.(value.Bool)
	if !ok {
		ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'if'"))
		return value.Nil
	}

	if vb {
		return ev.eval(t.Body, ss)
	} else if t.Else != nil {
		return ev.eval(t.Else, ss)
	}
	return value.Nil
}
