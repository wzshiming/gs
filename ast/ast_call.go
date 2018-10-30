package ast

import (
	"github.com/wzshiming/gs/position"
)

// Call the is a call expression.
type Call struct {
	position.Pos
	Name Expr
	Args Expr
}
