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
	for _, v := range v.List {
		buf.WriteString(v.String())
		buf.WriteString(", ")
	}
	return buf.String()
}

func (v *ValueTuple) Binary(t token.Token, y Value) (Value, error) {
	var vt *ValueTuple
	switch yy := y.(type) {
	case *ValueTuple:
		vt = yy
	default:
		return v, fmt.Errorf("Type to Tuple error")
	}

	if len(v.List) != len(vt.List) {
		return v, fmt.Errorf("Tuple The length is different")
	}

	vv := &ValueTuple{}
	for i, v := range v.List {
		ov, err := v.Binary(t, vt.List[i])
		if err != nil {
			return nil, err
		}
		vv.List = append(vv.List, ov)
	}
	return vv, nil
}

func (v *ValueTuple) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueTuple) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
