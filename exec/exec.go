package exec

import (
	"io/ioutil"

	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/evaluator"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/value"
)

type Exec struct {
	fset  *position.FileSet
	errs  *errors.Errors
	scope *value.Scope
}

func NewExec() *Exec {
	return &Exec{
		fset:  position.NewFileSet(),
		errs:  errors.NewErrors(),
		scope: value.NewScope(nil),
	}
}

func (e *Exec) Cmd(name string, expr []rune) (value.Value, error) {
	par := parser.NewParser(e.fset, e.errs, name, expr)
	exprs := par.Parse()
	if e.errs.Len() != 0 {
		return nil, e.errs
	}
	eval := evaluator.NewEvaluator(e.fset, e.errs)
	ret := eval.EvalBy(exprs, e.scope)
	if e.errs.Len() != 0 {
		return nil, e.errs
	}
	return ret, nil
}

func (e *Exec) File(filename string) (value.Value, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return e.Cmd(filename, []rune(string(b)))
}
