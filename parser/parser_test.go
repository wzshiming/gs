package parser

import (
	"testing"

	"github.com/wzshiming/gs/position"
	ffmt "gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {

	expr := `
func hello.String(a, b ) {
	s = (hello,a,b).String 1
	return s
}

`

	fset := position.NewFileSet()
	scan := NewParser(fset, "_", []rune(expr))
	out := scan.Parse()
	if scan.Err() != nil {
		ffmt.Puts(scan.Err())
	}
	for _, v := range out {
		ffmt.Puts(v)
	}

}
