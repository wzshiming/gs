package ast

import (
	"github.com/wzshiming/gs/position"
)

// Labeled the is loop control statement.
type Labeled struct {
	position.Pos
	Label *Literal
	Stmt  Expr
}
