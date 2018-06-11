package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// UnarySuf the is a unary suf expression.
type UnarySuf struct {
	position.Pos
	X  Expr
	Op token.Token
}

func (o *UnarySuf) String() string {
	if o == nil {
		return "<nil.UnarySuf>"
	}
	return fmt.Sprintf("%s%s ", o.X, o.Op)
}
