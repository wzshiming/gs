package ast

import (
	"bytes"
	"fmt"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

type Expr interface {
	String() string
}

type OperatorUnary struct {
	Pos position.Pos
	Op  token.Token
	X   Expr
}

func (o *OperatorUnary) String() string {
	if o == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s%s", o.Op, o.X)
}

type OperatorBinary struct {
	Pos position.Pos
	X   Expr
	Op  token.Token
	Y   Expr
}

func (o *OperatorBinary) String() string {
	if o == nil {
		return "<nil>"
	}
	return fmt.Sprintf("(%s %s %s)", o.X, o.Op, o.Y)
}

type Literal struct {
	Pos   position.Pos
	Type  token.Token
	Value string
}

func (l *Literal) String() string {
	if l == nil {
		return "<nil>"
	}
	return l.Value
}

// if 关键字
type IfExpr struct {
	Pos  position.Pos
	Cond Expr
	Body Expr
	Else Expr
}

func (l *IfExpr) String() string {
	if l == nil {
		return "<nil>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("if ")
	buf.WriteString(l.Cond.String())
	buf.WriteByte(' ')
	buf.WriteString(l.Body.String())
	if l.Else != nil {
		buf.WriteString("else ")
		buf.WriteString(l.Else.String())
	}
	return buf.String()
}

// {} 花括号表达式
type BraceExpr struct {
	Pos  position.Pos
	List []Expr
}

func (l *BraceExpr) String() string {
	if l == nil {
		return "<nil>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('{')

	for _, v := range l.List {
		buf.WriteString("\n  ")
		buf.WriteString(v.String())
	}
	buf.WriteByte('\n')
	buf.WriteString("} ")
	return buf.String()
}
