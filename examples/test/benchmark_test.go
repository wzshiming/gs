package test

import (
	"testing"

	"github.com/wzshiming/gs/exec"
)

func BenchmarkGoFloat(b *testing.B) {

	for i := 0; i != b.N; i++ {
		(func() {
			sum := 0.0
			for i := 0.0; i < 10000.0; i++ {
				sum += i + 1.0
			}
		})()
	}
}

func BenchmarkGoInt(b *testing.B) {

	for i := 0; i != b.N; i++ {
		(func() {
			sum := 0
			for i := 0; i < 10000; i++ {
				sum += i + 1
			}
		})()
	}
}

func BenchmarkGsFloat(b *testing.B) {

	expr := []rune(`
sum := 0.0
for i := 0.0; i < 10000.0; i ++ {
	sum += i + 1
}
`)
	exe := exec.NewExec()

	for i := 0; i != b.N; i++ {
		exe.Cmd("_", expr)
	}
}

func BenchmarkGsInt(b *testing.B) {

	expr := []rune(`
sum := 0
for i := 0; i < 10000; i ++ {
	sum += i + 1
}
`)
	exe := exec.NewExec()

	for i := 0; i != b.N; i++ {
		exe.Cmd("_", expr)
	}
}
