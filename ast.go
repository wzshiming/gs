package gs

import "fmt"

type Expr interface {
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
	return fmt.Sprintf("(%s%s%s)", o.X, o.Op, o.Y)
}

type Literal struct {
	Type  Token
	Value string
}

func (l *Literal) String() string {
	return l.Value
}
