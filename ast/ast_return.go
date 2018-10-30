package ast

import (
	"github.com/wzshiming/gs/position"
)

// Return the is a reurn expression.
// Used to end the function and return the result.
type Return struct {
	position.Pos
	Ret Expr
}
