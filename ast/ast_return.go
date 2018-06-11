package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Return the is a reurn expression.
// Used to end the function and return the result.
type Return struct {
	position.Pos
	Ret Expr
}

func (l *Return) String() string {
	if l == nil {
		return "<nil.Return>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("return ")
	buf.WriteString(l.Ret.String())
	return buf.String()
}
