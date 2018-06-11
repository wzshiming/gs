package ast

import (
	"bytes"

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

func (l *If) String() string {
	if l == nil {
		return "<nil.If>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("if ")
	if l.Init != nil {
		buf.WriteString(l.Init.String())
		buf.WriteString("; ")
	}
	buf.WriteString(l.Cond.String())
	buf.WriteByte(' ')
	if l.Body != nil {
		buf.WriteString(l.Body.String())
	}
	if l.Else != nil {
		buf.WriteString("else ")
		buf.WriteString(l.Else.String())
	}
	return buf.String()
}
