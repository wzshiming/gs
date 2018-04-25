package value

import (
	"bytes"
	"fmt"

	"github.com/wzshiming/gs/token"
)

type ValueTuple struct {
	List []Value
}

func (v *ValueTuple) String() string {
	if v == nil {
		return "<nil.ValueTuple>"
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('(')
	for k, v := range v.List {
		if k != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteString(")")
	return buf.String()
}

func (v *ValueTuple) Point() (Value, error) {
	return v, nil
}

func (v *ValueTuple) Len() int {
	return len(v.List)
}

func (v *ValueTuple) Index(i int) Value {
	return v.List[i]
}

func (v *ValueTuple) Binary(t token.Token, y Value) (Value, error) {
	var vt *ValueTuple
	switch yy := y.(type) {
	case *ValueTuple:
		vt = yy
	default:
		return ValueNil, fmt.Errorf("Type to Tuple error")
	}

	if len(v.List) != len(vt.List) {
		return ValueNil, fmt.Errorf("Tuple The length is different")
	}

	tmp := make([]Value, 0, len(vt.List))
	for _, v := range vt.List {
		yy, err := v.Point()
		if err != nil {
			return ValueNil, err
		}
		tmp = append(tmp, yy)
	}
	for i, v0 := range v.List {
		ov, err := v0.Binary(t, tmp[i])
		if err != nil {
			return ValueNil, err
		}
		tmp[i] = ov
	}

	return &ValueTuple{tmp}, nil
}

func (v *ValueTuple) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueTuple) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
