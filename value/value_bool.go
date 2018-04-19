package value

import (
	"github.com/wzshiming/gs/token"
)

var ValueTrue = &ValueBool{true}
var ValueFalse = &ValueBool{false}

type ValueBool struct {
	Val bool
}

func (v *ValueBool) String() string {
	if v.Val {
		return "true"
	} else {
		return "false"
	}
}

func (v *ValueBool) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v *ValueBool) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueBool) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}
