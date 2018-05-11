package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalReturn(t *ast.Return, s value.Assigner) value.Value {
	ev.stackRet--
	if t.Ret == nil {
		return value.Nil
	}
	return ev.eval(t.Ret, s)
}
