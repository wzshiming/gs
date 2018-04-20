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
	case *ast.Binary:
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
	case *ast.If:
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
	case *ast.For:
		ss := s.NewChildScope()
		ev.eval(t.Init, ss)
		i := 0

		var ex ast.Expr
		for {
			loop := ev.eval(t.Cond, ss)
			vb, ok := loop.(*value.ValueBool)
			if !ok {
				ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'for'."))
				break
			}

			if !vb.Val {
				break
			}

			ex = ev.eval(t.Body, ss)
			if t.Next != nil {
				ev.eval(t.Next, ss)
			}
		}

		if i != 0 {
			return ex
		}

		if t.Else != nil {
			return ev.eval(t.Else, ss)
		}

		return value.ValueNil
	case *ast.Brace:
		ss := s.NewChildScope()
		return ev.EvalBy(t.List, ss)
	case *ast.Call:

		switch t1 := t.Name.(type) {
		case *ast.Literal: // name a,b
			val, ok := s.Get(t1.Value)
			if !ok {
				ev.errorsPos(t.Pos, fmt.Errorf("未定义"))
				break
			}
			switch t2 := val.(type) {
			case *value.ValueFunc:
				ss := t2.Scope.NewChildScope()
				ti1, err := toExprList(t.Args)
				if err != nil {
					ev.errorsPos(t.Pos, err)
					break
				}
				for i := 0; i != len(ti1); i++ {
					ti1[i] = ev.eval(ti1[i], ss)
				}
				ti2, err := toIdentList(t2.Args)
				if err != nil {
					ev.errorsPos(t.Pos, err)
					break
				}
				for i := 0; i != len(ti1) && i != len(ti2); i++ {
					vv := ti1[i].(value.Value)
					ss.SetLocal(ti2[i].Value, vv)
				}

				return ev.eval(t2.Body, ss)
			}

		default: // typ.name a,b
		}
	case *ast.Func:
		fun := t.Func
		//ss := s.NewChildScope()
		vf := &value.ValueFunc{
			Scope: s,
			Body:  t.Body,
		}
		switch t0 := fun.(type) {
		case *ast.Call: // func name a,b
			switch t1 := t0.Name.(type) {
			case *ast.Literal: // func name a,b
				s.SetLocal(t1.Value, vf)
				//t1.Value
			default: // func typ.name a,b
			}
			vf.Args = t0.Args
		default: // func a,b
		}

		return vf
	}

	return e
}

func toIdentList(e ast.Expr) ([]*ast.Literal, error) {
	if e == nil {
		return nil, nil
	}
	switch t := e.(type) {
	case *ast.Literal:
		return []*ast.Literal{t}, nil
	case *ast.Tuple:
		at := make([]*ast.Literal, 0, len(t.List))
		for _, v := range t.List {
			v2, ok := v.(*ast.Literal)
			if !ok {
				return nil, fmt.Errorf("toIdentList not is ident error")
			}
			at = append(at, v2)
		}
		return at, nil
	}
	return nil, fmt.Errorf("toIdentList error")
}

func toExprList(e ast.Expr) ([]ast.Expr, error) {
	if e == nil {
		return nil, nil
	}
	switch t := e.(type) {
	case *ast.Tuple:
		return t.List, nil
	}
	return []ast.Expr{e}, nil
}
