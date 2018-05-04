package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/value"
)

type Evaluator struct {
	fset *position.FileSet
	errs *errors.Errors

	stackRet int
	stackFor int
	tableFor string
}

func NewEvaluator(fset *position.FileSet, errs *errors.Errors) *Evaluator {
	return &Evaluator{
		fset: fset,
		errs: errs,
	}
}

func (s *Evaluator) errorsPos(pos position.Pos, err error) {
	s.errs.Append(s.fset.Position(pos), err)
}

func (ev *Evaluator) Eval(es []ast.Expr) value.Value {
	s := value.NewScope(nil)
	return ev.EvalBy(es, s)
}

func (ev *Evaluator) EvalBy(es []ast.Expr, s *value.Scope) (ex value.Value) {
	sr := ev.stackRet
	sf := ev.stackFor
	for _, v := range es {
		ex = ev.eval(v, s)
		if ev.stackRet < sr || ev.stackFor < sf {
			return
		}
	}
	return
}

func (ev *Evaluator) eval(e ast.Expr, s *value.Scope) value.Value {
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
	case *ast.Break:
		ev.stackFor--
		return value.ValueNil
	case *ast.For:
		return ev.evalFor(t, s)
	case *ast.Brace:
		ss := s.NewChildScope()
		return ev.EvalBy(t.List, ss)
	case *ast.Call:
		return ev.evalCall(t, s)
	case *ast.Func:
		return ev.evalFunc(t, s)
	case *ast.Return:
		return ev.evalReturn(t, s)
	case *ast.Tuple:
		return ev.evalTuple(t, s)
	}
	if e == nil {
		return value.ValueNil
	}
	ev.errorsPos(e.GetPos(), fmt.Errorf("Undefined keyword processing"))
	return value.ValueNil
}

func (ev *Evaluator) toValues(e ast.Expr, s *value.Scope) value.Value {
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *ast.Tuple:
		vs := make([]value.Value, 0, len(t.List))
		for _, v := range t.List {
			vs = append(vs, ev.eval(v, s))
		}
		return &value.ValueTuple{vs, false}
	}
	return ev.eval(e, s)
}
