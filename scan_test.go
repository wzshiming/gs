package gs

import (
	"testing"
	ffmt "gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {

	expr := "-1+1-1+1"
	scan := NewScan(expr)
	out := scan.Scan()
	ffmt.Puts(out)
}
