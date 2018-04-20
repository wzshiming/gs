package value

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

type ValueFunc struct {
	Args  ast.Expr
	Scope *Scope
	Body  ast.Expr
}

func (v *ValueFunc) String() string {
	return v.Body.String()
}

func (v *ValueFunc) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v *ValueFunc) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueFunc) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}
