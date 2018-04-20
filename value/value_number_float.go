package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type valueNumberFloat float64

func newValueNumberFloat(f float64) valueNumberFloat {
	return valueNumberFloat(f)
}

func (v valueNumberFloat) String() string {
	return fmt.Sprint(float64(v))
}

func (v valueNumberFloat) Int() valueNumberInt {
	return valueNumberInt(v)
}

func (v valueNumberFloat) Float() valueNumberFloat {
	return v
}

func (v valueNumberFloat) BigInt() valueNumberBigInt {
	return newValueNumberBigInt(int64(v))
}

func (v valueNumberFloat) BigFloat() valueNumberBigFloat {
	return newValueNumberBigFloat(float64(v))
}

func (v valueNumberFloat) Binary(t token.Token, y Value) (vv Value, err error) {
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
		if v > maxFloat {
			return v.BigFloat().Binary(t, sum)
		} else {
			return v + sum.Float(), nil
		}
	case token.SUB:
		if v < minFloat {
			return v.BigFloat().Binary(t, sum)
		} else {
			return v - sum.Float(), nil
		}
	case token.MUL:
		if v > maxFloat {
			return v.BigFloat().Binary(t, sum)
		} else {
			return v * sum.Float(), nil
		}
	case token.QUO:
		if v < minFloat {
			return v.BigFloat().Binary(t, sum)
		} else {
			return v / sum.Float(), nil
		}
	case token.REM:
		return v.BigFloat().Binary(t, sum)

		// 比较
	case token.EQL:
		return ValueBool(v == sum.Float()), nil

	case token.LSS:
		return ValueBool(v < sum.Float()), nil

	case token.GTR:
		return ValueBool(v > sum.Float()), nil

	case token.NEQ:
		return ValueBool(v != sum.Float()), nil

	case token.LEQ:
		return ValueBool(v <= sum.Float()), nil

	case token.GEQ:
		return ValueBool(v >= sum.Float()), nil

	default:
		return v, undefined
	}
}

func (v valueNumberFloat) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		return -v, nil
	default:
		return v, undefined
	}
}

func (v valueNumberFloat) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		return v + 1, nil
	case token.DEC:
		return v - 1, nil
	default:
		return v, undefined
	}
}
