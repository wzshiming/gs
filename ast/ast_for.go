package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// for 关键字
type For struct {
	Pos  position.Pos
	Init Expr
	Cond Expr
	Next Expr
	Body Expr
	Else Expr
}

func (l *For) String() string {
	if l == nil {
		return "<nil.For>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("for ")

	if l.Init != nil {
		buf.WriteString(l.Init.String())
	}
	if l.Init != nil || l.Next != nil {
		buf.WriteString("; ")
	}
	buf.WriteString(l.Cond.String())

	if l.Init != nil || l.Next != nil {
		buf.WriteString("; ")
	}
	if l.Next != nil {
		buf.WriteString(l.Next.String())
	}
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
