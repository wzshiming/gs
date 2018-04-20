package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type ValueString struct {
	Val string
}

func (v *ValueString) String() string {
	return v.Val
}

func (v *ValueString) Clone() *ValueString {
	return &ValueString{
		Val: v.Val,
	}
}

func (v *ValueString) Binary(t token.Token, y Value) (Value, error) {

	sum := ""
	switch yy := y.(type) {
	case *ValueString:
		sum = yy.Val
	case *ValueVar:
		val, err := yy.Point()
		if err != nil {
			return v, err
		}
		return v.Binary(t, val)
	default:
		return v, fmt.Errorf("Type to string error")
	}

	v = v.Clone()
	switch t {
	case token.ADD:
		v.Val += sum
	default:
		return v, undefined
	}
	return v, nil
}

func (v *ValueString) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueString) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
