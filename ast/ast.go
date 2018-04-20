package ast

import (
	"fmt"

	"github.com/wzshiming/gs/position"
)

type Expr interface {
	fmt.Stringer
	GetPos() position.Pos
}
