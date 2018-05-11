package test

import (
	"testing"

	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/evaluator"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/value"
)

func BenchmarkGoFloat(b *testing.B) {

	for i := 0; i != b.N; i++ {
		(func() {
			sum := 0.0
			for i := 0.0; i < 100000.0; i++ {
				sum += i + 1.0
			}
		})()
	}
}

func BenchmarkGoInt(b *testing.B) {

	for i := 0; i != b.N; i++ {
		(func() {
			sum := 0
			for i := 0; i < 100000; i++ {
				sum += i + 1
			}
		})()
	}
}

func BenchmarkGs1(b *testing.B) {

	expr := `
sum := 0
for i := 0; i < 100000; i ++ {
	sum += i + 1
}
`

	fset := position.NewFileSet()
	errs := errors.NewErrors()
	scan := parser.NewParser(fset, errs, "_", []rune(expr))
	scope := value.NewScope(nil)
	out := scan.Parse()
	if errs.Len() != 0 {
		b.Error(errs)
	}

	ev := evaluator.NewEvaluator(fset, errs)

	for i := 0; i != b.N; i++ {
		ev.EvalBy(out, scope)
		if errs.Len() != 0 {
			b.Error(errs)
		}
	}
}
