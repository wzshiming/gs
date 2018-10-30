package ast

import (
	"github.com/wzshiming/gs/position"
)

// Break the is a brack expression.
// Used to jump out of the loop
type Break struct {
	position.Pos
	Label *Literal
}

func (l *Break) String() string {
	if l == nil {
		return "<nil.Break>"
	}
	return "break"
}
