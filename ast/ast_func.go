package ast

import (
	"github.com/wzshiming/gs/position"
)

// Func the is a func expression.
// Used to define functions or methods.
type Func struct {
	position.Pos
	Func Expr
	Body Expr
}
