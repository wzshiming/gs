package parser

import (
	"testing"

	"github.com/wzshiming/gs/position"
	ffmt "gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {

	expr := `
{
a +
if 1* -2 {
  3 + 4 
  44 +11
} else if 3 ** 4 {
  aa+1	
} else b+1

 a+ "123"
}

`

	fset := position.NewFileSet()
	scan := NewParser(fset, "_", []rune(expr))
	out := scan.ParseExpr()
	ffmt.Puts(out)

}
