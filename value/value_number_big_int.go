package value

import (
	"fmt"
	"math/big"

	"github.com/wzshiming/gs/token"
)

type numberBigInt struct {
	Val *big.Int
}

func newNumberBigInt(f int64) numberBigInt {
	return numberBigInt{
		Val: big.NewInt(f),
	}
}

func (v numberBigInt) String() string {
	return v.Val.String()
}

func (v numberBigInt) Point() Value {
	return v
}

func (v numberBigInt) Int() numberInt {
	return numberInt(v.Val.Int64())
}

func (v numberBigInt) Float() numberFloat {
	return numberFloat(v.Val.Int64())
}

func (v numberBigInt) BigInt() numberBigInt {
	return v
}

func (v numberBigInt) BigFloat() numberBigFloat {
	f, _, _ := big.ParseFloat(v.Val.String(), 0, 0, big.ToNearestEven)
	return numberBigFloat{f}
}

func (v numberBigInt) Binary(t token.Token, y Value) (vv Value, err error) {
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
			return ValueFalse, nil
		case token.NEQ:
			return ValueTrue, nil
		default:
			return Nil, fmt.Errorf("Type to number error")
		}
	default:
		return Nil, fmt.Errorf("Type to number error")
	}

	switch t {
	case token.ADD:
		v0 := big.NewInt(0)
		return numberBigInt{v0.Add(v.Val, sum.BigInt().Val)}, nil
	case token.SUB:
		v0 := big.NewInt(0)
		return numberBigInt{v0.Sub(v.Val, sum.BigInt().Val)}, nil
	case token.MUL:
		v0 := big.NewInt(0)
		return numberBigInt{v0.Mul(v.Val, sum.BigInt().Val)}, nil
	case token.QUO:
		return v.BigFloat().Binary(t, y)

		//	case token.POW:
		//		v0 := big.NewInt(1)
		//		vv = valueNumberBigInt{v.Val.Sqrt(v0.Quo(v0, sum))}
	case token.REM:
		return numberBigInt{v.Val.Rem(v.Val, sum.BigInt().Val)}, nil

		// 比较
	case token.EQL:
		return Bool(v.Val.Cmp(sum.BigInt().Val) == 0), nil

	case token.LSS:
		return Bool(v.Val.Cmp(sum.BigInt().Val) < 0), nil

	case token.GTR:
		return Bool(v.Val.Cmp(sum.BigInt().Val) > 0), nil

	case token.NEQ:
		return Bool(v.Val.Cmp(sum.BigInt().Val) != 0), nil

	case token.LEQ:
		return Bool(v.Val.Cmp(sum.BigInt().Val) <= 0), nil

	case token.GEQ:
		return Bool(v.Val.Cmp(sum.BigInt().Val) >= 0), nil

	default:
		return v, undefined
	}
}

func (v numberBigInt) UnaryPre(t token.Token) (Value, error) {

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

func (v numberBigInt) UnarySuf(t token.Token) (Value, error) {
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
