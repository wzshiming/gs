package value

import (
	"fmt"
	"reflect"

	"github.com/wzshiming/gs/token"
)

type FuncBuiltin struct {
	Val reflect.Value
}

func NewFuncBuiltin(body interface{}) *FuncBuiltin {
	return &FuncBuiltin{
		Val: reflect.ValueOf(body),
	}
}

func (v *FuncBuiltin) Call(v0 Value) (rr Value, err0 error) {
	in := []Value{}
	if b, ok := v0.(*Tuple); ok {
		in = b.List
	} else {
		in = []Value{v0}
	}
	rin := make([]reflect.Value, 0, len(in))
	for _, v := range in {
		rin = append(rin, reflect.ValueOf(v))
	}

	defer func() {

		if x := recover(); x != nil {
			err0 = fmt.Errorf("error: Builtin function :%v", x)
		}
	}()
	rout := v.Val.Call(rin)
	out := []Value{}
	for _, v := range rout {
		out = append(out, toGsValue(v))
	}
	switch len(out) {
	case 0:
		return Nil, nil
	case 1:
		return out[0], nil
	default:
		return NewTuple(out, false), nil
	}

}

func (v *FuncBuiltin) String() string {
	return "<FuncBuiltin>"
}

func (v *FuncBuiltin) Point() Value {
	return v
}

func (v *FuncBuiltin) Binary(t token.Token, y Value) (Value, error) {
	return Nil, undefined
}

func (v *FuncBuiltin) UnaryPre(t token.Token) (Value, error) {
	return Nil, undefined
}

func (v *FuncBuiltin) UnarySuf(t token.Token) (Value, error) {
	return Nil, undefined
}

func toGsValue(v reflect.Value) Value {
	switch v.Kind() {
	case reflect.Bool:
		return Bool(v.Bool())
	case reflect.String:
		return String(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return numberInt(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return numberInt(int64(v.Uint()))
	case reflect.Float32, reflect.Float64:
		return numberFloat(v.Float())
	case reflect.Func:
		return &FuncBuiltin{v}
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return Nil
		}
		return toGsValue(v.Elem())
	default:
		return String(fmt.Sprint(v.Interface()))
	}
}
