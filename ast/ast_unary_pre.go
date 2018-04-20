package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// 前缀一元表达式
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
