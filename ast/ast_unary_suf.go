package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// 后缀一元表达式
type UnarySuf struct {
	X   Expr
	Pos position.Pos
	Op  token.Token
}

func (o *UnarySuf) String() string {
	if o == nil {
		return "<nil.UnarySuf>"
	}
	return fmt.Sprintf("%s%s ", o.X, o.Op)
}
