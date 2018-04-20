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

func (v *valueNil) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *valueNil) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
