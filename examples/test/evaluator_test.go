package test

import (
	"testing"

	"github.com/wzshiming/gs/exec"
)

func TestEvaluator(t *testing.T) {

	expr := `
println 10,2
`
	exe := exec.NewExec()
	val, err := exe.Cmd("_", []rune(expr))
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
