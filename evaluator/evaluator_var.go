package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalVar(e ast.Expr, s value.Assigner) value.Value {
	switch t := e.(type) {
	case *ast.Tuple:
		return ev.evalTupleVar(t, s)
	case *ast.Literal:
		if t.Type != token.IDENT {
			break
		}
		return &value.Var{
			Name:  value.String(t.Value),
			Scope: s,
		}
	case *ast.Binary:
		if t.Op != token.PERIOD {
			break
		}
		x := ev.evalVar(t.X, s)
		x = x.Point()
		ass, ok := x.(value.Assigner)
		if !ok {
			break
		}
		y := ev.eval(t.Y, s)
		return &value.Var{
			Name:  y,
			Scope: ass,
		}
	case *ast.Brack:
		x := ev.evalVar(t.X, s)
		x = x.Point()
		ass, ok := x.(value.Assigner)
		if !ok {
			break
		}
		y := ev.eval(t.Y, s)
		return &value.Var{
			Name:  y,
			Scope: ass,
		}
	case *ast.UnaryPre:
		if t.Op != token.ELLIPSIS {
			break
		}
		return ev.evalUnaryPre(t, s)
	}
	if e == nil {
		return value.Nil
	}

	ev.errorsPos(e.GetPos(), fmt.Errorf("Not assignable to non left values"))
	return value.Nil
}
