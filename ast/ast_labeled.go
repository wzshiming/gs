package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// Labeled the is loop control statement.
type Labeled struct {
	position.Pos
	Label *Literal
	Stmt  Expr
}

func (l *Labeled) String() string {
	if l == nil {
		return "<nil.Labeled>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(l.Label.String())
	buf.WriteString(": ")
	buf.WriteString(l.Stmt.String())
	return buf.String()
}
