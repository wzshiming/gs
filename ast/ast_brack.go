package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// [] 方括号表达式
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
