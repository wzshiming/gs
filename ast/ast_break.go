package ast

import (
	"github.com/wzshiming/gs/position"
)

// break 关键字
type Break struct {
	position.Pos
}

func (l *Break) String() string {
	if l == nil {
		return "<nil.Break>"
	}
	return "break"
}
