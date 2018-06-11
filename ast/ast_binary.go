package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// Binary the is a binary expression.
// Used for mathematical
type Binary struct {
	position.Pos
	X  Expr
	Op token.Token
	Y  Expr
}

func (o *Binary) String() string {
	if o == nil {
		return "<nil.Binary>"
	}
	return fmt.Sprintf("(%s %s %s)", o.X, o.Op, o.Y)
}
