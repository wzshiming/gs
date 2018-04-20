package ast

import (
	"github.com/wzshiming/gs/position"
)

// continue 关键字
type Continue struct {
	position.Pos
}

func (l *Continue) String() string {
	if l == nil {
		return "<nil.Continue>"
	}
	return "continue"
}
