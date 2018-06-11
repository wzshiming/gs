package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Brace the is a brace expression.
// Used for function body or map definition or structure body definition.
type Brace struct {
	position.Pos
	List []Expr
}

func (l *Brace) String() string {
	if l == nil {
		return "<nil.Brace>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('{')

	for _, v := range l.List {
		buf.WriteString("\n  ")
		buf.WriteString(v.String())
	}
	buf.WriteString("\n}\n")
	return buf.String()
}
