package value

import (
	"bytes"
	"fmt"

	"github.com/wzshiming/gs/token"
)

func NewTuple(vs []Value, ellip bool) Value {
	switch len(vs) {
	case 0:
		return Nil
	case 1:
		return vs[0]
	default:
		return &Tuple{
			List:     vs,
			Ellipsis: ellip,
		}
	}
}

type Tuple struct {
	List     []Value
	Ellipsis bool
}

func (v *Tuple) String() string {
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

func (v *Tuple) Point() Value {
	return v
}

func (v *Tuple) Len() int {
	return len(v.List)
}

func (v *Tuple) Index(i int) Value {
	return v.List[i]
}

func (v *Tuple) Binary(t token.Token, y Value) (Value, error) {
	var vt *Tuple
	switch yy := y.(type) {
	case *Tuple:
		vt = yy
	default:
		return Nil, fmt.Errorf("Type to Tuple error")
	}

	tmp := make([]Value, 0, len(vt.List))
	for _, v := range vt.List {
		yy := v.Point()
		tmp = append(tmp, yy)
	}
	es := []int{}
	for i, v0 := range v.List {
		switch t := v0.(type) {
		case *Var:
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
		tmp0 = append(tmp0, NewTuple(m, false))
		tmp0 = append(tmp0, r...)
		tmp = tmp0
	default:
		return Nil, fmt.Errorf("Only one omitted parameter is allowed for the left value")
	}

	if len(v.List) != len(tmp) {
		return Nil, fmt.Errorf("Tuple The length is different")
	}
	for i, v0 := range v.List {
		ov, err := v0.Binary(t, tmp[i])
		if err != nil {
			return Nil, err
		}
		tmp[i] = ov
	}
	return NewTuple(tmp, false), nil

}

func (v *Tuple) UnaryPre(t token.Token) (Value, error) {
	return Nil, undefined
}

func (v *Tuple) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.ELLIPSIS:
		v.Ellipsis = true
		return v, nil
	}
	return Nil, undefined
}
