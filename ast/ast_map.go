package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// map 关键字
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
