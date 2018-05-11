package value

import (
	"bytes"

	"github.com/wzshiming/gs/token"
)

var (
	_ Assigner = Map{}
)

type Map map[Value]Value

func (m Map) String() string {
	if m == nil {
		return "<nil.ValueMap>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("{\n")
	for k, v := range m {
		buf.WriteString(k.String())
		buf.WriteString(": ")
		buf.WriteString(v.String())
		buf.WriteString(",\n")
	}
	buf.WriteString("}")
	return buf.String()
}

func (m Map) Point() Value {
	return m
}

func (m Map) Set(k Value, v Value) {
	m[k] = v
}

func (m Map) SetLocal(k Value, v Value) {
	m[k] = v
}

func (m Map) Get(k Value) (Value, bool) {
	v, ok := m[k]
	return v, ok
}

func (m Map) Child() Assigner {
	return m
}

func (m Map) Binary(t token.Token, y Value) (Value, error) {
	return Nil, undefined
}

func (m Map) UnaryPre(t token.Token) (Value, error) {
	return Nil, undefined
}

func (m Map) UnarySuf(t token.Token) (Value, error) {
	return Nil, undefined
}
