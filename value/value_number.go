package value

import (
	"fmt"
	"math"

	"github.com/wzshiming/gs/token"
)

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
	case *valueNil:
		if t == token.EQL {
			return ValueFalse, nil
		}
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

func (v *ValueNumber) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		v = v.Clone()
		v.Val = -v.Val
		return v, nil
	default:
		return v, undefined
	}
}

func (v *ValueNumber) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		v.Val += 1
		return v, nil
	case token.DEC:
		v.Val -= 1
		return v, nil
	default:
		return v, undefined
	}
}
