package value

import (
	"fmt"
	"math"

	"github.com/wzshiming/gs/token"
)

var undefined = fmt.Errorf("Undefined operation")

type Value interface {
	fmt.Stringer
	Binary(t token.Token, y Value) (Value, error)
	PreUnary(t token.Token) (Value, error)
	SufUnary(t token.Token) (Value, error)
}

//type ValueFunc struct {
//	Val ast.Expr
//}

//func (v *ValueFunc) String() string {
//	return v.Val.String()
//}

//func (v *ValueFunc) Binary(t token.Token, y Value) ( Value,error) {
//	return v, undefined
//}

//func (v *ValueFunc) PreUnary(t token.Token) ( Value,error) {
//	return v, undefined
//}

//func (v *ValueFunc) SufUnary(t token.Token) ( Value,error) {
//	return v, undefined
//}

type ValueVar struct {
	Scope *Scope
	Name  string
}

func (v *ValueVar) String() string {
	if v == nil {
		return "<ValueVar.nil>"
	}
	val, ok := v.Scope.Get(v.Name)
	if !ok || val == nil {
		return "<ValueVar.nil>"
	}
	return val.String()
}

func (v *ValueVar) Point() (Value, error) {
	val, ok := v.Scope.Get(v.Name)
	if !ok {
		return v, fmt.Errorf("Variable is empty '%v'", v.Name)
	}
	return val, nil
}

func (v *ValueVar) Binary(t token.Token, y Value) (Value, error) {

	switch t {
	case token.ASSIGN:
		v.Scope.Set(v.Name, y)
		return v, nil
	}

	val, err := v.Point()
	if err != nil {
		return v, err
	}

	return val.Binary(t, y)
}

func (v *ValueVar) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueVar) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}

type ValueNumber struct {
	Val float64
}

func (v *ValueNumber) String() string {
	return fmt.Sprint(v.Val)
}

func (v *ValueNumber) Clone() *ValueNumber {
	return &ValueNumber{
		Val: v.Val,
	}
}

func (v *ValueNumber) Binary(t token.Token, y Value) (vv Value, err error) {
	sum := 0.0
	switch yy := y.(type) {
	case *ValueNumber:
		sum = yy.Val
	case *ValueVar:
		val, err := yy.Point()
		if err != nil {
			return v, err
		}
		return v.Binary(t, val)
	default:
		return v, fmt.Errorf("Type to number error")
	}

	switch t {
	case token.ADD:
		v = v.Clone()
		v.Val += sum
		vv = v
	case token.SUB:
		v = v.Clone()
		v.Val -= sum
		vv = v
	case token.MUL:
		v = v.Clone()
		v.Val *= sum
		vv = v
	case token.QUO:
		v = v.Clone()
		v.Val /= sum
		vv = v
	case token.POW:
		v = v.Clone()
		v.Val = math.Pow(v.Val, sum)
		vv = v
	case token.REM:
		v = v.Clone()
		v.Val = float64(int64(v.Val) % int64(sum))
		vv = v

		// 比较
	case token.EQL:
		vv = &ValueBool{Val: v.Val == sum}

	case token.LSS:
		vv = &ValueBool{Val: v.Val < sum}

	case token.GTR:
		vv = &ValueBool{Val: v.Val > sum}

	case token.NEQ:
		vv = &ValueBool{Val: v.Val != sum}

	case token.LEQ:
		vv = &ValueBool{Val: v.Val <= sum}

	case token.GEQ:
		vv = &ValueBool{Val: v.Val >= sum}

	default:
		return v, undefined
	}
	return vv, nil
}

func (v *ValueNumber) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueNumber) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}

type ValueString struct {
	Val string
}

func (v *ValueString) String() string {
	return v.Val
}

func (v *ValueString) Clone() *ValueString {
	return &ValueString{
		Val: v.Val,
	}
}

func (v *ValueString) Binary(t token.Token, y Value) (Value, error) {

	sum := ""
	switch yy := y.(type) {
	case *ValueString:
		sum = yy.Val
	case *ValueVar:
		val, err := yy.Point()
		if err != nil {
			return v, err
		}
		return v.Binary(t, val)
	default:
		return v, fmt.Errorf("Type to string error")
	}

	v = v.Clone()
	switch t {
	case token.ADD:
		v.Val += sum
	default:
		return v, undefined
	}
	return v, nil
}

func (v *ValueString) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueString) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}

type ValueBool struct {
	Val bool
}

func (v *ValueBool) String() string {
	if v.Val {
		return "true"
	} else {
		return "false"
	}
}

func (v *ValueBool) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v *ValueBool) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueBool) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}

type ValueNil struct{}

func (v *ValueNil) String() string {
	return "nil"
}

func (v *ValueNil) Binary(t token.Token, y Value) (Value, error) {
	return v, undefined
}

func (v *ValueNil) PreUnary(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueNil) SufUnary(t token.Token) (Value, error) {
	return v, undefined
}
