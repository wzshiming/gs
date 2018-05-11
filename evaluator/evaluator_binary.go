package evaluator

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
	"github.com/wzshiming/gs/value"
)

func (ev *Evaluator) evalBinary(t *ast.Binary, s value.Assigner) value.Value {
	lx := value.Nil
	switch t.Op {
	case token.ASSIGN, token.DEFINE,
		token.ADD_ASSIGN, token.SUB_ASSIGN, token.MUL_ASSIGN, token.QUO_ASSIGN, token.POW_ASSIGN, token.REM_ASSIGN,
		token.AND_ASSIGN, token.OR_ASSIGN, token.XOR_ASSIGN, token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN:
		lx = ev.evalVar(t.X, s)
	default:
		lx = ev.eval(t.X, s)
	}

	ly := ev.eval(t.Y, s)
	z, err := ev.evalBinaryBy(lx, ly, t.Op)
	if err != nil {
		ev.errorsPos(t.Pos, err)
		return value.Nil
	}
	return z
}

func (ev *Evaluator) evalBinaryBy(x, y value.Value, op token.Token) (value.Value, error) {
	z, err := x.Binary(op, y)
	if err != nil {
		return value.Nil, err
	}
	return z, nil
}
