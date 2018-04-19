package eval

import (
	"testing"

	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
	ffmt "gopkg.in/ffmt.v1"
)

var defPuts = ffmt.NewOptional(5, ffmt.StyleP, ffmt.CanFilterDuplicate|ffmt.CanRowSpan)

func TestA(t *testing.T) {

	expr := `
a =	2 ** 10

b = a+a
b
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

	ev := NewEvaluator(fset, errs, out)
	ex := ev.Eval()

	if errs.Len() != 0 {
		ffmt.Puts(errs)
	}
	ffmt.Puts(ex)
}