package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// () 元组表达式
type Tuple struct {
	position.Pos
	List []Expr
}

func (l *Tuple) String() string {
	if l == nil {
		return "<nil.Tuple>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('(')

	for k, v := range l.List {
		if k != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteString(")")
	return buf.String()
}
