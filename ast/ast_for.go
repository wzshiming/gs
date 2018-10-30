package ast

import (
	"github.com/wzshiming/gs/position"
)

// For the is a for expression.
// Used for circulation
type For struct {
	position.Pos
	Init Expr
	Cond Expr
	Next Expr
	Body Expr
	Else Expr
}
