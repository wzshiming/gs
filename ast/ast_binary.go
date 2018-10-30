package ast

import (
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
