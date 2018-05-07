package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalTuple(t *ast.Tuple, s *value.Scope) value.Value {
	z := &value.Tuple{}
	for _, v := range t.List {
		b := ev.eval(v, s)
		switch t := b.(type) {
		case *value.Tuple:
			if t.Ellipsis {
				z.List = append(z.List, t.List...)
			} else {
				z.List = append(z.List, b)
			}
		default:
			z.List = append(z.List, b)
		}

	}
	return z
}
