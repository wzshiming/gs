package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// UnaryPre the is a unary pre expression.
type UnaryPre struct {
	position.Pos
	Op token.Token
	X  Expr
}

func (o *UnaryPre) String() string {
	if o == nil {
		return "<nil.UnaryPre>"
	}
	return fmt.Sprintf(" %s%s", o.Op, o.X)
}
