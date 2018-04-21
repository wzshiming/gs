package test

import (
	"testing"

	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/evaluator"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
	ffmt "gopkg.in/ffmt.v1"
)

func TestEvaluator(t *testing.T) {

	expr := `
func T i, j {
	if j >= 2 {
		return T i*i,j-1
	}
	return i
}

T 2,4
`

	fset := position.NewFileSet()
	errs := errors.NewErrors()
	scan := parser.NewParser(fset, errs, "_", []rune(expr))
	out := scan.Parse()
	if errs.Len() != 0 {
		ffmt.Puts(errs)
	}
	for _, v := range out {
		ffmt.Puts(v)
	}

	ev := evaluator.NewEvaluator(fset, errs)
	ex := ev.Eval(out)

	if errs.Len() != 0 {
		ffmt.Puts(errs)
	}
	ffmt.Puts(ex)
}
