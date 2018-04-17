package gs

import (
	"testing"

	ffmt "gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {
	expr := `
-1 + a.b - - 1 + 1
b+2
`
	scan := NewParser(expr)
	out := scan.ParseExprs()
	ffmt.Puts(out)
}
