package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalLiteral(t *ast.Literal, s *value.Scope) value.Value {
	switch t.Type {
	case token.NUMBER:
		return value.ParseValueNumber(t.Value)
	case token.STRING:
		return value.String(t.Value[1 : len(t.Value)-1])
	case token.BOOL:
		if t.Value == "true" {
			return value.ValueTrue
		}
		return value.ValueFalse
	case token.IDENT:
		return &value.Var{
			Name:  t.Value,
			Scope: s,
		}
	case token.NIL:
		return value.Nil
	default:
		ev.errorsPos(t.Pos, fmt.Errorf("Panic "+t.Type.String()))
		return value.Nil
	}
}
