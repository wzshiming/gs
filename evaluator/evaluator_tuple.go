package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalTuple(t *ast.Tuple, s *value.Scope) value.Value {
	z := &value.ValueTuple{}
	for _, v := range t.List {
		z.List = append(z.List, ev.eval(v, s))
	}
	return z
}
