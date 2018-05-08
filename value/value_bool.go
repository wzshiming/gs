package value

import (
	"github.com/wzshiming/gs/token"
)

// Boolean constant definition
var (
	True  Bool = true
	False Bool = false
)

type Bool bool

func (v Bool) String() string {
	if v {
		return "true"
	} else {
		return "false"
	}
}

func (v Bool) Point() Value {
	return v
}

func (v Bool) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v Bool) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v Bool) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
