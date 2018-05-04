package value

import (
	"github.com/wzshiming/gs/token"
)

type ValueVar struct {
	Scope    *Scope
	Ellipsis bool
	Name     string
}

func (v *ValueVar) String() string {
	if v == nil {
		return "<ValueVar.nil>"
	}
	val := v.Point()
	return "<" + v.Name + "." + val.String() + ">"
}

func (v *ValueVar) Point() Value {
	val, ok := v.Scope.Get(v.Name)
	if !ok {
		return ValueNil
	}
	return val
}

func (v *ValueVar) Binary(t token.Token, y Value) (Value, error) {

	switch t {
	case token.ASSIGN:
		yy := y.Point()
		v.Scope.Set(v.Name, yy)
		return v, nil
	case token.DEFINE, token.COLON:
		yy := y.Point()
		v.Scope.SetLocal(v.Name, yy)
		return v, nil
	}

	val := v.Point()

	switch t {
	case token.ADD_ASSIGN, token.SUB_ASSIGN, token.MUL_ASSIGN, token.QUO_ASSIGN, token.POW_ASSIGN, token.REM_ASSIGN,
		token.AND_ASSIGN, token.OR_ASSIGN, token.XOR_ASSIGN, token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN:
		t0 := t - (token.ADD_ASSIGN - token.ADD)
		val, err := val.Binary(t0, y)
		if err != nil {
			return ValueNil, err
		}
		v.Scope.Set(v.Name, val)
		return v, nil
	}
	return val.Binary(t, y)
}

func (v *ValueVar) UnaryPre(t token.Token) (Value, error) {
	switch t {
	case token.ELLIPSIS:
		v.Ellipsis = true
		return v, nil
	}

	val := v.Point()
	return val.UnaryPre(t)
}

func (v *ValueVar) UnarySuf(t token.Token) (Value, error) {
	val := v.Point()

	switch t {
	case token.INC:
		vv, err := val.Binary(token.ADD, newValueNumberBigInt(1))
		if err != nil {
			return nil, err
		}
		v.Scope.Set(v.Name, vv)
		return v, nil
	case token.DEC:
		vv, err := val.Binary(token.SUB, newValueNumberBigInt(1))
		if err != nil {
			return nil, err
		}
		v.Scope.Set(v.Name, vv)
		return v, nil
	}

	return val.UnarySuf(t)
}
