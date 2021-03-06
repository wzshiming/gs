package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/value"
)

// Evaluator evaluate ast
type Evaluator struct {
	fset *position.FileSet
	errs *errors.Errors

	stackRet int
}

// NewEvaluator Create a new evaluator
func NewEvaluator(fset *position.FileSet, errs *errors.Errors) *Evaluator {
	return &Evaluator{
		fset: fset,
		errs: errs,
	}
}

func (ev *Evaluator) errorsPos(pos position.Pos, err error) {
	ev.errs.Append(ev.fset.Position(pos), err)
}

// EvalBy evaluate expression in the scope
func (ev *Evaluator) EvalBy(es []ast.Expr, s value.Assigner) (ex value.Value) {
	sr := ev.stackRet
	for _, v := range es {
		ex = ev.eval(v, s)
		if ev.stackRet < sr {
			return
		}
	}
	ev.stackRet--
	return
}

func (ev *Evaluator) eval(e ast.Expr, s value.Assigner) value.Value {
	switch t := e.(type) {
	case *ast.Literal:
		return ev.evalLiteral(t, s)
	case *ast.Binary:
		return ev.evalBinary(t, s)
	case *ast.UnaryPre:
		return ev.evalUnaryPre(t, s)
	case *ast.UnarySuf:
		return ev.evalUnarySuf(t, s)
	case *ast.If:
		return ev.evalIf(t, s)
	case *ast.Brack:
		return ev.evalBrack(t, s)
	case *ast.Labeled:
		// TODO:
		return value.Nil
	case *ast.Break:
		// TODO:
		return value.Nil
	case *ast.Continue:
		// TODO:
		return value.Nil
	case *ast.For:
		return ev.evalFor(t, s)
	case *ast.Brace:
		ss := s.Child()
		return ev.EvalBy(t.List, ss)
	case *ast.Call:
		return ev.evalCall(t, s)
	case *ast.Func:
		return ev.evalFunc(t, s)
	case *ast.Return:
		return ev.evalReturn(t, s)
	case *ast.Tuple:
		return ev.evalTuple(t, s)
	case *ast.Map:
		return ev.evalMap(t, s)
	}
	ev.errorsPos(e.GetPos(), fmt.Errorf("Undefined keyword processing"))
	return value.Nil
}
