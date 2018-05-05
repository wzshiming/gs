package value

import (
	"bytes"
	"fmt"

	"github.com/wzshiming/gs/token"
)

func NewValueTuple(vs []Value, ellip bool) Value {
	switch len(vs) {
	case 0:
		return ValueNil
	case 1:
		return vs[0]
	default:
		return &ValueTuple{
			List:     vs,
			Ellipsis: ellip,
		}
	}
}

type ValueTuple struct {
	List     []Value
	Ellipsis bool
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

func (v *ValueTuple) Point() Value {
	return v
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

	tmp := make([]Value, 0, len(vt.List))
	for _, v := range vt.List {
		yy := v.Point()
		tmp = append(tmp, yy)
	}
	es := []int{}
	for i, v0 := range v.List {
		switch t := v0.(type) {
		case *ValueVar:
			if t.Ellipsis {
				t.Ellipsis = false
				es = append(es, i)
			}
		}
	}
	switch len(es) {
	case 0:
	case 1:
		e := es[0]
		ll := len(tmp) - len(v.List)
		l := tmp[:e]
		m := tmp[e : e+ll+1]
		r := tmp[e+ll+1:]
		tmp0 := make([]Value, 0, len(l)+len(r)+1)
		tmp0 = append(tmp0, l...)
		tmp0 = append(tmp0, NewValueTuple(m, false))
		tmp0 = append(tmp0, r...)
		tmp = tmp0
	default:
		return ValueNil, fmt.Errorf("Only one omitted parameter is allowed for the left value")
	}

	if len(v.List) != len(tmp) {
		return ValueNil, fmt.Errorf("Tuple The length is different")
	}
	for i, v0 := range v.List {
		ov, err := v0.Binary(t, tmp[i])
		if err != nil {
			return ValueNil, err
		}
		tmp[i] = ov
	}
	return NewValueTuple(tmp, false), nil

}

func (v *ValueTuple) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueTuple) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.ELLIPSIS:
		v.Ellipsis = true
		return v, nil
	}
	return v, undefined
}
