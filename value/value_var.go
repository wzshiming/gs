package value

import (
	"github.com/wzshiming/gs/token"
)

type Var struct {
	Scope    Assigner
	Ellipsis bool
	Name     Value
}

func (v *Var) String() string {
	if v == nil {
		return "<ValueVar.nil>"
	}
	val := v.Point()
	return "<" + v.Name.String() + "." + val.String() + ">"
}

func (v *Var) Point() Value {
	val, ok := v.Scope.Get(v.Name)
	if !ok {
		return Nil
	}
	return val
}

func (v *Var) Binary(t token.Token, y Value) (Value, error) {
	switch t {
	case token.ASSIGN:
		yy := y.Point()
		v.Scope.Set(v.Name, yy)
		return v, nil
	case token.DEFINE:
		yy := y.Point()
		v.Scope.SetLocal(v.Name, yy)
		return v, nil
	case token.ADD_ASSIGN, token.SUB_ASSIGN, token.MUL_ASSIGN, token.QUO_ASSIGN, token.POW_ASSIGN, token.REM_ASSIGN,
		token.AND_ASSIGN, token.OR_ASSIGN, token.XOR_ASSIGN, token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN:
		t0 := t - (token.ADD_ASSIGN - token.ADD)
		val := v.Point()
		val, err := val.Binary(t0, y)
		if err != nil {
			return Nil, err
		}
		v.Scope.Set(v.Name, val)
		return v, nil
	default:
		val := v.Point()
		return val.Binary(t, y)
	}
}

func (v *Var) UnaryPre(t token.Token) (Value, error) {
	switch t {
	case token.ELLIPSIS:
		v.Ellipsis = true
		return v, nil
	}

	val := v.Point()
	return val.UnaryPre(t)
}

func (v *Var) UnarySuf(t token.Token) (Value, error) {
	val := v.Point()

	switch t {
	case token.INC, token.DEC:
		vv, err := val.UnarySuf(t)
		if err != nil {
			return nil, err
		}
		v.Scope.Set(v.Name, vv)
		return v, nil
	}

	return val.UnarySuf(t)
}
