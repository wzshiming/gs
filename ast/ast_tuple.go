package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Tuple the is a tuple expression.
// Used to define tuple types.
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
