package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Brack the is a brack expression.
// Used to fetch indexes or slices.
type Brack struct {
	position.Pos
	X Expr
	Y Expr
}

func (l *Brack) String() string {
	if l == nil {
		return "<nil.Brace>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(l.X.String())
	buf.WriteString("[")
	buf.WriteString(l.Y.String())
	buf.WriteString("]")
	return buf.String()
}
