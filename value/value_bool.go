package value

import (
	"github.com/wzshiming/gs/token"
)

var (
	ValueTrue  ValueBool = true
	ValueFalse ValueBool = false
)

type ValueBool bool

func (v ValueBool) String() string {
	if v {
		return "true"
	} else {
		return "false"
	}
}

func (v ValueBool) Point() (Value, error) {
	return v, nil
}

func (v ValueBool) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v ValueBool) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v ValueBool) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
