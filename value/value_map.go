package value

import (
	"bytes"

	"github.com/wzshiming/gs/token"
)

type Map map[Value]Value

func (v Map) String() string {
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

func (v Map) Point() Value {
	return v
}

func (v Map) Binary(t token.Token, y Value) (Value, error) {
	return Nil, undefined
}

func (v Map) UnaryPre(t token.Token) (Value, error) {
	return Nil, undefined
}

func (v Map) UnarySuf(t token.Token) (Value, error) {
	return Nil, undefined
}
