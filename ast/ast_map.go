package ast

import (
	"github.com/wzshiming/gs/position"
)

// Map the is a map expression.
// Used to define map objects
type Map struct {
	position.Pos
	Body Expr
}
