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
		v0 := big.NewInt(0)
		vv = valueNumberBigInt{v0.Add(v.Val, sum.BigInt().Val)}
	case token.SUB:
		v0 := big.NewInt(0)
		vv = valueNumberBigInt{v0.Sub(v.Val, sum.BigInt().Val)}
	case token.MUL:
		v0 := big.NewInt(0)
		vv = valueNumberBigInt{v0.Mul(v.Val, sum.BigInt().Val)}
	case token.QUO:
		vv = valueNumberBigInt{v.Val.Quo(v.Val, sum.BigInt().Val)}
		//	case token.POW:
		//		v0 := big.NewInt(1)
		//		vv = valueNumberBigInt{v.Val.Sqrt(v0.Quo(v0, sum))}
	case token.REM:
		vv = valueNumberBigInt{v.Val.Rem(v.Val, sum.BigInt().Val)}

		// 比较
	case token.EQL:
		vv = ValueBool(v.Val.Cmp(sum.BigInt().Val) == 0)

	case token.LSS:
		vv = ValueBool(v.Val.Cmp(sum.BigInt().Val) < 0)

	case token.GTR:
		vv = ValueBool(v.Val.Cmp(sum.BigInt().Val) > 0)

	case token.NEQ:
		vv = ValueBool(v.Val.Cmp(sum.BigInt().Val) != 0)

	case token.LEQ:
		vv = ValueBool(v.Val.Cmp(sum.BigInt().Val) <= 0)

	case token.GEQ:
		vv = ValueBool(v.Val.Cmp(sum.BigInt().Val) >= 0)

	default:
		return v, undefined
	}
	return vv, nil
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
