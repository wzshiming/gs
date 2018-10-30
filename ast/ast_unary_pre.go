package ast

import (
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// UnaryPre the is a unary pre expression.
type UnaryPre struct {
	position.Pos
	Op token.Token
	X  Expr
}
