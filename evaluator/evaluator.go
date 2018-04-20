package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
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
		var vv value.Value
		switch t.Type {
		case token.NUMBER:
			vv = value.ParseValueNumber(t.Value)
		case token.STRING:
			vv = value.ValueString(t.Value[1 : len(t.Value)-1])
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
		lx := ev.eval(t.X, s)
		ly := ev.eval(t.Y, s)

		z, err := lx.Binary(t.Op, ly)
		if err != nil {
			ev.errorsPos(t.Pos, err)
			return value.ValueNil
		}
		return z
	case *ast.UnaryPre:
		lx := ev.eval(t.X, s)

		z, err := lx.UnaryPre(t.Op)
		if err != nil {
			ev.errorsPos(t.Pos, err)
			return value.ValueNil
		}
		return z
	case *ast.UnarySuf:
		lx := ev.eval(t.X, s)

		z, err := lx.UnarySuf(t.Op)
		if err != nil {
			ev.errorsPos(t.Pos, err)
			return value.ValueNil
		}
		return z
	case *ast.If:
		ss := s.NewChildScope()
		ev.eval(t.Init, ss)
		loop := ev.eval(t.Cond, ss)
		vb, ok := loop.(value.ValueBool)
		if !ok {
			ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'if'."))
			return value.ValueNil
		}

		if vb {
			return ev.eval(t.Body, ss)
		} else if t.Else != nil {
			return ev.eval(t.Else, ss)
		}
		return value.ValueNil
	case *ast.Break:
		ev.stackFor--
		return value.ValueNil
	case *ast.For:
		ss := s.NewChildScope()
		ev.eval(t.Init, ss)
		i := 0

		var ex value.Value
		for {
			loop := ev.eval(t.Cond, ss)
			vb, ok := loop.(value.ValueBool)
			if !ok {
				ev.errorsPos(t.Pos, fmt.Errorf("There are only Boolean values in the 'for'."))
				break
			}

			if !vb {
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

				ti1 := ev.toValues(t.Args, s)

				ti2 := ev.toIdents(t2.Args)

				ss := t2.Scope.NewChildScope()
				for i := 0; i != len(ti1) && i != len(ti2); i++ {
					ss.SetLocal(ti2[i].Value, ti1[i])
				}
				ev.stackRet++
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
	case *ast.Return:
		ev.stackRet--
		if t.Ret == nil {
			return value.ValueNil
		}
		return ev.eval(t.Ret, s)
	}
	if e == nil {
		return value.ValueNil
	}
	ev.errorsPos(e.GetPos(), fmt.Errorf("未定义关键字处理"))
	return value.ValueNil
}

func (ev *Evaluator) toValues(e ast.Expr, s *value.Scope) []value.Value {
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *ast.Tuple:
		vs := make([]value.Value, 0, len(t.List))
		for _, v := range t.List {
			vs = append(vs, ev.eval(v, s))
		}
		return vs
	}
	return []value.Value{ev.eval(e, s)}
}

func (ev *Evaluator) toIdents(e ast.Expr) []*ast.Literal {
	if e == nil {
		return nil
	}
	switch t := e.(type) {
	case *ast.Literal:
		return []*ast.Literal{t}
	case *ast.Tuple:
		at := make([]*ast.Literal, 0, len(t.List))
		for _, v := range t.List {
			v2, ok := v.(*ast.Literal)
			if !ok {
				ev.errorsPos(t.Pos, fmt.Errorf("toIdentList not is ident error"))
			}
			at = append(at, v2)
		}
		return at
	}

	ev.errorsPos(e.GetPos(), fmt.Errorf("toIdentList type error"))
	return nil
}
