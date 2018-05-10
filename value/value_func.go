package value

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

type Func struct {
	Args  ast.Expr
	Scope *Scope
	Body  ast.Expr
}

func (v *Func) String() string {
	return v.Body.String()
}

func (v *Func) Point() Value {
	return v
}

func (v *Func) Binary(t token.Token, y Value) (Value, error) {
	return Nil, undefined
}

func (v *Func) UnaryPre(t token.Token) (Value, error) {
	return Nil, undefined
}

func (v *Func) UnarySuf(t token.Token) (Value, error) {
	return Nil, undefined
}
