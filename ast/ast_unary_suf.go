package ast

import (
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

// UnarySuf the is a unary suf expression.
type UnarySuf struct {
	position.Pos
	X  Expr
	Op token.Token
}
