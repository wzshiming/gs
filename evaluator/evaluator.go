package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/value"
)

// Evaluator 执行 ast
type Evaluator struct {
	fset *position.FileSet
	errs *errors.Errors

	stackRet int
	stackFor int
	tableFor string
}

// NewEvaluator 新的执行
func NewEvaluator(fset *position.FileSet, errs *errors.Errors) *Evaluator {
	return &Evaluator{
		fset: fset,
		errs: errs,
	}
}

func (ev *Evaluator) errorsPos(pos position.Pos, err error) {
	ev.errs.Append(ev.fset.Position(pos), err)
}

// EvalBy 执行 expr 指定作用域
func (ev *Evaluator) EvalBy(es []ast.Expr, s value.Assigner) (ex value.Value) {
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
	case *ast.Break:
		ev.stackFor--
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
	if e == nil {
		return value.Nil
	}
	ev.errorsPos(e.GetPos(), fmt.Errorf("Undefined keyword processing"))
	return value.Nil
}

func (ev *Evaluator) toValues(e ast.Expr, s value.Assigner) value.Value {
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *ast.Tuple:
		vs := make([]value.Value, 0, len(t.List))
		for _, v := range t.List {
			vs = append(vs, ev.eval(v, s))
		}
		return value.NewTuple(vs, false)
	}
	return ev.eval(e, s)
}
