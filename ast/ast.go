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

// 前缀一元表达式
type OperatorPreUnary struct {
	Pos position.Pos
	Op  token.Token
	X   Expr
}

func (o *OperatorPreUnary) String() string {
	if o == nil {
		return "<nil.OperatorPreUnary>"
	}
	return fmt.Sprintf(" %s%s", o.Op, o.X)
}

// 后缀一元表达式
type OperatorSufUnary struct {
	X   Expr
	Pos position.Pos
	Op  token.Token
}

func (o *OperatorSufUnary) String() string {
	if o == nil {
		return "<nil.OperatorSufUnary>"
	}
	return fmt.Sprintf("%s%s ", o.X, o.Op)
}

// 二元表达式
type OperatorBinary struct {
	Pos position.Pos
	X   Expr
	Op  token.Token
	Y   Expr
}

func (o *OperatorBinary) String() string {
	if o == nil {
		return "<nil.OperatorBinary>"
	}
	return fmt.Sprintf("(%s %s %s)", o.X, o.Op, o.Y)
}

// 常量和符号
type Literal struct {
	Pos   position.Pos
	Type  token.Token
	Value string
}

func (l *Literal) String() string {
	if l == nil {
		return "<nil.Literal>"
	}
	return l.Value
}

// return 定义
type ReturnExpr struct {
	Pos position.Pos
	Ret Expr
}

func (l *ReturnExpr) String() string {
	if l == nil {
		return "<nil.ReturnExpr>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("return ")
	buf.WriteString(l.Ret.String())
	return buf.String()
}

// 函数定义
type FuncExpr struct {
	Pos  position.Pos
	Func Expr
	Body Expr
}

func (l *FuncExpr) String() string {
	if l == nil {
		return "<nil.FuncExpr>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("func ")

	if l.Func != nil {
		buf.WriteString(l.Func.String())
	} else {
		buf.WriteString("()")
	}
	buf.WriteString(" ")
	if l.Body != nil {
		buf.WriteString(l.Body.String())
	} else {
		buf.WriteString("{}")
	}
	return buf.String()
}

// call
type CallExpr struct {
	Pos  position.Pos
	Name Expr
	Args Expr
}

func (l *CallExpr) String() string {
	if l == nil {
		return "<nil.CallExpr>"
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

// if 关键字
type IfExpr struct {
	Pos  position.Pos
	Cond Expr
	Body Expr
	Else Expr
}

func (l *IfExpr) String() string {
	if l == nil {
		return "<nil.IfExpr>"
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

// () 元组表达式
type TupleExpr struct {
	Pos  position.Pos
	List []Expr
}

func (l *TupleExpr) String() string {
	if l == nil {
		return "<nil.TupleExpr>"
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

// {} 花括号表达式
type BraceExpr struct {
	Pos  position.Pos
	List []Expr
}

func (l *BraceExpr) String() string {
	if l == nil {
		return "<nil.BraceExpr>"
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
