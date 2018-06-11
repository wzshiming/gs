package ast

import (
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// Literal the is a literal expression.
// Used to define base types
type Literal struct {
	position.Pos
	Type  token.Token
	Value string
}

func (l *Literal) String() string {
	if l == nil {
		return "<nil.Literal>"
	}
	return l.Value
}
