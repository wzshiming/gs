package ast

import (
	"github.com/wzshiming/gs/position"
)

// Expr is expression.
// All syntax has to implement the expression.
type Expr interface {
	GetPos() position.Pos
}
