package value

import (
	"fmt"
	"math/big"

	"github.com/wzshiming/gs/token"
)

type numberBigFloat struct {
	Val *big.Float
}

func newNumberBigFloat(f float64) numberBigFloat {
	return numberBigFloat{
		Val: big.NewFloat(f),
	}
}

func (v numberBigFloat) String() string {
	return v.Val.String()
}

func (v numberBigFloat) Point() Value {
	return v
}

func (v numberBigFloat) Int() numberInt {
	val, _ := v.Val.Int64()
	return numberInt(val)
}

func (v numberBigFloat) Float() numberFloat {
	val, _ := v.Val.Float64()
	return numberFloat(val)
}

func (v numberBigFloat) BigInt() numberBigInt {
	z, _ := v.Val.Int(nil)
	return numberBigInt{z}
}

func (v numberBigFloat) BigFloat() numberBigFloat {
	return v
}

func (v numberBigFloat) Binary(t token.Token, y Value) (vv Value, err error) {
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
		v0 := big.NewFloat(0)
		return numberBigFloat{v0.Add(v.Val, sum.BigFloat().Val)}, nil
	case token.SUB:
		v0 := big.NewFloat(0)
		return numberBigFloat{v0.Sub(v.Val, sum.BigFloat().Val)}, nil
	case token.MUL:
		v0 := big.NewFloat(0)
		return numberBigFloat{v0.Mul(v.Val, sum.BigFloat().Val)}, nil
	case token.QUO:
		v0 := big.NewFloat(0)
		return numberBigFloat{v0.Quo(v.Val, sum.BigFloat().Val)}, nil
		//	case token.POW:
		//		v0 := big.NewFloat(1)
		//		vv = valueNumberBigFloat{v.Val.Sqrt(v0.Quo(v0, sum))}
	case token.REM:
		return v.BigInt().Binary(t, y)

		// 比较
	case token.EQL:
		return Bool(v.Val.Cmp(sum.BigFloat().Val) == 0), nil

	case token.LSS:
		return Bool(v.Val.Cmp(sum.BigFloat().Val) < 0), nil

	case token.GTR:
		return Bool(v.Val.Cmp(sum.BigFloat().Val) > 0), nil

	case token.NEQ:
		return Bool(v.Val.Cmp(sum.BigFloat().Val) != 0), nil

	case token.LEQ:
		return Bool(v.Val.Cmp(sum.BigFloat().Val) <= 0), nil

	case token.GEQ:
		return Bool(v.Val.Cmp(sum.BigFloat().Val) >= 0), nil

	default:
		return Nil, undefined
	}
}

func (v numberBigFloat) UnaryPre(t token.Token) (Value, error) {

	switch t {
	case token.ADD:
		return v, nil
	case token.SUB:
		v.Val.Mul(v.Val, big.NewFloat(-1))
		return v, nil
	default:
		return Nil, undefined
	}
}

func (v numberBigFloat) UnarySuf(t token.Token) (Value, error) {
	switch t {
	case token.INC:
		v.Val.Add(v.Val, big.NewFloat(1))
		return v, nil
	case token.DEC:
		v.Val.Sub(v.Val, big.NewFloat(1))
		return v, nil
	default:
		return Nil, undefined
	}
}
