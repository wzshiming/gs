package evaluator

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalLiteral(t *ast.Literal, s value.Assigner) value.Value {
	switch t.Type {
	case token.NUMBER:
		return value.ParseNumber(t.Value)
	case token.STRING:
		return value.String(t.Value[1 : len(t.Value)-1])
	case token.BOOL:
		if t.Value == "true" {
			return value.True
		}
		return value.False
	case token.IDENT:
		return &value.Var{
			Name:  value.String(t.Value),
			Scope: s,
		}
	case token.NIL:
		return value.Nil
	default:
		ev.errorsPos(t.Pos, fmt.Errorf("Panic "+t.Type.String()))
		return value.Nil
	}
}
