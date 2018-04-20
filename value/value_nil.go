package value

import (
	"github.com/wzshiming/gs/token"
)

var ValueNil *valueNil

type valueNil struct{}

func (v *valueNil) String() string {
	return "nil"
}

func (v *valueNil) Binary(t token.Token, y Value) (Value, error) {
	if t != token.EQL {
		return v, undefined
	}

	switch yy := y.(type) {
	case *ValueVar:
		val, err := yy.Point()
		if err != nil {
			return v, err
		}
		return v.Binary(t, val)
	case *valueNil:
		return ValueTrue, nil
	}

	return ValueFalse, nil
}

func (v *valueNil) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *valueNil) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
