package ast

import (
	"github.com/wzshiming/gs/position"
)

// Brack the is a brack expression.
// Used to fetch indexes or slices.
type Brack struct {
	position.Pos
	X Expr
	Y Expr
}
