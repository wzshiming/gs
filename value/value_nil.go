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
	return v, undefined
}

func (v *valueNil) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *valueNil) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}
