package ast

import (
	"bytes"

	"github.com/wzshiming/gs/position"
)

// return 定义
type Return struct {
	position.Pos
	Ret Expr
}

func (l *Return) String() string {
	if l == nil {
		return "<nil.Return>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("return ")
	buf.WriteString(l.Ret.String())
	return buf.String()
}
