package exec

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/builtin"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/evaluator"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/value"
)

type Exec struct {
	fset  *position.FileSet
	errs  *errors.Errors
	scope value.Assigner
}

func NewExec() *Exec {
	scope := value.NewScope(nil)
	for k, v := range builtin.Func {
		scope.SetLocal(value.String(k), value.NewFuncBuiltin(v))
	}
	return &Exec{
		fset:  position.NewFileSet(),
		errs:  errors.NewErrors(),
		scope: scope,
	}
}

func (e *Exec) Parse(name string, expr []rune) ([]ast.Expr, error) {
	par := parser.NewParser(e.fset, e.errs, name, expr)
	exprs := par.Parse()
	if e.errs.Len() != 0 {
		return nil, e.errs
	}
	return exprs, nil
}

func (e *Exec) Cmd(name string, expr []rune) (value.Value, error) {
	exprs, err := e.Parse(name, expr)
	if err != nil {
		return nil, err
	}
	eval := evaluator.NewEvaluator(e.fset, e.errs)
	ret := eval.EvalBy(exprs, e.scope)
	if e.errs.Len() != 0 {
		return nil, e.errs
	}
	return ret, nil
}
