package ast

import (
	"github.com/wzshiming/gs/position"
)

// Brace the is a brace expression.
// Used for function body or map definition or structure body definition.
type Brace struct {
	position.Pos
	List []Expr
}
