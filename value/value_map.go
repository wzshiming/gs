package value

import (
	"bytes"

	"github.com/wzshiming/gs/token"
)

type ValueMap map[Value]Value

func (v ValueMap) String() string {
	if v == nil {
		return "<nil.ValueMap>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("{\n")
	for k, v := range v {
		buf.WriteString(k.String())
		buf.WriteString(": ")
		buf.WriteString(v.String())
		buf.WriteString(",\n")
	}
	buf.WriteString("}")
	return buf.String()
}

func (v ValueMap) Point() Value {
	return v
}

func (v ValueMap) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v ValueMap) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v ValueMap) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
