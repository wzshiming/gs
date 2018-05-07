package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type String string

func (v String) String() string {
	return `"` + string(v) + `"`
}

func (v String) Point() Value {
	return v
}

func (v String) Clone() String {
	return v
}

func (v String) Binary(t token.Token, y Value) (Value, error) {

	var sum String
	switch yy := y.(type) {
	case String:
		sum = yy
	case *Var:
		val := yy.Point()
		return v.Binary(t, val)
	default:
		return Nil, fmt.Errorf("Type to string error")
	}

	switch t {
	case token.ADD:
		return v + String(sum), nil
	}
	return v, undefined
}

func (v String) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v String) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
