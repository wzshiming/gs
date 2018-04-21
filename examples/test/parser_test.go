package test

import (
	"testing"

	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
	ffmt "gopkg.in/ffmt.v1"
)

var defPuts = ffmt.NewOptional(5, ffmt.StyleP, ffmt.CanFilterDuplicate|ffmt.CanRowSpan)

func TestParser(t *testing.T) {

	expr := `
func echo.Hello a {
	b = 0
	for i = 0; i != a ;i++ {
		b = i
	} else {
		return nil
	}
	return b
}
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
		//defPuts.Print(v)
	}

}
