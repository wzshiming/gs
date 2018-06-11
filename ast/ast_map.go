package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Map the is a map expression.
// Used to define map objects
type Map struct {
	position.Pos
	Body Expr
}

func (l *Map) String() string {
	if l == nil {
		return "<nil.Map>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("map ")
	buf.WriteString(l.Body.String())
	return buf.String()
}
