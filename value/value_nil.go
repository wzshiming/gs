package value

import (
	"github.com/wzshiming/gs/token"
)

var ValueNil valueNil

type valueNil struct{}

func (v valueNil) String() string {
	return "nil"
}

func (v valueNil) Point() Value {
	return v
}

func (v valueNil) Binary(t token.Token, y Value) (Value, error) {
	b := false
	switch t {
	case token.EQL:
		b = true
	case token.NEQ:
		b = false
	default:
		return v, undefined
	}

	y0 := y.Point()

	switch y0.(type) {
	case valueNil:
		return ValueBool(b), nil
	default:
		return ValueBool(!b), nil
	}
}

func (v valueNil) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v valueNil) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
