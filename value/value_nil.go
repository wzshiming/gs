package value

import (
	"github.com/wzshiming/gs/token"
)

var Nil _Nil

type _Nil struct{}

func (v _Nil) String() string {
	return "nil"
}

func (v _Nil) Point() Value {
	return v
}

func (v _Nil) Binary(t token.Token, y Value) (Value, error) {
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
	case _Nil:
		return Bool(b), nil
	default:
		return Bool(!b), nil
	}
}

func (v _Nil) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v _Nil) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
