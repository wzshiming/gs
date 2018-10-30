package ast

import (
	"github.com/wzshiming/gs/position"
)

// Continue the is a continue expression.
// Used to jump out of the current
type Continue struct {
	position.Pos
	Label *Literal
}

func (l *Continue) String() string {
	if l == nil {
		return "<nil.Continue>"
	}
	return "continue"
}
