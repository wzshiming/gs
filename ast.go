package gs

import (
	"bytes"
	"fmt"
)

type Expr interface {
	String() string
}

type OperatorUnary struct {
	Op Token
	X  Expr
}

func (o *OperatorUnary) String() string {
	return fmt.Sprintf("%s%s", o.Op, o.X)
}

type OperatorBinary struct {
	X  Expr
	Op Token
	Y  Expr
}

func (o *OperatorBinary) String() string {
	return fmt.Sprintf("(%s %s %s)", o.X, o.Op, o.Y)
}

type Literal struct {
	Type  Token
	Value string
}

func (l *Literal) String() string {
	return l.Value
}

// if 关键字
type IfExpr struct {
	Cond Expr
	Body Expr
	Else Expr
}

func (l *IfExpr) String() string {
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
	List []Expr
}

func (l *BraceExpr) String() string {
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
