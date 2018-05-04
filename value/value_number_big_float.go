package value

import (
	"fmt"
	"math/big"

	"github.com/wzshiming/gs/token"
)

type valueNumberBigFloat struct {
	Val *big.Float
}

func newValueNumberBigFloat(f float64) valueNumberBigFloat {
	return valueNumberBigFloat{
		Val: big.NewFloat(f),
	}
}

func (v valueNumberBigFloat) String() string {
	return v.Val.String()
}

func (v valueNumberBigFloat) Point() Value {
	return v
}

func (v valueNumberBigFloat) Int() valueNumberInt {
	val, _ := v.Val.Int64()
	return valueNumberInt(val)
}

func (v valueNumberBigFloat) Float() valueNumberFloat {
	val, _ := v.Val.Float64()
	return valueNumberFloat(val)
}

func (v valueNumberBigFloat) BigInt() valueNumberBigInt {
	z, _ := v.Val.Int(nil)
	return valueNumberBigInt{z}
}

func (v valueNumberBigFloat) BigFloat() valueNumberBigFloat {
	return v
}

func (v valueNumberBigFloat) Binary(t token.Token, y Value) (vv Value, err error) {
	var sum ValueNumber
	switch yy := y.(type) {
	case ValueNumber:
		sum = yy
	case *ValueVar:
		val := yy.Point()
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
		v0 := big.NewFloat(0)
		return valueNumberBigFloat{v0.Add(v.Val, sum.BigFloat().Val)}, nil
	case token.SUB:
		v0 := big.NewFloat(0)
		return valueNumberBigFloat{v0.Sub(v.Val, sum.BigFloat().Val)}, nil
	case token.MUL:
		v0 := big.NewFloat(0)
		return valueNumberBigFloat{v0.Mul(v.Val, sum.BigFloat().Val)}, nil
	case token.QUO:
		v0 := big.NewFloat(0)
		return valueNumberBigFloat{v0.Quo(v.Val, sum.BigFloat().Val)}, nil
		//	case token.POW:
		//		v0 := big.NewFloat(1)
		//		vv = valueNumberBigFloat{v.Val.Sqrt(v0.Quo(v0, sum))}
	case token.REM:
		return v.BigInt().Binary(t, y)

		// 比较
	case token.EQL:
		return ValueBool(v.Val.Cmp(sum.BigFloat().Val) == 0), nil

	case token.LSS:
		return ValueBool(v.Val.Cmp(sum.BigFloat().Val) < 0), nil

	case token.GTR:
		return ValueBool(v.Val.Cmp(sum.BigFloat().Val) > 0), nil

	case token.NEQ:
		return ValueBool(v.Val.Cmp(sum.BigFloat().Val) != 0), nil

	case token.LEQ:
		return ValueBool(v.Val.Cmp(sum.BigFloat().Val) <= 0), nil

	case token.GEQ:
		return ValueBool(v.Val.Cmp(sum.BigFloat().Val) >= 0), nil

	default:
		return v, undefined
	}
}

func (v valueNumberBigFloat) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		v.Val.Mul(v.Val, big.NewFloat(-1))
		return v, nil
	default:
		return v, undefined
	}
}

func (v valueNumberBigFloat) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		v.Val.Add(v.Val, big.NewFloat(1))
		return v, nil
	case token.DEC:
		v.Val.Sub(v.Val, big.NewFloat(1))
		return v, nil
	default:
		return v, undefined
	}
}
