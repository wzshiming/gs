package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Func the is a func expression.
// Used to define functions or methods.
type Func struct {
	position.Pos
	Func Expr
	Body Expr
}

func (l *Func) String() string {
	if l == nil {
		return "<nil.Func>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("func ")

	if l.Func != nil {
		buf.WriteString(l.Func.String())
	} else {
		buf.WriteString("()")
	}
	buf.WriteString(" ")
	if l.Body != nil {
		buf.WriteString(l.Body.String())
	} else {
		buf.WriteString("{}")
	}
	return buf.String()
}
