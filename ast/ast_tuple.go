package ast

import (
	"github.com/wzshiming/gs/position"
)

// Tuple the is a tuple expression.
// Used to define tuple types.
type Tuple struct {
	position.Pos
	List []Expr
}
