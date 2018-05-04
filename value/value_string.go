package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type ValueString string

func (v ValueString) String() string {
	return `"` + string(v) + `"`
}

func (v ValueString) Point() Value {
	return v
}

func (v ValueString) Clone() ValueString {
	return v
}

func (v ValueString) Binary(t token.Token, y Value) (Value, error) {

	var sum ValueString
	switch yy := y.(type) {
	case ValueString:
		sum = yy
	case *ValueVar:
		val := yy.Point()
		return v.Binary(t, val)
	default:
		return ValueNil, fmt.Errorf("Type to string error")
	}

	switch t {
	case token.ADD:
		return v + ValueString(sum), nil
	}
	return v, undefined
}

func (v ValueString) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v ValueString) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
