package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type valueNumberInt int64

func newValueNumberInt(f int64) valueNumberInt {
	return valueNumberInt(f)
}

func (v valueNumberInt) String() string {
	return fmt.Sprint(int64(v))
}

func (v valueNumberInt) Int() valueNumberInt {
	return v
}

func (v valueNumberInt) Float() valueNumberFloat {
	return valueNumberFloat(v)
}

func (v valueNumberInt) BigInt() valueNumberBigInt {
	return newValueNumberBigInt(int64(v))
}

func (v valueNumberInt) BigFloat() valueNumberBigFloat {
	return newValueNumberBigFloat(float64(v))
}

func (v valueNumberInt) Binary(t token.Token, y Value) (vv Value, err error) {
	var sum ValueNumber
	switch yy := y.(type) {
	case ValueNumber:
		sum = yy
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
		if v > maxInt {
			return v.BigInt().Binary(t, sum)
		} else {
			return v + sum.Int(), nil
		}
	case token.SUB:
		if v < minInt {
			return v.BigInt().Binary(t, sum)
		} else {
			return v - sum.Int(), nil
		}
	case token.MUL:
		if v > maxInt {
			return v.BigInt().Binary(t, sum)
		} else {
			return v * sum.Int(), nil
		}
	case token.QUO:
		if v < minInt {
			return v.BigInt().Binary(t, sum)
		} else {
			return v / sum.Int(), nil
		}
	case token.REM:
		return v % sum.Int(), nil

		// 比较
	case token.EQL:
		return ValueBool(v == sum.Int()), nil

	case token.LSS:
		return ValueBool(v < sum.Int()), nil

	case token.GTR:
		return ValueBool(v > sum.Int()), nil

	case token.NEQ:
		return ValueBool(v != sum.Int()), nil

	case token.LEQ:
		return ValueBool(v <= sum.Int()), nil

	case token.GEQ:
		return ValueBool(v >= sum.Int()), nil

	default:
		return v, undefined
	}
}

func (v valueNumberInt) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		return -v, nil
	default:
		return v, undefined
	}
}

func (v valueNumberInt) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		return v + 1, nil
	case token.DEC:
		return v - 1, nil
	default:
		return v, undefined
	}
}
