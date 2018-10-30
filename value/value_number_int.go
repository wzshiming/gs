package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type numberInt int64

func (v numberInt) String() string {
	return fmt.Sprint(int64(v))
}

func (v numberInt) Point() Value {
	return v
}

func (v numberInt) Int() numberInt {
	return v
}

func (v numberInt) Float() numberFloat {
	return numberFloat(v)
}

func (v numberInt) Binary(t token.Token, y Value) (vv Value, err error) {
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
		return v + sum.Int(), nil
	case token.SUB:
		return v - sum.Int(), nil
	case token.MUL:
		return v * sum.Int(), nil
	case token.QUO:
		return v / sum.Int(), nil
	case token.REM:
		return v % sum.Int(), nil

		// 比较
	case token.EQL:
		return Bool(v == sum.Int()), nil

	case token.LSS:
		return Bool(v < sum.Int()), nil

	case token.GTR:
		return Bool(v > sum.Int()), nil

	case token.NEQ:
		return Bool(v != sum.Int()), nil

	case token.LEQ:
		return Bool(v <= sum.Int()), nil

	case token.GEQ:
		return Bool(v >= sum.Int()), nil

	default:
		return Nil, undefined
	}
}

func (v numberInt) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		return -v, nil
	default:
		return Nil, undefined
	}
}

func (v numberInt) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		return v + 1, nil
	case token.DEC:
		return v - 1, nil
	default:
		return Nil, undefined
	}
}
