package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
)

// Expr is expression.
// All syntax has to implement the expression.
type Expr interface {
	fmt.Stringer
	GetPos() position.Pos
}
