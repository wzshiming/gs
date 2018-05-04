package value

import (
	"fmt"
	"reflect"

	"github.com/wzshiming/gs/token"
)

type ValueFuncBuiltin struct {
	Val reflect.Value
}

func NewValueFuncBuiltin(body interface{}) *ValueFuncBuiltin {
	return &ValueFuncBuiltin{
		Val: reflect.ValueOf(body),
	}
}

func (v *ValueFuncBuiltin) Call(v0 Value) (rr Value, err0 error) {
	in := []Value{}
	if b, ok := v0.(*ValueTuple); ok {
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
		return ValueNil, nil
	case 1:
		return out[0], nil
	default:
		return &ValueTuple{out, false}, nil
	}

}

func (v *ValueFuncBuiltin) String() string {
	return "<FuncBuiltin>"
}

func (v *ValueFuncBuiltin) Point() Value {
	return v
}

func (v *ValueFuncBuiltin) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v *ValueFuncBuiltin) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueFuncBuiltin) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}

func toGsValue(v reflect.Value) Value {
	switch v.Kind() {
	case reflect.Bool:
		return ValueBool(v.Bool())
	case reflect.String:
		return ValueString(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return newValueNumberInt(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return newValueNumberInt(int64(v.Uint()))
	case reflect.Float32, reflect.Float64:
		return newValueNumberFloat(v.Float())
	case reflect.Func:
		return &ValueFuncBuiltin{v}
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return ValueNil
		}
		return toGsValue(v.Elem())
	default:
		return ValueString(fmt.Sprint(v.Interface()))
	}
}
