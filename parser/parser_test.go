package parser

import (
	"testing"

	"github.com/wzshiming/gs/position"
	ffmt "gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {

	expr := `
aaa 1+2**3,2+3
a.b aa, bb + 1 , cc
if 1* -2-- -- {
  3 + 4
  44 +11
} else if 3 ** 4 {
  aa+1
} else ...b + .1

 a + "123"...

`

	fset := position.NewFileSet()
	scan := NewParser(fset, "_", []rune(expr))
	out := scan.Parse()
	for _, v := range out {
		ffmt.Puts(v)
	}

}
