package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Call the is a call expression.
type Call struct {
	position.Pos
	Name Expr
	Args Expr
}

func (l *Call) String() string {
	if l == nil {
		return "<nil.Call>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(l.Name.String())
	buf.WriteString(" ")
	if l.Args != nil {
		buf.WriteString(l.Args.String())
	} else {
		buf.WriteString("()")
	}
	return buf.String()
}
