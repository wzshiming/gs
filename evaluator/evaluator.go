package eval

import (
	"fmt"
	"strconv"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

type Evaluator struct {
	fset *position.FileSet
	errs *errors.Errors
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

func (ev *Evaluator) Eval(es []ast.Expr) ast.Expr {
	s := value.NewScope(nil)
	return ev.EvalBy(es, s)
}

func (ev *Evaluator) EvalBy(es []ast.Expr, s *value.Scope) (ex ast.Expr) {
	for _, v := range es {
		ex = ev.eval(v, s)
	}
	return
}

func (ev *Evaluator) eval(e ast.Expr, s *value.Scope) ast.Expr {
	switch t := e.(type) {
	case *value.ValueVar:
		return t
	case *ast.Literal:
		var vv value.Value
		switch t.Type {
		case token.NUMBER:
			val, _ := strconv.ParseFloat(t.Value, 0)
			vv = &value.ValueNumber{val}
		case token.STRING:
			vv = &value.ValueString{t.Value[1 : len(t.Value)-1]}
		case token.BOOL:
			if t.Value == "true" {
				vv = value.ValueTrue
			} else {
				vv = value.ValueFalse
			}
		case token.NIL:
			vv = value.ValueNil
		case token.IDENT:
			vv = &value.ValueVar{
				Name:  t.Value,
				Scope: s,
			}
		default:
			ev.errorsPos(t.Pos, fmt.Errorf("Panic "+t.Type.String()))
		}
		return vv
	case *ast.OperatorBinary:
		x := ev.eval(t.X, s)
		y := ev.eval(t.Y, s)
		lx, ok := x.(value.Value)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("Execution error %v", x))
			return e
		}
		ly, ok := y.(value.Value)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("Execution error %v", y))
			return e
		}

		z, err := lx.Binary(t.Op, ly)
		if err != nil {
			ev.errorsPos(t.Pos, err)
			return e
		}

		return z
	case *ast.IfExpr:
		ss := s.NewChildScope()
		ev.eval(t.Init, ss)
		loop := ev.eval(t.Cond, ss)
		vb, ok := loop.(*value.ValueBool)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'if'."))
		}

		if vb.Val {
			return ev.eval(t.Body, ss)
		} else if t.Else != nil {
			return ev.eval(t.Else, ss)
		}
		return value.ValueNil
	case *ast.BraceExpr:
		ss := s.NewChildScope()
		return ev.EvalBy(t.List, ss)
	}

	return e
}
