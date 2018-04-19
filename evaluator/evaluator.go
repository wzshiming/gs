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
	es   []ast.Expr
}

func NewEvaluator(fset *position.FileSet, errs *errors.Errors, es []ast.Expr) *Evaluator {
	return &Evaluator{
		fset: fset,
		errs: errs,
		es:   es,
	}
}

func (s *Evaluator) errorsPos(pos position.Pos, err error) {
	s.errs.Append(s.fset.Position(pos), err)
}

func (ev *Evaluator) Eval() ast.Expr {
	s := value.NewScope(nil)
	return ev.EvalBy(s)
}

func (ev *Evaluator) EvalBy(s *value.Scope) (ex ast.Expr) {
	for _, v := range ev.es {
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
			vv = &value.ValueBool{t.Value == "true"}
		case token.NIL:
			vv = &value.ValueNil{}
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
	}

	return e
}
