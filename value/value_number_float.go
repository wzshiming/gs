package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type numberFloat float64

func newNumberFloat(f float64) numberFloat {
	return numberFloat(f)
}

func (v numberFloat) String() string {
	return fmt.Sprint(float64(v))
}

func (v numberFloat) Point() Value {
	return v
}

func (v numberFloat) Int() numberInt {
	return numberInt(v)
}

func (v numberFloat) Float() numberFloat {
	return v
}

func (v numberFloat) BigInt() numberBigInt {
	return newNumberBigInt(int64(v))
}

func (v numberFloat) BigFloat() numberBigFloat {
	return newNumberBigFloat(float64(v))
}

func (v numberFloat) Binary(t token.Token, y Value) (vv Value, err error) {
	var sum Number
	switch yy := y.(type) {
	case Number:
		sum = yy
	case *Var:
		val := yy.Point()
		return v.Binary(t, val)
	case *_Nil:
		switch t {
		case token.EQL:
			return False, nil
		case token.NEQ:
			return True, nil
		default:
			return Nil, fmt.Errorf("Type to number error")
		}
	default:
		return Nil, fmt.Errorf("Type to number error")
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
		return Bool(v == sum.Float()), nil

	case token.LSS:
		return Bool(v < sum.Float()), nil

	case token.GTR:
		return Bool(v > sum.Float()), nil

	case token.NEQ:
		return Bool(v != sum.Float()), nil

	case token.LEQ:
		return Bool(v <= sum.Float()), nil

	case token.GEQ:
		return Bool(v >= sum.Float()), nil

	default:
		return v, undefined
	}
}

func (v numberFloat) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		return -v, nil
	default:
		return v, undefined
	}
}

func (v numberFloat) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		return v + 1, nil
	case token.DEC:
		return v - 1, nil
	default:
		return v, undefined
	}
}
