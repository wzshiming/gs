package value

import (
	"fmt"
	"math/big"

	"github.com/wzshiming/gs/token"
)

type valueNumberBigInt struct {
	Val *big.Int
}

func newValueNumberBigInt(f int64) valueNumberBigInt {
	return valueNumberBigInt{
		Val: big.NewInt(f),
	}
}

func (v valueNumberBigInt) String() string {
	return v.Val.String()
}

func (v valueNumberBigInt) Point() (Value, error) {
	return v, nil
}

func (v valueNumberBigInt) Int() valueNumberInt {
	return valueNumberInt(v.Val.Int64())
}

func (v valueNumberBigInt) Float() valueNumberFloat {
	return valueNumberFloat(v.Val.Int64())
}

func (v valueNumberBigInt) BigInt() valueNumberBigInt {
	return v
}

func (v valueNumberBigInt) BigFloat() valueNumberBigFloat {
	f, _, _ := big.ParseFloat(v.Val.String(), 0, 0, big.ToNearestEven)
	return valueNumberBigFloat{f}
}

func (v valueNumberBigInt) Binary(t token.Token, y Value) (vv Value, err error) {
	var sum ValueNumber
	switch yy := y.(type) {
	case ValueNumber:
		sum = yy
	case *ValueVar:
		val, err := yy.Point()
		if err != nil {
			return ValueNil, err
		}
		return v.Binary(t, val)
	case *valueNil:
		switch t {
		case token.EQL:
			return ValueFalse, nil
		case token.NEQ:
			return ValueTrue, nil
		default:
			return ValueNil, fmt.Errorf("Type to number error")
		}
	default:
		return ValueNil, fmt.Errorf("Type to number error")
	}

	switch t {
	case token.ADD:
		v0 := big.NewInt(0)
		return valueNumberBigInt{v0.Add(v.Val, sum.BigInt().Val)}, nil
	case token.SUB:
		v0 := big.NewInt(0)
		return valueNumberBigInt{v0.Sub(v.Val, sum.BigInt().Val)}, nil
	case token.MUL:
		v0 := big.NewInt(0)
		return valueNumberBigInt{v0.Mul(v.Val, sum.BigInt().Val)}, nil
	case token.QUO:
		return v.BigFloat().Binary(t, y)

		//	case token.POW:
		//		v0 := big.NewInt(1)
		//		vv = valueNumberBigInt{v.Val.Sqrt(v0.Quo(v0, sum))}
	case token.REM:
		return valueNumberBigInt{v.Val.Rem(v.Val, sum.BigInt().Val)}, nil

		// 比较
	case token.EQL:
		return ValueBool(v.Val.Cmp(sum.BigInt().Val) == 0), nil

	case token.LSS:
		return ValueBool(v.Val.Cmp(sum.BigInt().Val) < 0), nil

	case token.GTR:
		return ValueBool(v.Val.Cmp(sum.BigInt().Val) > 0), nil

	case token.NEQ:
		return ValueBool(v.Val.Cmp(sum.BigInt().Val) != 0), nil

	case token.LEQ:
		return ValueBool(v.Val.Cmp(sum.BigInt().Val) <= 0), nil

	case token.GEQ:
		return ValueBool(v.Val.Cmp(sum.BigInt().Val) >= 0), nil

	default:
		return v, undefined
	}
}

func (v valueNumberBigInt) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		v.Val.Mul(v.Val, big.NewInt(-1))
		return v, nil
	default:
		return v, undefined
	}
}

func (v valueNumberBigInt) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		v.Val.Add(v.Val, big.NewInt(1))
		return v, nil
	case token.DEC:
		v.Val.Sub(v.Val, big.NewInt(1))
		return v, nil
	default:
		return v, undefined
	}
}
