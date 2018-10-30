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
