package ast

import (
	"github.com/wzshiming/gs/position"
)

// If the is a if expression.
// Used to logical judgment
type If struct {
	position.Pos
	Init Expr
	Cond Expr
	Body Expr
	Else Expr
}
