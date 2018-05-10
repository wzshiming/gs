package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalVar(e ast.Expr, s *value.Scope) value.Value {
	return ev.eval(e, s)
}
