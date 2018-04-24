package test

import (
	"testing"

	"github.com/wzshiming/gs/exec"
	ffmt "gopkg.in/ffmt.v1"
)

func TestEvaluator(t *testing.T) {

	expr := `
println 10,2
`
	exe := exec.NewExec()
	val, err := exe.Cmd("_", []rune(expr))
	if err != nil {
		ffmt.Mark(err)
	}
	ffmt.Mark(val)
}
